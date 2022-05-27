package types

import (
	"time"
)

type Tx struct {
	Hash            []byte     `db:"hash,omitempty"`
	BlockHash       []byte     `db:"block_hash,omitempty"`
	BlockNumber     int64      `db:"block_number,omitempty"`
	Gas             float64    `db:"gas,omitempty"`
	GasPrice        float64    `db:"gas_price,omitempty"`
	GasUsed         float64    `db:"gas_used,omitempty"`
	FromAddressHash []byte     `db:"from_address_hash,omitempty"`
	ToAddressHash   []byte     `db:"to_address_hash,omitempty"`
	Nonce           uint64     `db:"nonce,omitempty"`
	Index           uint       `db:"index,omitempty"`
	InsertedAt      time.Time  `db:"inserted_at"`
	UpdatedAt       *time.Time `db:"updated_at,omitempty"`
	Value           float64    `db:"value,omitempty"`
	Tx              []byte     `db:"tx,omitempty"`
	//web3
	CumulativeGasUsed            *float64   `db:"cumulative_gas_used,omitempty"`
	Error                        *string    `db:"error,omitempty"`
	Input                        []byte     `db:"input,omitempty"`
	R                            *float64   `db:"r,omitempty"`
	S                            *float64   `db:"s,omitempty"`
	V                            *float64   `db:"v,omitempty"`
	Status                       *uint64    `db:"status,omitempty"`
	CreatedContractAddressHash   []byte     `db:"created_contract_address_hash,omitempty"`
	CreatedContractCodeIndexedAt *time.Time `db:"created_contract_code_indexed_at,omitempty"`
	EarliestProcessingStart      *time.Time `db:"earliest_processing_start,omitempty"`
	OldBlockHash                 []byte     `db:"old_block_hash,omitempty"`
	RevertReason                 *string    `db:"revert_reason,omitempty"`
	MaxPriorityFeePerGas         *float64   `db:"max_priority_fee_per_gas,omitempty"`
	MaxFeePerGas                 *float64   `db:"max_fee_per_gas,omitempty"`
	EthTxType                    *uint8     `db:"eth_tx_type,omitempty"`
	HasErrorInInternalTxs        *bool      `db:"has_error_in_internal_txs,omitempty"`
	//cosmos
	Code       *int    `db:"code,omitempty"`
	Log        *string `db:"log,omitempty"`
	Codespace  *string `db:"codespace,omitempty"`
	Events     []byte  `db:"events,omitempty"`
	Info       *string `db:"info,omitempty"`
	Logs       []byte  `db:"logs"`
	Data       []byte  `db:"data,omitempty"`
	Memo       *string `db:"memo,omitempty"`
	FeePayer   []byte  `db:"fee_payer,omitempty"`
	FeeGranter []byte  `db:"fee_granter,omitempty"`
}
