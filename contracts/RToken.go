// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rtoken

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
)

// IRTokenMetaData contains all meta data concerning the IRToken contract.
var IRTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8ff8e15a": "addT(address)",
		"f7b7d6b6": "getTA(uint32)",
		"2df2685f": "getTI(address)",
		"3c7bdc19": "isValid(uint32)",
	},
}

// IRTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IRTokenMetaData.ABI instead.
var IRTokenABI = IRTokenMetaData.ABI

// Deprecated: Use IRTokenMetaData.Sigs instead.
// IRTokenFuncSigs maps the 4-byte function signature to its string representation.
var IRTokenFuncSigs = IRTokenMetaData.Sigs

// IRToken is an auto generated Go binding around an Ethereum contract.
type IRToken struct {
	IRTokenCaller     // Read-only binding to the contract
	IRTokenTransactor // Write-only binding to the contract
	IRTokenFilterer   // Log filterer for contract events
}

// IRTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRTokenSession struct {
	Contract     *IRToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRTokenCallerSession struct {
	Contract *IRTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IRTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRTokenTransactorSession struct {
	Contract     *IRTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IRTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRTokenRaw struct {
	Contract *IRToken // Generic contract binding to access the raw methods on
}

// IRTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRTokenCallerRaw struct {
	Contract *IRTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IRTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRTokenTransactorRaw struct {
	Contract *IRTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRToken creates a new instance of IRToken, bound to a specific deployed contract.
func NewIRToken(address common.Address, backend bind.ContractBackend) (*IRToken, error) {
	contract, err := bindIRToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRToken{IRTokenCaller: IRTokenCaller{contract: contract}, IRTokenTransactor: IRTokenTransactor{contract: contract}, IRTokenFilterer: IRTokenFilterer{contract: contract}}, nil
}

// NewIRTokenCaller creates a new read-only instance of IRToken, bound to a specific deployed contract.
func NewIRTokenCaller(address common.Address, caller bind.ContractCaller) (*IRTokenCaller, error) {
	contract, err := bindIRToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRTokenCaller{contract: contract}, nil
}

// NewIRTokenTransactor creates a new write-only instance of IRToken, bound to a specific deployed contract.
func NewIRTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IRTokenTransactor, error) {
	contract, err := bindIRToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRTokenTransactor{contract: contract}, nil
}

// NewIRTokenFilterer creates a new log filterer instance of IRToken, bound to a specific deployed contract.
func NewIRTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IRTokenFilterer, error) {
	contract, err := bindIRToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRTokenFilterer{contract: contract}, nil
}

// bindIRToken binds a generic wrapper to an already deployed contract.
func bindIRToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRToken *IRTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRToken.Contract.IRTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRToken *IRTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRToken.Contract.IRTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRToken *IRTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRToken.Contract.IRTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRToken *IRTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRToken *IRTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRToken *IRTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRToken.Contract.contract.Transact(opts, method, params...)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_IRToken *IRTokenCaller) GetTA(opts *bind.CallOpts, tIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _IRToken.contract.Call(opts, &out, "getTA", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_IRToken *IRTokenSession) GetTA(tIndex uint32) (common.Address, error) {
	return _IRToken.Contract.GetTA(&_IRToken.CallOpts, tIndex)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_IRToken *IRTokenCallerSession) GetTA(tIndex uint32) (common.Address, error) {
	return _IRToken.Contract.GetTA(&_IRToken.CallOpts, tIndex)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_IRToken *IRTokenCaller) GetTI(opts *bind.CallOpts, t common.Address) (uint32, bool, error) {
	var out []interface{}
	err := _IRToken.contract.Call(opts, &out, "getTI", t)

	if err != nil {
		return *new(uint32), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_IRToken *IRTokenSession) GetTI(t common.Address) (uint32, bool, error) {
	return _IRToken.Contract.GetTI(&_IRToken.CallOpts, t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_IRToken *IRTokenCallerSession) GetTI(t common.Address) (uint32, bool, error) {
	return _IRToken.Contract.GetTI(&_IRToken.CallOpts, t)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_IRToken *IRTokenCaller) IsValid(opts *bind.CallOpts, tIndex uint32) (bool, error) {
	var out []interface{}
	err := _IRToken.contract.Call(opts, &out, "isValid", tIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_IRToken *IRTokenSession) IsValid(tIndex uint32) (bool, error) {
	return _IRToken.Contract.IsValid(&_IRToken.CallOpts, tIndex)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_IRToken *IRTokenCallerSession) IsValid(tIndex uint32) (bool, error) {
	return _IRToken.Contract.IsValid(&_IRToken.CallOpts, tIndex)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_IRToken *IRTokenTransactor) AddT(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _IRToken.contract.Transact(opts, "addT", t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_IRToken *IRTokenSession) AddT(t common.Address) (*types.Transaction, error) {
	return _IRToken.Contract.AddT(&_IRToken.TransactOpts, t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_IRToken *IRTokenTransactorSession) AddT(t common.Address) (*types.Transaction, error) {
	return _IRToken.Contract.AddT(&_IRToken.TransactOpts, t)
}

// RTokenMetaData contains all meta data concerning the RToken contract.
var RTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"AddT\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"addT\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getTA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"getTI\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"isValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8ff8e15a": "addT(address)",
		"f7b7d6b6": "getTA(uint32)",
		"2df2685f": "getTI(address)",
		"b04c6c68": "getTNum()",
		"3c7bdc19": "isValid(uint32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600280546001600160a01b03191633179055610518806100326000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632df2685f1461005c5780633c7bdc19146100905780638ff8e15a146100b3578063b04c6c68146100db578063f7b7d6b6146100e3575b600080fd5b61006f61006a36600461042b565b61010e565b6040805163ffffffff90931683529015156020830152015b60405180910390f35b6100a361009e36600461045b565b6101e2565b6040519015158152602001610087565b6100c66100c136600461042b565b610259565b60405163ffffffff9091168152602001610087565b6000546100c6565b6100f66100f136600461045b565b61037c565b6040516001600160a01b039091168152602001610087565b6001600160a01b038116600090815260016020526040808220549051633c7bdc1960e01b815263ffffffff9091166004820181905282913090633c7bdc1990602401602060405180830381865afa15801561016d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101919190610481565b156101d85760008163ffffffff16815481106101af576101af6104a3565b6000918252602090912001546001600160a01b03858116911614156101d8579360019350915050565b9360009350915050565b6000805463ffffffff8316108015610244575060016000808463ffffffff1681548110610211576102116104a3565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff64010000000090910416155b1561025157506001919050565b506000919050565b6002546000906001600160a01b0316331461029f5760405162461bcd60e51b81526020600482015260026024820152614e4f60f01b604482015260640160405180910390fd5b6102a8826103c1565b156102cf57506001600160a01b031660009081526001602052604090205463ffffffff1690565b60008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563820180546001600160a01b0319166001600160a01b0387169081179091558084526020918252604093849020805463ffffffff191663ffffffff851690811790915584519182529181019190915290917f7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6910160405180910390a192915050565b6000805463ffffffff831610156102515760008263ffffffff16815481106103a6576103a66104a3565b6000918252602090912001546001600160a01b031692915050565b6000805b60005481101561042257826001600160a01b0316600082815481106103ec576103ec6104a3565b6000918252602090912001546001600160a01b031614156104105750600192915050565b8061041a816104b9565b9150506103c5565b50600092915050565b60006020828403121561043d57600080fd5b81356001600160a01b038116811461045457600080fd5b9392505050565b60006020828403121561046d57600080fd5b813563ffffffff8116811461045457600080fd5b60006020828403121561049357600080fd5b8151801515811461045457600080fd5b634e487b7160e01b600052603260045260246000fd5b60006000198214156104db57634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220bf1c4925813783305a650502cc68334133fdc4a8334f590ddd9a61173a5cc4e164736f6c634300080a0033",
}

// RTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RTokenMetaData.ABI instead.
var RTokenABI = RTokenMetaData.ABI

// Deprecated: Use RTokenMetaData.Sigs instead.
// RTokenFuncSigs maps the 4-byte function signature to its string representation.
var RTokenFuncSigs = RTokenMetaData.Sigs

// RTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RTokenMetaData.Bin instead.
var RTokenBin = RTokenMetaData.Bin

// DeployRToken deploys a new Ethereum contract, binding an instance of RToken to it.
func DeployRToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RToken, error) {
	parsed, err := RTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RToken{RTokenCaller: RTokenCaller{contract: contract}, RTokenTransactor: RTokenTransactor{contract: contract}, RTokenFilterer: RTokenFilterer{contract: contract}}, nil
}

// RToken is an auto generated Go binding around an Ethereum contract.
type RToken struct {
	RTokenCaller     // Read-only binding to the contract
	RTokenTransactor // Write-only binding to the contract
	RTokenFilterer   // Log filterer for contract events
}

// RTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RTokenSession struct {
	Contract     *RToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RTokenCallerSession struct {
	Contract *RTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RTokenTransactorSession struct {
	Contract     *RTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RTokenRaw struct {
	Contract *RToken // Generic contract binding to access the raw methods on
}

// RTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RTokenCallerRaw struct {
	Contract *RTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RTokenTransactorRaw struct {
	Contract *RTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRToken creates a new instance of RToken, bound to a specific deployed contract.
func NewRToken(address common.Address, backend bind.ContractBackend) (*RToken, error) {
	contract, err := bindRToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RToken{RTokenCaller: RTokenCaller{contract: contract}, RTokenTransactor: RTokenTransactor{contract: contract}, RTokenFilterer: RTokenFilterer{contract: contract}}, nil
}

// NewRTokenCaller creates a new read-only instance of RToken, bound to a specific deployed contract.
func NewRTokenCaller(address common.Address, caller bind.ContractCaller) (*RTokenCaller, error) {
	contract, err := bindRToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RTokenCaller{contract: contract}, nil
}

// NewRTokenTransactor creates a new write-only instance of RToken, bound to a specific deployed contract.
func NewRTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RTokenTransactor, error) {
	contract, err := bindRToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RTokenTransactor{contract: contract}, nil
}

// NewRTokenFilterer creates a new log filterer instance of RToken, bound to a specific deployed contract.
func NewRTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RTokenFilterer, error) {
	contract, err := bindRToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RTokenFilterer{contract: contract}, nil
}

// bindRToken binds a generic wrapper to an already deployed contract.
func bindRToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RToken *RTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RToken.Contract.RTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RToken *RTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RToken.Contract.RTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RToken *RTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RToken.Contract.RTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RToken *RTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RToken *RTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RToken *RTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RToken.Contract.contract.Transact(opts, method, params...)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_RToken *RTokenCaller) GetTA(opts *bind.CallOpts, tIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "getTA", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_RToken *RTokenSession) GetTA(tIndex uint32) (common.Address, error) {
	return _RToken.Contract.GetTA(&_RToken.CallOpts, tIndex)
}

// GetTA is a free data retrieval call binding the contract method 0xf7b7d6b6.
//
// Solidity: function getTA(uint32 tIndex) view returns(address)
func (_RToken *RTokenCallerSession) GetTA(tIndex uint32) (common.Address, error) {
	return _RToken.Contract.GetTA(&_RToken.CallOpts, tIndex)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_RToken *RTokenCaller) GetTI(opts *bind.CallOpts, t common.Address) (uint32, bool, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "getTI", t)

	if err != nil {
		return *new(uint32), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_RToken *RTokenSession) GetTI(t common.Address) (uint32, bool, error) {
	return _RToken.Contract.GetTI(&_RToken.CallOpts, t)
}

// GetTI is a free data retrieval call binding the contract method 0x2df2685f.
//
// Solidity: function getTI(address t) view returns(uint32, bool)
func (_RToken *RTokenCallerSession) GetTI(t common.Address) (uint32, bool, error) {
	return _RToken.Contract.GetTI(&_RToken.CallOpts, t)
}

// GetTNum is a free data retrieval call binding the contract method 0xb04c6c68.
//
// Solidity: function getTNum() view returns(uint32)
func (_RToken *RTokenCaller) GetTNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "getTNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTNum is a free data retrieval call binding the contract method 0xb04c6c68.
//
// Solidity: function getTNum() view returns(uint32)
func (_RToken *RTokenSession) GetTNum() (uint32, error) {
	return _RToken.Contract.GetTNum(&_RToken.CallOpts)
}

// GetTNum is a free data retrieval call binding the contract method 0xb04c6c68.
//
// Solidity: function getTNum() view returns(uint32)
func (_RToken *RTokenCallerSession) GetTNum() (uint32, error) {
	return _RToken.Contract.GetTNum(&_RToken.CallOpts)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_RToken *RTokenCaller) IsValid(opts *bind.CallOpts, tIndex uint32) (bool, error) {
	var out []interface{}
	err := _RToken.contract.Call(opts, &out, "isValid", tIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_RToken *RTokenSession) IsValid(tIndex uint32) (bool, error) {
	return _RToken.Contract.IsValid(&_RToken.CallOpts, tIndex)
}

// IsValid is a free data retrieval call binding the contract method 0x3c7bdc19.
//
// Solidity: function isValid(uint32 tIndex) view returns(bool)
func (_RToken *RTokenCallerSession) IsValid(tIndex uint32) (bool, error) {
	return _RToken.Contract.IsValid(&_RToken.CallOpts, tIndex)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_RToken *RTokenTransactor) AddT(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _RToken.contract.Transact(opts, "addT", t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_RToken *RTokenSession) AddT(t common.Address) (*types.Transaction, error) {
	return _RToken.Contract.AddT(&_RToken.TransactOpts, t)
}

// AddT is a paid mutator transaction binding the contract method 0x8ff8e15a.
//
// Solidity: function addT(address t) returns(uint32 index)
func (_RToken *RTokenTransactorSession) AddT(t common.Address) (*types.Transaction, error) {
	return _RToken.Contract.AddT(&_RToken.TransactOpts, t)
}

// RTokenAddTIterator is returned from FilterAddT and is used to iterate over the raw logs and unpacked data for AddT events raised by the RToken contract.
type RTokenAddTIterator struct {
	Event *RTokenAddT // Event containing the contract specifics and raw log

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
func (it *RTokenAddTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RTokenAddT)
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
		it.Event = new(RTokenAddT)
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
func (it *RTokenAddTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RTokenAddTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RTokenAddT represents a AddT event raised by the RToken contract.
type RTokenAddT struct {
	T      common.Address
	TIndex uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddT is a free log retrieval operation binding the contract event 0x7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6.
//
// Solidity: event AddT(address t, uint32 tIndex)
func (_RToken *RTokenFilterer) FilterAddT(opts *bind.FilterOpts) (*RTokenAddTIterator, error) {

	logs, sub, err := _RToken.contract.FilterLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return &RTokenAddTIterator{contract: _RToken.contract, event: "AddT", logs: logs, sub: sub}, nil
}

// WatchAddT is a free log subscription operation binding the contract event 0x7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6.
//
// Solidity: event AddT(address t, uint32 tIndex)
func (_RToken *RTokenFilterer) WatchAddT(opts *bind.WatchOpts, sink chan<- *RTokenAddT) (event.Subscription, error) {

	logs, sub, err := _RToken.contract.WatchLogs(opts, "AddT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RTokenAddT)
				if err := _RToken.contract.UnpackLog(event, "AddT", log); err != nil {
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

// ParseAddT is a log parse operation binding the contract event 0x7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6.
//
// Solidity: event AddT(address t, uint32 tIndex)
func (_RToken *RTokenFilterer) ParseAddT(log types.Log) (*RTokenAddT, error) {
	event := new(RTokenAddT)
	if err := _RToken.contract.UnpackLog(event, "AddT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
