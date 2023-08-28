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

// EntryPointMemoryUserOp is an auto generated low-level Go binding around an user-defined struct.
type EntryPointMemoryUserOp struct {
	Sender               common.Address
	Nonce                *big.Int
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	Paymaster            common.Address
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
}

// EntryPointUserOpInfo is an auto generated low-level Go binding around an user-defined struct.
type EntryPointUserOpInfo struct {
	MUserOp       EntryPointMemoryUserOp
	UserOpHash    [32]byte
	Prefund       *big.Int
	ContextOffset *big.Int
	PreOpGas      *big.Int
}

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

// EntryPointMetaData contains all meta data concerning the EntryPoint contract.
var EntryPointMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"targetSuccess\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"targetResult\",\"type\":\"bytes\"}],\"name\":\"ExecutionResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"FailedOp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderAddressResult\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureValidationFailed\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sigFailed\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"}],\"name\":\"ValidationResult\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sigFailed\",\"type\":\"bool\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"paymasterContext\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.ReturnInfo\",\"name\":\"returnInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"senderInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"factoryInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"paymasterInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.StakeInfo\",\"name\":\"stakeInfo\",\"type\":\"tuple\"}],\"internalType\":\"structIEntryPoint.AggregatorStakeInfo\",\"name\":\"aggregatorInfo\",\"type\":\"tuple\"}],\"name\":\"ValidationResultWithAggregation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"}],\"name\":\"AccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BeforeExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposit\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"SignatureAggregatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeDelaySec\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawTime\",\"type\":\"uint256\"}],\"name\":\"StakeUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualGasUsed\",\"type\":\"uint256\"}],\"name\":\"UserOperationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertReason\",\"type\":\"bytes\"}],\"name\":\"UserOperationRevertReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SIG_VALIDATION_FAILED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"}],\"name\":\"_validateSenderAndPaymaster\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"deposit\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"deposit\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"staked\",\"type\":\"bool\"},{\"internalType\":\"uint112\",\"name\":\"stake\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"},{\"internalType\":\"uint48\",\"name\":\"withdrawTime\",\"type\":\"uint48\"}],\"internalType\":\"structIStakeManager.DepositInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"}],\"name\":\"getSenderAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOpHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"userOps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIAggregator\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIEntryPoint.UserOpsPerAggregator[]\",\"name\":\"opsPerAggregator\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleAggregatedOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation[]\",\"name\":\"ops\",\"type\":\"tuple[]\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"handleOps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymaster\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.MemoryUserOp\",\"name\":\"mUserOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"prefund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contextOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preOpGas\",\"type\":\"uint256\"}],\"internalType\":\"structEntryPoint.UserOpInfo\",\"name\":\"opInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"innerHandleOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"name\":\"nonceSequenceNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"simulateHandleOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"simulateValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405260405162000012906200007a565b604051809103905ff0801580156200002c573d5f803e3d5ffd5b5073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156200006b575f80fd5b50600160028190555062000088565b61034080620071f683390190565b60805161714e620000a85f395f81816117eb0152613add015261714e5ff3fe608060405260043610610122575f3560e01c80638f41ec5a1161009f578063bb9fe6bf11610063578063bb9fe6bf146103f0578063c23a5cea14610406578063d6383f941461042e578063ee21942314610456578063fc7e286d1461047e57610132565b80638f41ec5a1461031e578063957122ab146103485780639b249f6914610370578063a619353114610398578063b760faf9146103d457610132565b8063205c2878116100e6578063205c28781461021a57806335567e1a146102425780634b1d7cf51461027e5780635287ce12146102a657806370a08231146102e257610132565b80630396cb60146101365780630bd28e3b146101525780631b2e01b81461017a5780631d732756146101b65780631fad948c146101f257610132565b3661013257610130336104be565b005b5f80fd5b610150600480360381019061014b9190613f08565b610575565b005b34801561015d575f80fd5b5061017860048036038101906101739190613f80565b610919565b005b348015610185575f80fd5b506101a0600480360381019061019b9190614005565b6109b1565b6040516101ad919061405b565b60405180910390f35b3480156101c1575f80fd5b506101dc60048036038101906101d791906143c3565b6109d1565b6040516101e9919061405b565b60405180910390f35b3480156101fd575f80fd5b50610218600480360381019061021391906144e2565b610b95565b005b348015610225575f80fd5b50610240600480360381019061023b919061453f565b610d21565b005b34801561024d575f80fd5b5061026860048036038101906102639190614005565b610f39565b604051610275919061405b565b60405180910390f35b348015610289575f80fd5b506102a4600480360381019061029f91906145d2565b610fe3565b005b3480156102b1575f80fd5b506102cc60048036038101906102c7919061462f565b611500565b6040516102d99190614731565b60405180910390f35b3480156102ed575f80fd5b506103086004803603810190610303919061462f565b611646565b604051610315919061405b565b60405180910390f35b348015610329575f80fd5b506103326116b6565b60405161033f919061405b565b60405180910390f35b348015610353575f80fd5b5061036e6004803603810190610369919061474a565b6116bb565b005b34801561037b575f80fd5b50610396600480360381019061039191906147db565b6117e8565b005b3480156103a3575f80fd5b506103be60048036038101906103b99190614849565b6118c3565b6040516103cb919061489f565b60405180910390f35b6103ee60048036038101906103e9919061462f565b6104be565b005b3480156103fb575f80fd5b506104046118fe565b005b348015610411575f80fd5b5061042c600480360381019061042791906148b8565b611aa4565b005b348015610439575f80fd5b50610454600480360381019061044f91906148e3565b611d94565b005b348015610461575f80fd5b5061047c60048036038101906104779190614849565b611ed9565b005b348015610489575f80fd5b506104a4600480360381019061049f919061462f565b612154565b6040516104b59594939291906149ac565b60405180910390f35b6104c881346121e4565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090508173ffffffffffffffffffffffffffffffffffffffff167f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4825f015f9054906101000a90046dffffffffffffffffffffffffffff166040516105699190614a36565b60405180910390a25050565b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f8263ffffffff16116105fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f390614aa9565b60405180910390fd5b806001015f9054906101000a900463ffffffff1663ffffffff168263ffffffff16101561065e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065590614b11565b60405180910390fd5b5f34825f01600f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff166106989190614b5c565b90505f81116106dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d390614bd9565b60405180910390fd5b6dffffffffffffffffffffffffffff801681111561072f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072690614c41565b60405180910390fd5b6040518060a00160405280835f015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff168152602001600115158152602001826dffffffffffffffffffffffffffff1681526020018463ffffffff1681526020015f65ffffffffffff168152505f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f015f6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff1602179055506020820151815f01600e6101000a81548160ff0219169083151502179055506040820151815f01600f6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff1602179055506060820151816001015f6101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160046101000a81548165ffffffffffff021916908365ffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167fa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01828560405161090c929190614c8f565b60405180910390a2505050565b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8277ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8154809291906109a990614cb6565b919050555050565b6001602052815f5260405f20602052805f5260405f205f91509150505481565b5f805a90503073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a44576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3b90614d47565b60405180910390fd5b5f855f015190505f8160400151905061138882606001518201015a1015610a8d577fdeaddead000000000000000000000000000000000000000000000000000000005f5260205ffd5b5f8089511115610b2a575f610aa7845f01515f8c866122ec565b905080610b28575f610aba610800612303565b90505f81511115610b2257845f015173ffffffffffffffffffffffffffffffffffffffff168a602001517f1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201876020015184604051610b19929190614ddf565b60405180910390a35b60019250505b505b5f88608001515a8603019050610b865f838b8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505085612332565b95505050505050949350505050565b610b9d612696565b5f8383905090505f8167ffffffffffffffff811115610bbf57610bbe61408c565b5b604051908082528060200260200182016040528015610bf857816020015b610be5613d6b565b815260200190600190039081610bdd5790505b5090505f5b82811015610c74575f828281518110610c1957610c18614e0d565b5b602002602001015190505f80610c54848a8a87818110610c3c57610c3b614e0d565b5b9050602002810190610c4e9190614e46565b856126e3565b91509150610c648483835f6128d7565b5050508080600101915050610bfd565b505f7fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97260405160405180910390a15f5b83811015610d0657610cf581888884818110610cc357610cc2614e0d565b5b9050602002810190610cd59190614e46565b858481518110610ce857610ce7614e0d565b5b6020026020010151612a67565b820191508080600101915050610ca4565b50610d118482612bc4565b505050610d1c612cdf565b505050565b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050805f015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16821115610dcf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dc690614eb8565b60405180910390fd5b81815f015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16610e079190614ed6565b815f015f6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb8484604051610e83929190614f5b565b60405180910390a25f8373ffffffffffffffffffffffffffffffffffffffff1683604051610eb090614faf565b5f6040518083038185875af1925050503d805f8114610eea576040519150601f19603f3d011682016040523d82523d5f602084013e610eef565b606091505b5050905080610f33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f2a9061500d565b60405180910390fd5b50505050565b5f60408277ffffffffffffffffffffffffffffffffffffffffffffffff16901b60015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8477ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205417905092915050565b610feb612696565b5f8383905090505f805b828110156111d0573686868381811061101157611010614e0d565b5b9050602002810190611023919061502b565b9050365f82805f01906110369190615052565b915091505f83602001602081019061104e91906150ef565b9050600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036110bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110b690615164565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146111a8578073ffffffffffffffffffffffffffffffffffffffff1663e3563a4f84848780604001906111209190615182565b6040518563ffffffff1660e01b815260040161113f9493929190615537565b5f6040518083038186803b158015611155575f80fd5b505afa925050508015611166575060015b6111a757806040517f86a9f75000000000000000000000000000000000000000000000000000000000815260040161119e919061557f565b60405180910390fd5b5b82829050866111b79190614b5c565b95505050505080806111c890614cb6565b915050610ff5565b505f8167ffffffffffffffff8111156111ec576111eb61408c565b5b60405190808252806020026020018201604052801561122557816020015b611212613d6b565b81526020019060019003908161120a5790505b5090507fbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f97260405160405180910390a15f805b8481101561135f573688888381811061127357611272614e0d565b5b9050602002810190611285919061502b565b9050365f82805f01906112989190615052565b915091505f8360200160208101906112b091906150ef565b90505f8383905090505f5b81811015611346575f8989815181106112d7576112d6614e0d565b5b602002602001015190505f806113128b8989878181106112fa576112f9614e0d565b5b905060200281019061130c9190614e46565b856126e3565b91509150611322848383896128d7565b8a8061132d90614cb6565b9b5050505050808061133e90614cb6565b9150506112bb565b505050505050808061135790614cb6565b915050611257565b505f8091505f5b858110156114a0573689898381811061138257611381614e0d565b5b9050602002810190611394919061502b565b90508060200160208101906113a991906150ef565b73ffffffffffffffffffffffffffffffffffffffff167f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d60405160405180910390a2365f82805f01906113fc9190615052565b915091505f8282905090505f5b818110156114885761145a8885858481811061142857611427614e0d565b5b905060200281019061143a9190614e46565b8b8b8151811061144d5761144c614e0d565b5b6020026020010151612a67565b876114659190614b5c565b9650878061147290614cb6565b985050808061148090614cb6565b915050611409565b5050505050808061149890614cb6565b915050611366565b505f73ffffffffffffffffffffffffffffffffffffffff167f575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d60405160405180910390a26114ee8682612bc4565b50505050506114fb612cdf565b505050565b611508613d9f565b5f808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060a00160405290815f82015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff1681526020015f8201600e9054906101000a900460ff161515151581526020015f8201600f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff168152602001600182015f9054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020016001820160049054906101000a900465ffffffffffff1665ffffffffffff1665ffffffffffff16815250509050919050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff169050919050565b600181565b5f858590501480156116e357505f8373ffffffffffffffffffffffffffffffffffffffff163b145b15611723576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161171a906155e2565b60405180910390fd5b601482829050106117ad575f82825f9060149261174293929190615608565b9061174d9190615683565b60601c90505f8173ffffffffffffffffffffffffffffffffffffffff163b036117ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117a29061572b565b60405180910390fd5b505b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117df90615769565b60405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663570e1a3684846040518363ffffffff1660e01b8152600401611844929190615787565b6020604051808303815f875af1158015611860573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061188491906157bd565b9050806040517f6ca7b8060000000000000000000000000000000000000000000000000000000081526004016118ba919061557f565b60405180910390fd5b5f6118cd82612ce9565b30466040516020016118e1939291906157e8565b604051602081830303815290604052805190602001209050919050565b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f816001015f9054906101000a900463ffffffff1663ffffffff1603611998576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161198f90615867565b60405180910390fd5b805f01600e9054906101000a900460ff166119e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119df906158cf565b60405180910390fd5b5f816001015f9054906101000a900463ffffffff1663ffffffff1642611a0e91906158ed565b9050808260010160046101000a81548165ffffffffffff021916908365ffffffffffff1602179055505f825f01600e6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167ffa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a82604051611a989190615956565b60405180910390a25050565b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f815f01600f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff1690505f8111611b56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b4d906159b9565b60405180910390fd5b5f8260010160049054906101000a900465ffffffffffff1665ffffffffffff1611611bb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bad90615a21565b60405180910390fd5b428260010160049054906101000a900465ffffffffffff1665ffffffffffff161115611c17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c0e90615a89565b60405180910390fd5b5f826001015f6101000a81548163ffffffff021916908363ffffffff1602179055505f8260010160046101000a81548165ffffffffffff021916908365ffffffffffff1602179055505f825f01600f6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda38483604051611cde929190614f5b565b60405180910390a25f8373ffffffffffffffffffffffffffffffffffffffff1682604051611d0b90614faf565b5f6040518083038185875af1925050503d805f8114611d45576040519150601f19603f3d011682016040523d82523d5f602084013e611d4a565b606091505b5050905080611d8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d8590615af1565b60405180910390fd5b50505050565b611d9c613d6b565b611da585612d01565b5f80611db25f88856126e3565b915091505f611dc18383612e22565b9050611dcb612f2b565b5f611dd75f8a87612a67565b9050611de1612f2b565b5f60605f73ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff1614611e86578973ffffffffffffffffffffffffffffffffffffffff168989604051611e3e929190615b33565b5f604051808303815f865af19150503d805f8114611e77576040519150601f19603f3d011682016040523d82523d5f602084013e611e7c565b606091505b5080925081935050505b8660800151838560200151866040015185856040517f8b7ac980000000000000000000000000000000000000000000000000000000008152600401611ed096959493929190615b4b565b60405180910390fd5b611ee1613d6b565b611eea82612d01565b5f80611ef75f85856126e3565b915091505f611f0c845f015160a00151612f30565b90505f611f1e855f01515f0151612f30565b9050611f28613df9565b365f888060400190611f3a9190615182565b915091505f6014838390501015611f51575f611f73565b82825f90601492611f6493929190615608565b90611f6f9190615683565b60601c5b9050611f7e81612f30565b93505050505f611f8e8686612e22565b90505f815f015190505f600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161490505f6040518060c001604052808b6080015181526020018b6040015181526020018315158152602001856020015165ffffffffffff168152602001856040015165ffffffffffff1681526020016120238c60600151612fd6565b81525090505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156120915750600173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614155b15612111575f60405180604001604052808573ffffffffffffffffffffffffffffffffffffffff1681526020016120c786612f30565b81525090508187878a846040517ffaecb4e4000000000000000000000000000000000000000000000000000000008152600401612108959493929190615cf6565b60405180910390fd5b808686896040517fe0cff05f00000000000000000000000000000000000000000000000000000000815260040161214b9493929190615d4f565b60405180910390fd5b5f602052805f5260405f205f91509050805f015f9054906101000a90046dffffffffffffffffffffffffffff1690805f01600e9054906101000a900460ff1690805f01600f9054906101000a90046dffffffffffffffffffffffffffff1690806001015f9054906101000a900463ffffffff16908060010160049054906101000a900465ffffffffffff16905085565b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f82825f015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff1661225c9190614b5c565b90506dffffffffffffffffffffffffffff80168111156122b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122a890615de3565b60405180910390fd5b80825f015f6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff16021790555050505050565b5f805f845160208601878987f19050949350505050565b60603d82811115612312578290505b604051602082018101604052818152815f602083013e8092505050919050565b5f805a90505f80865f015190505f61234982612fe0565b90505f8260a0015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361239157825f01519350612578565b8093505f885111156125775781870295506002808111156123b5576123b4615e01565b5b8a60028111156123c8576123c7615e01565b5b14612442578073ffffffffffffffffffffffffffffffffffffffff1663a9a2340984606001518c8b8a6040518563ffffffff1660e01b815260040161240f93929190615e74565b5f604051808303815f88803b158015612426575f80fd5b5087f1158015612438573d5f803e3d5ffd5b5050505050612576565b8073ffffffffffffffffffffffffffffffffffffffff1663a9a2340984606001518c8b8a6040518563ffffffff1660e01b815260040161248493929190615e74565b5f604051808303815f88803b15801561249b575f80fd5b5087f1935050505080156124ad575060015b612575576124b9615ebc565b806308c379a00361253657506124cd615edb565b806124d85750612538565b8b816040516020016124ea9190615fd4565b6040516020818303038152906040526040517f220266b600000000000000000000000000000000000000000000000000000000815260040161252d929190616031565b60405180910390fd5b505b8a6040517f220266b600000000000000000000000000000000000000000000000000000000815260040161256c91906160a9565b60405180910390fd5b5b5b5b5a850387019650818702955085896040015110156125cd578a6040517f220266b60000000000000000000000000000000000000000000000000000000081526004016125c4919061611f565b60405180910390fd5b5f868a604001510390506125e185826121e4565b5f8060028111156125f5576125f4615e01565b5b8c600281111561260857612607615e01565b5b1490508460a0015173ffffffffffffffffffffffffffffffffffffffff16855f015173ffffffffffffffffffffffffffffffffffffffff168c602001517f49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f8860200151858d8f60405161267e949392919061614b565b60405180910390a45050505050505095945050505050565b60028054036126da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126d1906161d8565b60405180910390fd5b60028081905550565b5f805f5a90505f845f015190506126fa8682613018565b612703866118c3565b8560200181815250505f8661010001358760e001358360400151846060015185608001511717171790506effffffffffffffffffffffffffffff8016811115612781576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161277890616240565b60405180910390fd5b5f8061278c846131ba565b905061279a8a8a8a84613228565b80985081935050506127b3845f0151856020015161353d565b6127f457896040517f220266b60000000000000000000000000000000000000000000000000000000081526004016127eb91906162a8565b60405180910390fd5b6127fc612f2b565b60605f73ffffffffffffffffffffffffffffffffffffffff168560a0015173ffffffffffffffffffffffffffffffffffffffff161461284b576128428b8b8b85876135f4565b80985081925050505b5f5a87039050808b60a00135101561289a578b6040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612891919061631e565b60405180910390fd5b828a60400181815250506128ad826138a2565b8a60600181815250508a60c001355a8803018a608001818152505050505050505050935093915050565b5f806128e2856138ab565b915091508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161461295657856040517f220266b600000000000000000000000000000000000000000000000000000000815260040161294d9190616394565b60405180910390fd5b801561299957856040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612990919061640a565b60405180910390fd5b5f6129a3856138ab565b80935081925050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612a1b57866040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612a129190616480565b60405180910390fd5b8115612a5e57866040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612a55919061651c565b60405180910390fd5b50505050505050565b5f805a90505f612a7a8460600151612fd6565b90503073ffffffffffffffffffffffffffffffffffffffff16631d732756868060600190612aa89190615182565b87856040518563ffffffff1660e01b8152600401612ac99493929190616662565b6020604051808303815f875af1925050508015612b0457506040513d601f19601f82011682018060405250810190612b0191906166bd565b60015b612bb6575f3d80602003612b1c5760205f803e5f5191505b507fdeaddead000000000000000000000000000000000000000000000000000000008103612b8157866040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612b789190616732565b60405180910390fd5b5f85608001515a85612b939190614ed6565b612b9d9190614b5c565b9050612bad886002888685612332565b94505050612bbb565b809350505b50509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612c32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c29906167a8565b60405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff1682604051612c5790614faf565b5f6040518083038185875af1925050503d805f8114612c91576040519150601f19603f3d011682016040523d82523d5f602084013e612c96565b606091505b5050905080612cda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cd190616810565b60405180910390fd5b505050565b6001600281905550565b5f612cf3826138ff565b805190602001209050919050565b3073ffffffffffffffffffffffffffffffffffffffff1663957122ab828060400190612d2d9190615182565b845f016020810190612d3f919061462f565b85806101200190612d509190615182565b6040518663ffffffff1660e01b8152600401612d7095949392919061682e565b5f6040518083038186803b158015612d86575f80fd5b505afa925050508015612d97575060015b612e1e57612da3615ebc565b806308c379a003612e0f5750612db7615edb565b80612dc25750612e11565b5f815114612e09575f816040517f220266b6000000000000000000000000000000000000000000000000000000008152600401612e009291906168ae565b60405180910390fd5b50612e19565b505b3d5f803e3d5ffd5b612e1f565b5b50565b612e2a613e11565b5f612e34846139d3565b90505f612e40846139d3565b90505f825f015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612e8357815f015190505b5f836020015190505f846040015190505f846020015190505f856040015190508165ffffffffffff168465ffffffffffff161015612ebf578193505b8065ffffffffffff168365ffffffffffff161115612edb578092505b60405180606001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018565ffffffffffff1681526020018465ffffffffffff1681525097505050505050505092915050565b435f52565b612f38613df9565b5f805f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050805f01600f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff16825f018181525050806001015f9054906101000a900463ffffffff1663ffffffff1682602001818152505050919050565b6060819050919050565b5f808260c0015190505f8360e001519050808203613002578192505050613013565b61300e82488301613a53565b925050505b919050565b815f01602081019061302a919061462f565b815f019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050816020013581602001818152505081608001358160400181815250508160a001358160600181815250508160c001358160800181815250508160e001358160c00181815250508161010001358160e0018181525050365f838061012001906130c89190615182565b915091505f82829050111561317b57601482829050101561311e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161311590616926565b60405180910390fd5b81815f9060149261313193929190615608565b9061313c9190615683565b60601c8360a0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506131b4565b5f8360a0019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250505b50505050565b5f805f73ffffffffffffffffffffffffffffffffffffffff168360a0015173ffffffffffffffffffffffffffffffffffffffff16036131fa5760016131fd565b60035b60ff1690505f8360800151828560600151028560400151010190508360c00151810292505050919050565b5f805f5a90505f855f015190505f815f0151905061325689888a80604001906132519190615182565b613a6b565b5f8260a001519050613266612f2b565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036132bb575f6132a484611646565b90508881116132b5578089036132b7565b5f5b9150505b8273ffffffffffffffffffffffffffffffffffffffff16633a871cdd85606001518c8c60200151856040518563ffffffff1660e01b815260040161330193929190616aa5565b6020604051808303815f8887f19350505050801561333d57506040513d601f19601f8201168201806040525081019061333a91906166bd565b60015b61340557613349615ebc565b806308c379a0036133c6575061335d615edb565b8061336857506133c8565b8b8160405160200161337a9190616b07565b6040516020818303038152906040526040517f220266b60000000000000000000000000000000000000000000000000000000081526004016133bd929190616031565b60405180910390fd5b505b8a6040517f220266b60000000000000000000000000000000000000000000000000000000081526004016133fc9190616b76565b60405180910390fd5b809650505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361352a575f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f815f015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff169050808a11156134f0578c6040517f220266b60000000000000000000000000000000000000000000000000000000081526004016134e79190616bec565b60405180910390fd5b898103825f015f6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff16021790555050505b5a85039650505050505094509492505050565b5f80604083901c90505f8390508067ffffffffffffffff1660015f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8477ffffffffffffffffffffffffffffffffffffffffffffffff1677ffffffffffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8154809291906135e590614cb6565b91905055149250505092915050565b60605f80855f015190505f81606001519050848111613648576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161363f90616c62565b60405180910390fd5b5f85820390505f8360a0015190505f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f815f015f9054906101000a90046dffffffffffffffffffffffffffff166dffffffffffffffffffffffffffff1690508981101561370a578c6040517f220266b60000000000000000000000000000000000000000000000000000000081526004016137019190616cca565b60405180910390fd5b898103825f015f6101000a8154816dffffffffffffffffffffffffffff02191690836dffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff1663f465c77e858e8e602001518e6040518563ffffffff1660e01b815260040161378393929190616aa5565b5f604051808303815f8887f1935050505080156137c257506040513d5f823e3d601f19601f820116820180604052508101906137bf9190616d64565b60015b61388a576137ce615ebc565b806308c379a00361384b57506137e2615edb565b806137ed575061384d565b8d816040516020016137ff9190616de4565b6040516020818303038152906040526040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613842929190616031565b60405180910390fd5b505b8c6040517f220266b60000000000000000000000000000000000000000000000000000000081526004016138819190616e53565b60405180910390fd5b81995080985050505050505050509550959350505050565b5f819050919050565b5f805f83036138bf575f80915091506138fa565b5f6138c9846139d3565b9050806040015165ffffffffffff164211806138f05750806020015165ffffffffffff1642105b9150805f01519250505b915091565b60605f61390b83613d46565b90505f836020013590505f61392e8580604001906139299190615182565b613d55565b90505f6139498680606001906139449190615182565b613d55565b90505f866080013590505f8760a0013590505f8860c0013590505f8960e0013590505f8a610100013590505f61398e8c8061012001906139899190615182565b613d55565b9050898989898989898989896040516020016139b39a99989796959493929190616e7f565b6040516020818303038152906040529a5050505050505050505050919050565b6139db613e11565b5f8290505f60a084901c90505f8165ffffffffffff1603613a005765ffffffffffff90505b5f60d085901c905060405180606001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018265ffffffffffff1681526020018365ffffffffffff168152509350505050919050565b5f818310613a615781613a63565b825b905092915050565b5f8282905014613d40575f835f01515f015190505f8173ffffffffffffffffffffffffffffffffffffffff163b14613ada57846040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613ad19190616f63565b60405180910390fd5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663570e1a36865f01516060015186866040518463ffffffff1660e01b8152600401613b3e929190615787565b6020604051808303815f8887f1158015613b5a573d5f803e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190613b7f91906157bd565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603613bf157856040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613be89190616fd9565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614613c6157856040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613c58919061704f565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff163b03613cbc57856040517f220266b6000000000000000000000000000000000000000000000000000000008152600401613cb391906170c5565b60405180910390fd5b5f84845f90601492613cd093929190615608565b90613cdb9190615683565b60601c90508273ffffffffffffffffffffffffffffffffffffffff1686602001517fd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d83895f015160a00151604051613d349291906170f1565b60405180910390a35050505b50505050565b5f808235905080915050919050565b5f60405182808583378082209250505092915050565b6040518060a00160405280613d7e613e55565b81526020015f80191681526020015f81526020015f81526020015f81525090565b6040518060a001604052805f6dffffffffffffffffffffffffffff1681526020015f151581526020015f6dffffffffffffffffffffffffffff1681526020015f63ffffffff1681526020015f65ffffffffffff1681525090565b60405180604001604052805f81526020015f81525090565b60405180606001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f65ffffffffffff1681526020015f65ffffffffffff1681525090565b6040518061010001604052805f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81526020015f81526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681526020015f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f63ffffffff82169050919050565b613ee781613ecf565b8114613ef1575f80fd5b50565b5f81359050613f0281613ede565b92915050565b5f60208284031215613f1d57613f1c613ec7565b5b5f613f2a84828501613ef4565b91505092915050565b5f77ffffffffffffffffffffffffffffffffffffffffffffffff82169050919050565b613f5f81613f33565b8114613f69575f80fd5b50565b5f81359050613f7a81613f56565b92915050565b5f60208284031215613f9557613f94613ec7565b5b5f613fa284828501613f6c565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f613fd482613fab565b9050919050565b613fe481613fca565b8114613fee575f80fd5b50565b5f81359050613fff81613fdb565b92915050565b5f806040838503121561401b5761401a613ec7565b5b5f61402885828601613ff1565b925050602061403985828601613f6c565b9150509250929050565b5f819050919050565b61405581614043565b82525050565b5f60208201905061406e5f83018461404c565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6140c28261407c565b810181811067ffffffffffffffff821117156140e1576140e061408c565b5b80604052505050565b5f6140f3613ebe565b90506140ff82826140b9565b919050565b5f67ffffffffffffffff82111561411e5761411d61408c565b5b6141278261407c565b9050602081019050919050565b828183375f83830152505050565b5f61415461414f84614104565b6140ea565b9050828152602081018484840111156141705761416f614078565b5b61417b848285614134565b509392505050565b5f82601f83011261419757614196614074565b5b81356141a7848260208601614142565b91505092915050565b5f80fd5b6141bd81614043565b81146141c7575f80fd5b50565b5f813590506141d8816141b4565b92915050565b5f61010082840312156141f4576141f36141b0565b5b6141ff6101006140ea565b90505f61420e84828501613ff1565b5f830152506020614221848285016141ca565b6020830152506040614235848285016141ca565b6040830152506060614249848285016141ca565b606083015250608061425d848285016141ca565b60808301525060a061427184828501613ff1565b60a08301525060c0614285848285016141ca565b60c08301525060e0614299848285016141ca565b60e08301525092915050565b5f819050919050565b6142b7816142a5565b81146142c1575f80fd5b50565b5f813590506142d2816142ae565b92915050565b5f61018082840312156142ee576142ed6141b0565b5b6142f860a06140ea565b90505f614307848285016141de565b5f8301525061010061431b848285016142c4565b602083015250610120614330848285016141ca565b604083015250610140614345848285016141ca565b60608301525061016061435a848285016141ca565b60808301525092915050565b5f80fd5b5f80fd5b5f8083601f84011261438357614382614074565b5b8235905067ffffffffffffffff8111156143a05761439f614366565b5b6020830191508360018202830111156143bc576143bb61436a565b5b9250929050565b5f805f806101c085870312156143dc576143db613ec7565b5b5f85013567ffffffffffffffff8111156143f9576143f8613ecb565b5b61440587828801614183565b9450506020614416878288016142d8565b9350506101a085013567ffffffffffffffff81111561443857614437613ecb565b5b6144448782880161436e565b925092505092959194509250565b5f8083601f84011261446757614466614074565b5b8235905067ffffffffffffffff81111561448457614483614366565b5b6020830191508360208202830111156144a05761449f61436a565b5b9250929050565b5f6144b182613fab565b9050919050565b6144c1816144a7565b81146144cb575f80fd5b50565b5f813590506144dc816144b8565b92915050565b5f805f604084860312156144f9576144f8613ec7565b5b5f84013567ffffffffffffffff81111561451657614515613ecb565b5b61452286828701614452565b93509350506020614535868287016144ce565b9150509250925092565b5f806040838503121561455557614554613ec7565b5b5f614562858286016144ce565b9250506020614573858286016141ca565b9150509250929050565b5f8083601f84011261459257614591614074565b5b8235905067ffffffffffffffff8111156145af576145ae614366565b5b6020830191508360208202830111156145cb576145ca61436a565b5b9250929050565b5f805f604084860312156145e9576145e8613ec7565b5b5f84013567ffffffffffffffff81111561460657614605613ecb565b5b6146128682870161457d565b93509350506020614625868287016144ce565b9150509250925092565b5f6020828403121561464457614643613ec7565b5b5f61465184828501613ff1565b91505092915050565b5f6dffffffffffffffffffffffffffff82169050919050565b61467c8161465a565b82525050565b5f8115159050919050565b61469681614682565b82525050565b6146a581613ecf565b82525050565b5f65ffffffffffff82169050919050565b6146c5816146ab565b82525050565b60a082015f8201516146df5f850182614673565b5060208201516146f2602085018261468d565b5060408201516147056040850182614673565b506060820151614718606085018261469c565b50608082015161472b60808501826146bc565b50505050565b5f60a0820190506147445f8301846146cb565b92915050565b5f805f805f6060868803121561476357614762613ec7565b5b5f86013567ffffffffffffffff8111156147805761477f613ecb565b5b61478c8882890161436e565b9550955050602061479f88828901613ff1565b935050604086013567ffffffffffffffff8111156147c0576147bf613ecb565b5b6147cc8882890161436e565b92509250509295509295909350565b5f80602083850312156147f1576147f0613ec7565b5b5f83013567ffffffffffffffff81111561480e5761480d613ecb565b5b61481a8582860161436e565b92509250509250929050565b5f80fd5b5f61016082840312156148405761483f614826565b5b81905092915050565b5f6020828403121561485e5761485d613ec7565b5b5f82013567ffffffffffffffff81111561487b5761487a613ecb565b5b6148878482850161482a565b91505092915050565b614899816142a5565b82525050565b5f6020820190506148b25f830184614890565b92915050565b5f602082840312156148cd576148cc613ec7565b5b5f6148da848285016144ce565b91505092915050565b5f805f80606085870312156148fb576148fa613ec7565b5b5f85013567ffffffffffffffff81111561491857614917613ecb565b5b6149248782880161482a565b945050602061493587828801613ff1565b935050604085013567ffffffffffffffff81111561495657614955613ecb565b5b6149628782880161436e565b925092505092959194509250565b6149798161465a565b82525050565b61498881614682565b82525050565b61499781613ecf565b82525050565b6149a6816146ab565b82525050565b5f60a0820190506149bf5f830188614970565b6149cc602083018761497f565b6149d96040830186614970565b6149e6606083018561498e565b6149f3608083018461499d565b9695505050505050565b5f819050919050565b5f614a20614a1b614a168461465a565b6149fd565b614043565b9050919050565b614a3081614a06565b82525050565b5f602082019050614a495f830184614a27565b92915050565b5f82825260208201905092915050565b7f6d757374207370656369667920756e7374616b652064656c61790000000000005f82015250565b5f614a93601a83614a4f565b9150614a9e82614a5f565b602082019050919050565b5f6020820190508181035f830152614ac081614a87565b9050919050565b7f63616e6e6f7420646563726561736520756e7374616b652074696d65000000005f82015250565b5f614afb601c83614a4f565b9150614b0682614ac7565b602082019050919050565b5f6020820190508181035f830152614b2881614aef565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f614b6682614043565b9150614b7183614043565b9250828201905080821115614b8957614b88614b2f565b5b92915050565b7f6e6f207374616b652073706563696669656400000000000000000000000000005f82015250565b5f614bc3601283614a4f565b9150614bce82614b8f565b602082019050919050565b5f6020820190508181035f830152614bf081614bb7565b9050919050565b7f7374616b65206f766572666c6f770000000000000000000000000000000000005f82015250565b5f614c2b600e83614a4f565b9150614c3682614bf7565b602082019050919050565b5f6020820190508181035f830152614c5881614c1f565b9050919050565b5f614c79614c74614c6f84613ecf565b6149fd565b614043565b9050919050565b614c8981614c5f565b82525050565b5f604082019050614ca25f83018561404c565b614caf6020830184614c80565b9392505050565b5f614cc082614043565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614cf257614cf1614b2f565b5b600182019050919050565b7f4141393220696e7465726e616c2063616c6c206f6e6c790000000000000000005f82015250565b5f614d31601783614a4f565b9150614d3c82614cfd565b602082019050919050565b5f6020820190508181035f830152614d5e81614d25565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015614d9c578082015181840152602081019050614d81565b5f8484015250505050565b5f614db182614d65565b614dbb8185614d6f565b9350614dcb818560208601614d7f565b614dd48161407c565b840191505092915050565b5f604082019050614df25f83018561404c565b8181036020830152614e048184614da7565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f80fd5b5f80fd5b5f8235600161016003833603038112614e6257614e61614e3a565b5b80830191505092915050565b7f576974686472617720616d6f756e7420746f6f206c61726765000000000000005f82015250565b5f614ea2601983614a4f565b9150614ead82614e6e565b602082019050919050565b5f6020820190508181035f830152614ecf81614e96565b9050919050565b5f614ee082614043565b9150614eeb83614043565b9250828203905081811115614f0357614f02614b2f565b5b92915050565b5f614f23614f1e614f1984613fab565b6149fd565b613fab565b9050919050565b5f614f3482614f09565b9050919050565b5f614f4582614f2a565b9050919050565b614f5581614f3b565b82525050565b5f604082019050614f6e5f830185614f4c565b614f7b602083018461404c565b9392505050565b5f81905092915050565b50565b5f614f9a5f83614f82565b9150614fa582614f8c565b5f82019050919050565b5f614fb982614f8f565b9150819050919050565b7f6661696c656420746f20776974686472617700000000000000000000000000005f82015250565b5f614ff7601283614a4f565b915061500282614fc3565b602082019050919050565b5f6020820190508181035f83015261502481614feb565b9050919050565b5f8235600160600383360303811261504657615045614e3a565b5b80830191505092915050565b5f808335600160200384360303811261506e5761506d614e3a565b5b80840192508235915067ffffffffffffffff8211156150905761508f614e3e565b5b6020830192506020820236038313156150ac576150ab614e42565b5b509250929050565b5f6150be82613fca565b9050919050565b6150ce816150b4565b81146150d8575f80fd5b50565b5f813590506150e9816150c5565b92915050565b5f6020828403121561510457615103613ec7565b5b5f615111848285016150db565b91505092915050565b7f4141393620696e76616c69642061676772656761746f720000000000000000005f82015250565b5f61514e601783614a4f565b91506151598261511a565b602082019050919050565b5f6020820190508181035f83015261517b81615142565b9050919050565b5f808335600160200384360303811261519e5761519d614e3a565b5b80840192508235915067ffffffffffffffff8211156151c0576151bf614e3e565b5b6020830192506001820236038313156151dc576151db614e42565b5b509250929050565b5f82825260208201905092915050565b5f819050919050565b5f61520b6020840184613ff1565b905092915050565b61521c81613fca565b82525050565b5f61523060208401846141ca565b905092915050565b61524181614043565b82525050565b5f80fd5b5f80fd5b5f80fd5b5f808335600160200384360303811261526f5761526e61524f565b5b83810192508235915060208301925067ffffffffffffffff82111561529757615296615247565b5b6001820236038313156152ad576152ac61524b565b5b509250929050565b5f82825260208201905092915050565b5f6152d083856152b5565b93506152dd838584614134565b6152e68361407c565b840190509392505050565b5f61016083016153035f8401846151fd565b61530f5f860182615213565b5061531d6020840184615222565b61532a6020860182615238565b506153386040840184615253565b858303604087015261534b8382846152c5565b9250505061535c6060840184615253565b858303606087015261536f8382846152c5565b925050506153806080840184615222565b61538d6080860182615238565b5061539b60a0840184615222565b6153a860a0860182615238565b506153b660c0840184615222565b6153c360c0860182615238565b506153d160e0840184615222565b6153de60e0860182615238565b506153ed610100840184615222565b6153fb610100860182615238565b5061540a610120840184615253565b85830361012087015261541e8382846152c5565b92505050615430610140840184615253565b8583036101408701526154448382846152c5565b925050508091505092915050565b5f61545d83836152f1565b905092915050565b5f82356001610160038336030381126154815761548061524f565b5b82810191505092915050565b5f602082019050919050565b5f6154a483856151e4565b9350836020840285016154b6846151f4565b805f5b878110156154f95784840389526154d08284615465565b6154da8582615452565b94506154e58361548d565b925060208a019950506001810190506154b9565b50829750879450505050509392505050565b5f6155168385614d6f565b9350615523838584614134565b61552c8361407c565b840190509392505050565b5f6040820190508181035f830152615550818688615499565b9050818103602083015261556581848661550b565b905095945050505050565b61557981613fca565b82525050565b5f6020820190506155925f830184615570565b92915050565b7f41413230206163636f756e74206e6f74206465706c6f796564000000000000005f82015250565b5f6155cc601983614a4f565b91506155d782615598565b602082019050919050565b5f6020820190508181035f8301526155f9816155c0565b9050919050565b5f80fd5b5f80fd5b5f808585111561561b5761561a615600565b5b8386111561562c5761562b615604565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b5f82821b905092915050565b5f61568e8383615642565b82615699813561564c565b925060148210156156d9576156d47fffffffffffffffffffffffffffffffffffffffff00000000000000000000000083601403600802615677565b831692505b505092915050565b7f41413330207061796d6173746572206e6f74206465706c6f79656400000000005f82015250565b5f615715601b83614a4f565b9150615720826156e1565b602082019050919050565b5f6020820190508181035f83015261574281615709565b9050919050565b5f6157545f83614a4f565b915061575f82614f8c565b5f82019050919050565b5f6020820190508181035f83015261578081615749565b9050919050565b5f6020820190508181035f8301526157a081848661550b565b90509392505050565b5f815190506157b781613fdb565b92915050565b5f602082840312156157d2576157d1613ec7565b5b5f6157df848285016157a9565b91505092915050565b5f6060820190506157fb5f830186614890565b6158086020830185615570565b615815604083018461404c565b949350505050565b7f6e6f74207374616b6564000000000000000000000000000000000000000000005f82015250565b5f615851600a83614a4f565b915061585c8261581d565b602082019050919050565b5f6020820190508181035f83015261587e81615845565b9050919050565b7f616c726561647920756e7374616b696e670000000000000000000000000000005f82015250565b5f6158b9601183614a4f565b91506158c482615885565b602082019050919050565b5f6020820190508181035f8301526158e6816158ad565b9050919050565b5f6158f7826146ab565b9150615902836146ab565b9250828201905065ffffffffffff8111156159205761591f614b2f565b5b92915050565b5f61594061593b615936846146ab565b6149fd565b614043565b9050919050565b61595081615926565b82525050565b5f6020820190506159695f830184615947565b92915050565b7f4e6f207374616b6520746f2077697468647261770000000000000000000000005f82015250565b5f6159a3601483614a4f565b91506159ae8261596f565b602082019050919050565b5f6020820190508181035f8301526159d081615997565b9050919050565b7f6d7573742063616c6c20756e6c6f636b5374616b6528292066697273740000005f82015250565b5f615a0b601d83614a4f565b9150615a16826159d7565b602082019050919050565b5f6020820190508181035f830152615a38816159ff565b9050919050565b7f5374616b65207769746864726177616c206973206e6f742064756500000000005f82015250565b5f615a73601b83614a4f565b9150615a7e82615a3f565b602082019050919050565b5f6020820190508181035f830152615aa081615a67565b9050919050565b7f6661696c656420746f207769746864726177207374616b6500000000000000005f82015250565b5f615adb601883614a4f565b9150615ae682615aa7565b602082019050919050565b5f6020820190508181035f830152615b0881615acf565b9050919050565b5f615b1a8385614f82565b9350615b27838584614134565b82840190509392505050565b5f615b3f828486615b0f565b91508190509392505050565b5f60c082019050615b5e5f83018961404c565b615b6b602083018861404c565b615b78604083018761499d565b615b85606083018661499d565b615b92608083018561497f565b81810360a0830152615ba48184614da7565b9050979650505050505050565b5f615bbb82614d65565b615bc581856152b5565b9350615bd5818560208601614d7f565b615bde8161407c565b840191505092915050565b5f60c083015f830151615bfe5f860182615238565b506020830151615c116020860182615238565b506040830151615c24604086018261468d565b506060830151615c3760608601826146bc565b506080830151615c4a60808601826146bc565b5060a083015184820360a0860152615c628282615bb1565b9150508091505092915050565b604082015f820151615c835f850182615238565b506020820151615c966020850182615238565b50505050565b604082015f820151615cb05f850182615238565b506020820151615cc36020850182615238565b50505050565b606082015f820151615cdd5f850182615213565b506020820151615cf06020850182615c9c565b50505050565b5f610140820190508181035f830152615d0f8188615be9565b9050615d1e6020830187615c6f565b615d2b6060830186615c6f565b615d3860a0830185615c6f565b615d4560e0830184615cc9565b9695505050505050565b5f60e0820190508181035f830152615d678187615be9565b9050615d766020830186615c6f565b615d836060830185615c6f565b615d9060a0830184615c6f565b95945050505050565b7f6465706f736974206f766572666c6f77000000000000000000000000000000005f82015250565b5f615dcd601083614a4f565b9150615dd882615d99565b602082019050919050565b5f6020820190508181035f830152615dfa81615dc1565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110615e3f57615e3e615e01565b5b50565b5f819050615e4f82615e2e565b919050565b5f615e5e82615e42565b9050919050565b615e6e81615e54565b82525050565b5f606082019050615e875f830186615e65565b8181036020830152615e998185614da7565b9050615ea8604083018461404c565b949350505050565b5f8160e01c9050919050565b5f60033d1115615ed85760045f803e615ed55f51615eb0565b90505b90565b5f60443d10615f6757615eec613ebe565b60043d036004823e80513d602482011167ffffffffffffffff82111715615f14575050615f67565b808201805167ffffffffffffffff811115615f325750505050615f67565b80602083010160043d038501811115615f4f575050505050615f67565b615f5e826020018501866140b9565b82955050505050505b90565b7f4141353020706f73744f702072657665727465643a2000000000000000000000815250565b5f81519050919050565b5f81905092915050565b5f615fae82615f90565b615fb88185615f9a565b9350615fc8818560208601614d7f565b80840191505092915050565b5f615fde82615f6a565b601682019150615fee8284615fa4565b915081905092915050565b5f61600382615f90565b61600d8185614a4f565b935061601d818560208601614d7f565b6160268161407c565b840191505092915050565b5f6040820190506160445f83018561404c565b81810360208301526160568184615ff9565b90509392505050565b7f4141353020706f73744f702072657665727400000000000000000000000000005f82015250565b5f616093601283614a4f565b915061609e8261605f565b602082019050919050565b5f6040820190506160bc5f83018461404c565b81810360208301526160cd81616087565b905092915050565b7f414135312070726566756e642062656c6f772061637475616c476173436f73745f82015250565b5f616109602083614a4f565b9150616114826160d5565b602082019050919050565b5f6040820190506161325f83018461404c565b8181036020830152616143816160fd565b905092915050565b5f60808201905061615e5f83018761404c565b61616b602083018661497f565b616178604083018561404c565b616185606083018461404c565b95945050505050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c005f82015250565b5f6161c2601f83614a4f565b91506161cd8261618e565b602082019050919050565b5f6020820190508181035f8301526161ef816161b6565b9050919050565b7f41413934206761732076616c756573206f766572666c6f7700000000000000005f82015250565b5f61622a601883614a4f565b9150616235826161f6565b602082019050919050565b5f6020820190508181035f8301526162578161621e565b9050919050565b7f4141323520696e76616c6964206163636f756e74206e6f6e63650000000000005f82015250565b5f616292601a83614a4f565b915061629d8261625e565b602082019050919050565b5f6040820190506162bb5f83018461404c565b81810360208301526162cc81616286565b905092915050565b7f41413430206f76657220766572696669636174696f6e4761734c696d697400005f82015250565b5f616308601e83614a4f565b9150616313826162d4565b602082019050919050565b5f6040820190506163315f83018461404c565b8181036020830152616342816162fc565b905092915050565b7f41413234207369676e6174757265206572726f720000000000000000000000005f82015250565b5f61637e601483614a4f565b91506163898261634a565b602082019050919050565b5f6040820190506163a75f83018461404c565b81810360208301526163b881616372565b905092915050565b7f414132322065787069726564206f72206e6f74206475650000000000000000005f82015250565b5f6163f4601783614a4f565b91506163ff826163c0565b602082019050919050565b5f60408201905061641d5f83018461404c565b818103602083015261642e816163e8565b905092915050565b7f41413334207369676e6174757265206572726f720000000000000000000000005f82015250565b5f61646a601483614a4f565b915061647582616436565b602082019050919050565b5f6040820190506164935f83018461404c565b81810360208301526164a48161645e565b905092915050565b7f41413332207061796d61737465722065787069726564206f72206e6f742064755f8201527f6500000000000000000000000000000000000000000000000000000000000000602082015250565b5f616506602183614a4f565b9150616511826164ac565b604082019050919050565b5f60408201905061652f5f83018461404c565b8181036020830152616540816164fa565b905092915050565b61010082015f82015161655d5f850182615213565b5060208201516165706020850182615238565b5060408201516165836040850182615238565b5060608201516165966060850182615238565b5060808201516165a96080850182615238565b5060a08201516165bc60a0850182615213565b5060c08201516165cf60c0850182615238565b5060e08201516165e260e0850182615238565b50505050565b6165f1816142a5565b82525050565b61018082015f82015161660c5f850182616548565b5060208201516166206101008501826165e8565b506040820151616634610120850182615238565b506060820151616648610140850182615238565b50608082015161665c610160850182615238565b50505050565b5f6101c0820190508181035f83015261667c81868861550b565b905061668b60208301856165f7565b8181036101a083015261669e8184614da7565b905095945050505050565b5f815190506166b7816141b4565b92915050565b5f602082840312156166d2576166d1613ec7565b5b5f6166df848285016166a9565b91505092915050565b7f41413935206f7574206f662067617300000000000000000000000000000000005f82015250565b5f61671c600f83614a4f565b9150616727826166e8565b602082019050919050565b5f6040820190506167455f83018461404c565b818103602083015261675681616710565b905092915050565b7f4141393020696e76616c69642062656e656669636961727900000000000000005f82015250565b5f616792601883614a4f565b915061679d8261675e565b602082019050919050565b5f6020820190508181035f8301526167bf81616786565b9050919050565b7f41413931206661696c65642073656e6420746f2062656e6566696369617279005f82015250565b5f6167fa601f83614a4f565b9150616805826167c6565b602082019050919050565b5f6020820190508181035f830152616827816167ee565b9050919050565b5f6060820190508181035f83015261684781878961550b565b90506168566020830186615570565b818103604083015261686981848661550b565b90509695505050505050565b5f819050919050565b5f61689861689361688e84616875565b6149fd565b614043565b9050919050565b6168a88161687e565b82525050565b5f6040820190506168c15f83018561689f565b81810360208301526168d38184615ff9565b90509392505050565b7f4141393320696e76616c6964207061796d6173746572416e64446174610000005f82015250565b5f616910601d83614a4f565b915061691b826168dc565b602082019050919050565b5f6020820190508181035f83015261693d81616904565b9050919050565b5f61016083016169565f8401846151fd565b6169625f860182615213565b506169706020840184615222565b61697d6020860182615238565b5061698b6040840184615253565b858303604087015261699e8382846152c5565b925050506169af6060840184615253565b85830360608701526169c28382846152c5565b925050506169d36080840184615222565b6169e06080860182615238565b506169ee60a0840184615222565b6169fb60a0860182615238565b50616a0960c0840184615222565b616a1660c0860182615238565b50616a2460e0840184615222565b616a3160e0860182615238565b50616a40610100840184615222565b616a4e610100860182615238565b50616a5d610120840184615253565b858303610120870152616a718382846152c5565b92505050616a83610140840184615253565b858303610140870152616a978382846152c5565b925050508091505092915050565b5f6060820190508181035f830152616abd8186616944565b9050616acc6020830185614890565b616ad9604083018461404c565b949350505050565b7f414132332072657665727465643a200000000000000000000000000000000000815250565b5f616b1182616ae1565b600f82019150616b218284615fa4565b915081905092915050565b7f4141323320726576657274656420286f72204f4f4729000000000000000000005f82015250565b5f616b60601683614a4f565b9150616b6b82616b2c565b602082019050919050565b5f604082019050616b895f83018461404c565b8181036020830152616b9a81616b54565b905092915050565b7f41413231206469646e2774207061792070726566756e640000000000000000005f82015250565b5f616bd6601783614a4f565b9150616be182616ba2565b602082019050919050565b5f604082019050616bff5f83018461404c565b8181036020830152616c1081616bca565b905092915050565b7f4141343120746f6f206c6974746c6520766572696669636174696f6e476173005f82015250565b5f616c4c601f83614a4f565b9150616c5782616c18565b602082019050919050565b5f6020820190508181035f830152616c7981616c40565b9050919050565b7f41413331207061796d6173746572206465706f73697420746f6f206c6f7700005f82015250565b5f616cb4601e83614a4f565b9150616cbf82616c80565b602082019050919050565b5f604082019050616cdd5f83018461404c565b8181036020830152616cee81616ca8565b905092915050565b5f616d08616d0384614104565b6140ea565b905082815260208101848484011115616d2457616d23614078565b5b616d2f848285614d7f565b509392505050565b5f82601f830112616d4b57616d4a614074565b5b8151616d5b848260208601616cf6565b91505092915050565b5f8060408385031215616d7a57616d79613ec7565b5b5f83015167ffffffffffffffff811115616d9757616d96613ecb565b5b616da385828601616d37565b9250506020616db4858286016166a9565b9150509250929050565b7f414133332072657665727465643a200000000000000000000000000000000000815250565b5f616dee82616dbe565b600f82019150616dfe8284615fa4565b915081905092915050565b7f4141333320726576657274656420286f72204f4f4729000000000000000000005f82015250565b5f616e3d601683614a4f565b9150616e4882616e09565b602082019050919050565b5f604082019050616e665f83018461404c565b8181036020830152616e7781616e31565b905092915050565b5f61014082019050616e935f83018d615570565b616ea0602083018c61404c565b616ead604083018b614890565b616eba606083018a614890565b616ec7608083018961404c565b616ed460a083018861404c565b616ee160c083018761404c565b616eee60e083018661404c565b616efc61010083018561404c565b616f0a610120830184614890565b9b9a5050505050505050505050565b7f414131302073656e64657220616c726561647920636f6e7374727563746564005f82015250565b5f616f4d601f83614a4f565b9150616f5882616f19565b602082019050919050565b5f604082019050616f765f83018461404c565b8181036020830152616f8781616f41565b905092915050565b7f4141313320696e6974436f6465206661696c6564206f72204f4f4700000000005f82015250565b5f616fc3601b83614a4f565b9150616fce82616f8f565b602082019050919050565b5f604082019050616fec5f83018461404c565b8181036020830152616ffd81616fb7565b905092915050565b7f4141313420696e6974436f6465206d7573742072657475726e2073656e6465725f82015250565b5f617039602083614a4f565b915061704482617005565b602082019050919050565b5f6040820190506170625f83018461404c565b81810360208301526170738161702d565b905092915050565b7f4141313520696e6974436f6465206d757374206372656174652073656e6465725f82015250565b5f6170af602083614a4f565b91506170ba8261707b565b602082019050919050565b5f6040820190506170d85f83018461404c565b81810360208301526170e9816170a3565b905092915050565b5f6040820190506171045f830185615570565b6171116020830184615570565b939250505056fea26469706673582212201fbc8d4180c25cf161c374b14f1b878cfcbbd5db784395de9382d213937bdb6764736f6c63430008150033608060405234801561000f575f80fd5b506103238061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063570e1a361461002d575b5f80fd5b61004760048036038101906100429190610169565b61005d565b60405161005491906101f3565b60405180910390f35b5f8083835f9060149261007293929190610214565b9061007d919061028f565b60601c90505f8484601490809261009693929190610214565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090505f60205f8351602085015f875af190505f519350806100f7575f93505b50505092915050565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261012957610128610108565b5b8235905067ffffffffffffffff8111156101465761014561010c565b5b60208301915083600182028301111561016257610161610110565b5b9250929050565b5f806020838503121561017f5761017e610100565b5b5f83013567ffffffffffffffff81111561019c5761019b610104565b5b6101a885828601610114565b92509250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101dd826101b4565b9050919050565b6101ed816101d3565b82525050565b5f6020820190506102065f8301846101e4565b92915050565b5f80fd5b5f80fd5b5f80858511156102275761022661020c565b5b8386111561023857610237610210565b5b6001850283019150848603905094509492505050565b5f82905092915050565b5f7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b5f82821b905092915050565b5f61029a838361024e565b826102a58135610258565b925060148210156102e5576102e07fffffffffffffffffffffffffffffffffffffffff00000000000000000000000083601403600802610283565b831692505b50509291505056fea2646970667358221220bf1a621a7c417b48ce558164fbd2b327d47b688f9d188dd51a42b1e3113eee8764736f6c63430008150033",
}

