package worker

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type testEthBackCall struct {
	txOpts *bind.TransactOpts
	method string
	caller common.Address
	args   []interface{}
}

type testEthBackend struct {
	callStack       []testEthBackCall
	balanceEth      *big.Int
	balanceContract *big.Int
}

func newTestEthBackend() *testEthBackend {
	return &testEthBackend{}
}

func (b *testEthBackend) CloseChannel(opts *bind.TransactOpts,
	agentAddr common.Address, block uint32, balance *big.Int,
	balanceMsgSig []byte, ClosingSig []byte) (*types.Transaction, error) {
	b.callStack = append(b.callStack, testEthBackCall{
		method: "PSCCooperativeClose",
		caller: opts.From,
		txOpts: opts,
		args: []interface{}{agentAddr, block, balance,
			balanceMsgSig, ClosingSig},
	})
	tx := types.NewTransaction(0, common.Address{}, big.NewInt(1), 1, big.NewInt(1), nil)
	return tx, nil
}

func (b *testEthBackend) EthBalanceAt(_ context.Context,
	addr common.Address) (*big.Int, error) {
	b.callStack = append(b.callStack, testEthBackCall{
		method: "EthBalanceAt",
		args:   []interface{}{addr},
	})
	return b.balanceEth, nil
}

func (b *testEthBackend) ContractBalanceOf(opts *bind.CallOpts,
	addr common.Address) (*big.Int, error) {
	b.callStack = append(b.callStack, testEthBackCall{
		method: "ContractBalanceOf",
		caller: opts.From,
		args:   []interface{}{addr},
	})
	return b.balanceContract, nil
}

func (b *testEthBackend) testCalled(t *testing.T, method string,
	caller common.Address, gasLimit uint64, args ...interface{}) {
	if len(b.callStack) == 0 {
		t.Fatalf("method %s not called. Callstack is empty", method)
	}
	for _, call := range b.callStack {
		if caller == call.caller && method == call.method &&
			reflect.DeepEqual(args, call.args) &&
			(call.txOpts == nil || call.txOpts.GasLimit == gasLimit) {
			return
		}
	}
	t.Logf("%+v\n", b.callStack)
	t.Fatalf("no call of %s from %v with args: %v", method, caller, args)
}

const (
	testTXNonce    uint64 = 1
	testTXGasLimit uint64 = 2
	testTXGasPrice int64  = 1
)

func (b *testEthBackend) CreateChannel(opts *bind.TransactOpts,
	agent common.Address, deposit *big.Int) (*types.Transaction, error) {
	b.callStack = append(b.callStack, testEthBackCall{
		method: "PSCCreateChannel",
		caller: opts.From,
		args:   []interface{}{agent, deposit},
	})

	tx := types.NewTransaction(
		testTXNonce, agent, deposit, testTXGasLimit,
		big.NewInt(testTXGasPrice), []byte{})

	return tx, nil
}
