package data

import (
	"encoding/base64"
	"fmt"
	"math/big"
)

//go:generate reform

// B64BigInt is a base64 of big.Int that implements json.Marshaler.
type B64BigInt string

// MarshalJSON marshals itself.
func (i B64BigInt) MarshalJSON() ([]byte, error) {
	buf, err := base64.URLEncoding.DecodeString(string(i))
	if err != nil {
		return nil, fmt.Errorf("could not decode base64: %v", err)
	}
	v := big.NewInt(0)
	v.SetBytes(buf)
	return []byte(v.String()), nil
}

// Account is an ethereum account.
//reform:accounts
type Account struct {
	ID         string    `reform:"id,pk"`
	EthAddr    string    `reform:"eth_addr"`
	PublicKey  string    `reform:"public_key"`
	PrivateKey string    `reform:"private_key"`
	PTCBalance uint64    `json:"ptcBalance" reform:"ptc_balance"`
	PSCBalance uint64    `json:"psc_balance" reform:"psc_balance"`
	EthBalance B64BigInt `json:"ethBalance" reform:"eth_balance"`
}
