// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PaymentContractABI is the input ABI used to generate the binding from.
const PaymentContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_agent\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_deposit\",\"type\":\"uint192\"},{\"indexed\":false,\"name\":\"_authentication_hash\",\"type\":\"bytes32\"}],\"name\":\"LogChannelCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_agent\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_open_block_number\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_balance\",\"type\":\"uint192\"}],\"name\":\"LogChannelClosed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_agent_address\",\"type\":\"address\"},{\"name\":\"_deposit\",\"type\":\"uint192\"}],\"name\":\"createChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_agent_address\",\"type\":\"address\"},{\"name\":\"_open_block_number\",\"type\":\"uint32\"},{\"name\":\"_balance\",\"type\":\"uint192\"},{\"name\":\"_balance_msg_sig\",\"type\":\"bytes\"},{\"name\":\"_closing_sig\",\"type\":\"bytes\"}],\"name\":\"closeChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PaymentContract is an auto generated Go binding around an Ethereum contract.
type PaymentContract struct {
	PaymentContractCaller     // Read-only binding to the contract
	PaymentContractTransactor // Write-only binding to the contract
	PaymentContractFilterer   // Log filterer for contract events
}

// PaymentContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentContractSession struct {
	Contract     *PaymentContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentContractCallerSession struct {
	Contract *PaymentContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PaymentContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentContractTransactorSession struct {
	Contract     *PaymentContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PaymentContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentContractRaw struct {
	Contract *PaymentContract // Generic contract binding to access the raw methods on
}

// PaymentContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentContractCallerRaw struct {
	Contract *PaymentContractCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentContractTransactorRaw struct {
	Contract *PaymentContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentContract creates a new instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContract(address common.Address, backend bind.ContractBackend) (*PaymentContract, error) {
	contract, err := bindPaymentContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentContract{PaymentContractCaller: PaymentContractCaller{contract: contract}, PaymentContractTransactor: PaymentContractTransactor{contract: contract}, PaymentContractFilterer: PaymentContractFilterer{contract: contract}}, nil
}

// NewPaymentContractCaller creates a new read-only instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContractCaller(address common.Address, caller bind.ContractCaller) (*PaymentContractCaller, error) {
	contract, err := bindPaymentContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentContractCaller{contract: contract}, nil
}

// NewPaymentContractTransactor creates a new write-only instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentContractTransactor, error) {
	contract, err := bindPaymentContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentContractTransactor{contract: contract}, nil
}

// NewPaymentContractFilterer creates a new log filterer instance of PaymentContract, bound to a specific deployed contract.
func NewPaymentContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentContractFilterer, error) {
	contract, err := bindPaymentContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentContractFilterer{contract: contract}, nil
}

// bindPaymentContract binds a generic wrapper to an already deployed contract.
func bindPaymentContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentContract *PaymentContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentContract.Contract.PaymentContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentContract *PaymentContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentContract.Contract.PaymentContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentContract *PaymentContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentContract.Contract.PaymentContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentContract *PaymentContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PaymentContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentContract *PaymentContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentContract *PaymentContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_address address) constant returns(uint256)
func (_PaymentContract *PaymentContractCaller) BalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PaymentContract.contract.Call(opts, out, "balanceOf", _address)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_address address) constant returns(uint256)
func (_PaymentContract *PaymentContractSession) BalanceOf(_address common.Address) (*big.Int, error) {
	return _PaymentContract.Contract.BalanceOf(&_PaymentContract.CallOpts, _address)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_address address) constant returns(uint256)
