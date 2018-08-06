-- Accounts are ethereum accounts.
CREATE TABLE accounts(
    id uuid PRIMARY KEY,
    eth_addr char(40) NOT NULL -- ethereum address
      CONSTRAINT unique_eth_addr UNIQUE,
    public_key text NOT NULL,
    private_key text NOT NULL
)
