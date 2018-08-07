package worker

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/furkhat/micropayments/data"
	reform "gopkg.in/reform.v1"
)

// EthBack defines a typed wrappers for the PSC, PTC contracts and ethereum methods used by the worker.
type EthBack interface {
	EthBalanceAt(context.Context, common.Address) (*big.Int, error)

	PSCBalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)

	PSCCooperativeClose(*bind.TransactOpts, common.Address, uint32,
		[common.HashLength]byte, *big.Int, []byte, []byte) (*types.Transaction, error)

	PSCAddBalanceERC20(*bind.TransactOpts, *big.Int) (*types.Transaction, error)

	PTCBalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)
}

// Worker defines methods for processing ethereum logs.
type Worker struct {
	db      *reform.DB
	ethBack EthBack
}

// NewWorker creates a Worker.
func NewWorker(db *reform.DB, ethBack EthBack) *Worker {
	return &Worker{db, ethBack}
}

// AfterApprove transfers all approved amount to the spender.
func (w *Worker) AfterApprove(approveLog *ethtypes.Log) error {
	accAddr := approveLog.Topics[1].Hex()
	acc := &data.Account{}
	if err := w.db.SelectOneTo(acc, "WHERE eth_addr=substr($1, 27)", accAddr); err != nil {
		return err
	}
	auth, err := accTransactOpts(acc)
	if err != nil {
		return err
	}
	auth.GasLimit = gasLimitAddBalanceERC20
	args, err := approvalNonIndexArgs.UnpackValues(approveLog.Data)
	if err != nil {
		return err
	}
	_, err = w.ethBack.PSCAddBalanceERC20(auth, args[0].(*big.Int))
	return err
}

// AfterTransfer update accounts PTC, PSC and ethereum balances.
func (w *Worker) AfterTransfer(transferLog *ethtypes.Log) error {
	return nil
}

// AfterChannelCreated registers channel in the system.
func (w *Worker) AfterChannelCreated(transferLog *ethtypes.Log) error {
	return nil
}

// AfterChannelClose update accounts PTC, PSC and ethereum balances.
func (w *Worker) AfterChannelClose(transferLog *ethtypes.Log) error {
	return nil
}