func (_PaymentContract *PaymentContractCallerSession) BalanceOf(_address common.Address) (*big.Int, error) {
	return _PaymentContract.Contract.BalanceOf(&_PaymentContract.CallOpts, _address)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x257cf25c.
//
// Solidity: function closeChannel(_agent_address address, _open_block_number uint32, _balance uint192, _balance_msg_sig bytes, _closing_sig bytes) returns()
func (_PaymentContract *PaymentContractTransactor) CloseChannel(opts *bind.TransactOpts, _agent_address common.Address, _open_block_number uint32, _balance *big.Int, _balance_msg_sig []byte, _closing_sig []byte) (*types.Transaction, error) {
	return _PaymentContract.contract.Transact(opts, "closeChannel", _agent_address, _open_block_number, _balance, _balance_msg_sig, _closing_sig)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x257cf25c.
//
// Solidity: function closeChannel(_agent_address address, _open_block_number uint32, _balance uint192, _balance_msg_sig bytes, _closing_sig bytes) returns()
func (_PaymentContract *PaymentContractSession) CloseChannel(_agent_address common.Address, _open_block_number uint32, _balance *big.Int, _balance_msg_sig []byte, _closing_sig []byte) (*types.Transaction, error) {
	return _PaymentContract.Contract.CloseChannel(&_PaymentContract.TransactOpts, _agent_address, _open_block_number, _balance, _balance_msg_sig, _closing_sig)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x257cf25c.
//
// Solidity: function closeChannel(_agent_address address, _open_block_number uint32, _balance uint192, _balance_msg_sig bytes, _closing_sig bytes) returns()
func (_PaymentContract *PaymentContractTransactorSession) CloseChannel(_agent_address common.Address, _open_block_number uint32, _balance *big.Int, _balance_msg_sig []byte, _closing_sig []byte) (*types.Transaction, error) {
	return _PaymentContract.Contract.CloseChannel(&_PaymentContract.TransactOpts, _agent_address, _open_block_number, _balance, _balance_msg_sig, _closing_sig)
}

// CreateChannel is a paid mutator transaction binding the contract method 0xa6d15963.
//
// Solidity: function createChannel(_agent_address address, _deposit uint192) returns()
func (_PaymentContract *PaymentContractTransactor) CreateChannel(opts *bind.TransactOpts, _agent_address common.Address, _deposit *big.Int) (*types.Transaction, error) {
	return _PaymentContract.contract.Transact(opts, "createChannel", _agent_address, _deposit)
}

// CreateChannel is a paid mutator transaction binding the contract method 0xa6d15963.
//
// Solidity: function createChannel(_agent_address address, _deposit uint192) returns()
func (_PaymentContract *PaymentContractSession) CreateChannel(_agent_address common.Address, _deposit *big.Int) (*types.Transaction, error) {
	return _PaymentContract.Contract.CreateChannel(&_PaymentContract.TransactOpts, _agent_address, _deposit)
}

// CreateChannel is a paid mutator transaction binding the contract method 0xa6d15963.
//
// Solidity: function createChannel(_agent_address address, _deposit uint192) returns()
func (_PaymentContract *PaymentContractTransactorSession) CreateChannel(_agent_address common.Address, _deposit *big.Int) (*types.Transaction, error) {
	return _PaymentContract.Contract.CreateChannel(&_PaymentContract.TransactOpts, _agent_address, _deposit)
}

// PaymentContractLogChannelClosedIterator is returned from FilterLogChannelClosed and is used to iterate over the raw logs and unpacked data for LogChannelClosed events raised by the PaymentContract contract.
type PaymentContractLogChannelClosedIterator struct {
	Event *PaymentContractLogChannelClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentContractLogChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentContractLogChannelClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentContractLogChannelClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentContractLogChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentContractLogChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentContractLogChannelClosed represents a LogChannelClosed event raised by the PaymentContract contract.
type PaymentContractLogChannelClosed struct {
	Agent           common.Address
	Client          common.Address
	OpenBlockNumber uint32
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogChannelClosed is a free log retrieval operation binding the contract event 0xd2963ebeb5ae5f633cc03dd638fa8614c15d4c28fce3db10e6b4a6326c92d6c8.
//
// Solidity: e LogChannelClosed(_agent indexed address, _client indexed address, _open_block_number uint32, _balance uint192)
func (_PaymentContract *PaymentContractFilterer) FilterLogChannelClosed(opts *bind.FilterOpts, _agent []common.Address, _client []common.Address) (*PaymentContractLogChannelClosedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}
	var _clientRule []interface{}
	for _, _clientItem := range _client {
		_clientRule = append(_clientRule, _clientItem)
	}

	logs, sub, err := _PaymentContract.contract.FilterLogs(opts, "LogChannelClosed", _agentRule, _clientRule)
	if err != nil {
		return nil, err
	}
	return &PaymentContractLogChannelClosedIterator{contract: _PaymentContract.contract, event: "LogChannelClosed", logs: logs, sub: sub}, nil
}

// WatchLogChannelClosed is a free log subscription operation binding the contract event 0xd2963ebeb5ae5f633cc03dd638fa8614c15d4c28fce3db10e6b4a6326c92d6c8.
//
// Solidity: e LogChannelClosed(_agent indexed address, _client indexed address, _open_block_number uint32, _balance uint192)
func (_PaymentContract *PaymentContractFilterer) WatchLogChannelClosed(opts *bind.WatchOpts, sink chan<- *PaymentContractLogChannelClosed, _agent []common.Address, _client []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}
	var _clientRule []interface{}
	for _, _clientItem := range _client {
		_clientRule = append(_clientRule, _clientItem)
	}

	logs, sub, err := _PaymentContract.contract.WatchLogs(opts, "LogChannelClosed", _agentRule, _clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentContractLogChannelClosed)
				if err := _PaymentContract.contract.UnpackLog(event, "LogChannelClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PaymentContractLogChannelCreatedIterator is returned from FilterLogChannelCreated and is used to iterate over the raw logs and unpacked data for LogChannelCreated events raised by the PaymentContract contract.
type PaymentContractLogChannelCreatedIterator struct {
	Event *PaymentContractLogChannelCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaymentContractLogChannelCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentContractLogChannelCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaymentContractLogChannelCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaymentContractLogChannelCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentContractLogChannelCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentContractLogChannelCreated represents a LogChannelCreated event raised by the PaymentContract contract.
type PaymentContractLogChannelCreated struct {
	Agent              common.Address
	Client             common.Address
	Deposit            *big.Int
	AuthenticationHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogChannelCreated is a free log retrieval operation binding the contract event 0x5b0f7ea5731202a04625655b219ed16bb526d1908967a4758fbbd863b24338a0.
//
// Solidity: e LogChannelCreated(_agent indexed address, _client indexed address, _deposit uint192, _authentication_hash bytes32)
func (_PaymentContract *PaymentContractFilterer) FilterLogChannelCreated(opts *bind.FilterOpts, _agent []common.Address, _client []common.Address) (*PaymentContractLogChannelCreatedIterator, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}
	var _clientRule []interface{}
	for _, _clientItem := range _client {
		_clientRule = append(_clientRule, _clientItem)
	}

	logs, sub, err := _PaymentContract.contract.FilterLogs(opts, "LogChannelCreated", _agentRule, _clientRule)
	if err != nil {
		return nil, err
	}
	return &PaymentContractLogChannelCreatedIterator{contract: _PaymentContract.contract, event: "LogChannelCreated", logs: logs, sub: sub}, nil
}

// WatchLogChannelCreated is a free log subscription operation binding the contract event 0x5b0f7ea5731202a04625655b219ed16bb526d1908967a4758fbbd863b24338a0.
//
// Solidity: e LogChannelCreated(_agent indexed address, _client indexed address, _deposit uint192, _authentication_hash bytes32)
func (_PaymentContract *PaymentContractFilterer) WatchLogChannelCreated(opts *bind.WatchOpts, sink chan<- *PaymentContractLogChannelCreated, _agent []common.Address, _client []common.Address) (event.Subscription, error) {

	var _agentRule []interface{}
	for _, _agentItem := range _agent {
		_agentRule = append(_agentRule, _agentItem)
	}
	var _clientRule []interface{}
	for _, _clientItem := range _client {
		_clientRule = append(_clientRule, _clientItem)
	}

	logs, sub, err := _PaymentContract.contract.WatchLogs(opts, "LogChannelCreated", _agentRule, _clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentContractLogChannelCreated)
				if err := _PaymentContract.contract.UnpackLog(event, "LogChannelCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
