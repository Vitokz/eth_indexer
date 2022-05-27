package db

import (
	"context"
	"time"

	"github.com/Vitokz/eth_indexer/db/types"
)

func (d *DB) InsertTransaction(ctx context.Context, tx types.Tx) error {
	query := `INSERT INTO transactions
			  (hash, block_hash, block_number, gas, gas_price, gas_used, from_address_hash, 
			   to_address_hash, nonce, "index", inserted_at, updated_at, value, tx, cumulative_gas_used, error, "input", 
			   r, s, status, v, created_contract_address_hash, created_contract_code_indexed_at, earliest_processing_start, 
			   old_block_hash, revert_reason, max_priority_fee_per_gas, max_fee_per_gas, eth_tx_type, has_error_in_internal_txs, 
			   code, log, codespace, events, info, logs, "data", memo, fee_payer, fee_granter)
			  VALUES($1, $2, $3, $4, $5, $6, $7, 
			         $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, 
			         $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, 
			         $31, $32, $33, $34, $35, $36, $37, $38, $39, $40);
    `
	_, err := d.DB.ExecContext(ctx,
		query,
		tx.Hash, tx.BlockHash, tx.BlockNumber, tx.Gas, tx.GasPrice, tx.GasUsed, tx.FromAddressHash,
		tx.ToAddressHash, tx.Nonce, tx.Index, tx.InsertedAt, time.Now(), tx.Value, tx.Tx, tx.CumulativeGasUsed, tx.Error, tx.Input,
		tx.R, tx.S, tx.Status, tx.V, tx.CreatedContractAddressHash, tx.CreatedContractCodeIndexedAt, tx.EarliestProcessingStart,
		tx.OldBlockHash, tx.RevertReason, tx.MaxPriorityFeePerGas, tx.MaxFeePerGas, tx.EthTxType, tx.HasErrorInInternalTxs,
		tx.Code, tx.Log, tx.Codespace, tx.Events, tx.Info, tx.Logs, tx.Data, tx.Memo, tx.FeePayer, tx.FeeGranter)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) GetTxByHash(ctx context.Context, hash []byte) (types.Tx, error) {
	var tx types.Tx
	query := `SELECT hash, block_hash, block_number, gas, gas_price, gas_used, from_address_hash, to_address_hash, 
			  nonce, "index", inserted_at, updated_at, value, tx, cumulative_gas_used, error, "input", r, s, status, v, 
			  created_contract_address_hash, created_contract_code_indexed_at, earliest_processing_start, old_block_hash, 
              revert_reason, max_priority_fee_per_gas, max_fee_per_gas, eth_tx_type, has_error_in_internal_txs, code, log, codespace, events, 	
			  info, logs, "data", memo, fee_payer, fee_granter
              FROM transactions t WHERE t.hash = ?;`

	err := d.DB.GetContext(ctx, &tx,
		query, hash,
	)
	if err != nil {
		return tx, err
	}

	return tx, nil
}
