CREATE TABLE transactions
(
    hash                             bytea        NOT NULL,
    block_hash                       bytea        NULL,
    block_number                     int4         NULL,
    gas                              numeric(100) NOT NULL,
    gas_price                        numeric(100) NOT NULL,
    gas_used                         numeric(100) NULL,
    from_address_hash                bytea        NOT NULL,
    to_address_hash                  bytea        NULL,
    nonce                            int4         NOT NULL,
    "index"                          int4         NULL,
    inserted_at                      timestamp    NOT NULL,
    updated_at                       timestamp    NOT NULL,
    value                            numeric(100) NOT NULL,
    tx                               bytea        NOT NULL,
    --web3 tx
    cumulative_gas_used              numeric(100) NULL,
    error                            varchar(255) NULL,
    "input"                          bytea        NULL,
    r                                numeric(100) NULL,
    s                                numeric(100) NULL,
    status                           int4         NULL,
    v                                numeric(100) NULL,
    created_contract_address_hash    bytea        NULL,
    created_contract_code_indexed_at timestamp    NULL,
    earliest_processing_start        timestamp    NULL,
    old_block_hash                   bytea        NULL,
    revert_reason                    text         NULL,
    max_priority_fee_per_gas         numeric(100) NULL,
    max_fee_per_gas                  numeric(100) NULL,
    eth_tx_type                      int4         NULL,
    has_error_in_internal_txs        bool         NULL,
    --cosmos tx
    code                             int,
    log                              varchar,
    codespace                        varchar,
    events                           bytea,
    info                             varchar,
    logs                             bytea,
    data                             bytea,
--     msg_types                        varchar[],
    memo                             varchar,
    fee_payer                        bytea,
    fee_granter                      bytea
--     CONSTRAINT collated_block_number CHECK (((block_hash IS NULL) OR (block_number IS NOT NULL))),
--     CONSTRAINT collated_cumalative_gas_used CHECK (((block_hash IS NULL) OR (cumulative_gas_used IS NOT NULL))),
--     CONSTRAINT collated_gas_used CHECK (((block_hash IS NULL) OR (gas_used IS NOT NULL))),
--     CONSTRAINT collated_index CHECK (((block_hash IS NULL) OR (index IS NOT NULL))),
--     CONSTRAINT error CHECK (((status = 0) OR ((status <> 0) AND (error IS NULL)))),
--     CONSTRAINT pending_block_number CHECK (((block_hash IS NOT NULL) OR (block_number IS NULL))),
--     CONSTRAINT pending_cumalative_gas_used CHECK (((block_hash IS NOT NULL) OR (cumulative_gas_used IS NULL))),
--     CONSTRAINT pending_gas_used CHECK (((block_hash IS NOT NULL) OR (gas_used IS NULL))),
--     CONSTRAINT pending_index CHECK (((block_hash IS NOT NULL) OR (index IS NULL))),
--     CONSTRAINT status CHECK ((((block_hash IS NULL) AND (status IS NULL)) OR (block_hash IS NOT NULL) OR ((status = 0) AND ((error)::text = 'dropped/replaced'::text)))),
--     CONSTRAINT transactions_pkey PRIMARY KEY (hash)
);

CREATE TABLE addresses
(
    hash                       bytea NOT NULL PRIMARY KEY,
    address_bech32             char(255),
    nonce                      int4,
    code_hash                  bytea,
    fetch_balance              numeric(100, 0),
    fetch_balance_block_number int8,
    verified                   bool,
    decompiled                 bool,
    contract_code              bytea,
    inserted_at                timestamp,
    updated_at                 timestamp
);

-- CREATE TABLE blocks
-- (
--
-- )
