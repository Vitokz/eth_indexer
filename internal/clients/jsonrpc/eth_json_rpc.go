package jsonrpc

import (
	"context"
	"fmt"

	"github.com/Vitokz/eth_indexer/internal/clients/jsonrpc/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthJsonClient struct {
	c *rpc.Client
}

func New(url string) (*EthJsonClient, error) {
	ethJsonrpc, err := rpc.DialHTTP(url)
	if err != nil {
		return nil, err
	}

	return &EthJsonClient{
		ethJsonrpc,
	}, nil
}

func (ec *EthJsonClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.RpcTransaction, isPending bool, err error) {
	var json *types.RpcTransaction
	err = ec.c.CallContext(ctx, &json, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, false, err
	} else if json == nil {
		return nil, false, ethereum.NotFound
	} else if _, r, _ := json.Tx.RawSignatureValues(); r == nil {
		return nil, false, fmt.Errorf("server returned transaction without signature")
	}

	if json.TxExtraInfo.From != nil && json.TxExtraInfo.BlockHash != nil {
		setSenderFromServer(json.Tx, *json.TxExtraInfo.From, *json.TxExtraInfo.BlockHash)
	}
	return json, json.TxExtraInfo.BlockNumber == nil, nil
}
