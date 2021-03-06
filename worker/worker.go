package worker

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/furkhat/micropayments/data"
	reform "gopkg.in/reform.v1"
)

// EthBack defines a typed wrappers for contract and ethereum methods used by the worker.
type EthBack interface {
	EthBalanceAt(context.Context, common.Address) (*big.Int, error)

	ContractBalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)

	CloseChannel(*bind.TransactOpts, common.Address, uint32,
		*big.Int, []byte, []byte) (*types.Transaction, error)
}

// Worker defines methods for processing ethereum logs.
type Worker struct {
	db      *reform.DB
	ethBack EthBack

	sessStart chan string
}

// NewWorker creates a Worker.
func NewWorker(db *reform.DB, ethBack EthBack) *Worker {
	return &Worker{db, ethBack, make(chan string)}
}

// AfterChannelCreated registers channel in the system.
func (w *Worker) AfterChannelCreated(log *ethtypes.Log) error {
	_, err := w.accountByHash(log.Topics[1])
	if err == sql.ErrNoRows {
		_, err = w.accountByHash(log.Topics[2])
	}
	if err != nil {
		return err
	}

	args, err := channelCreatedNonIndexArgs.NonIndexed().UnpackValues(log.Data)
	if err != nil {
		return err
	}
	amount := args[0].(*big.Int)

	agentAddr := hex.EncodeToString(log.Topics[1].Bytes())
	clientAddr := hex.EncodeToString(log.Topics[2].Bytes())
	id := data.NewUUID()
	err = w.db.Insert(&data.Channel{
		ID:           id,
		Agent:        strings.TrimLeft(agentAddr, "0"),
		Client:       strings.TrimLeft(clientAddr, "0"),
		TotalDeposit: amount.Uint64(),
	})
	if err != nil {
		return err
	}

	// Start selling the service.
	go func() { w.sessStart <- id }()

	return nil
}

// AfterChannelClose update accounts PTC, PSC and ethereum balances.
func (w *Worker) AfterChannelClose(closeLog *ethtypes.Log) error {
	err := w.updateAccountBalances(closeLog.Topics[1])
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func (w *Worker) updateAccountBalances(addrHash common.Hash) error {
	acc, err := w.accountByHash(addrHash)
	if err != nil {
		return err
	}

	accAddr := common.HexToAddress(acc.EthAddr)

	amount, err := w.ethBack.ContractBalanceOf(&bind.CallOpts{}, accAddr)
	if err != nil {
		return err
	}

	acc.ContractBalance = amount.Uint64()

	amount, err = w.ethBack.EthBalanceAt(context.Background(), accAddr)
	if err != nil {
		return err
	}

	acc.EthBalance = data.B64BigInt(base64.URLEncoding.EncodeToString(amount.Bytes()))

	return w.db.Update(acc)
}

func (w *Worker) accountByHash(addrHash common.Hash) (*data.Account, error) {
	accAddr := addrHash.Hex()
	acc := &data.Account{}
	err := w.db.SelectOneTo(acc, "WHERE eth_addr=substr($1, 27)", accAddr)
	return acc, err
}
