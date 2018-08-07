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

	PSCBalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)

	PSCCreateChannel(opts *bind.TransactOpts,
		agent common.Address, hash [common.HashLength]byte,
		deposit *big.Int) (*types.Transaction, error)

	PSCCooperativeClose(*bind.TransactOpts, common.Address, uint32,
		[common.HashLength]byte, *big.Int, []byte, []byte) (*types.Transaction, error)

	PTCIncreaseApproval(*bind.TransactOpts, common.Address, *big.Int) (*types.Transaction, error)

	PSCAddBalanceERC20(*bind.TransactOpts, *big.Int) (*types.Transaction, error)

	PTCBalanceOf(*bind.CallOpts, common.Address) (*big.Int, error)
}

type ethBackendInstance struct {
	psc     *contract.PrivatixServiceContract
	ptc     *contract.PrivatixTokenContract
	conn    *ethclient.Client
	timeout uint64
}

// NewEthBackend returns eth back implementation.
func NewEthBackend(psc *contract.PrivatixServiceContract,
	ptc *contract.PrivatixTokenContract, conn *ethclient.Client,
	timeout uint64) BackEnd {
	return &ethBackendInstance{
		psc:     psc,
		ptc:     ptc,
		conn:    conn,
		timeout: timeout,
	}
}

func (b *ethBackendInstance) PSCCooperativeClose(opts *bind.TransactOpts,
	agent common.Address, block uint32, offeringHash [common.HashLength]byte,
	balance *big.Int, balanceSig, closingSig []byte) (*types.Transaction, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	tx, err := b.psc.CooperativeClose(opts, agent, block, offeringHash,
		balance, balanceSig, closingSig)
	if err != nil {
		return nil, fmt.Errorf("failed to do cooperative close: %s", err)
	}
	return tx, nil
}

func (b *ethBackendInstance) PTCBalanceOf(opts *bind.CallOpts,
	owner common.Address) (*big.Int, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	val, err := b.ptc.BalanceOf(opts, owner)
	if err != nil {
		err = fmt.Errorf("failed to get PTC balance: %s", err)
	}
	return val, err
}

func (b *ethBackendInstance) PTCIncreaseApproval(opts *bind.TransactOpts,
	spender common.Address, addedVal *big.Int) (*types.Transaction, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	tx, err := b.ptc.IncreaseApproval(opts, spender, addedVal)
	if err != nil {
		return nil, fmt.Errorf("failed to PTC increase approval: %s", err)
	}
	return tx, nil
}

func (b *ethBackendInstance) PSCBalanceOf(opts *bind.CallOpts,
	owner common.Address) (*big.Int, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	val, err := b.psc.BalanceOf(opts, owner)
	if err != nil {
		err = fmt.Errorf("failed to get PSC balance: %s", err)
	}
	return val, err
}

func (b *ethBackendInstance) PSCAddBalanceERC20(opts *bind.TransactOpts,
	amount *big.Int) (*types.Transaction, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	tx, err := b.psc.AddBalanceERC20(opts, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to add ERC20 balance: %s", err)
	}
	return tx, nil
}

func (b *ethBackendInstance) PSCCreateChannel(opts *bind.TransactOpts,
	agent common.Address, hash [common.HashLength]byte,
	deposit *big.Int) (*types.Transaction, error) {
	ctx2, cancel := b.addTimeout(opts.Context)
	defer cancel()

	opts.Context = ctx2

	tx, err := b.psc.CreateChannel(opts, agent, hash, deposit, common.Hash{})
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
