package data

import (
	"crypto/ecdsa"
	cryptorand "crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	uuid "github.com/satori/go.uuid"
)

// NewUUID generates a new UUID.
func NewUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}

// NewAccount generates new account.
func NewAccount() *Account {
	priv, _ := ecdsa.GenerateKey(crypto.S256(), cryptorand.Reader)
	pub := base64.URLEncoding.EncodeToString(
		crypto.FromECDSAPub(&priv.PublicKey))
	addr := hex.EncodeToString(crypto.PubkeyToAddress(priv.PublicKey).Bytes())
	return &Account{
		ID:         NewUUID(),
		EthAddr:    addr,
		PublicKey:  pub,
		PrivateKey: base64.URLEncoding.EncodeToString(crypto.FromECDSA(priv)),
	}
}
