// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package superform

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

// InitSingleVaultData is an auto generated low-level Go binding around an user-defined struct.
type InitSingleVaultData struct {
	PayloadId       *big.Int
	SuperformId     *big.Int
	Amount          *big.Int
	OutputAmount    *big.Int
	MaxSlippage     *big.Int
	LiqData         LiqRequest
	HasDstSwap      bool
	Retain4626      bool
	ReceiverAddress common.Address
	ExtraFormData   []byte
}

// LiqRequest is an auto generated low-level Go binding around an user-defined struct.
type LiqRequest struct {
	TxData        []byte
	Token         common.Address
	InterimToken  common.Address
	BridgeId      uint8
	LiqDstChainId uint64
	NativeAmount  *big.Int
}

// SuperformMetaData contains all meta data concerning the Superform contract.
var SuperformMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"superRegistry_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BLOCK_CHAIN_ID_OUT_OF_BOUNDS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CANNOT_FORWARD_4646_TOKEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DIFFERENT_TOKENS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DIRECT_DEPOSIT_SWAP_FAILED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DIRECT_WITHDRAW_INVALID_LIQ_REQUEST\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"FAILED_TO_EXECUTE_TXDATA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_ALLOWANCE_FOR_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_BALANCE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INSUFFICIENT_NATIVE_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_CHAIN_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_CORE_STATE_REGISTRY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_EMERGENCY_QUEUE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_SUPERFORM_ROUTER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NOT_SUPER_REGISTRY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PAUSED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SUPERFORM_ID_NONEXISTENT\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VAULT_IMPLEMENTATION_FAILED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WITHDRAW_TOKEN_NOT_UPDATED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WITHDRAW_TX_DATA_NOT_UPDATED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WITHDRAW_ZERO_COLLATERAL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XCHAIN_WITHDRAW_INVALID_LIQ_REQUEST\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_AMOUNT\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdrawalProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FormDustForwardedToPaymaster\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"srcChainID\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcPayloadId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"Processed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"VaultAdded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"payloadId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superformId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSlippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interimToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"bridgeId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"liqDstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nativeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLiqRequest\",\"name\":\"liqData\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"hasDstSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retain4626\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraFormData\",\"type\":\"bytes\"}],\"internalType\":\"structInitSingleVaultData\",\"name\":\"singleVaultData_\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"srcSender_\",\"type\":\"address\"}],\"name\":\"directDepositIntoVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"payloadId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superformId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSlippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interimToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"bridgeId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"liqDstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nativeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLiqRequest\",\"name\":\"liqData\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"hasDstSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retain4626\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraFormData\",\"type\":\"bytes\"}],\"internalType\":\"structInitSingleVaultData\",\"name\":\"singleVaultData_\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"srcSender_\",\"type\":\"address\"}],\"name\":\"directWithdrawFromVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"forwardDustToPaymaster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPreviewPricePerVaultShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricePerVaultShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateRegistryId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultShareBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"superRegistry_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"previewDepositTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"previewRedeemFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"previewWithdrawFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superRegistry\",\"outputs\":[{\"internalType\":\"contractISuperRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superformYieldTokenName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superformYieldTokenSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId_\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"payloadId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superformId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSlippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interimToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"bridgeId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"liqDstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nativeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLiqRequest\",\"name\":\"liqData\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"hasDstSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retain4626\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraFormData\",\"type\":\"bytes\"}],\"internalType\":\"structInitSingleVaultData\",\"name\":\"singleVaultData_\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"srcSender_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId_\",\"type\":\"uint64\"}],\"name\":\"xChainDepositIntoVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"payloadId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superformId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSlippage\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"txData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interimToken\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"bridgeId\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"liqDstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nativeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLiqRequest\",\"name\":\"liqData\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"hasDstSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"retain4626\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraFormData\",\"type\":\"bytes\"}],\"internalType\":\"structInitSingleVaultData\",\"name\":\"singleVaultData_\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"srcSender_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId_\",\"type\":\"uint64\"}],\"name\":\"xChainWithdrawFromVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SuperformABI is the input ABI used to generate the binding from.
// Deprecated: Use SuperformMetaData.ABI instead.
var SuperformABI = SuperformMetaData.ABI

// Superform is an auto generated Go binding around an Ethereum contract.
type Superform struct {
	SuperformCaller     // Read-only binding to the contract
	SuperformTransactor // Write-only binding to the contract
	SuperformFilterer   // Log filterer for contract events
}

// SuperformCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuperformCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperformTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuperformTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperformFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuperformFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperformSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuperformSession struct {
	Contract     *Superform        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuperformCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuperformCallerSession struct {
	Contract *SuperformCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SuperformTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuperformTransactorSession struct {
	Contract     *SuperformTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SuperformRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuperformRaw struct {
	Contract *Superform // Generic contract binding to access the raw methods on
}

// SuperformCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuperformCallerRaw struct {
	Contract *SuperformCaller // Generic read-only contract binding to access the raw methods on
}

// SuperformTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuperformTransactorRaw struct {
	Contract *SuperformTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuperform creates a new instance of Superform, bound to a specific deployed contract.
func NewSuperform(address common.Address, backend bind.ContractBackend) (*Superform, error) {
	contract, err := bindSuperform(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Superform{SuperformCaller: SuperformCaller{contract: contract}, SuperformTransactor: SuperformTransactor{contract: contract}, SuperformFilterer: SuperformFilterer{contract: contract}}, nil
}

// NewSuperformCaller creates a new read-only instance of Superform, bound to a specific deployed contract.
func NewSuperformCaller(address common.Address, caller bind.ContractCaller) (*SuperformCaller, error) {
	contract, err := bindSuperform(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuperformCaller{contract: contract}, nil
}

// NewSuperformTransactor creates a new write-only instance of Superform, bound to a specific deployed contract.
func NewSuperformTransactor(address common.Address, transactor bind.ContractTransactor) (*SuperformTransactor, error) {
	contract, err := bindSuperform(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuperformTransactor{contract: contract}, nil
}

// NewSuperformFilterer creates a new log filterer instance of Superform, bound to a specific deployed contract.
func NewSuperformFilterer(address common.Address, filterer bind.ContractFilterer) (*SuperformFilterer, error) {
	contract, err := bindSuperform(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuperformFilterer{contract: contract}, nil
}

// bindSuperform binds a generic wrapper to an already deployed contract.
func bindSuperform(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SuperformMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Superform *SuperformRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Superform.Contract.SuperformCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Superform *SuperformRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Superform.Contract.SuperformTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Superform *SuperformRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Superform.Contract.SuperformTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Superform *SuperformCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Superform.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Superform *SuperformTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Superform.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Superform *SuperformTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Superform.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint64)
func (_Superform *SuperformCaller) CHAINID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint64)
func (_Superform *SuperformSession) CHAINID() (uint64, error) {
	return _Superform.Contract.CHAINID(&_Superform.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(uint64)
func (_Superform *SuperformCallerSession) CHAINID() (uint64, error) {
	return _Superform.Contract.CHAINID(&_Superform.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Superform *SuperformCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Superform *SuperformSession) Asset() (common.Address, error) {
	return _Superform.Contract.Asset(&_Superform.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Superform *SuperformCallerSession) Asset() (common.Address, error) {
	return _Superform.Contract.Asset(&_Superform.CallOpts)
}

// GetPreviewPricePerVaultShare is a free data retrieval call binding the contract method 0x38d92fd5.
//
// Solidity: function getPreviewPricePerVaultShare() view returns(uint256)
func (_Superform *SuperformCaller) GetPreviewPricePerVaultShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getPreviewPricePerVaultShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviewPricePerVaultShare is a free data retrieval call binding the contract method 0x38d92fd5.
//
// Solidity: function getPreviewPricePerVaultShare() view returns(uint256)
func (_Superform *SuperformSession) GetPreviewPricePerVaultShare() (*big.Int, error) {
	return _Superform.Contract.GetPreviewPricePerVaultShare(&_Superform.CallOpts)
}

// GetPreviewPricePerVaultShare is a free data retrieval call binding the contract method 0x38d92fd5.
//
// Solidity: function getPreviewPricePerVaultShare() view returns(uint256)
func (_Superform *SuperformCallerSession) GetPreviewPricePerVaultShare() (*big.Int, error) {
	return _Superform.Contract.GetPreviewPricePerVaultShare(&_Superform.CallOpts)
}

// GetPricePerVaultShare is a free data retrieval call binding the contract method 0xff5f3e48.
//
// Solidity: function getPricePerVaultShare() view returns(uint256)
func (_Superform *SuperformCaller) GetPricePerVaultShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getPricePerVaultShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerVaultShare is a free data retrieval call binding the contract method 0xff5f3e48.
//
// Solidity: function getPricePerVaultShare() view returns(uint256)
func (_Superform *SuperformSession) GetPricePerVaultShare() (*big.Int, error) {
	return _Superform.Contract.GetPricePerVaultShare(&_Superform.CallOpts)
}

// GetPricePerVaultShare is a free data retrieval call binding the contract method 0xff5f3e48.
//
// Solidity: function getPricePerVaultShare() view returns(uint256)
func (_Superform *SuperformCallerSession) GetPricePerVaultShare() (*big.Int, error) {
	return _Superform.Contract.GetPricePerVaultShare(&_Superform.CallOpts)
}

// GetStateRegistryId is a free data retrieval call binding the contract method 0x91deb882.
//
// Solidity: function getStateRegistryId() view returns(uint8)
func (_Superform *SuperformCaller) GetStateRegistryId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getStateRegistryId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStateRegistryId is a free data retrieval call binding the contract method 0x91deb882.
//
// Solidity: function getStateRegistryId() view returns(uint8)
func (_Superform *SuperformSession) GetStateRegistryId() (uint8, error) {
	return _Superform.Contract.GetStateRegistryId(&_Superform.CallOpts)
}

// GetStateRegistryId is a free data retrieval call binding the contract method 0x91deb882.
//
// Solidity: function getStateRegistryId() view returns(uint8)
func (_Superform *SuperformCallerSession) GetStateRegistryId() (uint8, error) {
	return _Superform.Contract.GetStateRegistryId(&_Superform.CallOpts)
}

// GetTotalAssets is a free data retrieval call binding the contract method 0x6e07302b.
//
// Solidity: function getTotalAssets() view returns(uint256)
func (_Superform *SuperformCaller) GetTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalAssets is a free data retrieval call binding the contract method 0x6e07302b.
//
// Solidity: function getTotalAssets() view returns(uint256)
func (_Superform *SuperformSession) GetTotalAssets() (*big.Int, error) {
	return _Superform.Contract.GetTotalAssets(&_Superform.CallOpts)
}

// GetTotalAssets is a free data retrieval call binding the contract method 0x6e07302b.
//
// Solidity: function getTotalAssets() view returns(uint256)
func (_Superform *SuperformCallerSession) GetTotalAssets() (*big.Int, error) {
	return _Superform.Contract.GetTotalAssets(&_Superform.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Superform *SuperformCaller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Superform *SuperformSession) GetTotalSupply() (*big.Int, error) {
	return _Superform.Contract.GetTotalSupply(&_Superform.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_Superform *SuperformCallerSession) GetTotalSupply() (*big.Int, error) {
	return _Superform.Contract.GetTotalSupply(&_Superform.CallOpts)
}

// GetVaultAddress is a free data retrieval call binding the contract method 0x65cacaa4.
//
// Solidity: function getVaultAddress() view returns(address)
func (_Superform *SuperformCaller) GetVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultAddress is a free data retrieval call binding the contract method 0x65cacaa4.
//
// Solidity: function getVaultAddress() view returns(address)
func (_Superform *SuperformSession) GetVaultAddress() (common.Address, error) {
	return _Superform.Contract.GetVaultAddress(&_Superform.CallOpts)
}

// GetVaultAddress is a free data retrieval call binding the contract method 0x65cacaa4.
//
// Solidity: function getVaultAddress() view returns(address)
func (_Superform *SuperformCallerSession) GetVaultAddress() (common.Address, error) {
	return _Superform.Contract.GetVaultAddress(&_Superform.CallOpts)
}

// GetVaultAsset is a free data retrieval call binding the contract method 0xb60262ca.
//
// Solidity: function getVaultAsset() view returns(address)
func (_Superform *SuperformCaller) GetVaultAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getVaultAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultAsset is a free data retrieval call binding the contract method 0xb60262ca.
//
// Solidity: function getVaultAsset() view returns(address)
func (_Superform *SuperformSession) GetVaultAsset() (common.Address, error) {
	return _Superform.Contract.GetVaultAsset(&_Superform.CallOpts)
}

// GetVaultAsset is a free data retrieval call binding the contract method 0xb60262ca.
//
// Solidity: function getVaultAsset() view returns(address)
func (_Superform *SuperformCallerSession) GetVaultAsset() (common.Address, error) {
	return _Superform.Contract.GetVaultAsset(&_Superform.CallOpts)
}

// GetVaultDecimals is a free data retrieval call binding the contract method 0xc32dcd89.
//
// Solidity: function getVaultDecimals() view returns(uint256)
func (_Superform *SuperformCaller) GetVaultDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getVaultDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultDecimals is a free data retrieval call binding the contract method 0xc32dcd89.
//
// Solidity: function getVaultDecimals() view returns(uint256)
func (_Superform *SuperformSession) GetVaultDecimals() (*big.Int, error) {
	return _Superform.Contract.GetVaultDecimals(&_Superform.CallOpts)
}

// GetVaultDecimals is a free data retrieval call binding the contract method 0xc32dcd89.
//
// Solidity: function getVaultDecimals() view returns(uint256)
func (_Superform *SuperformCallerSession) GetVaultDecimals() (*big.Int, error) {
	return _Superform.Contract.GetVaultDecimals(&_Superform.CallOpts)
}

// GetVaultName is a free data retrieval call binding the contract method 0x9f5376c1.
//
// Solidity: function getVaultName() view returns(string)
func (_Superform *SuperformCaller) GetVaultName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getVaultName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVaultName is a free data retrieval call binding the contract method 0x9f5376c1.
//
// Solidity: function getVaultName() view returns(string)
func (_Superform *SuperformSession) GetVaultName() (string, error) {
	return _Superform.Contract.GetVaultName(&_Superform.CallOpts)
}

// GetVaultName is a free data retrieval call binding the contract method 0x9f5376c1.
//
// Solidity: function getVaultName() view returns(string)
func (_Superform *SuperformCallerSession) GetVaultName() (string, error) {
	return _Superform.Contract.GetVaultName(&_Superform.CallOpts)
}

// GetVaultShareBalance is a free data retrieval call binding the contract method 0x35eda680.
//
// Solidity: function getVaultShareBalance() view returns(uint256)
func (_Superform *SuperformCaller) GetVaultShareBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getVaultShareBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultShareBalance is a free data retrieval call binding the contract method 0x35eda680.
//
// Solidity: function getVaultShareBalance() view returns(uint256)
func (_Superform *SuperformSession) GetVaultShareBalance() (*big.Int, error) {
	return _Superform.Contract.GetVaultShareBalance(&_Superform.CallOpts)
}

// GetVaultShareBalance is a free data retrieval call binding the contract method 0x35eda680.
//
// Solidity: function getVaultShareBalance() view returns(uint256)
func (_Superform *SuperformCallerSession) GetVaultShareBalance() (*big.Int, error) {
	return _Superform.Contract.GetVaultShareBalance(&_Superform.CallOpts)
}

// GetVaultSymbol is a free data retrieval call binding the contract method 0x77188067.
//
// Solidity: function getVaultSymbol() view returns(string)
func (_Superform *SuperformCaller) GetVaultSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "getVaultSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVaultSymbol is a free data retrieval call binding the contract method 0x77188067.
//
// Solidity: function getVaultSymbol() view returns(string)
func (_Superform *SuperformSession) GetVaultSymbol() (string, error) {
	return _Superform.Contract.GetVaultSymbol(&_Superform.CallOpts)
}

// GetVaultSymbol is a free data retrieval call binding the contract method 0x77188067.
//
// Solidity: function getVaultSymbol() view returns(string)
func (_Superform *SuperformCallerSession) GetVaultSymbol() (string, error) {
	return _Superform.Contract.GetVaultSymbol(&_Superform.CallOpts)
}

// PreviewDepositTo is a free data retrieval call binding the contract method 0x07c080f9.
//
// Solidity: function previewDepositTo(uint256 assets_) view returns(uint256)
func (_Superform *SuperformCaller) PreviewDepositTo(opts *bind.CallOpts, assets_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "previewDepositTo", assets_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDepositTo is a free data retrieval call binding the contract method 0x07c080f9.
//
// Solidity: function previewDepositTo(uint256 assets_) view returns(uint256)
func (_Superform *SuperformSession) PreviewDepositTo(assets_ *big.Int) (*big.Int, error) {
	return _Superform.Contract.PreviewDepositTo(&_Superform.CallOpts, assets_)
}

// PreviewDepositTo is a free data retrieval call binding the contract method 0x07c080f9.
//
// Solidity: function previewDepositTo(uint256 assets_) view returns(uint256)
func (_Superform *SuperformCallerSession) PreviewDepositTo(assets_ *big.Int) (*big.Int, error) {
	return _Superform.Contract.PreviewDepositTo(&_Superform.CallOpts, assets_)
}

// PreviewRedeemFrom is a free data retrieval call binding the contract method 0xb7ba28cd.
//
// Solidity: function previewRedeemFrom(uint256 shares_) view returns(uint256)
func (_Superform *SuperformCaller) PreviewRedeemFrom(opts *bind.CallOpts, shares_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "previewRedeemFrom", shares_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeemFrom is a free data retrieval call binding the contract method 0xb7ba28cd.
//
// Solidity: function previewRedeemFrom(uint256 shares_) view returns(uint256)
func (_Superform *SuperformSession) PreviewRedeemFrom(shares_ *big.Int) (*big.Int, error) {
	return _Superform.Contract.PreviewRedeemFrom(&_Superform.CallOpts, shares_)
}

// PreviewRedeemFrom is a free data retrieval call binding the contract method 0xb7ba28cd.
//
// Solidity: function previewRedeemFrom(uint256 shares_) view returns(uint256)
func (_Superform *SuperformCallerSession) PreviewRedeemFrom(shares_ *big.Int) (*big.Int, error) {
	return _Superform.Contract.PreviewRedeemFrom(&_Superform.CallOpts, shares_)
}

// PreviewWithdrawFrom is a free data retrieval call binding the contract method 0x37d25010.
//
// Solidity: function previewWithdrawFrom(uint256 assets_) view returns(uint256)
func (_Superform *SuperformCaller) PreviewWithdrawFrom(opts *bind.CallOpts, assets_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "previewWithdrawFrom", assets_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdrawFrom is a free data retrieval call binding the contract method 0x37d25010.
//
// Solidity: function previewWithdrawFrom(uint256 assets_) view returns(uint256)
func (_Superform *SuperformSession) PreviewWithdrawFrom(assets_ *big.Int) (*big.Int, error) {
	return _Superform.Contract.PreviewWithdrawFrom(&_Superform.CallOpts, assets_)
}

// PreviewWithdrawFrom is a free data retrieval call binding the contract method 0x37d25010.
//
// Solidity: function previewWithdrawFrom(uint256 assets_) view returns(uint256)
func (_Superform *SuperformCallerSession) PreviewWithdrawFrom(assets_ *big.Int) (*big.Int, error) {
	return _Superform.Contract.PreviewWithdrawFrom(&_Superform.CallOpts, assets_)
}

// SuperRegistry is a free data retrieval call binding the contract method 0x24c73dda.
//
// Solidity: function superRegistry() view returns(address)
func (_Superform *SuperformCaller) SuperRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "superRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperRegistry is a free data retrieval call binding the contract method 0x24c73dda.
//
// Solidity: function superRegistry() view returns(address)
func (_Superform *SuperformSession) SuperRegistry() (common.Address, error) {
	return _Superform.Contract.SuperRegistry(&_Superform.CallOpts)
}

// SuperRegistry is a free data retrieval call binding the contract method 0x24c73dda.
//
// Solidity: function superRegistry() view returns(address)
func (_Superform *SuperformCallerSession) SuperRegistry() (common.Address, error) {
	return _Superform.Contract.SuperRegistry(&_Superform.CallOpts)
}

// SuperformYieldTokenName is a free data retrieval call binding the contract method 0x20592d98.
//
// Solidity: function superformYieldTokenName() view returns(string)
func (_Superform *SuperformCaller) SuperformYieldTokenName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "superformYieldTokenName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SuperformYieldTokenName is a free data retrieval call binding the contract method 0x20592d98.
//
// Solidity: function superformYieldTokenName() view returns(string)
func (_Superform *SuperformSession) SuperformYieldTokenName() (string, error) {
	return _Superform.Contract.SuperformYieldTokenName(&_Superform.CallOpts)
}

// SuperformYieldTokenName is a free data retrieval call binding the contract method 0x20592d98.
//
// Solidity: function superformYieldTokenName() view returns(string)
func (_Superform *SuperformCallerSession) SuperformYieldTokenName() (string, error) {
	return _Superform.Contract.SuperformYieldTokenName(&_Superform.CallOpts)
}

// SuperformYieldTokenSymbol is a free data retrieval call binding the contract method 0x17a57e08.
//
// Solidity: function superformYieldTokenSymbol() view returns(string)
func (_Superform *SuperformCaller) SuperformYieldTokenSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "superformYieldTokenSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SuperformYieldTokenSymbol is a free data retrieval call binding the contract method 0x17a57e08.
//
// Solidity: function superformYieldTokenSymbol() view returns(string)
func (_Superform *SuperformSession) SuperformYieldTokenSymbol() (string, error) {
	return _Superform.Contract.SuperformYieldTokenSymbol(&_Superform.CallOpts)
}

// SuperformYieldTokenSymbol is a free data retrieval call binding the contract method 0x17a57e08.
//
// Solidity: function superformYieldTokenSymbol() view returns(string)
func (_Superform *SuperformCallerSession) SuperformYieldTokenSymbol() (string, error) {
	return _Superform.Contract.SuperformYieldTokenSymbol(&_Superform.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId_) view returns(bool)
func (_Superform *SuperformCaller) SupportsInterface(opts *bind.CallOpts, interfaceId_ [4]byte) (bool, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "supportsInterface", interfaceId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId_) view returns(bool)
func (_Superform *SuperformSession) SupportsInterface(interfaceId_ [4]byte) (bool, error) {
	return _Superform.Contract.SupportsInterface(&_Superform.CallOpts, interfaceId_)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId_) view returns(bool)
func (_Superform *SuperformCallerSession) SupportsInterface(interfaceId_ [4]byte) (bool, error) {
	return _Superform.Contract.SupportsInterface(&_Superform.CallOpts, interfaceId_)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Superform *SuperformCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Superform.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Superform *SuperformSession) Vault() (common.Address, error) {
	return _Superform.Contract.Vault(&_Superform.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Superform *SuperformCallerSession) Vault() (common.Address, error) {
	return _Superform.Contract.Vault(&_Superform.CallOpts)
}

// DirectDepositIntoVault is a paid mutator transaction binding the contract method 0xb9232775.
//
// Solidity: function directDepositIntoVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_) payable returns(uint256 shares)
func (_Superform *SuperformTransactor) DirectDepositIntoVault(opts *bind.TransactOpts, singleVaultData_ InitSingleVaultData, srcSender_ common.Address) (*types.Transaction, error) {
	return _Superform.contract.Transact(opts, "directDepositIntoVault", singleVaultData_, srcSender_)
}

// DirectDepositIntoVault is a paid mutator transaction binding the contract method 0xb9232775.
//
// Solidity: function directDepositIntoVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_) payable returns(uint256 shares)
func (_Superform *SuperformSession) DirectDepositIntoVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.DirectDepositIntoVault(&_Superform.TransactOpts, singleVaultData_, srcSender_)
}

// DirectDepositIntoVault is a paid mutator transaction binding the contract method 0xb9232775.
//
// Solidity: function directDepositIntoVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_) payable returns(uint256 shares)
func (_Superform *SuperformTransactorSession) DirectDepositIntoVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.DirectDepositIntoVault(&_Superform.TransactOpts, singleVaultData_, srcSender_)
}

// DirectWithdrawFromVault is a paid mutator transaction binding the contract method 0xcb829dc3.
//
// Solidity: function directWithdrawFromVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_) returns(uint256 assets)
func (_Superform *SuperformTransactor) DirectWithdrawFromVault(opts *bind.TransactOpts, singleVaultData_ InitSingleVaultData, srcSender_ common.Address) (*types.Transaction, error) {
	return _Superform.contract.Transact(opts, "directWithdrawFromVault", singleVaultData_, srcSender_)
}

// DirectWithdrawFromVault is a paid mutator transaction binding the contract method 0xcb829dc3.
//
// Solidity: function directWithdrawFromVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_) returns(uint256 assets)
func (_Superform *SuperformSession) DirectWithdrawFromVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.DirectWithdrawFromVault(&_Superform.TransactOpts, singleVaultData_, srcSender_)
}

// DirectWithdrawFromVault is a paid mutator transaction binding the contract method 0xcb829dc3.
//
// Solidity: function directWithdrawFromVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_) returns(uint256 assets)
func (_Superform *SuperformTransactorSession) DirectWithdrawFromVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.DirectWithdrawFromVault(&_Superform.TransactOpts, singleVaultData_, srcSender_)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address receiverAddress_, uint256 amount_) returns()
func (_Superform *SuperformTransactor) EmergencyWithdraw(opts *bind.TransactOpts, receiverAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Superform.contract.Transact(opts, "emergencyWithdraw", receiverAddress_, amount_)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address receiverAddress_, uint256 amount_) returns()
func (_Superform *SuperformSession) EmergencyWithdraw(receiverAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Superform.Contract.EmergencyWithdraw(&_Superform.TransactOpts, receiverAddress_, amount_)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x95ccea67.
//
// Solidity: function emergencyWithdraw(address receiverAddress_, uint256 amount_) returns()
func (_Superform *SuperformTransactorSession) EmergencyWithdraw(receiverAddress_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Superform.Contract.EmergencyWithdraw(&_Superform.TransactOpts, receiverAddress_, amount_)
}

// ForwardDustToPaymaster is a paid mutator transaction binding the contract method 0x4dcd03c0.
//
// Solidity: function forwardDustToPaymaster(address token_) returns()
func (_Superform *SuperformTransactor) ForwardDustToPaymaster(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _Superform.contract.Transact(opts, "forwardDustToPaymaster", token_)
}

// ForwardDustToPaymaster is a paid mutator transaction binding the contract method 0x4dcd03c0.
//
// Solidity: function forwardDustToPaymaster(address token_) returns()
func (_Superform *SuperformSession) ForwardDustToPaymaster(token_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.ForwardDustToPaymaster(&_Superform.TransactOpts, token_)
}

// ForwardDustToPaymaster is a paid mutator transaction binding the contract method 0x4dcd03c0.
//
// Solidity: function forwardDustToPaymaster(address token_) returns()
func (_Superform *SuperformTransactorSession) ForwardDustToPaymaster(token_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.ForwardDustToPaymaster(&_Superform.TransactOpts, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address superRegistry_, address vault_, address asset_) returns()
func (_Superform *SuperformTransactor) Initialize(opts *bind.TransactOpts, superRegistry_ common.Address, vault_ common.Address, asset_ common.Address) (*types.Transaction, error) {
	return _Superform.contract.Transact(opts, "initialize", superRegistry_, vault_, asset_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address superRegistry_, address vault_, address asset_) returns()
func (_Superform *SuperformSession) Initialize(superRegistry_ common.Address, vault_ common.Address, asset_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.Initialize(&_Superform.TransactOpts, superRegistry_, vault_, asset_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address superRegistry_, address vault_, address asset_) returns()
func (_Superform *SuperformTransactorSession) Initialize(superRegistry_ common.Address, vault_ common.Address, asset_ common.Address) (*types.Transaction, error) {
	return _Superform.Contract.Initialize(&_Superform.TransactOpts, superRegistry_, vault_, asset_)
}

// XChainDepositIntoVault is a paid mutator transaction binding the contract method 0x95e4f1da.
//
// Solidity: function xChainDepositIntoVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_, uint64 srcChainId_) returns(uint256 shares)
func (_Superform *SuperformTransactor) XChainDepositIntoVault(opts *bind.TransactOpts, singleVaultData_ InitSingleVaultData, srcSender_ common.Address, srcChainId_ uint64) (*types.Transaction, error) {
	return _Superform.contract.Transact(opts, "xChainDepositIntoVault", singleVaultData_, srcSender_, srcChainId_)
}

// XChainDepositIntoVault is a paid mutator transaction binding the contract method 0x95e4f1da.
//
// Solidity: function xChainDepositIntoVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_, uint64 srcChainId_) returns(uint256 shares)
func (_Superform *SuperformSession) XChainDepositIntoVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address, srcChainId_ uint64) (*types.Transaction, error) {
	return _Superform.Contract.XChainDepositIntoVault(&_Superform.TransactOpts, singleVaultData_, srcSender_, srcChainId_)
}

// XChainDepositIntoVault is a paid mutator transaction binding the contract method 0x95e4f1da.
//
// Solidity: function xChainDepositIntoVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_, uint64 srcChainId_) returns(uint256 shares)
func (_Superform *SuperformTransactorSession) XChainDepositIntoVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address, srcChainId_ uint64) (*types.Transaction, error) {
	return _Superform.Contract.XChainDepositIntoVault(&_Superform.TransactOpts, singleVaultData_, srcSender_, srcChainId_)
}

// XChainWithdrawFromVault is a paid mutator transaction binding the contract method 0xef164fef.
//
// Solidity: function xChainWithdrawFromVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_, uint64 srcChainId_) returns(uint256 assets)
func (_Superform *SuperformTransactor) XChainWithdrawFromVault(opts *bind.TransactOpts, singleVaultData_ InitSingleVaultData, srcSender_ common.Address, srcChainId_ uint64) (*types.Transaction, error) {
	return _Superform.contract.Transact(opts, "xChainWithdrawFromVault", singleVaultData_, srcSender_, srcChainId_)
}

// XChainWithdrawFromVault is a paid mutator transaction binding the contract method 0xef164fef.
//
// Solidity: function xChainWithdrawFromVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_, uint64 srcChainId_) returns(uint256 assets)
func (_Superform *SuperformSession) XChainWithdrawFromVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address, srcChainId_ uint64) (*types.Transaction, error) {
	return _Superform.Contract.XChainWithdrawFromVault(&_Superform.TransactOpts, singleVaultData_, srcSender_, srcChainId_)
}

// XChainWithdrawFromVault is a paid mutator transaction binding the contract method 0xef164fef.
//
// Solidity: function xChainWithdrawFromVault((uint256,uint256,uint256,uint256,uint256,(bytes,address,address,uint8,uint64,uint256),bool,bool,address,bytes) singleVaultData_, address srcSender_, uint64 srcChainId_) returns(uint256 assets)
func (_Superform *SuperformTransactorSession) XChainWithdrawFromVault(singleVaultData_ InitSingleVaultData, srcSender_ common.Address, srcChainId_ uint64) (*types.Transaction, error) {
	return _Superform.Contract.XChainWithdrawFromVault(&_Superform.TransactOpts, singleVaultData_, srcSender_, srcChainId_)
}

// SuperformEmergencyWithdrawalProcessedIterator is returned from FilterEmergencyWithdrawalProcessed and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawalProcessed events raised by the Superform contract.
type SuperformEmergencyWithdrawalProcessedIterator struct {
	Event *SuperformEmergencyWithdrawalProcessed // Event containing the contract specifics and raw log

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
func (it *SuperformEmergencyWithdrawalProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperformEmergencyWithdrawalProcessed)
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
		it.Event = new(SuperformEmergencyWithdrawalProcessed)
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
func (it *SuperformEmergencyWithdrawalProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperformEmergencyWithdrawalProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperformEmergencyWithdrawalProcessed represents a EmergencyWithdrawalProcessed event raised by the Superform contract.
type SuperformEmergencyWithdrawalProcessed struct {
	RefundAddress common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawalProcessed is a free log retrieval operation binding the contract event 0x83b8068554a495dbb4af07014f8171144d6670eb522b356f8d3f37cbd76ba116.
//
// Solidity: event EmergencyWithdrawalProcessed(address indexed refundAddress, uint256 indexed amount)
func (_Superform *SuperformFilterer) FilterEmergencyWithdrawalProcessed(opts *bind.FilterOpts, refundAddress []common.Address, amount []*big.Int) (*SuperformEmergencyWithdrawalProcessedIterator, error) {

	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Superform.contract.FilterLogs(opts, "EmergencyWithdrawalProcessed", refundAddressRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &SuperformEmergencyWithdrawalProcessedIterator{contract: _Superform.contract, event: "EmergencyWithdrawalProcessed", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawalProcessed is a free log subscription operation binding the contract event 0x83b8068554a495dbb4af07014f8171144d6670eb522b356f8d3f37cbd76ba116.
//
// Solidity: event EmergencyWithdrawalProcessed(address indexed refundAddress, uint256 indexed amount)
func (_Superform *SuperformFilterer) WatchEmergencyWithdrawalProcessed(opts *bind.WatchOpts, sink chan<- *SuperformEmergencyWithdrawalProcessed, refundAddress []common.Address, amount []*big.Int) (event.Subscription, error) {

	var refundAddressRule []interface{}
	for _, refundAddressItem := range refundAddress {
		refundAddressRule = append(refundAddressRule, refundAddressItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Superform.contract.WatchLogs(opts, "EmergencyWithdrawalProcessed", refundAddressRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperformEmergencyWithdrawalProcessed)
				if err := _Superform.contract.UnpackLog(event, "EmergencyWithdrawalProcessed", log); err != nil {
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

// ParseEmergencyWithdrawalProcessed is a log parse operation binding the contract event 0x83b8068554a495dbb4af07014f8171144d6670eb522b356f8d3f37cbd76ba116.
//
// Solidity: event EmergencyWithdrawalProcessed(address indexed refundAddress, uint256 indexed amount)
func (_Superform *SuperformFilterer) ParseEmergencyWithdrawalProcessed(log types.Log) (*SuperformEmergencyWithdrawalProcessed, error) {
	event := new(SuperformEmergencyWithdrawalProcessed)
	if err := _Superform.contract.UnpackLog(event, "EmergencyWithdrawalProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperformFormDustForwardedToPaymasterIterator is returned from FilterFormDustForwardedToPaymaster and is used to iterate over the raw logs and unpacked data for FormDustForwardedToPaymaster events raised by the Superform contract.
type SuperformFormDustForwardedToPaymasterIterator struct {
	Event *SuperformFormDustForwardedToPaymaster // Event containing the contract specifics and raw log

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
func (it *SuperformFormDustForwardedToPaymasterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperformFormDustForwardedToPaymaster)
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
		it.Event = new(SuperformFormDustForwardedToPaymaster)
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
func (it *SuperformFormDustForwardedToPaymasterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperformFormDustForwardedToPaymasterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperformFormDustForwardedToPaymaster represents a FormDustForwardedToPaymaster event raised by the Superform contract.
type SuperformFormDustForwardedToPaymaster struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFormDustForwardedToPaymaster is a free log retrieval operation binding the contract event 0xd34222ea8b5b095ec7a6f42ff87fa762ab480b9e8cea464915ee985f1510c10a.
//
// Solidity: event FormDustForwardedToPaymaster(address indexed token, uint256 indexed amount)
func (_Superform *SuperformFilterer) FilterFormDustForwardedToPaymaster(opts *bind.FilterOpts, token []common.Address, amount []*big.Int) (*SuperformFormDustForwardedToPaymasterIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Superform.contract.FilterLogs(opts, "FormDustForwardedToPaymaster", tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &SuperformFormDustForwardedToPaymasterIterator{contract: _Superform.contract, event: "FormDustForwardedToPaymaster", logs: logs, sub: sub}, nil
}

// WatchFormDustForwardedToPaymaster is a free log subscription operation binding the contract event 0xd34222ea8b5b095ec7a6f42ff87fa762ab480b9e8cea464915ee985f1510c10a.
//
// Solidity: event FormDustForwardedToPaymaster(address indexed token, uint256 indexed amount)
func (_Superform *SuperformFilterer) WatchFormDustForwardedToPaymaster(opts *bind.WatchOpts, sink chan<- *SuperformFormDustForwardedToPaymaster, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Superform.contract.WatchLogs(opts, "FormDustForwardedToPaymaster", tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperformFormDustForwardedToPaymaster)
				if err := _Superform.contract.UnpackLog(event, "FormDustForwardedToPaymaster", log); err != nil {
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

// ParseFormDustForwardedToPaymaster is a log parse operation binding the contract event 0xd34222ea8b5b095ec7a6f42ff87fa762ab480b9e8cea464915ee985f1510c10a.
//
// Solidity: event FormDustForwardedToPaymaster(address indexed token, uint256 indexed amount)
func (_Superform *SuperformFilterer) ParseFormDustForwardedToPaymaster(log types.Log) (*SuperformFormDustForwardedToPaymaster, error) {
	event := new(SuperformFormDustForwardedToPaymaster)
	if err := _Superform.contract.UnpackLog(event, "FormDustForwardedToPaymaster", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperformInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Superform contract.
type SuperformInitializedIterator struct {
	Event *SuperformInitialized // Event containing the contract specifics and raw log

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
func (it *SuperformInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperformInitialized)
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
		it.Event = new(SuperformInitialized)
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
func (it *SuperformInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperformInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperformInitialized represents a Initialized event raised by the Superform contract.
type SuperformInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Superform *SuperformFilterer) FilterInitialized(opts *bind.FilterOpts) (*SuperformInitializedIterator, error) {

	logs, sub, err := _Superform.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SuperformInitializedIterator{contract: _Superform.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Superform *SuperformFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SuperformInitialized) (event.Subscription, error) {

	logs, sub, err := _Superform.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperformInitialized)
				if err := _Superform.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Superform *SuperformFilterer) ParseInitialized(log types.Log) (*SuperformInitialized, error) {
	event := new(SuperformInitialized)
	if err := _Superform.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperformProcessedIterator is returned from FilterProcessed and is used to iterate over the raw logs and unpacked data for Processed events raised by the Superform contract.
type SuperformProcessedIterator struct {
	Event *SuperformProcessed // Event containing the contract specifics and raw log

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
func (it *SuperformProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperformProcessed)
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
		it.Event = new(SuperformProcessed)
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
func (it *SuperformProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperformProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperformProcessed represents a Processed event raised by the Superform contract.
type SuperformProcessed struct {
	SrcChainID   uint64
	DstChainId   uint64
	SrcPayloadId *big.Int
	Amount       *big.Int
	Vault        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProcessed is a free log retrieval operation binding the contract event 0x9664a7293fbeac5e42927bc5eb69c82d1fe2f0b17e510b38ebfba582d47923fe.
//
// Solidity: event Processed(uint64 indexed srcChainID, uint64 indexed dstChainId, uint256 indexed srcPayloadId, uint256 amount, address vault)
func (_Superform *SuperformFilterer) FilterProcessed(opts *bind.FilterOpts, srcChainID []uint64, dstChainId []uint64, srcPayloadId []*big.Int) (*SuperformProcessedIterator, error) {

	var srcChainIDRule []interface{}
	for _, srcChainIDItem := range srcChainID {
		srcChainIDRule = append(srcChainIDRule, srcChainIDItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var srcPayloadIdRule []interface{}
	for _, srcPayloadIdItem := range srcPayloadId {
		srcPayloadIdRule = append(srcPayloadIdRule, srcPayloadIdItem)
	}

	logs, sub, err := _Superform.contract.FilterLogs(opts, "Processed", srcChainIDRule, dstChainIdRule, srcPayloadIdRule)
	if err != nil {
		return nil, err
	}
	return &SuperformProcessedIterator{contract: _Superform.contract, event: "Processed", logs: logs, sub: sub}, nil
}

// WatchProcessed is a free log subscription operation binding the contract event 0x9664a7293fbeac5e42927bc5eb69c82d1fe2f0b17e510b38ebfba582d47923fe.
//
// Solidity: event Processed(uint64 indexed srcChainID, uint64 indexed dstChainId, uint256 indexed srcPayloadId, uint256 amount, address vault)
func (_Superform *SuperformFilterer) WatchProcessed(opts *bind.WatchOpts, sink chan<- *SuperformProcessed, srcChainID []uint64, dstChainId []uint64, srcPayloadId []*big.Int) (event.Subscription, error) {

	var srcChainIDRule []interface{}
	for _, srcChainIDItem := range srcChainID {
		srcChainIDRule = append(srcChainIDRule, srcChainIDItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var srcPayloadIdRule []interface{}
	for _, srcPayloadIdItem := range srcPayloadId {
		srcPayloadIdRule = append(srcPayloadIdRule, srcPayloadIdItem)
	}

	logs, sub, err := _Superform.contract.WatchLogs(opts, "Processed", srcChainIDRule, dstChainIdRule, srcPayloadIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperformProcessed)
				if err := _Superform.contract.UnpackLog(event, "Processed", log); err != nil {
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

// ParseProcessed is a log parse operation binding the contract event 0x9664a7293fbeac5e42927bc5eb69c82d1fe2f0b17e510b38ebfba582d47923fe.
//
// Solidity: event Processed(uint64 indexed srcChainID, uint64 indexed dstChainId, uint256 indexed srcPayloadId, uint256 amount, address vault)
func (_Superform *SuperformFilterer) ParseProcessed(log types.Log) (*SuperformProcessed, error) {
	event := new(SuperformProcessed)
	if err := _Superform.contract.UnpackLog(event, "Processed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SuperformVaultAddedIterator is returned from FilterVaultAdded and is used to iterate over the raw logs and unpacked data for VaultAdded events raised by the Superform contract.
type SuperformVaultAddedIterator struct {
	Event *SuperformVaultAdded // Event containing the contract specifics and raw log

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
func (it *SuperformVaultAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperformVaultAdded)
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
		it.Event = new(SuperformVaultAdded)
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
func (it *SuperformVaultAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperformVaultAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperformVaultAdded represents a VaultAdded event raised by the Superform contract.
type SuperformVaultAdded struct {
	Id    *big.Int
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultAdded is a free log retrieval operation binding the contract event 0xa3ccd9b56d18a571b67b97905e5ef425788000d31a490513f7cad937175beeeb.
//
// Solidity: event VaultAdded(uint256 indexed id, address indexed vault)
func (_Superform *SuperformFilterer) FilterVaultAdded(opts *bind.FilterOpts, id []*big.Int, vault []common.Address) (*SuperformVaultAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Superform.contract.FilterLogs(opts, "VaultAdded", idRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &SuperformVaultAddedIterator{contract: _Superform.contract, event: "VaultAdded", logs: logs, sub: sub}, nil
}

// WatchVaultAdded is a free log subscription operation binding the contract event 0xa3ccd9b56d18a571b67b97905e5ef425788000d31a490513f7cad937175beeeb.
//
// Solidity: event VaultAdded(uint256 indexed id, address indexed vault)
func (_Superform *SuperformFilterer) WatchVaultAdded(opts *bind.WatchOpts, sink chan<- *SuperformVaultAdded, id []*big.Int, vault []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Superform.contract.WatchLogs(opts, "VaultAdded", idRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperformVaultAdded)
				if err := _Superform.contract.UnpackLog(event, "VaultAdded", log); err != nil {
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

// ParseVaultAdded is a log parse operation binding the contract event 0xa3ccd9b56d18a571b67b97905e5ef425788000d31a490513f7cad937175beeeb.
//
// Solidity: event VaultAdded(uint256 indexed id, address indexed vault)
func (_Superform *SuperformFilterer) ParseVaultAdded(log types.Log) (*SuperformVaultAdded, error) {
	event := new(SuperformVaultAdded)
	if err := _Superform.contract.UnpackLog(event, "VaultAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