// EntryPointABI is the input ABI used to generate the binding from.
// Deprecated: Use EntryPointMetaData.ABI instead.
var EntryPointABI = EntryPointMetaData.ABI

// EntryPointBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EntryPointMetaData.Bin instead.
var EntryPointBin = EntryPointMetaData.Bin

// DeployEntryPoint deploys a new Ethereum contract, binding an instance of EntryPoint to it.
func DeployEntryPoint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EntryPoint, error) {
	parsed, err := EntryPointMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EntryPointBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EntryPoint{EntryPointCaller: EntryPointCaller{contract: contract}, EntryPointTransactor: EntryPointTransactor{contract: contract}, EntryPointFilterer: EntryPointFilterer{contract: contract}}, nil
}

// EntryPoint is an auto generated Go binding around an Ethereum contract.
type EntryPoint struct {
	EntryPointCaller     // Read-only binding to the contract
	EntryPointTransactor // Write-only binding to the contract
	EntryPointFilterer   // Log filterer for contract events
}

// EntryPointCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntryPointCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntryPointTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntryPointFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryPointSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntryPointSession struct {
	Contract     *EntryPoint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntryPointCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntryPointCallerSession struct {
	Contract *EntryPointCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EntryPointTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntryPointTransactorSession struct {
	Contract     *EntryPointTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EntryPointRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntryPointRaw struct {
	Contract *EntryPoint // Generic contract binding to access the raw methods on
}

// EntryPointCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntryPointCallerRaw struct {
	Contract *EntryPointCaller // Generic read-only contract binding to access the raw methods on
}

// EntryPointTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntryPointTransactorRaw struct {
	Contract *EntryPointTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntryPoint creates a new instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPoint(address common.Address, backend bind.ContractBackend) (*EntryPoint, error) {
	contract, err := bindEntryPoint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EntryPoint{EntryPointCaller: EntryPointCaller{contract: contract}, EntryPointTransactor: EntryPointTransactor{contract: contract}, EntryPointFilterer: EntryPointFilterer{contract: contract}}, nil
}

// NewEntryPointCaller creates a new read-only instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPointCaller(address common.Address, caller bind.ContractCaller) (*EntryPointCaller, error) {
	contract, err := bindEntryPoint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntryPointCaller{contract: contract}, nil
}

