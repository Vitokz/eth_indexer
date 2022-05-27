package indexer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/Vitokz/eth_indexer/config"
	"github.com/Vitokz/eth_indexer/db"
	"github.com/Vitokz/eth_indexer/helpers"
	"github.com/Vitokz/eth_indexer/workers/indexer/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	web3types "github.com/ethereum/go-ethereum/core/types"
	web3 "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/tendermint/tendermint/libs/log"
	tmRpc "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const (
	timeToCheckLatestBlock = 1 * time.Second
)

// Worker is are indexer structure
type Worker struct {
	mut sync.Mutex
	wg  sync.WaitGroup

	errch  chan error
	stopch chan int

	web3JsonRpc *web3.Client
	ethJsonRpc  *rpc.Client
	tmJsonRpc   *tmRpc.HTTP
	db          *db.DB
	cdc         *params.EncodingConfig
	ctx         context.Context
	txHandler   TxHandler
	logger      log.Logger

	currentHeight uint64
	latestHeight  uint64
}

func New(ctx context.Context, cdc *params.EncodingConfig, currentBLock *uint64, cfg config.ConfigI) (*Worker, error) {
	httpClient := &http.Client{}
	tmRpcClient, err := tmRpc.NewWithClient(cfg.GetTendermintRPC(), cfg.GetTendermintRPC(), httpClient)
	if err != nil {
		return nil, err
	}

	db, err := db.NewDB(cfg.GetPgDbDsn())
	if err != nil {
		return nil, err
	}

	web3Client, err := web3.Dial(cfg.GetEthereumJsonRPC())
	if err != nil {
		return nil, err
	}

	ethJsonrpc, err := rpc.DialHTTP(cfg.GetEthereumJsonRPC())
	if err != nil {
		return nil, err
	}

	w := &Worker{
		mut:    sync.Mutex{},
		errch:  make(chan error),
		stopch: make(chan int),

		web3JsonRpc: web3Client,
		ethJsonRpc:  ethJsonrpc,
		tmJsonRpc:   tmRpcClient,
		cdc:         cdc,
		ctx:         ctx,
		db:          db,

		currentHeight: 1,
		latestHeight:  1,
	}

	w.txHandler = NewTxHandler(ctx, w)
	if currentBLock != nil {
		w.currentHeight = *currentBLock
	}

	return w, err
}

func (w *Worker) Start() {
	w.ReCheckLatestBlock()
}

func (w *Worker) Shutdown() {

}

// ReCheckLatestBlock in an eternal loop checks the number of the last block
func (w *Worker) ReCheckLatestBlock() {
	var (
		ticker = time.NewTicker(timeToCheckLatestBlock)
		ctx    = context.Background()
	)

	go func() {
		for {
			select {
			case <-ticker.C:
				block, err := w.web3JsonRpc.BlockByNumber(ctx, big.NewInt(-1))
				if err != nil {
					w.errch <- errors.New("failed to get latest block height")
				}
				number := block.Header().Number
				if err != nil {
					w.errch <- err
				}
				w.mut.Lock()
				w.latestHeight = number.Uint64()
				w.mut.Unlock()
			}
		}
	}()

	return
}

