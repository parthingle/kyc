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

// AccountManagerFactoryMetaData contains all meta data concerning the AccountManagerFactory contract.
var AccountManagerFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_masterWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractAccountManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createManager\",\"outputs\":[{\"internalType\":\"contractAccountManager\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051611f94380380611f9483398181016040528101906100319190610106565b8060405161003e9061009b565b6100489190610140565b604051809103905ff080158015610061573d5f803e3d5ffd5b5073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050610159565b610fa380610ff183390190565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100d5826100ac565b9050919050565b6100e5816100cb565b81146100ef575f80fd5b50565b5f81519050610100816100dc565b92915050565b5f6020828403121561011b5761011a6100a8565b5b5f610128848285016100f2565b91505092915050565b61013a816100cb565b82525050565b5f6020820190506101535f830184610131565b92915050565b608051610e7361017e5f395f818160d80152818161013e015261022a0152610e735ff3fe608060405234801562000010575f80fd5b506004361062000044575f3560e01c806311464fbe14620000485780632e6d2944146200006a5780638cb84e1814620000a0575b5f80fd5b62000052620000d6565b604051620000619190620003ce565b60405180910390f35b62000088600480360381019062000082919062000467565b620000fa565b604051620000979190620003ce565b60405180910390f35b620000be6004803603810190620000b8919062000467565b620001fa565b604051620000cd9190620004bd565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b5f80620001088484620001fa565b90505f8173ffffffffffffffffffffffffffffffffffffffff163b90505f81111562000139578192505050620001f4565b835f1b7f000000000000000000000000000000000000000000000000000000000000000086604051602401620001709190620004bd565b60405160208183030381529060405263c4d66de860e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051620001c3906200033a565b620001d09291906200056c565b8190604051809103905ff5905080158015620001ee573d5f803e3d5ffd5b50925050505b92915050565b5f620002f2825f1b6040518060200162000214906200033a565b6020820181038252601f19601f820116604052507f0000000000000000000000000000000000000000000000000000000000000000866040516024016200025c9190620004bd565b60405160208183030381529060405263c4d66de860e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051602001620002b49291906200056c565b604051602081830303815290604052604051602001620002d6929190620005de565b60405160208183030381529060405280519060200120620002fa565b905092915050565b5f6200030883833062000310565b905092915050565b5f604051836040820152846020820152828152600b810160ff815360558120925050509392505050565b610838806200060683390190565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f620003906200038a620003848462000348565b62000367565b62000348565b9050919050565b5f620003a38262000370565b9050919050565b5f620003b68262000397565b9050919050565b620003c881620003aa565b82525050565b5f602082019050620003e35f830184620003bd565b92915050565b5f80fd5b5f620003f98262000348565b9050919050565b6200040b81620003ed565b811462000416575f80fd5b50565b5f81359050620004298162000400565b92915050565b5f819050919050565b62000443816200042f565b81146200044e575f80fd5b50565b5f81359050620004618162000438565b92915050565b5f806040838503121562000480576200047f620003e9565b5b5f6200048f8582860162000419565b9250506020620004a28582860162000451565b9150509250929050565b620004b781620003ed565b82525050565b5f602082019050620004d25f830184620004ac565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101562000511578082015181840152602081019050620004f4565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6200053882620004d8565b620005448185620004e2565b935062000556818560208601620004f2565b62000561816200051c565b840191505092915050565b5f604082019050620005815f830185620004ac565b81810360208301526200059581846200052c565b90509392505050565b5f81905092915050565b5f620005b482620004d8565b620005c081856200059e565b9350620005d2818560208601620004f2565b80840191505092915050565b5f620005eb8285620005a8565b9150620005f98284620005a8565b9150819050939250505056fe608060405260405161083838038061083883398181016040528101906100259190610501565b61003682825f61003d60201b60201c565b505061071d565b61004c8361007460201b60201c565b5f825111806100585750805b1561006f5761006d83836100c960201b60201c565b505b505050565b610083816100fc60201b60201c565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60606100f48383604051806060016040528060278152602001610811602791396101be60201b60201c565b905092915050565b61010b8161024660201b60201c565b61014a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610141906105db565b60405180910390fd5b8061017c7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b61026860201b60201c565b5f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60605f808573ffffffffffffffffffffffffffffffffffffffff16856040516101e7919061063d565b5f60405180830381855af49150503d805f811461021f576040519150601f19603f3d011682016040523d82523d5f602084013e610224565b606091505b509150915061023b8683838761027160201b60201c565b925050509392505050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f819050919050565b606083156102d8575f8351036102d0576102908561024660201b60201c565b6102cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c69061069d565b60405180910390fd5b5b8290506102e9565b6102e883836102f160201b60201c565b5b949350505050565b5f825111156103035781518083602001fd5b806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033791906106fd565b60405180910390fd5b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61037a82610351565b9050919050565b61038a81610370565b8114610394575f80fd5b50565b5f815190506103a581610381565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6103f9826103b3565b810181811067ffffffffffffffff82111715610418576104176103c3565b5b80604052505050565b5f61042a610340565b905061043682826103f0565b919050565b5f67ffffffffffffffff821115610455576104546103c3565b5b61045e826103b3565b9050602081019050919050565b5f5b8381101561048857808201518184015260208101905061046d565b5f8484015250505050565b5f6104a56104a08461043b565b610421565b9050828152602081018484840111156104c1576104c06103af565b5b6104cc84828561046b565b509392505050565b5f82601f8301126104e8576104e76103ab565b5b81516104f8848260208601610493565b91505092915050565b5f806040838503121561051757610516610349565b5b5f61052485828601610397565b925050602083015167ffffffffffffffff8111156105455761054461034d565b5b610551858286016104d4565b9150509250929050565b5f82825260208201905092915050565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e5f8201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b5f6105c5602d8361055b565b91506105d08261056b565b604082019050919050565b5f6020820190508181035f8301526105f2816105b9565b9050919050565b5f81519050919050565b5f81905092915050565b5f610617826105f9565b6106218185610603565b935061063181856020860161046b565b80840191505092915050565b5f610648828461060d565b915081905092915050565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000005f82015250565b5f610687601d8361055b565b915061069282610653565b602082019050919050565b5f6020820190508181035f8301526106b48161067b565b9050919050565b5f81519050919050565b5f6106cf826106bb565b6106d9818561055b565b93506106e981856020860161046b565b6106f2816103b3565b840191505092915050565b5f6020820190508181035f83015261071581846106c5565b905092915050565b60e8806107295f395ff3fe608060405236601057600e6018565b005b60166018565b005b601e602c565b602a6026602e565b603a565b565b565b5f60356058565b905090565b365f80375f80365f845af43d5f803e805f81146054573d5ff35b3d5ffd5b5f60827f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5f1b60a9565b5f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f81905091905056fea2646970667358221220c99191dec8618c1e178449667c193bed9abe299e7b2994aacd63e83f3e74692264736f6c63430008150033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212200b6451c5c7a8c60bccec845c2cd5d62dccb4ad47b72071a304d918c11269a09164736f6c63430008150033608060405234801562000010575f80fd5b5060405162000fa338038062000fa38339818101604052810190620000369190620001b8565b80600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000876200008e60201b60201c565b50620002c4565b60018054906101000a900460ff1615620000df576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000d6906200026c565b60405180910390fd5b60ff801660015f9054906101000a900460ff1660ff1614620001515760ff60015f6101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620001489190620002a9565b60405180910390a15b565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620001828262000157565b9050919050565b620001948162000176565b81146200019f575f80fd5b50565b5f81519050620001b28162000189565b92915050565b5f60208284031215620001d057620001cf62000153565b5b5f620001df84828501620001a2565b91505092915050565b5f82825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e6974695f8201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b5f62000254602783620001e8565b91506200026182620001f8565b604082019050919050565b5f6020820190508181035f830152620002858162000246565b9050919050565b5f60ff82169050919050565b620002a3816200028c565b82525050565b5f602082019050620002be5f83018462000298565b92915050565b610cd180620002d25f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80639021ce1e116100645780639021ce1e14610141578063c4d66de81461015d578063cfef9c9714610179578063e461f00d146101a9578063fc0d0117146101d957610091565b80630bd28e3b146100955780631828983a146100b15780631b2e01b8146100e157806335567e1a14610111575b5f80fd5b6100af60048036038101906100aa919061086e565b6101f7565b005b6100cb60048036038101906100c691906108f3565b61028e565b6040516100d89190610938565b60405180910390f35b6100fb60048036038101906100f69190610951565b6102ab565b60405161010891906109a7565b60405180910390f35b61012b60048036038101906101269190610951565b6102ca565b60405161013891906109a7565b60405180910390f35b61015b600480360381019061015691906109e3565b610373565b005b610177600480360381019061017291906108f3565b610410565b005b610193600480360381019061018e91906108f3565b61054b565b6040516101a09190610938565b60405180910390f35b6101c360048036038101906101be91906108f3565b610639565b6040516101d09190610938565b60405180910390f35b6101e1610728565b6040516101ee9190610a39565b60405180910390f35b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8277ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f81548092919061028690610a7f565b919050555050565b6002602052805f5260405f205f915054906101000a900460ff1681565b5f602052815f5260405f20602052805f5260405f205f91509150505481565b5f60408277ffffffffffffffffffffffffffffffffffffffffffffffff16901b5f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8477ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205417905092915050565b60025f825f01602081019061038891906108f3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1661040d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040490610b20565b60405180910390fd5b50565b5f60018054906101000a900460ff1615905080801561043f57506001805f9054906101000a900460ff1660ff16105b8061046c575061044e3061074e565b15801561046b57506001805f9054906101000a900460ff1660ff16145b5b6104ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a290610bae565b60405180910390fd5b6001805f6101000a81548160ff021916908360ff16021790555080156104e65760018060016101000a81548160ff0219169083151502179055505b6104ef82610770565b8015610547575f6001806101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161053e9190610c1a565b60405180910390a15b5050565b5f600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d390610c7d565b60405180910390fd5b5f60025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060019050919050565b5f600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c190610c7d565b60405180910390fd5b600160025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060019050919050565b600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b80600160026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f2cf2b0218b8dadc614f3075ff795e104c70361f38eaf7694cf580f354d4fce4560405160405180910390a250565b5f80fd5b5f80fd5b5f77ffffffffffffffffffffffffffffffffffffffffffffffff82169050919050565b61084d81610821565b8114610857575f80fd5b50565b5f8135905061086881610844565b92915050565b5f6020828403121561088357610882610819565b5b5f6108908482850161085a565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108c282610899565b9050919050565b6108d2816108b8565b81146108dc575f80fd5b50565b5f813590506108ed816108c9565b92915050565b5f6020828403121561090857610907610819565b5b5f610915848285016108df565b91505092915050565b5f8115159050919050565b6109328161091e565b82525050565b5f60208201905061094b5f830184610929565b92915050565b5f806040838503121561096757610966610819565b5b5f610974858286016108df565b92505060206109858582860161085a565b9150509250929050565b5f819050919050565b6109a18161098f565b82525050565b5f6020820190506109ba5f830184610998565b92915050565b5f80fd5b5f61016082840312156109da576109d96109c0565b5b81905092915050565b5f602082840312156109f8576109f7610819565b5b5f82013567ffffffffffffffff811115610a1557610a1461081d565b5b610a21848285016109c4565b91505092915050565b610a33816108b8565b82525050565b5f602082019050610a4c5f830184610a2a565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610a898261098f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610abb57610aba610a52565b5b600182019050919050565b5f82825260208201905092915050565b7f55736572206e6f7420617574686f72697a6564000000000000000000000000005f82015250565b5f610b0a601383610ac6565b9150610b1582610ad6565b602082019050919050565b5f6020820190508181035f830152610b3781610afe565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f610b98602e83610ac6565b9150610ba382610b3e565b604082019050919050565b5f6020820190508181035f830152610bc581610b8c565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f610c04610bff610bfa84610bcc565b610be1565b610bd5565b9050919050565b610c1481610bea565b82525050565b5f602082019050610c2d5f830184610c0b565b92915050565b7f4f6e6c79206d61737465722077616c6c65742063616e2063616c6c00000000005f82015250565b5f610c67601b83610ac6565b9150610c7282610c33565b602082019050919050565b5f6020820190508181035f830152610c9481610c5b565b905091905056fea264697066735822122097b1f674dbdb39a7485a30e1c7d59768af827ec6fde1478f835f07a4e8c96cd064736f6c63430008150033",
}

// AccountManagerFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountManagerFactoryMetaData.ABI instead.
var AccountManagerFactoryABI = AccountManagerFactoryMetaData.ABI

// AccountManagerFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccountManagerFactoryMetaData.Bin instead.
var AccountManagerFactoryBin = AccountManagerFactoryMetaData.Bin

// DeployAccountManagerFactory deploys a new Ethereum contract, binding an instance of AccountManagerFactory to it.
func DeployAccountManagerFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _masterWallet common.Address) (common.Address, *types.Transaction, *AccountManagerFactory, error) {
	parsed, err := AccountManagerFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccountManagerFactoryBin), backend, _masterWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountManagerFactory{AccountManagerFactoryCaller: AccountManagerFactoryCaller{contract: contract}, AccountManagerFactoryTransactor: AccountManagerFactoryTransactor{contract: contract}, AccountManagerFactoryFilterer: AccountManagerFactoryFilterer{contract: contract}}, nil
}

// AccountManagerFactory is an auto generated Go binding around an Ethereum contract.
type AccountManagerFactory struct {
	AccountManagerFactoryCaller     // Read-only binding to the contract
	AccountManagerFactoryTransactor // Write-only binding to the contract
	AccountManagerFactoryFilterer   // Log filterer for contract events
}

// AccountManagerFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerFactorySession struct {
	Contract     *AccountManagerFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AccountManagerFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerFactoryCallerSession struct {
	Contract *AccountManagerFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AccountManagerFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerFactoryTransactorSession struct {
	Contract     *AccountManagerFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AccountManagerFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerFactoryRaw struct {
	Contract *AccountManagerFactory // Generic contract binding to access the raw methods on
}

// AccountManagerFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerFactoryCallerRaw struct {
	Contract *AccountManagerFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerFactoryTransactorRaw struct {
	Contract *AccountManagerFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManagerFactory creates a new instance of AccountManagerFactory, bound to a specific deployed contract.
func NewAccountManagerFactory(address common.Address, backend bind.ContractBackend) (*AccountManagerFactory, error) {
	contract, err := bindAccountManagerFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFactory{AccountManagerFactoryCaller: AccountManagerFactoryCaller{contract: contract}, AccountManagerFactoryTransactor: AccountManagerFactoryTransactor{contract: contract}, AccountManagerFactoryFilterer: AccountManagerFactoryFilterer{contract: contract}}, nil
}

// NewAccountManagerFactoryCaller creates a new read-only instance of AccountManagerFactory, bound to a specific deployed contract.
func NewAccountManagerFactoryCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerFactoryCaller, error) {
	contract, err := bindAccountManagerFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFactoryCaller{contract: contract}, nil
}

// NewAccountManagerFactoryTransactor creates a new write-only instance of AccountManagerFactory, bound to a specific deployed contract.
func NewAccountManagerFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerFactoryTransactor, error) {
	contract, err := bindAccountManagerFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFactoryTransactor{contract: contract}, nil
}

// NewAccountManagerFactoryFilterer creates a new log filterer instance of AccountManagerFactory, bound to a specific deployed contract.
func NewAccountManagerFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerFactoryFilterer, error) {
	contract, err := bindAccountManagerFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFactoryFilterer{contract: contract}, nil
}

