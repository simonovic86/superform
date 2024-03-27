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

// IPayloadHelperDecodedDstPayload is an auto generated low-level Go binding around an user-defined struct.
type IPayloadHelperDecodedDstPayload struct {
	TxType          uint8
	CallbackType    uint8
	SrcSender       common.Address
	SrcChainId      uint64
	Amounts         []*big.Int
	OutputAmounts   []*big.Int
	Slippages       []*big.Int
	SuperformIds    []*big.Int
	HasDstSwaps     []bool
	ReceiverAddress common.Address
	SrcPayloadId    *big.Int
	ExtraFormData   []byte
	Multi           uint8
}

// PayloadHelperMetaData contains all meta data concerning the PayloadHelper contract.
var PayloadHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"superRegistry_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"INVALID_PAYLOAD\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_PAYLOAD_ID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstPayloadId_\",\"type\":\"uint256\"}],\"name\":\"decodeCoreStateRegistryPayload\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"callbackType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"srcSender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"outputAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"slippages\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"superformIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"hasDstSwaps\",\"type\":\"bool[]\"},{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcPayloadId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraFormData\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"multi\",\"type\":\"uint8\"}],\"internalType\":\"structIPayloadHelper.DecodedDstPayload\",\"name\":\"v\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstPayloadId_\",\"type\":\"uint256\"}],\"name\":\"decodeCoreStateRegistryPayloadLiqData\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"txDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"interimTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"bridgeIds\",\"type\":\"uint8[]\"},{\"internalType\":\"uint64[]\",\"name\":\"liqDstChainIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nativeAmounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcPayloadId_\",\"type\":\"uint256\"}],\"name\":\"decodePayloadHistory\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"txType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"callbackType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"multi\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"srcSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiverAddressSP\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"payloadId_\",\"type\":\"uint256\"}],\"name\":\"decodeTimeLockFailedPayload\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"srcSender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcPayloadId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superformId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timelockPayloadId_\",\"type\":\"uint256\"}],\"name\":\"decodeTimeLockPayload\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcPayloadId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"superformId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstPayloadId_\",\"type\":\"uint256\"}],\"name\":\"getDstPayloadProof\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superRegistry\",\"outputs\":[{\"internalType\":\"contractISuperRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PayloadHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use PayloadHelperMetaData.ABI instead.
var PayloadHelperABI = PayloadHelperMetaData.ABI

// PayloadHelper is an auto generated Go binding around an Ethereum contract.
type PayloadHelper struct {
	PayloadHelperCaller     // Read-only binding to the contract
	PayloadHelperTransactor // Write-only binding to the contract
	PayloadHelperFilterer   // Log filterer for contract events
}

// PayloadHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayloadHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayloadHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayloadHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayloadHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayloadHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayloadHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PayloadHelperSession struct {
	Contract     *PayloadHelper    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayloadHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayloadHelperCallerSession struct {
	Contract *PayloadHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PayloadHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayloadHelperTransactorSession struct {
	Contract     *PayloadHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PayloadHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayloadHelperRaw struct {
	Contract *PayloadHelper // Generic contract binding to access the raw methods on
}

// PayloadHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayloadHelperCallerRaw struct {
	Contract *PayloadHelperCaller // Generic read-only contract binding to access the raw methods on
}

// PayloadHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayloadHelperTransactorRaw struct {
	Contract *PayloadHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayloadHelper creates a new instance of PayloadHelper, bound to a specific deployed contract.
func NewPayloadHelper(address common.Address, backend bind.ContractBackend) (*PayloadHelper, error) {
	contract, err := bindPayloadHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PayloadHelper{PayloadHelperCaller: PayloadHelperCaller{contract: contract}, PayloadHelperTransactor: PayloadHelperTransactor{contract: contract}, PayloadHelperFilterer: PayloadHelperFilterer{contract: contract}}, nil
}

// NewPayloadHelperCaller creates a new read-only instance of PayloadHelper, bound to a specific deployed contract.
func NewPayloadHelperCaller(address common.Address, caller bind.ContractCaller) (*PayloadHelperCaller, error) {
	contract, err := bindPayloadHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayloadHelperCaller{contract: contract}, nil
}

// NewPayloadHelperTransactor creates a new write-only instance of PayloadHelper, bound to a specific deployed contract.
func NewPayloadHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*PayloadHelperTransactor, error) {
	contract, err := bindPayloadHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayloadHelperTransactor{contract: contract}, nil
}

