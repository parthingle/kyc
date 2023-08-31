// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gen

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

// TestCounterMetaData contains all meta data concerning the TestCounter contract.
var TestCounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decrementCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805534801561001457600080fd5b506101d7806100246000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635b34b96614610046578063a87d942c14610050578063f5c5ad831461006e575b600080fd5b61004e610078565b005b610058610093565b60405161006591906100d0565b60405180910390f35b61007661009c565b005b600160008082825461008a919061011a565b92505081905550565b60008054905090565b60016000808282546100ae919061015e565b92505081905550565b6000819050919050565b6100ca816100b7565b82525050565b60006020820190506100e560008301846100c1565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610125826100b7565b9150610130836100b7565b925082820190508281121560008312168382126000841215161715610158576101576100eb565b5b92915050565b6000610169826100b7565b9150610174836100b7565b925082820390508181126000841216828213600085121516171561019b5761019a6100eb565b5b9291505056fea2646970667358221220a1129142de372de068bb6797773b37f4f614479b655f8d9e34f175f6af5a4f8664736f6c63430008130033",
}

// TestCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use TestCounterMetaData.ABI instead.
var TestCounterABI = TestCounterMetaData.ABI

// TestCounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestCounterMetaData.Bin instead.
var TestCounterBin = TestCounterMetaData.Bin

// DeployTestCounter deploys a new Ethereum contract, binding an instance of TestCounter to it.
func DeployTestCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestCounter, error) {
	parsed, err := TestCounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestCounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestCounter{TestCounterCaller: TestCounterCaller{contract: contract}, TestCounterTransactor: TestCounterTransactor{contract: contract}, TestCounterFilterer: TestCounterFilterer{contract: contract}}, nil
}

// TestCounter is an auto generated Go binding around an Ethereum contract.
type TestCounter struct {
	TestCounterCaller     // Read-only binding to the contract
	TestCounterTransactor // Write-only binding to the contract
	TestCounterFilterer   // Log filterer for contract events
}

// TestCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestCounterSession struct {
	Contract     *TestCounter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCounterCallerSession struct {
	Contract *TestCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestCounterTransactorSession struct {
	Contract     *TestCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestCounterRaw struct {
	Contract *TestCounter // Generic contract binding to access the raw methods on
}

// TestCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCounterCallerRaw struct {
	Contract *TestCounterCaller // Generic read-only contract binding to access the raw methods on
}

// TestCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestCounterTransactorRaw struct {
	Contract *TestCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestCounter creates a new instance of TestCounter, bound to a specific deployed contract.
func NewTestCounter(address common.Address, backend bind.ContractBackend) (*TestCounter, error) {
	contract, err := bindTestCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestCounter{TestCounterCaller: TestCounterCaller{contract: contract}, TestCounterTransactor: TestCounterTransactor{contract: contract}, TestCounterFilterer: TestCounterFilterer{contract: contract}}, nil
}

// NewTestCounterCaller creates a new read-only instance of TestCounter, bound to a specific deployed contract.
func NewTestCounterCaller(address common.Address, caller bind.ContractCaller) (*TestCounterCaller, error) {
	contract, err := bindTestCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCounterCaller{contract: contract}, nil
}

// NewTestCounterTransactor creates a new write-only instance of TestCounter, bound to a specific deployed contract.
func NewTestCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*TestCounterTransactor, error) {
	contract, err := bindTestCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestCounterTransactor{contract: contract}, nil
}

// NewTestCounterFilterer creates a new log filterer instance of TestCounter, bound to a specific deployed contract.
func NewTestCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*TestCounterFilterer, error) {
	contract, err := bindTestCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestCounterFilterer{contract: contract}, nil
}

// bindTestCounter binds a generic wrapper to an already deployed contract.
func bindTestCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestCounter *TestCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestCounter.Contract.TestCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestCounter *TestCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.Contract.TestCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestCounter *TestCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestCounter.Contract.TestCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestCounter *TestCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestCounter *TestCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestCounter *TestCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestCounter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_TestCounter *TestCounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestCounter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_TestCounter *TestCounterSession) GetCount() (*big.Int, error) {
	return _TestCounter.Contract.GetCount(&_TestCounter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_TestCounter *TestCounterCallerSession) GetCount() (*big.Int, error) {
	return _TestCounter.Contract.GetCount(&_TestCounter.CallOpts)
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_TestCounter *TestCounterTransactor) DecrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.contract.Transact(opts, "decrementCounter")
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_TestCounter *TestCounterSession) DecrementCounter() (*types.Transaction, error) {
	return _TestCounter.Contract.DecrementCounter(&_TestCounter.TransactOpts)
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_TestCounter *TestCounterTransactorSession) DecrementCounter() (*types.Transaction, error) {
	return _TestCounter.Contract.DecrementCounter(&_TestCounter.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_TestCounter *TestCounterTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestCounter.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_TestCounter *TestCounterSession) IncrementCounter() (*types.Transaction, error) {
	return _TestCounter.Contract.IncrementCounter(&_TestCounter.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_TestCounter *TestCounterTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _TestCounter.Contract.IncrementCounter(&_TestCounter.TransactOpts)
}
