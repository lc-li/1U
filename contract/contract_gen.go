// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RandomNumberRandomRequest is an auto generated low-level Go binding around an user-defined struct.
type RandomNumberRandomRequest struct {
	RoundId       *big.Int
	RandomNumbers []*big.Int
	Fulfilled     bool
	Timestamp     uint64
}

// RandomNumberMetaData contains all meta data concerning the RandomNumber contract.
var RandomNumberMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subscriptionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"RandomNumber__InvalidCallbackGasLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinator\",\"type\":\"address\"}],\"name\":\"RandomNumber__InvalidCoordinator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"RandomNumber__InvalidSubscriptionId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RandomNumber__RequestAlreadyFulfilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RandomNumber__RequestNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RandomNumber__RequestPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RandomNumber__ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint96\",\"name\":\"roundId\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RequestedRandomness\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint96\",\"name\":\"roundId\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"randomNumbers\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RandomnessFulfilled\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCurrentRound\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getLatestRandomNumber\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getRandomRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"roundId\",\"type\":\"uint96\"},{\"internalType\":\"uint256[]\",\"name\":\"randomNumbers\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structRandomNumber.RandomRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"s_requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_subscriptionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RandomNumberABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomNumberMetaData.ABI instead.
var RandomNumberABI = RandomNumberMetaData.ABI

// RandomNumber is an auto generated Go binding around an Ethereum contract.
type RandomNumber struct {
	RandomNumberCaller     // Read-only binding to the contract
	RandomNumberTransactor // Write-only binding to the contract
	RandomNumberFilterer   // Log filterer for contract events
}

// RandomNumberCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomNumberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomNumberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomNumberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomNumberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomNumberSession struct {
	Contract     *RandomNumber     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomNumberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomNumberCallerSession struct {
	Contract *RandomNumberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RandomNumberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomNumberTransactorSession struct {
	Contract     *RandomNumberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RandomNumberRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomNumberRaw struct {
	Contract *RandomNumber // Generic contract binding to access the raw methods on
}

// RandomNumberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomNumberCallerRaw struct {
	Contract *RandomNumberCaller // Generic read-only contract binding to access the raw methods on
}

// RandomNumberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomNumberTransactorRaw struct {
	Contract *RandomNumberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomNumber creates a new instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumber(address common.Address, backend bind.ContractBackend) (*RandomNumber, error) {
	contract, err := bindRandomNumber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomNumber{RandomNumberCaller: RandomNumberCaller{contract: contract}, RandomNumberTransactor: RandomNumberTransactor{contract: contract}, RandomNumberFilterer: RandomNumberFilterer{contract: contract}}, nil
}

// NewRandomNumberCaller creates a new read-only instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumberCaller(address common.Address, caller bind.ContractCaller) (*RandomNumberCaller, error) {
	contract, err := bindRandomNumber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomNumberCaller{contract: contract}, nil
}

// NewRandomNumberTransactor creates a new write-only instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumberTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomNumberTransactor, error) {
	contract, err := bindRandomNumber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomNumberTransactor{contract: contract}, nil
}

// NewRandomNumberFilterer creates a new log filterer instance of RandomNumber, bound to a specific deployed contract.
func NewRandomNumberFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomNumberFilterer, error) {
	contract, err := bindRandomNumber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomNumberFilterer{contract: contract}, nil
}

