-- +migrate Up

CREATE TABLE block (
    rowid SERIAL PRIMARY KEY,
    height INT NOT NULL, -- 0 for the genesis block.
    block_id BYTEA NOT NULL,
    parent_id BYTEA,  -- Null only for the genesis block.
    raw BYTEA NOT NULL
);

-- Well-formed and properly signed transactions that have not yet been included in a block.
CREATE TABLE mempool_transaction (
    rowid SERIAL PRIMARY KEY,
    txid BYTEA NOT NULL,
    raw BYTEA NOT NULL
);

CREATE TYPE transaction_status AS ENUM ('pending', 'mined', 'rejected');

CREATE TABLE transaction_auditlog (
    rowid SERIAL PRIMARY KEY,
    txid BYTEA NOT NULL,
    raw BYTEA NOT NULL,
    status transaction_status NOT NULL
);

CREATE TABLE utxo (
    rowid SERIAL PRIMARY KEY,
    txid BYTEA NOT NULL,
    output_idx INT NOT NULL,
    value INT NOT NULL,
    script_pubkey BYTEA NOT NULL

    -- we don't really need this field
    -- block_id INT NOT NULL REFERENCES block(rowid)
);

-- @KK: Eventually, we probably want a table that indexes mined transactions by ID.  We'll still need to store the raw
--      (serialized) blocks, though.  Also, this gets a little bit tricky, because shallow forks may cause the same
--      transaction to be part of multiple blocks at similar heights.

-- +migrate Down
DROP TABLE utxo;
DROP TABLE mempool_transaction;
DROP TABLE block;
DROP TABLE transaction_auditlog;
DROP TYPE transaction_status;
