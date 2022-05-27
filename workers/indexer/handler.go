package indexer

import (
	"context"
	"sync"
	"time"

	dbTypes "github.com/Vitokz/eth_indexer/db/types"
	"github.com/Vitokz/eth_indexer/workers/indexer/types"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type TxHandler struct {
	block chan types.BlockData
	wg    sync.WaitGroup
	w     *Worker

	ctx context.Context

	msgHandlers types.MsgsTypes
}

func NewTxHandler(ctx context.Context, w *Worker) TxHandler {
	return TxHandler{
		block: make(chan types.BlockData),
		wg:    sync.WaitGroup{},
		w:     w,
		ctx:   ctx,
	}
}

func (t *TxHandler) Handle() {
	for {
		select {
		case block := <-t.block:
			_ = block
		}
	}
}

func (t *TxHandler) CheckAddresses() {

}

func (t *TxHandler) HandleTxs(block types.BlockData) {
	var (
		tmTxs           = block.Data.Txs
		tmTxsResults    = block.TxsResults
		web3Block       = block.Web3Block
		web3Txs         = block.Web3Block.Transactions()
		web3TxsReceipts = block.Receipts
	)

	web3Block.Header().
	for _, tx := range tmTxs {
		ethHash := common.BytesToHash(tx.Hash)
		if ethTx := block.Web3Block.Transaction(ethHash); ethTx != nil {
			recp := web3TxsReceipts[ethTx.Hash()]
			// handle eth tx
		}

		for _, msg := range tx.TxInfo.Msgs {
			handler := t.msgHandlers.GetMsgHandler(msg.Type)
			_ = handler // handle cosmos tx
			//err := handler()
		}
	}
}

func (t *TxHandler) handleEthTx(block *types.BlockData, ethTx ethTypes.Transaction, receipt ethTypes.Receipt) {
	var (
		cumulativeGasUsed   = float64(receipt.CumulativeGasUsed)
		r, s, v             = ethTx.RawSignatureValues()
		rVal, sVal, vVal    = float64(r.Int64()), float64(s.Int64()), float64(v.Int64())
		status              = receipt.Status
		ethTxType           = ethTx.Type()
		maxProrityFeePerGas = float64(ethTx.GasTipCap().Int64())
		maxFeePerGas        = float64(ethTx.GasFeeCap().Int64())
		value =float64(ethTx.Value().Int64())
	)

	tx := dbTypes.Tx{
		Hash:            ethTx.Hash().Bytes(),
		BlockHash:       block.BlockID.Hash,
		BlockNumber:     block.Header.Height,
		Gas:             float64(ethTx.Gas()),
		GasPrice:        float64(ethTx.GasPrice().Int64()),
		GasUsed:         float64(receipt.GasUsed),
		FromAddressHash: nil, // 0?,
		ToAddressHash:   ethTx.To().Bytes(),
		Nonce:           ethTx.Nonce(),
		Index:           receipt.TransactionIndex,
		InsertedAt:      time.Time{},  //?
		UpdatedAt:       &time.Time{}, //?
		Value:           value,
		//Tx:                           nil, //?
		CumulativeGasUsed:            &cumulativeGasUsed,
		Error:                        nil, //?
		Input:                        ethTx.In,
		R:                            &rVal,
		S:                            &sVal,
		V:                            &vVal,
		Status:                       &status,
		CreatedContractAddressHash:   receipt.ContractAddress.Bytes(),
		CreatedContractCodeIndexedAt: nil, //?
		EarliestProcessingStart:      nil, //?
		OldBlockHash:                 nil, //?
		RevertReason:                 nil, //?
		MaxPriorityFeePerGas:         &maxProrityFeePerGas,
		MaxFeePerGas:                 &maxFeePerGas,
		EthTxType:                    &ethTxType,
		HasErrorInInternalTxs:        nil,
		//Code:                         nil,
		//Log:                          nil,
		//Codespace:                    nil,
		//Events:                       nil,
		//Info:                         nil,
		//Logs:                         nil,
		//Data:                         nil,
		//Memo:                         nil,
		//FeePayer:                     nil,
		//FeeGranter:                   nil,
	}
}

//for _, msg := range tx.GetMsgs() {
////check is eth tx
//ethMsg, ok := msg.(*evmtypes.MsgEthereumTx)
//if ok {
//tx := ethMsg.AsTransaction()
//
//fmt.Printf("its evm tx: %s", tx.Hash().String())
//continue
//}
//
////check is cosmos tx
//msgType := sdkTypes.MsgTypeURL(msg)
//handler := t.msgHandlers.GetMsgHandler(msgType)
//
//err := handler(msg)
//if err != nil {
////t.AddTxToDbWithPendingState()
//}
//
//fmt.Println(msgType)
//}