// CheckCurrentBlock get all information about the results of block execution
func (w *Worker) CheckCurrentBlock() {
	go func() {
		for {
			if w.GetCurrentHeight() > w.GetLatestHeight() {
				continue
			}

			var (
				height  = int64(w.GetCurrentHeight())
				tmBlock = w.fetchTendermintBlock(height)

				resultBlockChan   = make(chan *ctypes.ResultBlockResults)
				txsChan           = make(chan types.Txs)
				metaChan          = make(chan *types.BlockMeta)
				web3BlockChan     = make(chan *web3types.Block)
				web3BlockReceipts = make(chan web3types.Receipts)
			)

			// get block info fom tendermint
			go w.fetchTendermintBlockTxResults(height, resultBlockChan)
			go w.fetchTendermintBLockTxs(height, len(tmBlock.Block.Txs), txsChan)
			go w.fetchTendermintBlockMeta(height, metaChan)
			txs := <-txsChan
			results := <-resultBlockChan
			meta := <-metaChan

			// get block info from eth
			go w.fetchBlockWeb3(height, web3BlockChan)
			web3Block := <-web3BlockChan
			go w.fetchBlockTxReceiptsWeb3(web3Block, web3BlockReceipts)
			web3receipts := <-web3BlockReceipts

			receipts := make(map[common.Hash]ethTypes.Receipt)
			for _, v := range web3receipts {
				receipts[v.TxHash] = *v
			}

			block := types.BlockData{
				BlockID:               tmBlock.BlockID,
				Header:                tmBlock.Block.Header,
				Data:                  types.Data{Txs: txs},
				Evidence:              tmBlock.Block.Evidence,
				LastCommit:            tmBlock.Block.LastCommit,
				TxsResults:            results.TxsResults,
				BeginBlockEvents:      results.BeginBlockEvents,
				EndBlockEvents:        results.EndBlockEvents,
				ValidatorUpdates:      results.ValidatorUpdates,
				ConsensusParamUpdates: results.ConsensusParamUpdates,
				Web3Block:             *web3Block,
				Receipts:              receipts,
				BlockMeta:             *meta,
			}
			_ = block
			// block results, tx result
			// make two global goroutines: 1 read all data, 2 build a bussines logic and push db transaction in db

			//txs := make([]sdkTypes.Tx, len(block.Block.Txs))
			//for i, txBz := range block.Block.Txs {
			//	txs[i], err = w.cdc.TxConfig.TxDecoder()(txBz)
			//	if err != nil {
			//		panic(err)
			//	}
			//}
			//
			//if len(txs) > 0 {
			//	w.txHandler.WriteTxs(txs)
			//}

			w.mut.Lock()
			w.currentHeight++
			w.mut.Unlock()
		}
	}()
}

func (w *Worker) GetLatestHeight() uint64  { return w.latestHeight }
func (w *Worker) GetCurrentHeight() uint64 { return w.currentHeight }

// fetchTendermintBlock gets block data from tendermint
func (w *Worker) fetchTendermintBlock(height int64) *ctypes.ResultBlock {
	start := time.Now()

	// Request until get block
	for first := true; true; first = false {
		// Request block
		result, err := w.tmJsonRpc.Block(w.ctx, &height)
		if err == nil {
			if !first {
				w.logger.Info(
					fmt.Sprintf("Fetched block (after %s)", helpers.DurationToString(time.Since(start))),
					"block", height,
				)
			} else {
				w.logger.Info(
					fmt.Sprintf("Fetched block (%s)", helpers.DurationToString(time.Since(start))),
					"block", height,
				)
			}
			return result
		}
	}

	return nil
}

// fetchTendermintBlockTxResults gets block txs results from tendermint
func (w *Worker) fetchTendermintBlockTxResults(height int64, ch chan *ctypes.ResultBlockResults) {

	// Request until get block results
	for {
		// Request block results
		result, err := w.tmJsonRpc.BlockResults(w.ctx, &height)
		if err == nil { // len(result.EndBlockEvents) != 0
			ch <- result
			break
		}
	}
}

// fetchTendermintBLockTxs gets txs payload from tendermint
func (w *Worker) fetchTendermintBLockTxs(height int64, total int, ch chan types.Txs) {
	query := fmt.Sprintf("tx.height=%d", height)
	page, perPage := 1, 100

	var results []types.Tx
	for len(results) < total {

		// Request transactions
		result, err := w.tmJsonRpc.TxSearch(w.ctx, query, true, &page, &perPage, "")
		w.panicError(err)

		for _, tx := range result.Txs {
			var result types.Tx
			var txLog []interface{}

			// Recover messages from raw transaction bytes
			recoveredTx, err := w.cdc.TxConfig.TxDecoder()(tx.Tx)
			w.panicError(err)

			// Parse transaction results logs
			err = json.Unmarshal([]byte(tx.TxResult.Log), &txLog)
			if err != nil {
				w.panicError(err)
			} else {
				result.Logs = txLog
			}

			result.TxInfo = w.parseTxInfo(recoveredTx)
			result.TxsResults = tx.TxResult
			result.Hash = tx.Hash
			result.Height = tx.Height
			result.Index = tx.Index

			results = append(results, result)
		}

		// read next page
		if len(result.Txs) > 0 {
			page++
		}
	}

	// Send results to the channel
	ch <- results
}

