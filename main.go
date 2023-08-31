package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/parthingle/kyc/gen"
)

const (
	EntryPointContractAddress = "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"
	MasterWalletPrivKey       = "<Private Key>"

	devAccountManagerContractAddress       = "0x327EAE52d61C5105dE7e521fDEe50AfcE12626B1"
	devEntrypointContractAddress           = "0xd03807FE0a4A2D05ce80D8B2fa85cB317A01eaCd"
	devSimpleAccountFactoryContractAddress = "0x46eE8aCED7540D78741cf68d02004C6A6b862d77"

	devSampleWallet        = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	devSampleWalletPrivKey = "<Private Key>"

	goerliPermissionedAccountFactory = "0x638CB7b533884700fc9a04Dd765F8309F2390f00"
	goerliAccountManager             = "0x746F1967c9640aBa29Cd0dC587E7499F8A89C80e"
	goerliCounter                    = "0x422D3218512D17B4c35159072924d1B31f87d3A4"

	local      = "http://127.0.0.1:8545"
	baseGoerli = "https://goerli.base.org"
)

func main() {
	// Geth node must be running with RPC enabled
	client, err := ethclient.Dial(baseGoerli)

	if err != nil {
		log.Fatal(err)
	}

	// Print network ID to confirm connection
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to network:", networkID)

	// Can now use client to interact with node
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(devSampleWallet), nil)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get balance: %v", err))
	}

	fmt.Println("Balance is:", balance)
	// deployTestCounter(client, networkID, MasterWalletPrivKey)
	// deployAccountManagerContract(client, networkID, MasterWalletPrivKey)
	deployPermissionedAccountFactoryContract(client, networkID, MasterWalletPrivKey)
	// getBytecodeAt(client, common.HexToAddress(goerliAccountFactory))

	// counter, err := gen.NewTestCounter(common.HexToAddress(goerliCounter), client)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to instantiate counter contract: %v", err))
	// }

	// fmt.Println("Counter contract is:", counter)

	// ks, mmAccount3Wallet, err := getKey(MasterWalletPrivKey)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to parse private key: %v", err))
	// }

	// auth := getAuth(ks, mmAccount3Wallet, client, networkID)

	// tx, err := counter.IncrementCounter(auth)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to increment counter: %v", err))
	// }

	// fmt.Println("Transaction hash: ", tx.Hash().Hex())

	// // Wait for transaction to be mined
	// _, err = bind.WaitMined(context.Background(), client, tx)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to wait for transaction to be mined: %v", err))
	// }

	// fmt.Println("Counter incremented")

	// // Get counter value

	// value, err := counter.GetCount(nil)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to get counter value: %v", err))
	// }

	// fmt.Println("Counter value is:", value)
}

func getBytecodeAt(client *ethclient.Client, add common.Address) {

	// 1. Check bytecode
	bytecode, err := client.CodeAt(context.Background(), add, nil)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get bytecode: %v", err))
	} else {
		fmt.Println("Bytecode is:", bytecode)
	}

}

func getKey(privateKeyString string) (*keystore.KeyStore, accounts.Account, error) {
	// Deploy AccountManager contract
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse private key: %v", err))
	}

	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	mmAccount3Wallet, err := ks.ImportECDSA(privateKey, "")
	if err != nil {
		if err != keystore.ErrAccountAlreadyExists {
			log.Fatal(fmt.Errorf("failed to import private key: %v", err))
		}
	}

	err = ks.Unlock(mmAccount3Wallet, "")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to unlock account: %v", err))
	}

	return ks, mmAccount3Wallet, nil
}

func deployPermissionedAccountFactoryContract(client *ethclient.Client, networkID *big.Int, masterWalletPriv string) (common.Address, *types.Transaction, *gen.PermissionedAccountFactory, error) {

	// Deploy AccountManager contract
	ks, mmAccount3Wallet, err := getKey(masterWalletPriv)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse private key: %v", err))
	}

	auth := getAuth(ks, mmAccount3Wallet, client, networkID)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create authorized transactor: %v", err))
	}
	auth.GasLimit = uint64(10000000)
	permissionedAccountFactoryContractAddress, tx, permissionedAccountFactoryContract, err :=
		gen.DeployPermissionedAccountFactory(auth, client, common.HexToAddress(EntryPointContractAddress), mmAccount3Wallet.Address)
	if err != nil {
		log.Fatalf("Failed to deploy new simple account factory contract: %v", err)
	}
	fmt.Println("Permissioned Account Factory contract address: ", permissionedAccountFactoryContractAddress.Hex())
	fmt.Println("Contract: ", permissionedAccountFactoryContract)
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	return permissionedAccountFactoryContractAddress, tx, nil, nil
}

