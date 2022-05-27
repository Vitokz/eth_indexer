package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Vitokz/eth_indexer/internal/clients/jsonrpc"

	"github.com/Vitokz/eth_indexer/config"
	"github.com/Vitokz/eth_indexer/internal/multiclient"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	ethApp "github.com/tharsis/ethermint/app"
	ethEncoding "github.com/tharsis/ethermint/encoding"
)

func main() {
	cfg := config.ParseConfigInConfigFile()

	cdc := ethEncoding.MakeConfig(ethApp.ModuleBasics)

	cm, err := multiclient.NewMultiClient(cfg, &cdc)
	if err != nil {
		panic(err)
	}
	_ = cm

	ctx := context.Background()

	fmt.Println(vesting.MsgCreateVestingAccount{}.Type())

	//db, err := db.NewDB(cfg.GetPgDbDsn())
	//if err != nil {
	//	panic(err)
	//}

	//s, err := hex.DecodeString("014d7d16d7af4fefb61bd95b765c8cab")
	//
	//err = db.InsertAddressData(ctx, types.Addresses{
	//	Hash:                    s,
	//	AddressBech32:           "asd",
	//	Nonce:                   1,
	//	CodeHash:                hexutils.HexToBytes("013d7d16d7ad4fefb61bd95b765c8ceb"),
	//	FetchBalance:            10,
	//	FetchBalanceBlockNUmber: 10,
	//	Verified:                false,
	//	Decompiled:              false,
	//	ContractCode:            hexutils.HexToBytes("013d7d16d7ad4fefb61bd95b765c8ceb"),
	//})
	//if err != nil {
	//	panic(err)
	//}

	//_, err = db.GetAddressByHexHash(ctx, "013d7d16d7ad4fefb61bd95b765c8cab")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(u)
	eth, err := jsonrpc.New(cfg.GetEthereumJsonRPC())
	if err != nil {
		panic(err)
	}
	tx, _, err := eth.TransactionByHash(ctx, common.HexToHash("0x068624610f98264fe5f0a2772962104b6490ad87704c60a9ec1b6efd3222cc40"))
	fmt.Println(tx)
	//ind := indexer.New(ctx,cm.Jsonrpc.GetEth(), &cdc, nil)
	//ind.Start()
	//for {
	//	time.Sleep(10 * time.Second)
	//	fmt.Println(ind.GetLatestHeight())
	//}

	//fmt.Println(ind.GetLatestHeight())

	//for _, txBz := range resp.Block.Txs {
	//	tx, err := cdc.TxConfig.TxDecoder()(txBz)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	// check txs
	//	for _, msg := range tx.GetMsgs() {
	//
	//		//check is eth tx
	//		ethMsg, ok := msg.(*evmtypes.MsgEthereumTx)
	//		if ok {
	//			tx := ethMsg.AsTransaction()
	//
	//			fmt.Printf("its evm tx: %s", tx.Hash().String())
	//			//continue
	//		}
	//
	//		//check is cosmos tx
	//		msgType := cosmSdk.MsgTypeURL(msg)
	//		fmt.Println(lst)
	//	}
	//}

	//txs, err := resp.GetTxsHash()
	//if err != nil {
	//	panic(err)
	//}
	//
	//height, err := resp.GetHeight()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(tx)
}