// NewEntryPointTransactor creates a new write-only instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPointTransactor(address common.Address, transactor bind.ContractTransactor) (*EntryPointTransactor, error) {
	contract, err := bindEntryPoint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntryPointTransactor{contract: contract}, nil
}

// NewEntryPointFilterer creates a new log filterer instance of EntryPoint, bound to a specific deployed contract.
func NewEntryPointFilterer(address common.Address, filterer bind.ContractFilterer) (*EntryPointFilterer, error) {
	contract, err := bindEntryPoint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntryPointFilterer{contract: contract}, nil
}

// bindEntryPoint binds a generic wrapper to an already deployed contract.
func bindEntryPoint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EntryPointMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntryPoint *EntryPointRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntryPoint.Contract.EntryPointCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntryPoint *EntryPointRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.Contract.EntryPointTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntryPoint *EntryPointRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntryPoint.Contract.EntryPointTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntryPoint *EntryPointCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntryPoint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntryPoint *EntryPointTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntryPoint *EntryPointTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntryPoint.Contract.contract.Transact(opts, method, params...)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_EntryPoint *EntryPointCaller) SIGVALIDATIONFAILED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "SIG_VALIDATION_FAILED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_EntryPoint *EntryPointSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _EntryPoint.Contract.SIGVALIDATIONFAILED(&_EntryPoint.CallOpts)
}

