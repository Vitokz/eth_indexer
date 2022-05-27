package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type RpcTransaction struct {
	Tx          *ethTypes.Transaction
	TxExtraInfo txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *RpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.Tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.TxExtraInfo)
}
