package data

//go:generate reform

// Account is an ethereum account.
//reform:accounts
type Account struct {
	ID         string `reform:"id,pk"`
	EthAddr    string `reform:"eth_addr"`
	PublicKey  string `reform:"public_key"`
	PrivateKey string `reform:"private_key"`
}