func getAuth(ks *keystore.KeyStore, wallet accounts.Account, client *ethclient.Client, networkID *big.Int) *bind.TransactOpts {
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, wallet, networkID)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create authorized transactor: %v", err))
	}

	nonce, err := client.PendingNonceAt(context.Background(), wallet.Address)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(1000000000)
	auth.GasLimit = uint64(10000000)

	return auth
}

func deployTestCounter(client *ethclient.Client, networkID *big.Int, masterWalletPriv string) (common.Address, *types.Transaction, *gen.AccountManager, error) {
	// Deploy AccountManager contract
	ks, mmAccount3Wallet, err := getKey(masterWalletPriv)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse private key: %v", err))
	}

	auth := getAuth(ks, mmAccount3Wallet, client, networkID)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create authorized transactor: %v", err))
	}

	testCounterFactoryContractAddress, tx, testCounterFactoryContract, err := gen.DeployTestCounter(auth, client)
	if err != nil {
		log.Fatalf("Failed to deploy new simple account factory contract: %v", err)
	}
	fmt.Println("Test Counter contract address: ", testCounterFactoryContractAddress.Hex())
	fmt.Println("Contract: ", testCounterFactoryContract)
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	return testCounterFactoryContractAddress, tx, nil, nil

}

func deployAccountManagerContract(client *ethclient.Client, networkID *big.Int, masterWalletPriv string) (common.Address, *types.Transaction, *gen.AccountManager, error) {

	// // 1. Check bytecode
	// bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(entrypointContract), nil)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to get bytecode: %v", err))
	// } else {
	// 	fmt.Println("Bytecode is:", bytecode)
	// }

	entryPoint, err := gen.NewEntryPoint(common.HexToAddress(EntryPointContractAddress), client)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to instantiate entry point contract: %v", err))
	}

	fmt.Println("Entry point contract is:", entryPoint)

	// Deploy AccountManager contract
	privateKey, err := crypto.HexToECDSA(masterWalletPriv)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse private key: %v", err))
	}

	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	mmAccount3Wallet, err := ks.ImportECDSA(privateKey, "")
	if err != nil {
		if err != keystore.ErrAccountAlreadyExists {
			log.Fatal(fmt.Errorf("failed to import private key: %v", err))
		}
	}

	err = ks.Unlock(mmAccount3Wallet, "")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to unlock account: %v", err))
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, mmAccount3Wallet, networkID)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to create authorized transactor: %v", err))
	}

	nonce, err := client.PendingNonceAt(context.Background(), mmAccount3Wallet.Address)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(1000000000)
	auth.GasLimit = uint64(4000000)

	accountManagerFactoryContractAddress, tx, accountManagerFactoryContract, err := gen.DeployAccountManager(auth, client, common.HexToAddress(devEntrypointContractAddress))
	if err != nil {
		log.Fatalf("Failed to deploy new simple account factory contract: %v", err)
	}
	fmt.Println("Account manager contract address: ", accountManagerFactoryContractAddress.Hex())
	fmt.Println("Contract: ", accountManagerFactoryContract)
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	return accountManagerFactoryContractAddress, tx, nil, nil
	// accountManagerContractAddress, tx, accountManagerContract, err := gen.DeployAccountManager(auth, client, mmAccount3Wallet.Address)
	// if err != nil {
	// 	log.Fatalf("Failed to deploy new account manager contract: %v", err)
	// }

	// fmt.Println("Account manager contract address: ", accountManagerContractAddress.Hex())
	// fmt.Println("Contract: ", accountManagerContract)
	// fmt.Println("Transaction hash: ", tx.Hash().Hex())

	// return accountManagerContractAddress, tx, accountManagerContract, nil
}