// SIGVALIDATIONFAILED is a free data retrieval call binding the contract method 0x8f41ec5a.
//
// Solidity: function SIG_VALIDATION_FAILED() view returns(uint256)
func (_EntryPoint *EntryPointCallerSession) SIGVALIDATIONFAILED() (*big.Int, error) {
	return _EntryPoint.Contract.SIGVALIDATIONFAILED(&_EntryPoint.CallOpts)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_EntryPoint *EntryPointCaller) ValidateSenderAndPaymaster(opts *bind.CallOpts, initCode []byte, sender common.Address, paymasterAndData []byte) error {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "_validateSenderAndPaymaster", initCode, sender, paymasterAndData)

	if err != nil {
		return err
	}

	return err

}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_EntryPoint *EntryPointSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _EntryPoint.Contract.ValidateSenderAndPaymaster(&_EntryPoint.CallOpts, initCode, sender, paymasterAndData)
}

// ValidateSenderAndPaymaster is a free data retrieval call binding the contract method 0x957122ab.
//
// Solidity: function _validateSenderAndPaymaster(bytes initCode, address sender, bytes paymasterAndData) view returns()
func (_EntryPoint *EntryPointCallerSession) ValidateSenderAndPaymaster(initCode []byte, sender common.Address, paymasterAndData []byte) error {
	return _EntryPoint.Contract.ValidateSenderAndPaymaster(&_EntryPoint.CallOpts, initCode, sender, paymasterAndData)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPoint *EntryPointCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPoint *EntryPointSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.BalanceOf(&_EntryPoint.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EntryPoint *EntryPointCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EntryPoint.Contract.BalanceOf(&_EntryPoint.CallOpts, account)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPoint *EntryPointCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Deposit         *big.Int
		Staked          bool
		Stake           *big.Int
		UnstakeDelaySec uint32
		WithdrawTime    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Staked = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnstakeDelaySec = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.WithdrawTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPoint *EntryPointSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _EntryPoint.Contract.Deposits(&_EntryPoint.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint112 deposit, bool staked, uint112 stake, uint32 unstakeDelaySec, uint48 withdrawTime)
func (_EntryPoint *EntryPointCallerSession) Deposits(arg0 common.Address) (struct {
	Deposit         *big.Int
	Staked          bool
	Stake           *big.Int
	UnstakeDelaySec uint32
	WithdrawTime    *big.Int
}, error) {
	return _EntryPoint.Contract.Deposits(&_EntryPoint.CallOpts, arg0)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_EntryPoint *EntryPointCaller) GetDepositInfo(opts *bind.CallOpts, account common.Address) (IStakeManagerDepositInfo, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getDepositInfo", account)

	if err != nil {
		return *new(IStakeManagerDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerDepositInfo)).(*IStakeManagerDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_EntryPoint *EntryPointSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _EntryPoint.Contract.GetDepositInfo(&_EntryPoint.CallOpts, account)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5287ce12.
//
// Solidity: function getDepositInfo(address account) view returns((uint112,bool,uint112,uint32,uint48) info)
func (_EntryPoint *EntryPointCallerSession) GetDepositInfo(account common.Address) (IStakeManagerDepositInfo, error) {
	return _EntryPoint.Contract.GetDepositInfo(&_EntryPoint.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPoint *EntryPointCaller) GetNonce(opts *bind.CallOpts, sender common.Address, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getNonce", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPoint *EntryPointSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.GetNonce(&_EntryPoint.CallOpts, sender, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x35567e1a.
//
// Solidity: function getNonce(address sender, uint192 key) view returns(uint256 nonce)
func (_EntryPoint *EntryPointCallerSession) GetNonce(sender common.Address, key *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.GetNonce(&_EntryPoint.CallOpts, sender, key)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointCaller) GetUserOpHash(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "getUserOpHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointSession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// GetUserOpHash is a free data retrieval call binding the contract method 0xa6193531.
//
// Solidity: function getUserOpHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) view returns(bytes32)
func (_EntryPoint *EntryPointCallerSession) GetUserOpHash(userOp UserOperation) ([32]byte, error) {
	return _EntryPoint.Contract.GetUserOpHash(&_EntryPoint.CallOpts, userOp)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPoint *EntryPointCaller) NonceSequenceNumber(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EntryPoint.contract.Call(opts, &out, "nonceSequenceNumber", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPoint *EntryPointSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.NonceSequenceNumber(&_EntryPoint.CallOpts, arg0, arg1)
}

// NonceSequenceNumber is a free data retrieval call binding the contract method 0x1b2e01b8.
//
// Solidity: function nonceSequenceNumber(address , uint192 ) view returns(uint256)
func (_EntryPoint *EntryPointCallerSession) NonceSequenceNumber(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _EntryPoint.Contract.NonceSequenceNumber(&_EntryPoint.CallOpts, arg0, arg1)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPoint *EntryPointTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPoint *EntryPointSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddStake(&_EntryPoint.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_EntryPoint *EntryPointTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _EntryPoint.Contract.AddStake(&_EntryPoint.TransactOpts, unstakeDelaySec)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_EntryPoint *EntryPointTransactor) DepositTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "depositTo", account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_EntryPoint *EntryPointSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.DepositTo(&_EntryPoint.TransactOpts, account)
}

// DepositTo is a paid mutator transaction binding the contract method 0xb760faf9.
//
// Solidity: function depositTo(address account) payable returns()
func (_EntryPoint *EntryPointTransactorSession) DepositTo(account common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.DepositTo(&_EntryPoint.TransactOpts, account)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPoint *EntryPointTransactor) GetSenderAddress(opts *bind.TransactOpts, initCode []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "getSenderAddress", initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPoint *EntryPointSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.GetSenderAddress(&_EntryPoint.TransactOpts, initCode)
}

// GetSenderAddress is a paid mutator transaction binding the contract method 0x9b249f69.
//
// Solidity: function getSenderAddress(bytes initCode) returns()
func (_EntryPoint *EntryPointTransactorSession) GetSenderAddress(initCode []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.GetSenderAddress(&_EntryPoint.TransactOpts, initCode)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactor) HandleAggregatedOps(opts *bind.TransactOpts, opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleAggregatedOps", opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPoint *EntryPointSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleAggregatedOps(&_EntryPoint.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleAggregatedOps is a paid mutator transaction binding the contract method 0x4b1d7cf5.
//
// Solidity: function handleAggregatedOps(((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[],address,bytes)[] opsPerAggregator, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleAggregatedOps(opsPerAggregator []IEntryPointUserOpsPerAggregator, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleAggregatedOps(&_EntryPoint.TransactOpts, opsPerAggregator, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactor) HandleOps(opts *bind.TransactOpts, ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "handleOps", ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_EntryPoint *EntryPointSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary)
}

// HandleOps is a paid mutator transaction binding the contract method 0x1fad948c.
//
// Solidity: function handleOps((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] ops, address beneficiary) returns()
func (_EntryPoint *EntryPointTransactorSession) HandleOps(ops []UserOperation, beneficiary common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.HandleOps(&_EntryPoint.TransactOpts, ops, beneficiary)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPoint *EntryPointTransactor) IncrementNonce(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "incrementNonce", key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPoint *EntryPointSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.IncrementNonce(&_EntryPoint.TransactOpts, key)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x0bd28e3b.
//
// Solidity: function incrementNonce(uint192 key) returns()
func (_EntryPoint *EntryPointTransactorSession) IncrementNonce(key *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.IncrementNonce(&_EntryPoint.TransactOpts, key)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactor) InnerHandleOp(opts *bind.TransactOpts, callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "innerHandleOp", callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// InnerHandleOp is a paid mutator transaction binding the contract method 0x1d732756.
//
// Solidity: function innerHandleOp(bytes callData, ((address,uint256,uint256,uint256,uint256,address,uint256,uint256),bytes32,uint256,uint256,uint256) opInfo, bytes context) returns(uint256 actualGasCost)
func (_EntryPoint *EntryPointTransactorSession) InnerHandleOp(callData []byte, opInfo EntryPointUserOpInfo, context []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.InnerHandleOp(&_EntryPoint.TransactOpts, callData, opInfo, context)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_EntryPoint *EntryPointTransactor) SimulateHandleOp(opts *bind.TransactOpts, op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "simulateHandleOp", op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_EntryPoint *EntryPointSession) SimulateHandleOp(op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.SimulateHandleOp(&_EntryPoint.TransactOpts, op, target, targetCallData)
}

// SimulateHandleOp is a paid mutator transaction binding the contract method 0xd6383f94.
//
// Solidity: function simulateHandleOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) op, address target, bytes targetCallData) returns()
func (_EntryPoint *EntryPointTransactorSession) SimulateHandleOp(op UserOperation, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _EntryPoint.Contract.SimulateHandleOp(&_EntryPoint.TransactOpts, op, target, targetCallData)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_EntryPoint *EntryPointTransactor) SimulateValidation(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "simulateValidation", userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_EntryPoint *EntryPointSession) SimulateValidation(userOp UserOperation) (*types.Transaction, error) {
	return _EntryPoint.Contract.SimulateValidation(&_EntryPoint.TransactOpts, userOp)
}

// SimulateValidation is a paid mutator transaction binding the contract method 0xee219423.
//
// Solidity: function simulateValidation((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp) returns()
func (_EntryPoint *EntryPointTransactorSession) SimulateValidation(userOp UserOperation) (*types.Transaction, error) {
	return _EntryPoint.Contract.SimulateValidation(&_EntryPoint.TransactOpts, userOp)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPoint *EntryPointTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPoint *EntryPointSession) UnlockStake() (*types.Transaction, error) {
	return _EntryPoint.Contract.UnlockStake(&_EntryPoint.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_EntryPoint *EntryPointTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _EntryPoint.Contract.UnlockStake(&_EntryPoint.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPoint *EntryPointTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPoint *EntryPointSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.WithdrawStake(&_EntryPoint.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_EntryPoint *EntryPointTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _EntryPoint.Contract.WithdrawStake(&_EntryPoint.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_EntryPoint *EntryPointTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.contract.Transact(opts, "withdrawTo", withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_EntryPoint *EntryPointSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.WithdrawTo(&_EntryPoint.TransactOpts, withdrawAddress, withdrawAmount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 withdrawAmount) returns()
func (_EntryPoint *EntryPointTransactorSession) WithdrawTo(withdrawAddress common.Address, withdrawAmount *big.Int) (*types.Transaction, error) {
	return _EntryPoint.Contract.WithdrawTo(&_EntryPoint.TransactOpts, withdrawAddress, withdrawAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EntryPoint *EntryPointTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntryPoint.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EntryPoint *EntryPointSession) Receive() (*types.Transaction, error) {
	return _EntryPoint.Contract.Receive(&_EntryPoint.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EntryPoint *EntryPointTransactorSession) Receive() (*types.Transaction, error) {
	return _EntryPoint.Contract.Receive(&_EntryPoint.TransactOpts)
}

// EntryPointAccountDeployedIterator is returned from FilterAccountDeployed and is used to iterate over the raw logs and unpacked data for AccountDeployed events raised by the EntryPoint contract.
type EntryPointAccountDeployedIterator struct {
	Event *EntryPointAccountDeployed // Event containing the contract specifics and raw log

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
func (it *EntryPointAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointAccountDeployed)
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
		it.Event = new(EntryPointAccountDeployed)
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
func (it *EntryPointAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointAccountDeployed represents a AccountDeployed event raised by the EntryPoint contract.
type EntryPointAccountDeployed struct {
	UserOpHash [32]byte
	Sender     common.Address
	Factory    common.Address
	Paymaster  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDeployed is a free log retrieval operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPoint *EntryPointFilterer) FilterAccountDeployed(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointAccountDeployedIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointAccountDeployedIterator{contract: _EntryPoint.contract, event: "AccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAccountDeployed is a free log subscription operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPoint *EntryPointFilterer) WatchAccountDeployed(opts *bind.WatchOpts, sink chan<- *EntryPointAccountDeployed, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "AccountDeployed", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointAccountDeployed)
				if err := _EntryPoint.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
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

// ParseAccountDeployed is a log parse operation binding the contract event 0xd51a9c61267aa6196961883ecf5ff2da6619c37dac0fa92122513fb32c032d2d.
//
// Solidity: event AccountDeployed(bytes32 indexed userOpHash, address indexed sender, address factory, address paymaster)
func (_EntryPoint *EntryPointFilterer) ParseAccountDeployed(log types.Log) (*EntryPointAccountDeployed, error) {
	event := new(EntryPointAccountDeployed)
	if err := _EntryPoint.contract.UnpackLog(event, "AccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointBeforeExecutionIterator is returned from FilterBeforeExecution and is used to iterate over the raw logs and unpacked data for BeforeExecution events raised by the EntryPoint contract.
type EntryPointBeforeExecutionIterator struct {
	Event *EntryPointBeforeExecution // Event containing the contract specifics and raw log

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
func (it *EntryPointBeforeExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointBeforeExecution)
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
		it.Event = new(EntryPointBeforeExecution)
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
func (it *EntryPointBeforeExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointBeforeExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointBeforeExecution represents a BeforeExecution event raised by the EntryPoint contract.
type EntryPointBeforeExecution struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeforeExecution is a free log retrieval operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPoint *EntryPointFilterer) FilterBeforeExecution(opts *bind.FilterOpts) (*EntryPointBeforeExecutionIterator, error) {

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return &EntryPointBeforeExecutionIterator{contract: _EntryPoint.contract, event: "BeforeExecution", logs: logs, sub: sub}, nil
}

// WatchBeforeExecution is a free log subscription operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPoint *EntryPointFilterer) WatchBeforeExecution(opts *bind.WatchOpts, sink chan<- *EntryPointBeforeExecution) (event.Subscription, error) {

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "BeforeExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointBeforeExecution)
				if err := _EntryPoint.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
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

// ParseBeforeExecution is a log parse operation binding the contract event 0xbb47ee3e183a558b1a2ff0874b079f3fc5478b7454eacf2bfc5af2ff5878f972.
//
// Solidity: event BeforeExecution()
func (_EntryPoint *EntryPointFilterer) ParseBeforeExecution(log types.Log) (*EntryPointBeforeExecution, error) {
	event := new(EntryPointBeforeExecution)
	if err := _EntryPoint.contract.UnpackLog(event, "BeforeExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the EntryPoint contract.
type EntryPointDepositedIterator struct {
	Event *EntryPointDeposited // Event containing the contract specifics and raw log

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
func (it *EntryPointDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointDeposited)
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
		it.Event = new(EntryPointDeposited)
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
func (it *EntryPointDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointDeposited represents a Deposited event raised by the EntryPoint contract.
type EntryPointDeposited struct {
	Account      common.Address
	TotalDeposit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPoint *EntryPointFilterer) FilterDeposited(opts *bind.FilterOpts, account []common.Address) (*EntryPointDepositedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointDepositedIterator{contract: _EntryPoint.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPoint *EntryPointFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EntryPointDeposited, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "Deposited", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointDeposited)
				if err := _EntryPoint.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed account, uint256 totalDeposit)
func (_EntryPoint *EntryPointFilterer) ParseDeposited(log types.Log) (*EntryPointDeposited, error) {
	event := new(EntryPointDeposited)
	if err := _EntryPoint.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointSignatureAggregatorChangedIterator is returned from FilterSignatureAggregatorChanged and is used to iterate over the raw logs and unpacked data for SignatureAggregatorChanged events raised by the EntryPoint contract.
type EntryPointSignatureAggregatorChangedIterator struct {
	Event *EntryPointSignatureAggregatorChanged // Event containing the contract specifics and raw log

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
func (it *EntryPointSignatureAggregatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointSignatureAggregatorChanged)
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
		it.Event = new(EntryPointSignatureAggregatorChanged)
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
func (it *EntryPointSignatureAggregatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointSignatureAggregatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointSignatureAggregatorChanged represents a SignatureAggregatorChanged event raised by the EntryPoint contract.
type EntryPointSignatureAggregatorChanged struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignatureAggregatorChanged is a free log retrieval operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPoint *EntryPointFilterer) FilterSignatureAggregatorChanged(opts *bind.FilterOpts, aggregator []common.Address) (*EntryPointSignatureAggregatorChangedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointSignatureAggregatorChangedIterator{contract: _EntryPoint.contract, event: "SignatureAggregatorChanged", logs: logs, sub: sub}, nil
}

// WatchSignatureAggregatorChanged is a free log subscription operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPoint *EntryPointFilterer) WatchSignatureAggregatorChanged(opts *bind.WatchOpts, sink chan<- *EntryPointSignatureAggregatorChanged, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "SignatureAggregatorChanged", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointSignatureAggregatorChanged)
				if err := _EntryPoint.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
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

// ParseSignatureAggregatorChanged is a log parse operation binding the contract event 0x575ff3acadd5ab348fe1855e217e0f3678f8d767d7494c9f9fefbee2e17cca4d.
//
// Solidity: event SignatureAggregatorChanged(address indexed aggregator)
func (_EntryPoint *EntryPointFilterer) ParseSignatureAggregatorChanged(log types.Log) (*EntryPointSignatureAggregatorChanged, error) {
	event := new(EntryPointSignatureAggregatorChanged)
	if err := _EntryPoint.contract.UnpackLog(event, "SignatureAggregatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the EntryPoint contract.
type EntryPointStakeLockedIterator struct {
	Event *EntryPointStakeLocked // Event containing the contract specifics and raw log

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
func (it *EntryPointStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointStakeLocked)
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
		it.Event = new(EntryPointStakeLocked)
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
func (it *EntryPointStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointStakeLocked represents a StakeLocked event raised by the EntryPoint contract.
type EntryPointStakeLocked struct {
	Account         common.Address
	TotalStaked     *big.Int
	UnstakeDelaySec *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPoint *EntryPointFilterer) FilterStakeLocked(opts *bind.FilterOpts, account []common.Address) (*EntryPointStakeLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointStakeLockedIterator{contract: _EntryPoint.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPoint *EntryPointFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *EntryPointStakeLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "StakeLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointStakeLocked)
				if err := _EntryPoint.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed account, uint256 totalStaked, uint256 unstakeDelaySec)
func (_EntryPoint *EntryPointFilterer) ParseStakeLocked(log types.Log) (*EntryPointStakeLocked, error) {
	event := new(EntryPointStakeLocked)
	if err := _EntryPoint.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointStakeUnlockedIterator is returned from FilterStakeUnlocked and is used to iterate over the raw logs and unpacked data for StakeUnlocked events raised by the EntryPoint contract.
type EntryPointStakeUnlockedIterator struct {
	Event *EntryPointStakeUnlocked // Event containing the contract specifics and raw log

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
func (it *EntryPointStakeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointStakeUnlocked)
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
		it.Event = new(EntryPointStakeUnlocked)
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
func (it *EntryPointStakeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointStakeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointStakeUnlocked represents a StakeUnlocked event raised by the EntryPoint contract.
type EntryPointStakeUnlocked struct {
	Account      common.Address
	WithdrawTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakeUnlocked is a free log retrieval operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPoint *EntryPointFilterer) FilterStakeUnlocked(opts *bind.FilterOpts, account []common.Address) (*EntryPointStakeUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointStakeUnlockedIterator{contract: _EntryPoint.contract, event: "StakeUnlocked", logs: logs, sub: sub}, nil
}

// WatchStakeUnlocked is a free log subscription operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPoint *EntryPointFilterer) WatchStakeUnlocked(opts *bind.WatchOpts, sink chan<- *EntryPointStakeUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "StakeUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointStakeUnlocked)
				if err := _EntryPoint.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
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

// ParseStakeUnlocked is a log parse operation binding the contract event 0xfa9b3c14cc825c412c9ed81b3ba365a5b459439403f18829e572ed53a4180f0a.
//
// Solidity: event StakeUnlocked(address indexed account, uint256 withdrawTime)
func (_EntryPoint *EntryPointFilterer) ParseStakeUnlocked(log types.Log) (*EntryPointStakeUnlocked, error) {
	event := new(EntryPointStakeUnlocked)
	if err := _EntryPoint.contract.UnpackLog(event, "StakeUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the EntryPoint contract.
type EntryPointStakeWithdrawnIterator struct {
	Event *EntryPointStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *EntryPointStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointStakeWithdrawn)
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
		it.Event = new(EntryPointStakeWithdrawn)
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
func (it *EntryPointStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointStakeWithdrawn represents a StakeWithdrawn event raised by the EntryPoint contract.
type EntryPointStakeWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, account []common.Address) (*EntryPointStakeWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointStakeWithdrawnIterator{contract: _EntryPoint.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *EntryPointStakeWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "StakeWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointStakeWithdrawn)
				if err := _EntryPoint.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0xb7c918e0e249f999e965cafeb6c664271b3f4317d296461500e71da39f0cbda3.
//
// Solidity: event StakeWithdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) ParseStakeWithdrawn(log types.Log) (*EntryPointStakeWithdrawn, error) {
	event := new(EntryPointStakeWithdrawn)
	if err := _EntryPoint.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointUserOperationEventIterator is returned from FilterUserOperationEvent and is used to iterate over the raw logs and unpacked data for UserOperationEvent events raised by the EntryPoint contract.
type EntryPointUserOperationEventIterator struct {
	Event *EntryPointUserOperationEvent // Event containing the contract specifics and raw log

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
func (it *EntryPointUserOperationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointUserOperationEvent)
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
		it.Event = new(EntryPointUserOperationEvent)
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
func (it *EntryPointUserOperationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointUserOperationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointUserOperationEvent represents a UserOperationEvent event raised by the EntryPoint contract.
type EntryPointUserOperationEvent struct {
	UserOpHash    [32]byte
	Sender        common.Address
	Paymaster     common.Address
	Nonce         *big.Int
	Success       bool
	ActualGasCost *big.Int
	ActualGasUsed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserOperationEvent is a free log retrieval operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) FilterUserOperationEvent(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (*EntryPointUserOperationEventIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointUserOperationEventIterator{contract: _EntryPoint.contract, event: "UserOperationEvent", logs: logs, sub: sub}, nil
}

// WatchUserOperationEvent is a free log subscription operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) WatchUserOperationEvent(opts *bind.WatchOpts, sink chan<- *EntryPointUserOperationEvent, userOpHash [][32]byte, sender []common.Address, paymaster []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var paymasterRule []interface{}
	for _, paymasterItem := range paymaster {
		paymasterRule = append(paymasterRule, paymasterItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "UserOperationEvent", userOpHashRule, senderRule, paymasterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointUserOperationEvent)
				if err := _EntryPoint.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
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

// ParseUserOperationEvent is a log parse operation binding the contract event 0x49628fd1471006c1482da88028e9ce4dbb080b815c9b0344d39e5a8e6ec1419f.
//
// Solidity: event UserOperationEvent(bytes32 indexed userOpHash, address indexed sender, address indexed paymaster, uint256 nonce, bool success, uint256 actualGasCost, uint256 actualGasUsed)
func (_EntryPoint *EntryPointFilterer) ParseUserOperationEvent(log types.Log) (*EntryPointUserOperationEvent, error) {
	event := new(EntryPointUserOperationEvent)
	if err := _EntryPoint.contract.UnpackLog(event, "UserOperationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointUserOperationRevertReasonIterator is returned from FilterUserOperationRevertReason and is used to iterate over the raw logs and unpacked data for UserOperationRevertReason events raised by the EntryPoint contract.
type EntryPointUserOperationRevertReasonIterator struct {
	Event *EntryPointUserOperationRevertReason // Event containing the contract specifics and raw log

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
func (it *EntryPointUserOperationRevertReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointUserOperationRevertReason)
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
		it.Event = new(EntryPointUserOperationRevertReason)
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
func (it *EntryPointUserOperationRevertReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointUserOperationRevertReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointUserOperationRevertReason represents a UserOperationRevertReason event raised by the EntryPoint contract.
type EntryPointUserOperationRevertReason struct {
	UserOpHash   [32]byte
	Sender       common.Address
	Nonce        *big.Int
	RevertReason []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserOperationRevertReason is a free log retrieval operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) FilterUserOperationRevertReason(opts *bind.FilterOpts, userOpHash [][32]byte, sender []common.Address) (*EntryPointUserOperationRevertReasonIterator, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointUserOperationRevertReasonIterator{contract: _EntryPoint.contract, event: "UserOperationRevertReason", logs: logs, sub: sub}, nil
}

// WatchUserOperationRevertReason is a free log subscription operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) WatchUserOperationRevertReason(opts *bind.WatchOpts, sink chan<- *EntryPointUserOperationRevertReason, userOpHash [][32]byte, sender []common.Address) (event.Subscription, error) {

	var userOpHashRule []interface{}
	for _, userOpHashItem := range userOpHash {
		userOpHashRule = append(userOpHashRule, userOpHashItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "UserOperationRevertReason", userOpHashRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointUserOperationRevertReason)
				if err := _EntryPoint.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
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

// ParseUserOperationRevertReason is a log parse operation binding the contract event 0x1c4fada7374c0a9ee8841fc38afe82932dc0f8e69012e927f061a8bae611a201.
//
// Solidity: event UserOperationRevertReason(bytes32 indexed userOpHash, address indexed sender, uint256 nonce, bytes revertReason)
func (_EntryPoint *EntryPointFilterer) ParseUserOperationRevertReason(log types.Log) (*EntryPointUserOperationRevertReason, error) {
	event := new(EntryPointUserOperationRevertReason)
	if err := _EntryPoint.contract.UnpackLog(event, "UserOperationRevertReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntryPointWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the EntryPoint contract.
type EntryPointWithdrawnIterator struct {
	Event *EntryPointWithdrawn // Event containing the contract specifics and raw log

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
func (it *EntryPointWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntryPointWithdrawn)
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
		it.Event = new(EntryPointWithdrawn)
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
func (it *EntryPointWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntryPointWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntryPointWithdrawn represents a Withdrawn event raised by the EntryPoint contract.
type EntryPointWithdrawn struct {
	Account         common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) FilterWithdrawn(opts *bind.FilterOpts, account []common.Address) (*EntryPointWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.FilterLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &EntryPointWithdrawnIterator{contract: _EntryPoint.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *EntryPointWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EntryPoint.contract.WatchLogs(opts, "Withdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntryPointWithdrawn)
				if err := _EntryPoint.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd1c19fbcd4551a5edfb66d43d2e337c04837afda3482b42bdf569a8fccdae5fb.
//
// Solidity: event Withdrawn(address indexed account, address withdrawAddress, uint256 amount)
func (_EntryPoint *EntryPointFilterer) ParseWithdrawn(log types.Log) (*EntryPointWithdrawn, error) {
	event := new(EntryPointWithdrawn)
	if err := _EntryPoint.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
