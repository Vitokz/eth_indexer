package types

import "time"

type Addresses struct {
	Hash                    []byte    `db:"hash"`
	AddressBech32           string    `db:"address_bech32"`
	Nonce                   int       `db:"nonce"`
	CodeHash                []byte    `db:"code_hash"`
	FetchBalance            float64   `db:"fetch_balance"`
	FetchBalanceBlockNUmber int       `db:"fetch_balance_block_number"`
	Verified                bool      `db:"verified"`
	Decompiled              bool      `db:"decompiled"`
	ContractCode            []byte    `db:"contract_code"`
	InsertedAt              time.Time `db:"inserted_at"`
	UpdatedAt               time.Time `db:"updated_at"`
}
