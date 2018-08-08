package eth

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/furkhat/micropayments/contract"
)

// BackEnd defines a typed wrappers for the PSC, PTC contracts and ethereum methods used by the worker.
type BackEnd interface {
	EthBalanceAt(context.Context, common.Address) (*big.Int, error)

	ContractBalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)

	CreateChannel(opts *bind.TransactOpts, agent common.Address,
		deposit *big.Int) (*types.Transaction, error)

	CloseChannel(*bind.TransactOpts, common.Address, uint32,
		*big.Int, []byte, []byte) (*types.Transaction, error)
}

type ethBackendInstance struct {
	payCntr *contract.PaymentContract
	conn    *ethclient.Client
	timeout uint64
}

// NewEthBackend returns eth back implementation.
func NewEthBackend(payCntr *contract.PaymentContract, conn *ethclient.Client,
	timeout uint64) BackEnd {
	return &ethBackendInstance{
		payCntr: payCntr,
		conn:    conn,
		timeout: timeout,
	}
}

func (b *ethBackendInstance) CloseChannel(opts *bind.TransactOpts,
	agent common.Address, block uint32, balance *big.Int,
	balanceSig, closingSig []byte) (*types.Transaction, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	tx, err := b.payCntr.CloseChannel(opts, agent, block,
		balance, balanceSig, closingSig)
	if err != nil {
		return nil, fmt.Errorf("failed to close: %s", err)
	}
	return tx, nil
}

func (b *ethBackendInstance) ContractBalanceOf(opts *bind.CallOpts,
	owner common.Address) (*big.Int, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	val, err := b.payCntr.BalanceOf(opts, owner)
	if err != nil {
		err = fmt.Errorf("failed to get balance: %s", err)
	}
	return val, err
}

func (b *ethBackendInstance) CreateChannel(opts *bind.TransactOpts,
	agent common.Address, deposit *big.Int) (*types.Transaction, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	tx, err := b.payCntr.CreateChannel(opts, agent, deposit)
	if err != nil {
		err = fmt.Errorf("failed to create PSC channel: %s", err)
	}
	return tx, err
}

func (b *ethBackendInstance) EthBalanceAt(ctx context.Context,
	owner common.Address) (*big.Int, error) {
	ctx2, cancel := b.addTimeout(ctx)
	defer cancel()

	return b.conn.BalanceAt(ctx2, owner, nil)
}

func (b *ethBackendInstance) addTimeout(
	ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithTimeout(ctx,
		time.Duration(b.timeout)*time.Second)
}