// NewPayloadHelperFilterer creates a new log filterer instance of PayloadHelper, bound to a specific deployed contract.
func NewPayloadHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*PayloadHelperFilterer, error) {
	contract, err := bindPayloadHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayloadHelperFilterer{contract: contract}, nil
}

// bindPayloadHelper binds a generic wrapper to an already deployed contract.
func bindPayloadHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PayloadHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayloadHelper *PayloadHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayloadHelper.Contract.PayloadHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayloadHelper *PayloadHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayloadHelper.Contract.PayloadHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayloadHelper *PayloadHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayloadHelper.Contract.PayloadHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayloadHelper *PayloadHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PayloadHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayloadHelper *PayloadHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayloadHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayloadHelper *PayloadHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayloadHelper.Contract.contract.Transact(opts, method, params...)
}

// DecodeCoreStateRegistryPayload is a free data retrieval call binding the contract method 0xcd7a066a.
//
// Solidity: function decodeCoreStateRegistryPayload(uint256 dstPayloadId_) view returns((uint8,uint8,address,uint64,uint256[],uint256[],uint256[],uint256[],bool[],address,uint256,bytes,uint8) v)
func (_PayloadHelper *PayloadHelperCaller) DecodeCoreStateRegistryPayload(opts *bind.CallOpts, dstPayloadId_ *big.Int) (IPayloadHelperDecodedDstPayload, error) {
	var out []interface{}
	err := _PayloadHelper.contract.Call(opts, &out, "decodeCoreStateRegistryPayload", dstPayloadId_)

	if err != nil {
		return *new(IPayloadHelperDecodedDstPayload), err
	}

	out0 := *abi.ConvertType(out[0], new(IPayloadHelperDecodedDstPayload)).(*IPayloadHelperDecodedDstPayload)

	return out0, err

}

// DecodeCoreStateRegistryPayload is a free data retrieval call binding the contract method 0xcd7a066a.
//
// Solidity: function decodeCoreStateRegistryPayload(uint256 dstPayloadId_) view returns((uint8,uint8,address,uint64,uint256[],uint256[],uint256[],uint256[],bool[],address,uint256,bytes,uint8) v)
func (_PayloadHelper *PayloadHelperSession) DecodeCoreStateRegistryPayload(dstPayloadId_ *big.Int) (IPayloadHelperDecodedDstPayload, error) {
	return _PayloadHelper.Contract.DecodeCoreStateRegistryPayload(&_PayloadHelper.CallOpts, dstPayloadId_)
}

// DecodeCoreStateRegistryPayload is a free data retrieval call binding the contract method 0xcd7a066a.
//
// Solidity: function decodeCoreStateRegistryPayload(uint256 dstPayloadId_) view returns((uint8,uint8,address,uint64,uint256[],uint256[],uint256[],uint256[],bool[],address,uint256,bytes,uint8) v)
func (_PayloadHelper *PayloadHelperCallerSession) DecodeCoreStateRegistryPayload(dstPayloadId_ *big.Int) (IPayloadHelperDecodedDstPayload, error) {
	return _PayloadHelper.Contract.DecodeCoreStateRegistryPayload(&_PayloadHelper.CallOpts, dstPayloadId_)
}