// bindRandomNumber binds a generic wrapper to an already deployed contract.
func bindRandomNumber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RandomNumberMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomNumber *RandomNumberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomNumber.Contract.RandomNumberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomNumber *RandomNumberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomNumber.Contract.RandomNumberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomNumber *RandomNumberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomNumber.Contract.RandomNumberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomNumber *RandomNumberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomNumber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomNumber *RandomNumberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomNumber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomNumber *RandomNumberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomNumber.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint96)
func (_RandomNumber *RandomNumberCaller) GetCurrentRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomNumber.contract.Call(opts, &out, "getCurrentRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint96)
func (_RandomNumber *RandomNumberSession) GetCurrentRound() (*big.Int, error) {
	return _RandomNumber.Contract.GetCurrentRound(&_RandomNumber.CallOpts)
}

// GetCurrentRound is a free data retrieval call binding the contract method 0xa32bf597.
//
// Solidity: function getCurrentRound() view returns(uint96)
func (_RandomNumber *RandomNumberCallerSession) GetCurrentRound() (*big.Int, error) {
	return _RandomNumber.Contract.GetCurrentRound(&_RandomNumber.CallOpts)
}

// GetLatestRandomNumber is a free data retrieval call binding the contract method 0x82f28d18.
//
// Solidity: function getLatestRandomNumber(uint256 requestId) view returns(uint256[])
func (_RandomNumber *RandomNumberCaller) GetLatestRandomNumber(opts *bind.CallOpts, requestId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _RandomNumber.contract.Call(opts, &out, "getLatestRandomNumber", requestId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLatestRandomNumber is a free data retrieval call binding the contract method 0x82f28d18.
//
// Solidity: function getLatestRandomNumber(uint256 requestId) view returns(uint256[])
func (_RandomNumber *RandomNumberSession) GetLatestRandomNumber(requestId *big.Int) ([]*big.Int, error) {
	return _RandomNumber.Contract.GetLatestRandomNumber(&_RandomNumber.CallOpts, requestId)
}

// GetLatestRandomNumber is a free data retrieval call binding the contract method 0x82f28d18.
//
// Solidity: function getLatestRandomNumber(uint256 requestId) view returns(uint256[])
func (_RandomNumber *RandomNumberCallerSession) GetLatestRandomNumber(requestId *big.Int) ([]*big.Int, error) {
	return _RandomNumber.Contract.GetLatestRandomNumber(&_RandomNumber.CallOpts, requestId)
}

// GetRandomRequest is a free data retrieval call binding the contract method 0xaae31df3.
//
// Solidity: function getRandomRequest(uint256 requestId) view returns((uint96,uint256[],bool,uint64))
func (_RandomNumber *RandomNumberCaller) GetRandomRequest(opts *bind.CallOpts, requestId *big.Int) (RandomNumberRandomRequest, error) {
	var out []interface{}
	err := _RandomNumber.contract.Call(opts, &out, "getRandomRequest", requestId)

	if err != nil {
		return *new(RandomNumberRandomRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(RandomNumberRandomRequest)).(*RandomNumberRandomRequest)

	return out0, err

}

// GetRandomRequest is a free data retrieval call binding the contract method 0xaae31df3.
//
// Solidity: function getRandomRequest(uint256 requestId) view returns((uint96,uint256[],bool,uint64))
func (_RandomNumber *RandomNumberSession) GetRandomRequest(requestId *big.Int) (RandomNumberRandomRequest, error) {
	return _RandomNumber.Contract.GetRandomRequest(&_RandomNumber.CallOpts, requestId)
}

// GetRandomRequest is a free data retrieval call binding the contract method 0xaae31df3.
//
// Solidity: function getRandomRequest(uint256 requestId) view returns((uint96,uint256[],bool,uint64))
func (_RandomNumber *RandomNumberCallerSession) GetRandomRequest(requestId *big.Int) (RandomNumberRandomRequest, error) {
	return _RandomNumber.Contract.GetRandomRequest(&_RandomNumber.CallOpts, requestId)
}

// SSubscriptionId is a free data retrieval call binding the contract method 0x8ac00021.
//
// Solidity: function s_subscriptionId() view returns(uint256)
func (_RandomNumber *RandomNumberCaller) SSubscriptionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomNumber.contract.Call(opts, &out, "s_subscriptionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SSubscriptionId is a free data retrieval call binding the contract method 0x8ac00021.
//
// Solidity: function s_subscriptionId() view returns(uint256)
func (_RandomNumber *RandomNumberSession) SSubscriptionId() (*big.Int, error) {
	return _RandomNumber.Contract.SSubscriptionId(&_RandomNumber.CallOpts)
}

// SSubscriptionId is a free data retrieval call binding the contract method 0x8ac00021.
//
// Solidity: function s_subscriptionId() view returns(uint256)
func (_RandomNumber *RandomNumberCallerSession) SSubscriptionId() (*big.Int, error) {
	return _RandomNumber.Contract.SSubscriptionId(&_RandomNumber.CallOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x09e26daf.
//
// Solidity: function requestRandomWords(uint32 numWords, uint32 callbackGasLimit, uint16 requestConfirmations) returns(uint256 s_requestId)
func (_RandomNumber *RandomNumberTransactor) RequestRandomWords(opts *bind.TransactOpts, numWords uint32, callbackGasLimit uint32, requestConfirmations uint16) (*types.Transaction, error) {
	return _RandomNumber.contract.Transact(opts, "requestRandomWords", numWords, callbackGasLimit, requestConfirmations)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x09e26daf.
//
// Solidity: function requestRandomWords(uint32 numWords, uint32 callbackGasLimit, uint16 requestConfirmations) returns(uint256 s_requestId)
func (_RandomNumber *RandomNumberSession) RequestRandomWords(numWords uint32, callbackGasLimit uint32, requestConfirmations uint16) (*types.Transaction, error) {
	return _RandomNumber.Contract.RequestRandomWords(&_RandomNumber.TransactOpts, numWords, callbackGasLimit, requestConfirmations)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x09e26daf.
//
// Solidity: function requestRandomWords(uint32 numWords, uint32 callbackGasLimit, uint16 requestConfirmations) returns(uint256 s_requestId)
func (_RandomNumber *RandomNumberTransactorSession) RequestRandomWords(numWords uint32, callbackGasLimit uint32, requestConfirmations uint16) (*types.Transaction, error) {
	return _RandomNumber.Contract.RequestRandomWords(&_RandomNumber.TransactOpts, numWords, callbackGasLimit, requestConfirmations)
}

// RandomNumberRandomnessFulfilledIterator is returned from FilterRandomnessFulfilled and is used to iterate over the raw logs and unpacked data for RandomnessFulfilled events raised by the RandomNumber contract.
type RandomNumberRandomnessFulfilledIterator struct {
	Event *RandomNumberRandomnessFulfilled // Event containing the contract specifics and raw log

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
func (it *RandomNumberRandomnessFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberRandomnessFulfilled)
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
		it.Event = new(RandomNumberRandomnessFulfilled)
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
func (it *RandomNumberRandomnessFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberRandomnessFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberRandomnessFulfilled represents a RandomnessFulfilled event raised by the RandomNumber contract.
type RandomNumberRandomnessFulfilled struct {
	RequestId     *big.Int
	RoundId       *big.Int
	RandomNumbers []*big.Int
	Timestamp     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRandomnessFulfilled is a free log retrieval operation binding the contract event 0x9cef3d202ba79b1ca1780bc42bcf30044a325f09f828d5b7bf750c38f4415f0d.
//
// Solidity: event RandomnessFulfilled(uint256 indexed requestId, uint96 indexed roundId, uint256[] randomNumbers, uint64 timestamp)
func (_RandomNumber *RandomNumberFilterer) FilterRandomnessFulfilled(opts *bind.FilterOpts, requestId []*big.Int, roundId []*big.Int) (*RandomNumberRandomnessFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _RandomNumber.contract.FilterLogs(opts, "RandomnessFulfilled", requestIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomNumberRandomnessFulfilledIterator{contract: _RandomNumber.contract, event: "RandomnessFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomnessFulfilled is a free log subscription operation binding the contract event 0x9cef3d202ba79b1ca1780bc42bcf30044a325f09f828d5b7bf750c38f4415f0d.
//
// Solidity: event RandomnessFulfilled(uint256 indexed requestId, uint96 indexed roundId, uint256[] randomNumbers, uint64 timestamp)
func (_RandomNumber *RandomNumberFilterer) WatchRandomnessFulfilled(opts *bind.WatchOpts, sink chan<- *RandomNumberRandomnessFulfilled, requestId []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _RandomNumber.contract.WatchLogs(opts, "RandomnessFulfilled", requestIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberRandomnessFulfilled)
				if err := _RandomNumber.contract.UnpackLog(event, "RandomnessFulfilled", log); err != nil {
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

// ParseRandomnessFulfilled is a log parse operation binding the contract event 0x9cef3d202ba79b1ca1780bc42bcf30044a325f09f828d5b7bf750c38f4415f0d.
//
// Solidity: event RandomnessFulfilled(uint256 indexed requestId, uint96 indexed roundId, uint256[] randomNumbers, uint64 timestamp)
func (_RandomNumber *RandomNumberFilterer) ParseRandomnessFulfilled(log types.Log) (*RandomNumberRandomnessFulfilled, error) {
	event := new(RandomNumberRandomnessFulfilled)
	if err := _RandomNumber.contract.UnpackLog(event, "RandomnessFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomNumberRequestedRandomnessIterator is returned from FilterRequestedRandomness and is used to iterate over the raw logs and unpacked data for RequestedRandomness events raised by the RandomNumber contract.
type RandomNumberRequestedRandomnessIterator struct {
	Event *RandomNumberRequestedRandomness // Event containing the contract specifics and raw log

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
func (it *RandomNumberRequestedRandomnessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomNumberRequestedRandomness)
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
		it.Event = new(RandomNumberRequestedRandomness)
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
func (it *RandomNumberRequestedRandomnessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomNumberRequestedRandomnessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomNumberRequestedRandomness represents a RequestedRandomness event raised by the RandomNumber contract.
type RandomNumberRequestedRandomness struct {
	RequestId *big.Int
	RoundId   *big.Int
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestedRandomness is a free log retrieval operation binding the contract event 0xd2f18edafb3c7b9ff10065d14d6c01bbebdb1c77ddaef191f589b6e0b1ed3d07.
//
// Solidity: event RequestedRandomness(uint256 indexed requestId, uint96 indexed roundId, uint64 timestamp)
func (_RandomNumber *RandomNumberFilterer) FilterRequestedRandomness(opts *bind.FilterOpts, requestId []*big.Int, roundId []*big.Int) (*RandomNumberRequestedRandomnessIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _RandomNumber.contract.FilterLogs(opts, "RequestedRandomness", requestIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomNumberRequestedRandomnessIterator{contract: _RandomNumber.contract, event: "RequestedRandomness", logs: logs, sub: sub}, nil
}

// WatchRequestedRandomness is a free log subscription operation binding the contract event 0xd2f18edafb3c7b9ff10065d14d6c01bbebdb1c77ddaef191f589b6e0b1ed3d07.
//
// Solidity: event RequestedRandomness(uint256 indexed requestId, uint96 indexed roundId, uint64 timestamp)
func (_RandomNumber *RandomNumberFilterer) WatchRequestedRandomness(opts *bind.WatchOpts, sink chan<- *RandomNumberRequestedRandomness, requestId []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _RandomNumber.contract.WatchLogs(opts, "RequestedRandomness", requestIdRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomNumberRequestedRandomness)
				if err := _RandomNumber.contract.UnpackLog(event, "RequestedRandomness", log); err != nil {
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

// ParseRequestedRandomness is a log parse operation binding the contract event 0xd2f18edafb3c7b9ff10065d14d6c01bbebdb1c77ddaef191f589b6e0b1ed3d07.
//
// Solidity: event RequestedRandomness(uint256 indexed requestId, uint96 indexed roundId, uint64 timestamp)
func (_RandomNumber *RandomNumberFilterer) ParseRequestedRandomness(log types.Log) (*RandomNumberRequestedRandomness, error) {
	event := new(RandomNumberRequestedRandomness)
	if err := _RandomNumber.contract.UnpackLog(event, "RequestedRandomness", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
