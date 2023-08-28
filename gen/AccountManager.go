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

// AccountManagerMetaData contains all meta data concerning the AccountManager contract.
var AccountManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_masterWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"authorizeUserOp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610758380380610758833981810160405281019061003191906100c9565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506100f4565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100988261006f565b9050919050565b6100a88161008e565b81146100b2575f80fd5b50565b5f815190506100c38161009f565b92915050565b5f602082840312156100de576100dd61006b565b5b5f6100eb848285016100b5565b91505092915050565b60805161063e61011a5f395f81816101de015281816102c901526103b4015261063e5ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80631828983a146100595780639021ce1e14610089578063cfef9c97146100a5578063e461f00d146100d5578063fc0d011714610105575b5f80fd5b610073600480360381019061006e9190610438565b610123565b604051610080919061047d565b60405180910390f35b6100a3600480360381019061009e91906104b9565b61013f565b005b6100bf60048036038101906100ba9190610438565b6101db565b6040516100cc919061047d565b60405180910390f35b6100ef60048036038101906100ea9190610438565b6102c6565b6040516100fc919061047d565b60405180910390f35b61010d6103b2565b60405161011a919061050f565b60405180910390f35b5f602052805f5260405f205f915054906101000a900460ff1681565b5f80825f0160208101906101539190610438565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff166101d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101cf90610582565b60405180910390fd5b50565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461026a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610261906105ea565b60405180910390fd5b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060019050919050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610355576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034c906105ea565b60405180910390fd5b60015f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060019050919050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610407826103de565b9050919050565b610417816103fd565b8114610421575f80fd5b50565b5f813590506104328161040e565b92915050565b5f6020828403121561044d5761044c6103d6565b5b5f61045a84828501610424565b91505092915050565b5f8115159050919050565b61047781610463565b82525050565b5f6020820190506104905f83018461046e565b92915050565b5f80fd5b5f61016082840312156104b0576104af610496565b5b81905092915050565b5f602082840312156104ce576104cd6103d6565b5b5f82013567ffffffffffffffff8111156104eb576104ea6103da565b5b6104f78482850161049a565b91505092915050565b610509816103fd565b82525050565b5f6020820190506105225f830184610500565b92915050565b5f82825260208201905092915050565b7f55736572206e6f7420617574686f72697a6564000000000000000000000000005f82015250565b5f61056c601383610528565b915061057782610538565b602082019050919050565b5f6020820190508181035f83015261059981610560565b9050919050565b7f4f6e6c79206d61737465722077616c6c65742063616e2063616c6c00000000005f82015250565b5f6105d4601b83610528565b91506105df826105a0565b602082019050919050565b5f6020820190508181035f830152610601816105c8565b905091905056fea26469706673582212207273c2b6eed9d9367c4b7cbf01675aa9a87b5981059a126bfedc73fdc749399d64736f6c63430008150033",
}

// AccountManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountManagerMetaData.ABI instead.
var AccountManagerABI = AccountManagerMetaData.ABI

// AccountManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountManagerMetaData.Bin instead.
var AccountManagerBin = AccountManagerMetaData.Bin

// DeployAccountManager deploys a new Ethereum contract, binding an instance of AccountManager to it.
func DeployAccountManager(auth *bind.TransactOpts, backend bind.ContractBackend, _masterWallet common.Address) (common.Address, *types.Transaction, *AccountManager, error) {
	parsed, err := AccountManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountManagerBin), backend, _masterWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// AccountManager is an auto generated Go binding around an Ethereum contract.
type AccountManager struct {
	AccountManagerCaller     // Read-only binding to the contract
	AccountManagerTransactor // Write-only binding to the contract
	AccountManagerFilterer   // Log filterer for contract events
}

// AccountManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerSession struct {
	Contract     *AccountManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerCallerSession struct {
	Contract *AccountManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerTransactorSession struct {
	Contract     *AccountManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerRaw struct {
	Contract *AccountManager // Generic contract binding to access the raw methods on
}

// AccountManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerCallerRaw struct {
	Contract *AccountManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerTransactorRaw struct {
	Contract *AccountManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManager creates a new instance of AccountManager, bound to a specific deployed contract.
func NewAccountManager(address common.Address, backend bind.ContractBackend) (*AccountManager, error) {
	contract, err := bindAccountManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// NewAccountManagerCaller creates a new read-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerCaller, error) {
	contract, err := bindAccountManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCaller{contract: contract}, nil
}

// NewAccountManagerTransactor creates a new write-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerTransactor, error) {
	contract, err := bindAccountManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerTransactor{contract: contract}, nil
}

// NewAccountManagerFilterer creates a new log filterer instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerFilterer, error) {
	contract, err := bindAccountManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFilterer{contract: contract}, nil
}

// bindAccountManager binds a generic wrapper to an already deployed contract.
func bindAccountManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.AccountManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transact(opts, method, params...)
}

