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

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// AccountManagerMetaData contains all meta data concerning the AccountManager contract.
var AccountManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_masterWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterWallet\",\"type\":\"address\"}],\"name\":\"AccountManagerInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"authorizeUserOp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_masterWallet\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001011380380620010118339818101604052810190620000379190620001c0565b80600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000886200008f60201b60201c565b50620002d6565b60018054906101000a900460ff1615620000e0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000d79062000279565b60405180910390fd5b60ff8016600160009054906101000a900460ff1660ff1614620001545760ff600160006101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200014b9190620002b9565b60405180910390a15b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000188826200015b565b9050919050565b6200019a816200017b565b8114620001a657600080fd5b50565b600081519050620001ba816200018f565b92915050565b600060208284031215620001d957620001d862000156565b5b6000620001e984828501620001a9565b91505092915050565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b600062000261602783620001f2565b91506200026e8262000203565b604082019050919050565b60006020820190508181036000830152620002948162000252565b9050919050565b600060ff82169050919050565b620002b3816200029b565b82525050565b6000602082019050620002d06000830184620002a8565b92915050565b610d2b80620002e66000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80639021ce1e116100665780639021ce1e14610144578063c4d66de814610160578063cfef9c971461017c578063e461f00d146101ac578063fc0d0117146101dc57610093565b80630bd28e3b146100985780631828983a146100b45780631b2e01b8146100e457806335567e1a14610114575b600080fd5b6100b260048036038101906100ad919061089c565b6101fa565b005b6100ce60048036038101906100c99190610927565b610296565b6040516100db919061096f565b60405180910390f35b6100fe60048036038101906100f9919061098a565b6102b6565b60405161010b91906109e3565b60405180910390f35b61012e6004803603810190610129919061098a565b6102db565b60405161013b91906109e3565b60405180910390f35b61015e60048036038101906101599190610a23565b610389565b005b61017a60048036038101906101759190610927565b61042a565b005b61019660048036038101906101919190610927565b61056a565b6040516101a3919061096f565b60405180910390f35b6101c660048036038101906101c19190610927565b61065d565b6040516101d3919061096f565b60405180910390f35b6101e4610750565b6040516101f19190610a7b565b60405180910390f35b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008277ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600081548092919061028e90610ac5565b919050555050565b60026020528060005260406000206000915054906101000a900460ff1681565b6000602052816000526040600020602052806000526040600020600091509150505481565b600060408277ffffffffffffffffffffffffffffffffffffffffffffffff16901b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008477ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205417905092915050565b600260008260000160208101906103a09190610927565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610427576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041e90610b6a565b60405180910390fd5b50565b600060018054906101000a900460ff1615905080801561045b575060018060009054906101000a900460ff1660ff16105b80610489575061046a30610776565b158015610488575060018060009054906101000a900460ff1660ff16145b5b6104c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bf90610bfc565b60405180910390fd5b60018060006101000a81548160ff021916908360ff16021790555080156105045760018060016101000a81548160ff0219169083151502179055505b61050d82610799565b80156105665760006001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161055d9190610c6e565b60405180910390a15b5050565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f390610cd5565b60405180910390fd5b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b6000600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e690610cd5565b60405180910390fd5b6001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b80600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f2cf2b0218b8dadc614f3075ff795e104c70361f38eaf7694cf580f354d4fce4560405160405180910390a250565b600080fd5b600080fd5b600077ffffffffffffffffffffffffffffffffffffffffffffffff82169050919050565b6108798161084c565b811461088457600080fd5b50565b60008135905061089681610870565b92915050565b6000602082840312156108b2576108b1610842565b5b60006108c084828501610887565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108f4826108c9565b9050919050565b610904816108e9565b811461090f57600080fd5b50565b600081359050610921816108fb565b92915050565b60006020828403121561093d5761093c610842565b5b600061094b84828501610912565b91505092915050565b60008115159050919050565b61096981610954565b82525050565b60006020820190506109846000830184610960565b92915050565b600080604083850312156109a1576109a0610842565b5b60006109af85828601610912565b92505060206109c085828601610887565b9150509250929050565b6000819050919050565b6109dd816109ca565b82525050565b60006020820190506109f860008301846109d4565b92915050565b600080fd5b60006101608284031215610a1a57610a196109fe565b5b81905092915050565b600060208284031215610a3957610a38610842565b5b600082013567ffffffffffffffff811115610a5757610a56610847565b5b610a6384828501610a03565b91505092915050565b610a75816108e9565b82525050565b6000602082019050610a906000830184610a6c565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ad0826109ca565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b0257610b01610a96565b5b600182019050919050565b600082825260208201905092915050565b7f55736572206e6f7420617574686f72697a656400000000000000000000000000600082015250565b6000610b54601383610b0d565b9150610b5f82610b1e565b602082019050919050565b60006020820190508181036000830152610b8381610b47565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000610be6602e83610b0d565b9150610bf182610b8a565b604082019050919050565b60006020820190508181036000830152610c1581610bd9565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000610c58610c53610c4e84610c1c565b610c33565b610c26565b9050919050565b610c6881610c3d565b82525050565b6000602082019050610c836000830184610c5f565b92915050565b7f4f6e6c79206d61737465722077616c6c65742063616e2063616c6c0000000000600082015250565b6000610cbf601b83610b0d565b9150610cca82610c89565b602082019050919050565b60006020820190508181036000830152610cee81610cb2565b905091905056fea26469706673582212203fdab5327e04addf4e6cde6fd9752d9206fd45a93bd02847e6573d83e5662a5d64736f6c63430008130033",
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

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_AccountManager *AccountManagerCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_AccountManager *AccountManagerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _AccountManager.Contract.GetNonce(&_AccountManager.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_AccountManager *AccountManagerCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _AccountManager.Contract.GetNonce(&_AccountManager.CallOpts, sender, key)
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

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_AccountManager *AccountManagerCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AccountManager.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_AccountManager *AccountManagerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AccountManager.Contract.NonceSequenceNumber(&_AccountManager.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_AccountManager *AccountManagerCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AccountManager.Contract.NonceSequenceNumber(&_AccountManager.CallOpts, arg0, arg1)
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

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_AccountManager *AccountManagerTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_AccountManager *AccountManagerSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.IncrementNonce(&_AccountManager.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_AccountManager *AccountManagerTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _AccountManager.Contract.IncrementNonce(&_AccountManager.TransactOpts, key)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _masterWallet) returns()
func (_AccountManager *AccountManagerTransactor) Initialize(opts *bind.TransactOpts, _masterWallet common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "initialize", _masterWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _masterWallet) returns()
func (_AccountManager *AccountManagerSession) Initialize(_masterWallet common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts, _masterWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _masterWallet) returns()
func (_AccountManager *AccountManagerTransactorSession) Initialize(_masterWallet common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts, _masterWallet)
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

// AccountManagerAccountManagerInitializedIterator is returned from FilterAccountManagerInitialized and is used to iterate over the raw logs and unpacked data for AccountManagerInitialized events raised by the AccountManager contract.
type AccountManagerAccountManagerInitializedIterator struct {
	Event *AccountManagerAccountManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AccountManagerAccountManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAccountManagerInitialized)
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
		it.Event = new(AccountManagerAccountManagerInitialized)
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
func (it *AccountManagerAccountManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAccountManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAccountManagerInitialized represents a AccountManagerInitialized event raised by the AccountManager contract.
type AccountManagerAccountManagerInitialized struct {
	MasterWallet common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountManagerInitialized is a free log retrieval operation binding the contract event 0x2cf2b0218b8dadc614f3075ff795e104c70361f38eaf7694cf580f354d4fce45.
//
// Solidity: event AccountManagerInitialized(address indexed masterWallet)
func (_AccountManager *AccountManagerFilterer) FilterAccountManagerInitialized(opts *bind.FilterOpts, masterWallet []common.Address) (*AccountManagerAccountManagerInitializedIterator, error) {

	var masterWalletRule []interface{}
	for _, masterWalletItem := range masterWallet {
		masterWalletRule = append(masterWalletRule, masterWalletItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "AccountManagerInitialized", masterWalletRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerAccountManagerInitializedIterator{contract: _AccountManager.contract, event: "AccountManagerInitialized", logs: logs, sub: sub}, nil
}

// WatchAccountManagerInitialized is a free log subscription operation binding the contract event 0x2cf2b0218b8dadc614f3075ff795e104c70361f38eaf7694cf580f354d4fce45.
//
// Solidity: event AccountManagerInitialized(address indexed masterWallet)
func (_AccountManager *AccountManagerFilterer) WatchAccountManagerInitialized(opts *bind.WatchOpts, sink chan<- *AccountManagerAccountManagerInitialized, masterWallet []common.Address) (event.Subscription, error) {

	var masterWalletRule []interface{}
	for _, masterWalletItem := range masterWallet {
		masterWalletRule = append(masterWalletRule, masterWalletItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "AccountManagerInitialized", masterWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAccountManagerInitialized)
				if err := _AccountManager.contract.UnpackLog(event, "AccountManagerInitialized", log); err != nil {
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

// ParseAccountManagerInitialized is a log parse operation binding the contract event 0x2cf2b0218b8dadc614f3075ff795e104c70361f38eaf7694cf580f354d4fce45.
//
// Solidity: event AccountManagerInitialized(address indexed masterWallet)
func (_AccountManager *AccountManagerFilterer) ParseAccountManagerInitialized(log types.Log) (*AccountManagerAccountManagerInitialized, error) {
	event := new(AccountManagerAccountManagerInitialized)
	if err := _AccountManager.contract.UnpackLog(event, "AccountManagerInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AccountManager contract.
type AccountManagerInitializedIterator struct {
	Event *AccountManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AccountManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerInitialized)
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
		it.Event = new(AccountManagerInitialized)
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
func (it *AccountManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerInitialized represents a Initialized event raised by the AccountManager contract.
type AccountManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AccountManager *AccountManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountManagerInitializedIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountManagerInitializedIterator{contract: _AccountManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AccountManager *AccountManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerInitialized)
				if err := _AccountManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AccountManager *AccountManagerFilterer) ParseInitialized(log types.Log) (*AccountManagerInitialized, error) {
	event := new(AccountManagerInitialized)
	if err := _AccountManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
