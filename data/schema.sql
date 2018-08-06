-- Accounts are ethereum accounts.
CREATE TABLE accounts(
    id uuid PRIMARY KEY,
    eth_addr char(40) NOT NULL -- ethereum address
      CONSTRAINT unique_eth_addr UNIQUE,
    public_key text NOT NULL,
    private_key text NOT NULL,

    ptc_balance bigint NOT NULL -- PTC balance
        CONSTRAINT positive_ptc_balance CHECK (accounts.ptc_balance >= 0),

    psc_balance bigint NOT NULL -- PSC balance
        CONSTRAINT positive_psc_balance CHECK (accounts.psc_balance >= 0),

    eth_balance char(32) NOT NULL -- ethereum balance up to 99999 ETH in WEI. Ethereum's uint192 in base64 (RFC-4648).
)