// AuthorizeUserOp is a free data retrieval call binding the contract method 0x9021ce1e.
//
// Solidity: function authorizeUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns()
func (_AccountManager *AccountManagerCaller) AuthorizeUserOp(opts *bind.CallOpts, userOp UserOperation) error {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "authorizeUserOp", userOp)

	if err != nil {
		return err
	}

	return err

}

// AuthorizeUserOp is a free data retrieval call binding the contract method 0x9021ce1e.
//
// Solidity: function authorizeUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns()
func (_AccountManager *AccountManagerSession) AuthorizeUserOp(userOp UserOperation) error {
	return _AccountManager.Contract.AuthorizeUserOp(&_AccountManager.CallOpts, userOp)
}

// AuthorizeUserOp is a free data retrieval call binding the contract method 0x9021ce1e.
//
// Solidity: function authorizeUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns()
func (_AccountManager *AccountManagerCallerSession) AuthorizeUserOp(userOp UserOperation) error {
	return _AccountManager.Contract.AuthorizeUserOp(&_AccountManager.CallOpts, userOp)
}

// AuthorizedUsers is a free data retrieval call binding the contract method 0x1828983a.
//
// Solidity: function authorizedUsers(address ) view returns(bool)
func (_AccountManager *AccountManagerCaller) AuthorizedUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "authorizedUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedUsers is a free data retrieval call binding the contract method 0x1828983a.
//
// Solidity: function authorizedUsers(address ) view returns(bool)
func (_AccountManager *AccountManagerSession) AuthorizedUsers(arg0 common.Address) (bool, error) {
	return _AccountManager.Contract.AuthorizedUsers(&_AccountManager.CallOpts, arg0)
}

// AuthorizedUsers is a free data retrieval call binding the contract method 0x1828983a.
//
// Solidity: function authorizedUsers(address ) view returns(bool)
func (_AccountManager *AccountManagerCallerSession) AuthorizedUsers(arg0 common.Address) (bool, error) {
	return _AccountManager.Contract.AuthorizedUsers(&_AccountManager.CallOpts, arg0)
}

// MasterWallet is a free data retrieval call binding the contract method 0xfc0d0117.
//
// Solidity: function masterWallet() view returns(address)
func (_AccountManager *AccountManagerCaller) MasterWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "masterWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterWallet is a free data retrieval call binding the contract method 0xfc0d0117.
//
// Solidity: function masterWallet() view returns(address)
func (_AccountManager *AccountManagerSession) MasterWallet() (common.Address, error) {
	return _AccountManager.Contract.MasterWallet(&_AccountManager.CallOpts)
}

// MasterWallet is a free data retrieval call binding the contract method 0xfc0d0117.
//
// Solidity: function masterWallet() view returns(address)
func (_AccountManager *AccountManagerCallerSession) MasterWallet() (common.Address, error) {
	return _AccountManager.Contract.MasterWallet(&_AccountManager.CallOpts)
}

// AddUserAuthorization is a paid mutator transaction binding the contract method 0xe461f00d.
//
// Solidity: function addUserAuthorization(address user) returns(bool)
func (_AccountManager *AccountManagerTransactor) AddUserAuthorization(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "addUserAuthorization", user)
}

// AddUserAuthorization is a paid mutator transaction binding the contract method 0xe461f00d.
//
// Solidity: function addUserAuthorization(address user) returns(bool)
func (_AccountManager *AccountManagerSession) AddUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AddUserAuthorization(&_AccountManager.TransactOpts, user)
}

// AddUserAuthorization is a paid mutator transaction binding the contract method 0xe461f00d.
//
// Solidity: function addUserAuthorization(address user) returns(bool)
func (_AccountManager *AccountManagerTransactorSession) AddUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AddUserAuthorization(&_AccountManager.TransactOpts, user)
}

// RemoveUserAuthorization is a paid mutator transaction binding the contract method 0xcfef9c97.
//
// Solidity: function removeUserAuthorization(address user) returns(bool)
func (_AccountManager *AccountManagerTransactor) RemoveUserAuthorization(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "removeUserAuthorization", user)
}

// RemoveUserAuthorization is a paid mutator transaction binding the contract method 0xcfef9c97.
//
// Solidity: function removeUserAuthorization(address user) returns(bool)
func (_AccountManager *AccountManagerSession) RemoveUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.RemoveUserAuthorization(&_AccountManager.TransactOpts, user)
}

// RemoveUserAuthorization is a paid mutator transaction binding the contract method 0xcfef9c97.
//
// Solidity: function removeUserAuthorization(address user) returns(bool)
func (_AccountManager *AccountManagerTransactorSession) RemoveUserAuthorization(user common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.RemoveUserAuthorization(&_AccountManager.TransactOpts, user)
}
