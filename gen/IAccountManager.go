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

// IAccountManagerMetaData contains all meta data concerning the IAccountManager contract.
var IAccountManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"authorizeUserOp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAccountManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountManagerMetaData.ABI instead.
var IAccountManagerABI = IAccountManagerMetaData.ABI

// IAccountManager is an auto generated Go binding around an Ethereum contract.
type IAccountManager struct {
	IAccountManagerCaller     // Read-only binding to the contract
	IAccountManagerTransactor // Write-only binding to the contract
	IAccountManagerFilterer   // Log filterer for contract events
}

// IAccountManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountManagerSession struct {
	Contract     *IAccountManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountManagerCallerSession struct {
	Contract *IAccountManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IAccountManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountManagerTransactorSession struct {
	Contract     *IAccountManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAccountManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountManagerRaw struct {
	Contract *IAccountManager // Generic contract binding to access the raw methods on
}

// IAccountManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountManagerCallerRaw struct {
	Contract *IAccountManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountManagerTransactorRaw struct {
	Contract *IAccountManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountManager creates a new instance of IAccountManager, bound to a specific deployed contract.
func NewIAccountManager(address common.Address, backend bind.ContractBackend) (*IAccountManager, error) {
	contract, err := bindIAccountManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountManager{IAccountManagerCaller: IAccountManagerCaller{contract: contract}, IAccountManagerTransactor: IAccountManagerTransactor{contract: contract}, IAccountManagerFilterer: IAccountManagerFilterer{contract: contract}}, nil
}

// NewIAccountManagerCaller creates a new read-only instance of IAccountManager, bound to a specific deployed contract.
func NewIAccountManagerCaller(address common.Address, caller bind.ContractCaller) (*IAccountManagerCaller, error) {
	contract, err := bindIAccountManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountManagerCaller{contract: contract}, nil
}

// NewIAccountManagerTransactor creates a new write-only instance of IAccountManager, bound to a specific deployed contract.
func NewIAccountManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountManagerTransactor, error) {
	contract, err := bindIAccountManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountManagerTransactor{contract: contract}, nil
}

// NewIAccountManagerFilterer creates a new log filterer instance of IAccountManager, bound to a specific deployed contract.
func NewIAccountManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountManagerFilterer, error) {
	contract, err := bindIAccountManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountManagerFilterer{contract: contract}, nil
}

// bindIAccountManager binds a generic wrapper to an already deployed contract.
func bindIAccountManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccountManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountManager *IAccountManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountManager.Contract.IAccountManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountManager *IAccountManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountManager.Contract.IAccountManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountManager *IAccountManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountManager.Contract.IAccountManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountManager *IAccountManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountManager *IAccountManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountManager *IAccountManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountManager.Contract.contract.Transact(opts, method, params...)
}

// AuthorizeUserOp is a free data retrieval call binding the contract method 0x9021ce1e.
//
// Solidity: function authorizeUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns()
func (_IAccountManager *IAccountManagerCaller) AuthorizeUserOp(opts *bind.CallOpts, userOp UserOperation) error {
	var out []interface{}
	err := _IAccountManager.contract.Call(opts, &out, "authorizeUserOp", userOp)

	if err != nil {
		return err
	}

	return err

}

// AuthorizeUserOp is a free data retrieval call binding the contract method 0x9021ce1e.
//
// Solidity: function authorizeUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns()
func (_IAccountManager *IAccountManagerSession) AuthorizeUserOp(userOp UserOperation) error {
	return _IAccountManager.Contract.AuthorizeUserOp(&_IAccountManager.CallOpts, userOp)
}

// AuthorizeUserOp is a free data retrieval call binding the contract method 0x9021ce1e.
//
// Solidity: function authorizeUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns()
func (_IAccountManager *IAccountManagerCallerSession) AuthorizeUserOp(userOp UserOperation) error {
	return _IAccountManager.Contract.AuthorizeUserOp(&_IAccountManager.CallOpts, userOp)
}

// AddUserAuthorization is a paid mutator transaction binding the contract method 0xe461f00d.
//
// Solidity: function addUserAuthorization(address user) returns(bool)
func (_IAccountManager *IAccountManagerTransactor) AddUserAuthorization(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IAccountManager.contract.Transact(opts, "addUserAuthorization", user)
}

// AddUserAuthorization is a paid mutator transaction binding the contract method 0xe461f00d.
//
// Solidity: function addUserAuthorization(address user) returns(bool)
func (_IAccountManager *IAccountManagerSession) AddUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _IAccountManager.Contract.AddUserAuthorization(&_IAccountManager.TransactOpts, user)
}

// AddUserAuthorization is a paid mutator transaction binding the contract method 0xe461f00d.
//
// Solidity: function addUserAuthorization(address user) returns(bool)
func (_IAccountManager *IAccountManagerTransactorSession) AddUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _IAccountManager.Contract.AddUserAuthorization(&_IAccountManager.TransactOpts, user)
}

// RemoveUserAuthorization is a paid mutator transaction binding the contract method 0xcfef9c97.
//
// Solidity: function removeUserAuthorization(address user) returns(bool)
func (_IAccountManager *IAccountManagerTransactor) RemoveUserAuthorization(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IAccountManager.contract.Transact(opts, "removeUserAuthorization", user)
}

// RemoveUserAuthorization is a paid mutator transaction binding the contract method 0xcfef9c97.
//
// Solidity: function removeUserAuthorization(address user) returns(bool)
func (_IAccountManager *IAccountManagerSession) RemoveUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _IAccountManager.Contract.RemoveUserAuthorization(&_IAccountManager.TransactOpts, user)
}

// RemoveUserAuthorization is a paid mutator transaction binding the contract method 0xcfef9c97.
//
// Solidity: function removeUserAuthorization(address user) returns(bool)
func (_IAccountManager *IAccountManagerTransactorSession) RemoveUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _IAccountManager.Contract.RemoveUserAuthorization(&_IAccountManager.TransactOpts, user)
}