// bindAccountManagerFactory binds a generic wrapper to an already deployed contract.
func bindAccountManagerFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccountManagerFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerFactory *AccountManagerFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerFactory.Contract.AccountManagerFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerFactory *AccountManagerFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerFactory.Contract.AccountManagerFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerFactory *AccountManagerFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerFactory.Contract.AccountManagerFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManagerFactory *AccountManagerFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccountManagerFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManagerFactory *AccountManagerFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManagerFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManagerFactory *AccountManagerFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManagerFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_AccountManagerFactory *AccountManagerFactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerFactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_AccountManagerFactory *AccountManagerFactorySession) AccountImplementation() (common.Address, error) {
	return _AccountManagerFactory.Contract.AccountImplementation(&_AccountManagerFactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_AccountManagerFactory *AccountManagerFactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _AccountManagerFactory.Contract.AccountImplementation(&_AccountManagerFactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_AccountManagerFactory *AccountManagerFactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccountManagerFactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_AccountManagerFactory *AccountManagerFactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _AccountManagerFactory.Contract.GetAddress(&_AccountManagerFactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_AccountManagerFactory *AccountManagerFactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _AccountManagerFactory.Contract.GetAddress(&_AccountManagerFactory.CallOpts, owner, salt)
}

// CreateManager is a paid mutator transaction binding the contract method 0x2e6d2944.
//
// Solidity: function createManager(address owner, uint256 salt) returns(address ret)
func (_AccountManagerFactory *AccountManagerFactoryTransactor) CreateManager(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _AccountManagerFactory.contract.Transact(opts, "createManager", owner, salt)
}

// CreateManager is a paid mutator transaction binding the contract method 0x2e6d2944.
//
// Solidity: function createManager(address owner, uint256 salt) returns(address ret)
func (_AccountManagerFactory *AccountManagerFactorySession) CreateManager(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _AccountManagerFactory.Contract.CreateManager(&_AccountManagerFactory.TransactOpts, owner, salt)
}

// CreateManager is a paid mutator transaction binding the contract method 0x2e6d2944.
//
// Solidity: function createManager(address owner, uint256 salt) returns(address ret)
func (_AccountManagerFactory *AccountManagerFactoryTransactorSession) CreateManager(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _AccountManagerFactory.Contract.CreateManager(&_AccountManagerFactory.TransactOpts, owner, salt)
}
