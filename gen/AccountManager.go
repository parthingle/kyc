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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_masterWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"masterWallet\",\"type\":\"address\"}],\"name\":\"AccountManagerInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"authorizeUserOp\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"anOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeUserAuthorization\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162000fa338038062000fa38339818101604052810190620000369190620001b8565b80600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000876200008e60201b60201c565b50620002c4565b60018054906101000a900460ff1615620000df576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000d6906200026c565b60405180910390fd5b60ff801660015f9054906101000a900460ff1660ff1614620001515760ff60015f6101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620001489190620002a9565b60405180910390a15b565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001828262000157565b9050919050565b620001948162000176565b81146200019f575f80fd5b50565b5f81519050620001b28162000189565b92915050565b5f60208284031215620001d057620001cf62000153565b5b5f620001df84828501620001a2565b91505092915050565b5f82825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e6974695f8201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b5f62000254602783620001e8565b91506200026182620001f8565b604082019050919050565b5f6020820190508181035f830152620002858162000246565b9050919050565b5f60ff82169050919050565b620002a3816200028c565b82525050565b5f602082019050620002be5f83018462000298565b92915050565b610cd180620002d25f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80639021ce1e116100645780639021ce1e14610141578063c4d66de81461015d578063cfef9c9714610179578063e461f00d146101a9578063fc0d0117146101d957610091565b80630bd28e3b146100955780631828983a146100b15780631b2e01b8146100e157806335567e1a14610111575b5f80fd5b6100af60048036038101906100aa919061086e565b6101f7565b005b6100cb60048036038101906100c691906108f3565b61028e565b6040516100d89190610938565b60405180910390f35b6100fb60048036038101906100f69190610951565b6102ab565b60405161010891906109a7565b60405180910390f35b61012b60048036038101906101269190610951565b6102ca565b60405161013891906109a7565b60405180910390f35b61015b600480360381019061015691906109e3565b610373565b005b610177600480360381019061017291906108f3565b610410565b005b610193600480360381019061018e91906108f3565b61054b565b6040516101a09190610938565b60405180910390f35b6101c360048036038101906101be91906108f3565b610639565b6040516101d09190610938565b60405180910390f35b6101e1610728565b6040516101ee9190610a39565b60405180910390f35b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8277ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f81548092919061028690610a7f565b919050555050565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f602052815f5260405f20602052805f5260405f205f91509150505481565b5f60408277ffffffffffffffffffffffffffffffffffffffffffffffff16901b5f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8477ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205417905092915050565b60025f825f01602081019061038891906108f3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661040d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040490610b20565b60405180910390fd5b50565b5f60018054906101000a900460ff1615905080801561043f57506001805f9054906101000a900460ff1660ff16105b8061046c575061044e3061074e565b15801561046b57506001805f9054906101000a900460ff1660ff16145b5b6104ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a290610bae565b60405180910390fd5b6001805f6101000a81548160ff021916908360ff16021790555080156104e65760018060016101000a81548160ff0219169083151502179055505b6104ef82610770565b8015610547575f6001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161053e9190610c1a565b60405180910390a15b5050565b5f600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d390610c7d565b60405180910390fd5b5f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060019050919050565b5f600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c190610c7d565b60405180910390fd5b600160025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060019050919050565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b80600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f2cf2b0218b8dadc614f3075ff795e104c70361f38eaf7694cf580f354d4fce4560405160405180910390a250565b5f80fd5b5f80fd5b5f77ffffffffffffffffffffffffffffffffffffffffffffffff82169050919050565b61084d81610821565b8114610857575f80fd5b50565b5f8135905061086881610844565b92915050565b5f6020828403121561088357610882610819565b5b5f6108908482850161085a565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108c282610899565b9050919050565b6108d2816108b8565b81146108dc575f80fd5b50565b5f813590506108ed816108c9565b92915050565b5f6020828403121561090857610907610819565b5b5f610915848285016108df565b91505092915050565b5f8115159050919050565b6109328161091e565b82525050565b5f60208201905061094b5f830184610929565b92915050565b5f806040838503121561096757610966610819565b5b5f610974858286016108df565b92505060206109858582860161085a565b9150509250929050565b5f819050919050565b6109a18161098f565b82525050565b5f6020820190506109ba5f830184610998565b92915050565b5f80fd5b5f61016082840312156109da576109d96109c0565b5b81905092915050565b5f602082840312156109f8576109f7610819565b5b5f82013567ffffffffffffffff811115610a1557610a1461081d565b5b610a21848285016109c4565b91505092915050565b610a33816108b8565b82525050565b5f602082019050610a4c5f830184610a2a565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610a898261098f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610abb57610aba610a52565b5b600182019050919050565b5f82825260208201905092915050565b7f55736572206e6f7420617574686f72697a6564000000000000000000000000005f82015250565b5f610b0a601383610ac6565b9150610b1582610ad6565b602082019050919050565b5f6020820190508181035f830152610b3781610afe565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f610b98602e83610ac6565b9150610ba382610b3e565b604082019050919050565b5f6020820190508181035f830152610bc581610b8c565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f610c04610bff610bfa84610bcc565b610be1565b610bd5565b9050919050565b610c1481610bea565b82525050565b5f602082019050610c2d5f830184610c0b565b92915050565b7f4f6e6c79206d61737465722077616c6c65742063616e2063616c6c00000000005f82015250565b5f610c67601b83610ac6565b9150610c7282610c33565b602082019050919050565b5f6020820190508181035f830152610c9481610c5b565b905091905056fea2646970667358221220db7fe0040f8e054bd94d2134272cb85a18d3532a2ebecd7310eb00ba4c4d210e64736f6c63430008150033",
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
// Solidity: function initialize(address anOwner) returns()
func (_AccountManager *AccountManagerTransactor) Initialize(opts *bind.TransactOpts, anOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "initialize", anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_AccountManager *AccountManagerSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts, anOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address anOwner) returns()
func (_AccountManager *AccountManagerTransactorSession) Initialize(anOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Initialize(&_AccountManager.TransactOpts, anOwner)
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
