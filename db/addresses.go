package db

import (
	"context"
	"time"

	"github.com/Vitokz/eth_indexer/db/types"
)

func (d *DB) InsertAddressData(ctx context.Context, a types.Addresses) error {
	query := `INSERT INTO addresses 
 			  (hash, address_bech32, nonce, code_hash, fetch_balance,fetch_balance_block_number, 
              verified, decompiled, contract_code, inserted_at, updated_at)
		      VALUES 
			  ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);`

	_, err := d.DB.ExecContext(ctx,
		query,
		a.Hash, a.AddressBech32, a.Nonce, a.CodeHash, a.FetchBalance, a.FetchBalanceBlockNUmber,
		a.Verified, a.Decompiled, a.ContractCode, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetAddressByHexHash(ctx context.Context, hash []byte) (types.Addresses, error) {
	query := `SELECT hash, address_bech32, nonce, code_hash, fetch_balance, fetch_balance_block_number, verified, 
       		  decompiled, contract_code, inserted_at, updated_at
			  FROM addresses a
	          WHERE a.hash=$1`
	var u types.Addresses

	err := d.DB.GetContext(ctx, &u,
		query, hash,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

//update params
