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
	ID              string    `reform:"id,pk"`
	EthAddr         string    `reform:"eth_addr"`
	PublicKey       string    `reform:"public_key"`
	PrivateKey      string    `reform:"private_key"`
	ContractBalance uint64    `reform:"contract_balance"`
	EthBalance      B64BigInt `reform:"eth_balance"`
}

// Channel is a state channel.
//reform:channels
type Channel struct {
	ID               string  `reform:"id,pk"`
	Agent            string  `reform:"agent"`
	Client           string  `reform:"client"`
	Closed           bool    `reform:"closed"`
	TotalDeposit     uint64  `reform:"total_deposit"`
	ReceiptBalance   uint64  `reform:"receipt_balance"`
	ReceiptSignature *string `reform:"receipt_signature"`
}

// System settings keys.
const (
	SettingsLastSeenBlock = "system.monitor.lastseenblock"
)

// Setting is a system settings.
//reform:settings
type Setting struct {
	Key   string `reform:"key,pk"`
	Value string `reform:"value"`
}
