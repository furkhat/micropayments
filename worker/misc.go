package worker

import (
	"encoding/base64"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/furkhat/micropayments/data"
)

func accTransactOpts(acc *data.Account) (*bind.TransactOpts, error) {
	keyBytes, err := base64.URLEncoding.DecodeString(acc.PrivateKey)
	if err != nil {
		return nil, err
	}
	key, err := crypto.ToECDSA(keyBytes)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactor(key), nil
}
