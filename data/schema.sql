BEGIN TRANSACTION;

-- Etehereum address in hex
CREATE DOMAIN eth_addr AS char(40);

-- Accounts are ethereum accounts.
CREATE TABLE accounts(
    id uuid PRIMARY KEY,
    eth_addr eth_addr NOT NULL -- ethereum address
      CONSTRAINT unique_eth_addr UNIQUE,
    public_key text NOT NULL,
    private_key text NOT NULL,

    ptc_balance bigint NOT NULL -- PTC balance
        CONSTRAINT positive_ptc_balance CHECK (accounts.ptc_balance >= 0),

    psc_balance bigint NOT NULL -- PSC balance
        CONSTRAINT positive_psc_balance CHECK (accounts.psc_balance >= 0),

    eth_balance char(32) NOT NULL -- ethereum balance up to 99999 ETH in WEI. Ethereum's uint192 in base64 (RFC-4648).
);

-- State channels.
CREATE TABLE channels (
    id uuid PRIMARY KEY,
    agent eth_addr NOT NULL,
    client eth_addr NOT NULL,
    closed boolean NOT NULL,
    total_deposit bigint NOT NULL
        CONSTRAINT positive_total_deposit CHECK (channels.total_deposit >= 0),
    receipt_balance bigint NOT NULL -- last payment amount received
        CONSTRAINT positive_receipt_balance CHECK (channels.receipt_balance >= 0),
    receipt_signature text -- signature corresponding to last payment
);

CREATE TABLE settings (
    key text PRIMARY KEY,
    value text NOT NULL
);

END TRANSACTION;
