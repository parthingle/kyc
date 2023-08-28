package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/parthingle/kyc/gen"
)

const (
	entrypointContract = "0xd03807FE0a4A2D05ce80D8B2fa85cB317A01eaCd"

	masterWalletAddr    = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	masterWalletPrivKey = "<PRIVATE KEY>"

	mmAccount3Addr    = "0xd85fadc9936a0069ad68bd322e464ead34d53583"
	mmAccount3PrivKey = "<PRIVATE KEY>"

	accountManagerAddr = "0x327EAE52d61C5105dE7e521fDEe50AfcE12626B1"

	simpleAccountFactoryAddr = "0x46eE8aCED7540D78741cf68d02004C6A6b862d77"

	userEOA        = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
	userEOAPrivKey = "<PRIVATE KEY>"

	permissionedAccountFactoryAddr = "0x7D8c94C6E6b12bf7b7351AF08BC3367d2e3f8C4b"
)

func main() {

	// Geth node must be running with RPC enabled
	client, err := ethclient.Dial("http://127.0.0.1:8545/")

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
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"), nil)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get balance: %v", err))
	}

	fmt.Println("Balance is:", balance)

	// // 1. Check bytecode
	bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(entrypointContract), nil)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get bytecode: %v", err))
	} else {
		fmt.Println("Bytecode is:", bytecode)
	}

	mgr, err := gen.NewAccountManager(common.HexToAddress(accountManagerAddr), client)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to instantiate account manager contract: %v", err))
	}
	fmt.Println("Account manager contract is:", mgr.AccountManagerCaller)

	// res, err := mgr.MasterWallet(&bind.CallOpts{
	// 	From: common.HexToAddress(masterWalletAddr),
	// })
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to get master wallet: %v", err))
	// }
	// fmt.Println("Master wallet is:", res.Hex())

	// // Deploy AccountManager contract
	// privateKey, err := crypto.HexToECDSA(masterWalletPrivKey)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to parse private key: %v", err))
	// }

	// ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	// masterWallet, err := ks.ImportECDSA(privateKey, "")
	// if err != nil {
	// 	if err != keystore.ErrAccountAlreadyExists {
	// 		log.Fatal(fmt.Errorf("failed to import private key: %v", err))
	// 	}
	// }

	// err = ks.Unlock(masterWallet, "")
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to unlock account: %v", err))
	// }

	// auth, err := bind.NewKeyStoreTransactorWithChainID(ks, masterWallet, networkID)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to create authorized transactor: %v", err))
	// }
	// nonce, err := client.PendingNonceAt(context.Background(), masterWallet.Address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)
	// auth.GasPrice = big.NewInt(1000000000)
	// auth.GasLimit = uint64(3000000)

	// fmt.Printf("auth: %v\n\n", auth.Signer)

	// tx, err := mgr.AddUserAuthorization(auth, common.HexToAddress(userEOA))
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to add user authorization: %v", err))
	// }

	// fmt.Println("Transaction hash: ", tx.Hash().Hex())

	// nonce, err = client.PendingNonceAt(context.Background(), masterWallet.Address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// auth.Nonce = big.NewInt(int64(nonce))

	// isUserAuthorized, err := mgr.AuthorizedUsers(nil, common.HexToAddress(userEOA))
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to get user authorization: %v", err))
	// }

	// if isUserAuthorized {
	// 	fmt.Println("User is authorized")
	// } else {
	// 	fmt.Println("User is not authorized")
	// }

	// tx, err = mgr.RemoveUserAuthorization(auth, common.HexToAddress(userEOA))
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to add user authorization: %v", err))
	// }

	// fmt.Println("Transaction hash: ", tx.Hash().Hex())

	isUserAuthorized, err := mgr.AuthorizedUsers(nil, common.HexToAddress(userEOA))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get user authorization: %v", err))
	}

	if isUserAuthorized {
		fmt.Println("User is authorized")
	} else {
		fmt.Println("User is not authorized")
	}
}

func deployPermissionedAccountFactoryContract(client *ethclient.Client, networkID *big.Int) (common.Address, *types.Transaction, *gen.PermissionedAccountFactory, error) {

	// // 1. Check bytecode
	// bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(entrypointContract), nil)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to get bytecode: %v", err))
	// } else {
	// 	fmt.Println("Bytecode is:", bytecode)
	// }

	entryPoint, err := gen.NewEntryPoint(common.HexToAddress(entrypointContract), client)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to instantiate entry point contract: %v", err))
	}

	fmt.Println("Entry point contract is:", entryPoint)

	// Deploy PermissionedAccountFactory contract
	privateKey, err := crypto.HexToECDSA(mmAccount3PrivKey)
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
	auth.GasLimit = uint64(3000000)

	fmt.Printf("auth: %v\n\n", auth.Signer)

	permissionedAccountFactoryContractAddress, tx, permissionedAccountFactoryContract, err := gen.DeployPermissionedAccountFactory(auth, client, common.HexToAddress(entrypointContract), common.HexToAddress(mmAccount3Addr))
	if err != nil {
		log.Fatalf("Failed to deploy new Permissioned Account Factory contract: %v", err)
	}

	fmt.Println("Permissioned Account Factory contract address: ", permissionedAccountFactoryContractAddress.Hex())
	fmt.Println("Contract: ", permissionedAccountFactoryContract)
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	return permissionedAccountFactoryContractAddress, tx, permissionedAccountFactoryContract, nil
}

func deployAccountManagerContract(client *ethclient.Client, networkID *big.Int) (common.Address, *types.Transaction, *gen.AccountManager, error) {

	// // 1. Check bytecode
	// bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(entrypointContract), nil)
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("failed to get bytecode: %v", err))
	// } else {
	// 	fmt.Println("Bytecode is:", bytecode)
	// }

	entryPoint, err := gen.NewEntryPoint(common.HexToAddress(entrypointContract), client)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to instantiate entry point contract: %v", err))
	}
	bytecode, err := client.CodeAt(context.Background(), common.HexToAddress(permissionedAccountFactoryAddr), nil)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to get bytecode: %v", err))
	} else {
		fmt.Println("Bytecode is:", bytecode)
	}

	fmt.Println("Entry point contract is:", entryPoint)

	// Deploy AccountManager contract
	privateKey, err := crypto.HexToECDSA(masterWalletPrivKey)
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
	auth.GasLimit = uint64(3000000)

	fmt.Printf("auth: %v\n\n", auth.Signer)

	accountManagerContractAddress, tx, accountManagerContract, err := gen.DeployAccountManager(auth, client, common.HexToAddress(mmAccount3Addr))
	if err != nil {
		log.Fatalf("Failed to deploy new account manager contract: %v", err)
	}

	fmt.Println("Account manager contract address: ", accountManagerContractAddress.Hex())
	fmt.Println("Contract: ", accountManagerContract)
	fmt.Println("Transaction hash: ", tx.Hash().Hex())

	return accountManagerContractAddress, tx, accountManagerContract, nil
}
