package types

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	web3types "github.com/ethereum/go-ethereum/core/types"

	//ctypes "github.com/tendermint/tendermint/rpc/core/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/types"
)

type BlockData struct {
	// Tendermint Block Payload
	BlockID    types.BlockID
	Header     types.Header
	Data       Data
	Evidence   types.EvidenceData
	LastCommit *types.Commit

	// Tendermint Block Results Payload
	TxsResults            []*abci.ResponseDeliverTx
	BeginBlockEvents      []abci.Event           `json:"begin_block_events"`
	EndBlockEvents        []abci.Event           `json:"end_block_events"`
	ValidatorUpdates      []abci.ValidatorUpdate `json:"validator_updates"`
	ConsensusParamUpdates *abci.ConsensusParams  `json:"consensus_param_updates"`

	// Web3 Block
	Web3Block web3types.Block
	Receipts  Receipts

	// Meta
	BlockMeta BlockMeta
}

type BlockMeta struct {
	Size int
}

type Data struct {
	Txs Txs `json:"txs"`
}

type Receipts map[common.Hash]ethTypes.Receipt