// DecodeCoreStateRegistryPayloadLiqData is a free data retrieval call binding the contract method 0x6408fe7e.
//
// Solidity: function decodeCoreStateRegistryPayloadLiqData(uint256 dstPayloadId_) view returns(bytes[] txDatas, address[] tokens, address[] interimTokens, uint8[] bridgeIds, uint64[] liqDstChainIds, uint256[] amountsIn, uint256[] nativeAmounts)
func (_PayloadHelper *PayloadHelperCaller) DecodeCoreStateRegistryPayloadLiqData(opts *bind.CallOpts, dstPayloadId_ *big.Int) (struct {
	TxDatas        [][]byte
	Tokens         []common.Address
	InterimTokens  []common.Address
	BridgeIds      []uint8
	LiqDstChainIds []uint64
	AmountsIn      []*big.Int
	NativeAmounts  []*big.Int
}, error) {
	var out []interface{}
	err := _PayloadHelper.contract.Call(opts, &out, "decodeCoreStateRegistryPayloadLiqData", dstPayloadId_)

	outstruct := new(struct {
		TxDatas        [][]byte
		Tokens         []common.Address
		InterimTokens  []common.Address
		BridgeIds      []uint8
		LiqDstChainIds []uint64
		AmountsIn      []*big.Int
		NativeAmounts  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxDatas = *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	outstruct.Tokens = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.InterimTokens = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)
	outstruct.BridgeIds = *abi.ConvertType(out[3], new([]uint8)).(*[]uint8)
	outstruct.LiqDstChainIds = *abi.ConvertType(out[4], new([]uint64)).(*[]uint64)
	outstruct.AmountsIn = *abi.ConvertType(out[5], new([]*big.Int)).(*[]*big.Int)
	outstruct.NativeAmounts = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// DecodeCoreStateRegistryPayloadLiqData is a free data retrieval call binding the contract method 0x6408fe7e.
//
// Solidity: function decodeCoreStateRegistryPayloadLiqData(uint256 dstPayloadId_) view returns(bytes[] txDatas, address[] tokens, address[] interimTokens, uint8[] bridgeIds, uint64[] liqDstChainIds, uint256[] amountsIn, uint256[] nativeAmounts)
func (_PayloadHelper *PayloadHelperSession) DecodeCoreStateRegistryPayloadLiqData(dstPayloadId_ *big.Int) (struct {
	TxDatas        [][]byte
	Tokens         []common.Address
	InterimTokens  []common.Address
	BridgeIds      []uint8
	LiqDstChainIds []uint64
	AmountsIn      []*big.Int
	NativeAmounts  []*big.Int
}, error) {
	return _PayloadHelper.Contract.DecodeCoreStateRegistryPayloadLiqData(&_PayloadHelper.CallOpts, dstPayloadId_)
}

// DecodeCoreStateRegistryPayloadLiqData is a free data retrieval call binding the contract method 0x6408fe7e.
//
// Solidity: function decodeCoreStateRegistryPayloadLiqData(uint256 dstPayloadId_) view returns(bytes[] txDatas, address[] tokens, address[] interimTokens, uint8[] bridgeIds, uint64[] liqDstChainIds, uint256[] amountsIn, uint256[] nativeAmounts)
func (_PayloadHelper *PayloadHelperCallerSession) DecodeCoreStateRegistryPayloadLiqData(dstPayloadId_ *big.Int) (struct {
	TxDatas        [][]byte
	Tokens         []common.Address
	InterimTokens  []common.Address
	BridgeIds      []uint8
	LiqDstChainIds []uint64
	AmountsIn      []*big.Int
	NativeAmounts  []*big.Int
}, error) {
	return _PayloadHelper.Contract.DecodeCoreStateRegistryPayloadLiqData(&_PayloadHelper.CallOpts, dstPayloadId_)
}

// DecodePayloadHistory is a free data retrieval call binding the contract method 0x7fab2df1.
//
// Solidity: function decodePayloadHistory(uint256 srcPayloadId_) view returns(uint8 txType, uint8 callbackType, uint8 multi, address srcSender, address receiverAddressSP, uint64 srcChainId)
func (_PayloadHelper *PayloadHelperCaller) DecodePayloadHistory(opts *bind.CallOpts, srcPayloadId_ *big.Int) (struct {
	TxType            uint8
	CallbackType      uint8
	Multi             uint8
	SrcSender         common.Address
	ReceiverAddressSP common.Address
	SrcChainId        uint64
}, error) {
	var out []interface{}
	err := _PayloadHelper.contract.Call(opts, &out, "decodePayloadHistory", srcPayloadId_)

	outstruct := new(struct {
		TxType            uint8
		CallbackType      uint8
		Multi             uint8
		SrcSender         common.Address
		ReceiverAddressSP common.Address
		SrcChainId        uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TxType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.CallbackType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Multi = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.SrcSender = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ReceiverAddressSP = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.SrcChainId = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// DecodePayloadHistory is a free data retrieval call binding the contract method 0x7fab2df1.
//
// Solidity: function decodePayloadHistory(uint256 srcPayloadId_) view returns(uint8 txType, uint8 callbackType, uint8 multi, address srcSender, address receiverAddressSP, uint64 srcChainId)
func (_PayloadHelper *PayloadHelperSession) DecodePayloadHistory(srcPayloadId_ *big.Int) (struct {
	TxType            uint8
	CallbackType      uint8
	Multi             uint8
	SrcSender         common.Address
	ReceiverAddressSP common.Address
	SrcChainId        uint64
}, error) {
	return _PayloadHelper.Contract.DecodePayloadHistory(&_PayloadHelper.CallOpts, srcPayloadId_)
}

// DecodePayloadHistory is a free data retrieval call binding the contract method 0x7fab2df1.
//
// Solidity: function decodePayloadHistory(uint256 srcPayloadId_) view returns(uint8 txType, uint8 callbackType, uint8 multi, address srcSender, address receiverAddressSP, uint64 srcChainId)
func (_PayloadHelper *PayloadHelperCallerSession) DecodePayloadHistory(srcPayloadId_ *big.Int) (struct {
	TxType            uint8
	CallbackType      uint8
	Multi             uint8
	SrcSender         common.Address
	ReceiverAddressSP common.Address
	SrcChainId        uint64
}, error) {
	return _PayloadHelper.Contract.DecodePayloadHistory(&_PayloadHelper.CallOpts, srcPayloadId_)
}

// DecodeTimeLockFailedPayload is a free data retrieval call binding the contract method 0xef86871f.
//
// Solidity: function decodeTimeLockFailedPayload(uint256 payloadId_) view returns(address srcSender, uint64 srcChainId, uint256 srcPayloadId, uint256 superformId, uint256 amount)
func (_PayloadHelper *PayloadHelperCaller) DecodeTimeLockFailedPayload(opts *bind.CallOpts, payloadId_ *big.Int) (struct {
	SrcSender    common.Address
	SrcChainId   uint64
	SrcPayloadId *big.Int
	SuperformId  *big.Int
	Amount       *big.Int
}, error) {
	var out []interface{}
	err := _PayloadHelper.contract.Call(opts, &out, "decodeTimeLockFailedPayload", payloadId_)

	outstruct := new(struct {
		SrcSender    common.Address
		SrcChainId   uint64
		SrcPayloadId *big.Int
		SuperformId  *big.Int
		Amount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SrcSender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SrcChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.SrcPayloadId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SuperformId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DecodeTimeLockFailedPayload is a free data retrieval call binding the contract method 0xef86871f.
//
// Solidity: function decodeTimeLockFailedPayload(uint256 payloadId_) view returns(address srcSender, uint64 srcChainId, uint256 srcPayloadId, uint256 superformId, uint256 amount)
func (_PayloadHelper *PayloadHelperSession) DecodeTimeLockFailedPayload(payloadId_ *big.Int) (struct {
	SrcSender    common.Address
	SrcChainId   uint64
	SrcPayloadId *big.Int
	SuperformId  *big.Int
	Amount       *big.Int
}, error) {
	return _PayloadHelper.Contract.DecodeTimeLockFailedPayload(&_PayloadHelper.CallOpts, payloadId_)
}

// DecodeTimeLockFailedPayload is a free data retrieval call binding the contract method 0xef86871f.
//
// Solidity: function decodeTimeLockFailedPayload(uint256 payloadId_) view returns(address srcSender, uint64 srcChainId, uint256 srcPayloadId, uint256 superformId, uint256 amount)
func (_PayloadHelper *PayloadHelperCallerSession) DecodeTimeLockFailedPayload(payloadId_ *big.Int) (struct {
	SrcSender    common.Address
	SrcChainId   uint64
	SrcPayloadId *big.Int
	SuperformId  *big.Int
	Amount       *big.Int
}, error) {
	return _PayloadHelper.Contract.DecodeTimeLockFailedPayload(&_PayloadHelper.CallOpts, payloadId_)
}

// DecodeTimeLockPayload is a free data retrieval call binding the contract method 0x8ecd95dd.
//
// Solidity: function decodeTimeLockPayload(uint256 timelockPayloadId_) view returns(address receiverAddress, uint64 srcChainId, uint256 srcPayloadId, uint256 superformId, uint256 amount)
func (_PayloadHelper *PayloadHelperCaller) DecodeTimeLockPayload(opts *bind.CallOpts, timelockPayloadId_ *big.Int) (struct {
	ReceiverAddress common.Address
	SrcChainId      uint64
	SrcPayloadId    *big.Int
	SuperformId     *big.Int
	Amount          *big.Int
}, error) {
	var out []interface{}
	err := _PayloadHelper.contract.Call(opts, &out, "decodeTimeLockPayload", timelockPayloadId_)

	outstruct := new(struct {
		ReceiverAddress common.Address
		SrcChainId      uint64
		SrcPayloadId    *big.Int
		SuperformId     *big.Int
		Amount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiverAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SrcChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.SrcPayloadId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SuperformId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DecodeTimeLockPayload is a free data retrieval call binding the contract method 0x8ecd95dd.
//
// Solidity: function decodeTimeLockPayload(uint256 timelockPayloadId_) view returns(address receiverAddress, uint64 srcChainId, uint256 srcPayloadId, uint256 superformId, uint256 amount)
func (_PayloadHelper *PayloadHelperSession) DecodeTimeLockPayload(timelockPayloadId_ *big.Int) (struct {
	ReceiverAddress common.Address
	SrcChainId      uint64
	SrcPayloadId    *big.Int
	SuperformId     *big.Int
	Amount          *big.Int
}, error) {
	return _PayloadHelper.Contract.DecodeTimeLockPayload(&_PayloadHelper.CallOpts, timelockPayloadId_)
}

// DecodeTimeLockPayload is a free data retrieval call binding the contract method 0x8ecd95dd.
//
// Solidity: function decodeTimeLockPayload(uint256 timelockPayloadId_) view returns(address receiverAddress, uint64 srcChainId, uint256 srcPayloadId, uint256 superformId, uint256 amount)
func (_PayloadHelper *PayloadHelperCallerSession) DecodeTimeLockPayload(timelockPayloadId_ *big.Int) (struct {
	ReceiverAddress common.Address
	SrcChainId      uint64
	SrcPayloadId    *big.Int
	SuperformId     *big.Int
	Amount          *big.Int
}, error) {
	return _PayloadHelper.Contract.DecodeTimeLockPayload(&_PayloadHelper.CallOpts, timelockPayloadId_)
}

// GetDstPayloadProof is a free data retrieval call binding the contract method 0xc0ca67a9.
//
// Solidity: function getDstPayloadProof(uint256 dstPayloadId_) view returns(bytes32)
func (_PayloadHelper *PayloadHelperCaller) GetDstPayloadProof(opts *bind.CallOpts, dstPayloadId_ *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PayloadHelper.contract.Call(opts, &out, "getDstPayloadProof", dstPayloadId_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDstPayloadProof is a free data retrieval call binding the contract method 0xc0ca67a9.
//
// Solidity: function getDstPayloadProof(uint256 dstPayloadId_) view returns(bytes32)
func (_PayloadHelper *PayloadHelperSession) GetDstPayloadProof(dstPayloadId_ *big.Int) ([32]byte, error) {
	return _PayloadHelper.Contract.GetDstPayloadProof(&_PayloadHelper.CallOpts, dstPayloadId_)
}

// GetDstPayloadProof is a free data retrieval call binding the contract method 0xc0ca67a9.
//
// Solidity: function getDstPayloadProof(uint256 dstPayloadId_) view returns(bytes32)
func (_PayloadHelper *PayloadHelperCallerSession) GetDstPayloadProof(dstPayloadId_ *big.Int) ([32]byte, error) {
	return _PayloadHelper.Contract.GetDstPayloadProof(&_PayloadHelper.CallOpts, dstPayloadId_)
}

// SuperRegistry is a free data retrieval call binding the contract method 0x24c73dda.
//
// Solidity: function superRegistry() view returns(address)
func (_PayloadHelper *PayloadHelperCaller) SuperRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PayloadHelper.contract.Call(opts, &out, "superRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperRegistry is a free data retrieval call binding the contract method 0x24c73dda.
//
// Solidity: function superRegistry() view returns(address)
func (_PayloadHelper *PayloadHelperSession) SuperRegistry() (common.Address, error) {
	return _PayloadHelper.Contract.SuperRegistry(&_PayloadHelper.CallOpts)
}

// SuperRegistry is a free data retrieval call binding the contract method 0x24c73dda.
//
// Solidity: function superRegistry() view returns(address)
func (_PayloadHelper *PayloadHelperCallerSession) SuperRegistry() (common.Address, error) {
	return _PayloadHelper.Contract.SuperRegistry(&_PayloadHelper.CallOpts)
}
