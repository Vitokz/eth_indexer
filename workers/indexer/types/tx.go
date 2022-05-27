package types

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/bytes"
)

type Txs []Tx

type Tx struct {
	Hash       bytes.HexBytes
	TxsResults abci.ResponseDeliverTx
	TxInfo     TxInfo
	Logs       []interface{}
	Height     int64
	Index      uint32
}

type TxMsg struct {
	Type string
	Msg  sdkTypes.Msg
	From []string
}

type TxFee struct {
	Gas        uint64
	Amount     sdkTypes.Coins
	FeePayer   sdkTypes.Address
	FeeGranter sdkTypes.Address
}

type TxInfo struct {
	Msgs []TxMsg `json:"msgs"`
	Memo string  `json:"memo"`
	Fee  TxFee   `json:"fee"`
}