// fetchTendermintBlockMeta gets block meta from tendermint
func (w *Worker) fetchTendermintBlockMeta(height int64, ch chan *types.BlockMeta) {

	// Request blockchain info
	metas, err := w.tmJsonRpc.BlockchainInfo(w.ctx, height, height)
	w.panicError(err)

	// Send result to the channel
	ch <- &types.BlockMeta{
		Size: metas.BlockMetas[0].BlockSize,
	}
}

// fetchBlockWeb3 gets block info from evm
func (w *Worker) fetchBlockWeb3(height int64, ch chan *web3types.Block) {

	// Request block by number
	result, err := w.web3JsonRpc.BlockByNumber(w.ctx, big.NewInt(height))
	w.panicError(err)

	// Send result to the channel
	ch <- result
}

// fetchBlockTxReceiptsWeb3 gets block txs receipts from evm
func (w *Worker) fetchBlockTxReceiptsWeb3(block *web3types.Block, ch chan web3types.Receipts) {

	// Request transaction receipts by hashes in parallel
	results := make(web3types.Receipts, len(block.Transactions()))
	wg := &sync.WaitGroup{}
	wg.Add(len(block.Transactions()))
	for i, tx := range block.Transactions() {
		go func(i int) {
			defer wg.Done()
			result, err := w.web3JsonRpc.TransactionReceipt(w.ctx, tx.Hash())
			w.panicError(err)
			results[i] = result
		}(i)
	}
	wg.Wait()

	// Send results to the channel
	ch <- results
}

// fetchBlockTxsWeb3 gets block txs receipts from evm
func (w *Worker) fetchBlockTxsWeb3(block *web3types.Block, ch chan web3types.Transactions) {

	// Request transaction receipts by hashes in parallel
	results := make(web3types.Transactions, len(block.Transactions()))
	wg := &sync.WaitGroup{}
	wg.Add(len(block.Transactions()))
	for i, tx := range block.Transactions() {
		go func(i int) {
			defer wg.Done()
			result, _, err := w.web3JsonRpc.TransactionByHash(w.ctx, tx.Hash())
			w.panicError(err)
			results[i] = result
		}(i)
	}
	wg.Wait()

	// Send results to the channel
	ch <- results
}

// parseTxInfo reads data from a transaction and msgs
func (w *Worker) parseTxInfo(tx sdkTypes.Tx) (txInfo types.TxInfo) {
	if tx == nil {
		return
	}
	// read msgs
	txInfo.Msgs = make([]types.TxMsg, len(tx.GetMsgs()))
	for i, rawMsg := range tx.GetMsgs() {
		var msg types.TxMsg
		msg.Type = sdkTypes.MsgTypeURL(rawMsg)
		msg.Msg = rawMsg

		for _, signer := range rawMsg.GetSigners() {
			msg.From = append(msg.From, signer.String())
		}

		txInfo.Msgs[i] = msg
	}

	feeTx := tx.(sdkTypes.FeeTx)
	txInfo.Fee.Gas = feeTx.GetGas()
	txInfo.Fee.Amount = feeTx.GetFee()
	txInfo.Fee.FeePayer = feeTx.FeePayer()
	txInfo.Fee.FeeGranter = feeTx.FeeGranter()
	txInfo.Memo = tx.(sdkTypes.TxWithMemo).GetMemo()

	return
}

// panicError panic if got err
func (w *Worker) panicError(err error) {
	if err != nil {
		w.logger.Error(fmt.Sprintf("Error: %v", err))
		panic(err)
	}
}
