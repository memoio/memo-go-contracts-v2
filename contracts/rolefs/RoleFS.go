// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolefs

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

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addr\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AirDrop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"airDrop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"fd1fc4a0": "airDrop(address[],uint256)",
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"42966c68": "burn(uint256)",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"f0141d84": "getDecimals()",
		"17d7de7c": "getName()",
		"15070401": "getSymbol()",
		"c4e41b22": "getTotalSupply()",
		"39509351": "increaseAllowance(address,uint256)",
		"79c65068": "mintToken(address,uint256)",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_IERC20 *IERC20Caller) GetDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_IERC20 *IERC20Session) GetDecimals() (uint8, error) {
	return _IERC20.Contract.GetDecimals(&_IERC20.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) GetDecimals() (uint8, error) {
	return _IERC20.Contract.GetDecimals(&_IERC20.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IERC20 *IERC20Caller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IERC20 *IERC20Session) GetName() (string, error) {
	return _IERC20.Contract.GetName(&_IERC20.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IERC20 *IERC20CallerSession) GetName() (string, error) {
	return _IERC20.Contract.GetName(&_IERC20.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_IERC20 *IERC20Caller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_IERC20 *IERC20Session) GetSymbol() (string, error) {
	return _IERC20.Contract.GetSymbol(&_IERC20.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_IERC20 *IERC20CallerSession) GetSymbol() (string, error) {
	return _IERC20.Contract.GetSymbol(&_IERC20.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "getTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) GetTotalSupply() (*big.Int, error) {
	return _IERC20.Contract.GetTotalSupply(&_IERC20.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) GetTotalSupply() (*big.Int, error) {
	return _IERC20.Contract.GetTotalSupply(&_IERC20.CallOpts)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfd1fc4a0.
//
// Solidity: function airDrop(address[] addrs, uint256 money) returns(bool)
func (_IERC20 *IERC20Transactor) AirDrop(opts *bind.TransactOpts, addrs []common.Address, money *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "airDrop", addrs, money)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfd1fc4a0.
//
// Solidity: function airDrop(address[] addrs, uint256 money) returns(bool)
func (_IERC20 *IERC20Session) AirDrop(addrs []common.Address, money *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.AirDrop(&_IERC20.TransactOpts, addrs, money)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfd1fc4a0.
//
// Solidity: function airDrop(address[] addrs, uint256 money) returns(bool)
func (_IERC20 *IERC20TransactorSession) AirDrop(addrs []common.Address, money *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.AirDrop(&_IERC20.TransactOpts, addrs, money)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_IERC20 *IERC20Transactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "burn", burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_IERC20 *IERC20Session) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Burn(&_IERC20.TransactOpts, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Burn(&_IERC20.TransactOpts, burnAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IERC20 *IERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IERC20 *IERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.DecreaseAllowance(&_IERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_IERC20 *IERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.DecreaseAllowance(&_IERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IERC20 *IERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IERC20 *IERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.IncreaseAllowance(&_IERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_IERC20 *IERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.IncreaseAllowance(&_IERC20.TransactOpts, spender, addedValue)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address target, uint256 mintedAmount) returns(bool)
func (_IERC20 *IERC20Transactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "mintToken", target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address target, uint256 mintedAmount) returns(bool)
func (_IERC20 *IERC20Session) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.MintToken(&_IERC20.TransactOpts, target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address target, uint256 mintedAmount) returns(bool)
func (_IERC20 *IERC20TransactorSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.MintToken(&_IERC20.TransactOpts, target, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20AirDropIterator is returned from FilterAirDrop and is used to iterate over the raw logs and unpacked data for AirDrop events raised by the IERC20 contract.
type IERC20AirDropIterator struct {
	Event *IERC20AirDrop // Event containing the contract specifics and raw log

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
func (it *IERC20AirDropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20AirDrop)
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
		it.Event = new(IERC20AirDrop)
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
func (it *IERC20AirDropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20AirDropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20AirDrop represents a AirDrop event raised by the IERC20 contract.
type IERC20AirDrop struct {
	Addr  []common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAirDrop is a free log retrieval operation binding the contract event 0xeeb9052f86d5d21d099b20a1a39748440e720a832b3b99482afa7152fa6bd906.
//
// Solidity: event AirDrop(address[] addr, uint256 value)
func (_IERC20 *IERC20Filterer) FilterAirDrop(opts *bind.FilterOpts) (*IERC20AirDropIterator, error) {

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "AirDrop")
	if err != nil {
		return nil, err
	}
	return &IERC20AirDropIterator{contract: _IERC20.contract, event: "AirDrop", logs: logs, sub: sub}, nil
}

// WatchAirDrop is a free log subscription operation binding the contract event 0xeeb9052f86d5d21d099b20a1a39748440e720a832b3b99482afa7152fa6bd906.
//
// Solidity: event AirDrop(address[] addr, uint256 value)
func (_IERC20 *IERC20Filterer) WatchAirDrop(opts *bind.WatchOpts, sink chan<- *IERC20AirDrop) (event.Subscription, error) {

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "AirDrop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20AirDrop)
				if err := _IERC20.contract.UnpackLog(event, "AirDrop", log); err != nil {
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

// ParseAirDrop is a log parse operation binding the contract event 0xeeb9052f86d5d21d099b20a1a39748440e720a832b3b99482afa7152fa6bd906.
//
// Solidity: event AirDrop(address[] addr, uint256 value)
func (_IERC20 *IERC20Filterer) ParseAirDrop(log types.Log) (*IERC20AirDrop, error) {
	event := new(IERC20AirDrop)
	if err := _IERC20.contract.UnpackLog(event, "AirDrop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileSysMetaData contains all meta data concerning the IFileSys contract.
var IFileSysMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newPro\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"payToken\",\"type\":\"uint32\"}],\"name\":\"createFs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"token\",\"type\":\"uint32\"}],\"name\":\"getChannelInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"}],\"name\":\"getFsInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"}],\"name\":\"getFsInfoAggOrder\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"}],\"name\":\"getFsPSum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getFsPro\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getSettleInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"token\",\"type\":\"uint32\"}],\"name\":\"getStoreInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"pAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newPro\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"subRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"roleType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"50cbb46f": "addKeeper(uint64)",
		"dee68733": "addOrder(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256)",
		"0f60c7b3": "addRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256)",
		"8026d040": "createFs(uint64,uint32)",
		"761966b7": "getBalance(uint64,uint32)",
		"cc122253": "getChannelInfo(uint64,uint64,uint32)",
		"324ab551": "getFsInfo(uint64)",
		"03eb18aa": "getFsInfoAggOrder(uint64,uint64)",
		"650c6a89": "getFsPSum(uint64)",
		"25d931e8": "getFsPro(uint64,uint64)",
		"7b31a24d": "getSettleInfo(uint64,uint32)",
		"3f5f363e": "getStoreInfo(uint64,uint64,uint32)",
		"69eafd02": "proWithdraw(uint64,uint32,address,address,uint256,uint256)",
		"e04f98ed": "recharge(uint64,uint32,address,address,uint256)",
		"8f6c3f03": "subOrder(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256)",
		"75818519": "subRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256)",
		"bab8d63a": "withdraw(uint64,uint32,uint8,address,address,uint256)",
	},
}

// IFileSysABI is the input ABI used to generate the binding from.
// Deprecated: Use IFileSysMetaData.ABI instead.
var IFileSysABI = IFileSysMetaData.ABI

// Deprecated: Use IFileSysMetaData.Sigs instead.
// IFileSysFuncSigs maps the 4-byte function signature to its string representation.
var IFileSysFuncSigs = IFileSysMetaData.Sigs

// IFileSys is an auto generated Go binding around an Ethereum contract.
type IFileSys struct {
	IFileSysCaller     // Read-only binding to the contract
	IFileSysTransactor // Write-only binding to the contract
	IFileSysFilterer   // Log filterer for contract events
}

// IFileSysCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileSysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileSysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileSysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileSysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileSysSession struct {
	Contract     *IFileSys         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileSysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileSysCallerSession struct {
	Contract *IFileSysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IFileSysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileSysTransactorSession struct {
	Contract     *IFileSysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFileSysRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileSysRaw struct {
	Contract *IFileSys // Generic contract binding to access the raw methods on
}

// IFileSysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileSysCallerRaw struct {
	Contract *IFileSysCaller // Generic read-only contract binding to access the raw methods on
}

// IFileSysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileSysTransactorRaw struct {
	Contract *IFileSysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileSys creates a new instance of IFileSys, bound to a specific deployed contract.
func NewIFileSys(address common.Address, backend bind.ContractBackend) (*IFileSys, error) {
	contract, err := bindIFileSys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileSys{IFileSysCaller: IFileSysCaller{contract: contract}, IFileSysTransactor: IFileSysTransactor{contract: contract}, IFileSysFilterer: IFileSysFilterer{contract: contract}}, nil
}

// NewIFileSysCaller creates a new read-only instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysCaller(address common.Address, caller bind.ContractCaller) (*IFileSysCaller, error) {
	contract, err := bindIFileSys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysCaller{contract: contract}, nil
}

// NewIFileSysTransactor creates a new write-only instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileSysTransactor, error) {
	contract, err := bindIFileSys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileSysTransactor{contract: contract}, nil
}

// NewIFileSysFilterer creates a new log filterer instance of IFileSys, bound to a specific deployed contract.
func NewIFileSysFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileSysFilterer, error) {
	contract, err := bindIFileSys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileSysFilterer{contract: contract}, nil
}

// bindIFileSys binds a generic wrapper to an already deployed contract.
func bindIFileSys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileSysABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSys *IFileSysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSys.Contract.IFileSysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSys *IFileSysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSys.Contract.IFileSysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSys *IFileSysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSys.Contract.IFileSysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileSys *IFileSysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileSys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileSys *IFileSysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileSys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileSys *IFileSysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileSys.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256, uint256)
func (_IFileSys *IFileSysCaller) GetBalance(opts *bind.CallOpts, index uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getBalance", index, tIndex)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256, uint256)
func (_IFileSys *IFileSysSession) GetBalance(index uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetBalance(&_IFileSys.CallOpts, index, tIndex)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256, uint256)
func (_IFileSys *IFileSysCallerSession) GetBalance(index uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetBalance(&_IFileSys.CallOpts, index, tIndex)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xcc122253.
//
// Solidity: function getChannelInfo(uint64 user, uint64 provider, uint32 token) view returns(uint256, uint64, uint64)
func (_IFileSys *IFileSysCaller) GetChannelInfo(opts *bind.CallOpts, user uint64, provider uint64, token uint32) (*big.Int, uint64, uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getChannelInfo", user, provider, token)

	if err != nil {
		return *new(*big.Int), *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return out0, out1, out2, err

}

// GetChannelInfo is a free data retrieval call binding the contract method 0xcc122253.
//
// Solidity: function getChannelInfo(uint64 user, uint64 provider, uint32 token) view returns(uint256, uint64, uint64)
func (_IFileSys *IFileSysSession) GetChannelInfo(user uint64, provider uint64, token uint32) (*big.Int, uint64, uint64, error) {
	return _IFileSys.Contract.GetChannelInfo(&_IFileSys.CallOpts, user, provider, token)
}

// GetChannelInfo is a free data retrieval call binding the contract method 0xcc122253.
//
// Solidity: function getChannelInfo(uint64 user, uint64 provider, uint32 token) view returns(uint256, uint64, uint64)
func (_IFileSys *IFileSysCallerSession) GetChannelInfo(user uint64, provider uint64, token uint32) (*big.Int, uint64, uint64, error) {
	return _IFileSys.Contract.GetChannelInfo(&_IFileSys.CallOpts, user, provider, token)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool, uint32)
func (_IFileSys *IFileSysCaller) GetFsInfo(opts *bind.CallOpts, user uint64) (bool, uint32, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsInfo", user)

	if err != nil {
		return *new(bool), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool, uint32)
func (_IFileSys *IFileSysSession) GetFsInfo(user uint64) (bool, uint32, error) {
	return _IFileSys.Contract.GetFsInfo(&_IFileSys.CallOpts, user)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool, uint32)
func (_IFileSys *IFileSysCallerSession) GetFsInfo(user uint64) (bool, uint32, error) {
	return _IFileSys.Contract.GetFsInfo(&_IFileSys.CallOpts, user)
}

// GetFsInfoAggOrder is a free data retrieval call binding the contract method 0x03eb18aa.
//
// Solidity: function getFsInfoAggOrder(uint64 user, uint64 provider) view returns(uint64, uint64)
func (_IFileSys *IFileSysCaller) GetFsInfoAggOrder(opts *bind.CallOpts, user uint64, provider uint64) (uint64, uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsInfoAggOrder", user, provider)

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetFsInfoAggOrder is a free data retrieval call binding the contract method 0x03eb18aa.
//
// Solidity: function getFsInfoAggOrder(uint64 user, uint64 provider) view returns(uint64, uint64)
func (_IFileSys *IFileSysSession) GetFsInfoAggOrder(user uint64, provider uint64) (uint64, uint64, error) {
	return _IFileSys.Contract.GetFsInfoAggOrder(&_IFileSys.CallOpts, user, provider)
}

// GetFsInfoAggOrder is a free data retrieval call binding the contract method 0x03eb18aa.
//
// Solidity: function getFsInfoAggOrder(uint64 user, uint64 provider) view returns(uint64, uint64)
func (_IFileSys *IFileSysCallerSession) GetFsInfoAggOrder(user uint64, provider uint64) (uint64, uint64, error) {
	return _IFileSys.Contract.GetFsInfoAggOrder(&_IFileSys.CallOpts, user, provider)
}

// GetFsPSum is a free data retrieval call binding the contract method 0x650c6a89.
//
// Solidity: function getFsPSum(uint64 user) view returns(uint64)
func (_IFileSys *IFileSysCaller) GetFsPSum(opts *bind.CallOpts, user uint64) (uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsPSum", user)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetFsPSum is a free data retrieval call binding the contract method 0x650c6a89.
//
// Solidity: function getFsPSum(uint64 user) view returns(uint64)
func (_IFileSys *IFileSysSession) GetFsPSum(user uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPSum(&_IFileSys.CallOpts, user)
}

// GetFsPSum is a free data retrieval call binding the contract method 0x650c6a89.
//
// Solidity: function getFsPSum(uint64 user) view returns(uint64)
func (_IFileSys *IFileSysCallerSession) GetFsPSum(user uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPSum(&_IFileSys.CallOpts, user)
}

// GetFsPro is a free data retrieval call binding the contract method 0x25d931e8.
//
// Solidity: function getFsPro(uint64 user, uint64 i) view returns(uint64)
func (_IFileSys *IFileSysCaller) GetFsPro(opts *bind.CallOpts, user uint64, i uint64) (uint64, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsPro", user, i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetFsPro is a free data retrieval call binding the contract method 0x25d931e8.
//
// Solidity: function getFsPro(uint64 user, uint64 i) view returns(uint64)
func (_IFileSys *IFileSysSession) GetFsPro(user uint64, i uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPro(&_IFileSys.CallOpts, user, i)
}

// GetFsPro is a free data retrieval call binding the contract method 0x25d931e8.
//
// Solidity: function getFsPro(uint64 user, uint64 i) view returns(uint64)
func (_IFileSys *IFileSysCallerSession) GetFsPro(user uint64, i uint64) (uint64, error) {
	return _IFileSys.Contract.GetFsPro(&_IFileSys.CallOpts, user, i)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x7b31a24d.
//
// Solidity: function getSettleInfo(uint64 index, uint32 tIndex) view returns(uint64, uint64, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_IFileSys *IFileSysCaller) GetSettleInfo(opts *bind.CallOpts, index uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getSettleInfo", index, tIndex)

	if err != nil {
		return *new(uint64), *new(uint64), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, err

}

// GetSettleInfo is a free data retrieval call binding the contract method 0x7b31a24d.
//
// Solidity: function getSettleInfo(uint64 index, uint32 tIndex) view returns(uint64, uint64, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_IFileSys *IFileSysSession) GetSettleInfo(index uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetSettleInfo(&_IFileSys.CallOpts, index, tIndex)
}

// GetSettleInfo is a free data retrieval call binding the contract method 0x7b31a24d.
//
// Solidity: function getSettleInfo(uint64 index, uint32 tIndex) view returns(uint64, uint64, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_IFileSys *IFileSysCallerSession) GetSettleInfo(index uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IFileSys.Contract.GetSettleInfo(&_IFileSys.CallOpts, index, tIndex)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x3f5f363e.
//
// Solidity: function getStoreInfo(uint64 user, uint64 provider, uint32 token) view returns(uint64, uint64, uint256)
func (_IFileSys *IFileSysCaller) GetStoreInfo(opts *bind.CallOpts, user uint64, provider uint64, token uint32) (uint64, uint64, *big.Int, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getStoreInfo", user, provider, token)

	if err != nil {
		return *new(uint64), *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetStoreInfo is a free data retrieval call binding the contract method 0x3f5f363e.
//
// Solidity: function getStoreInfo(uint64 user, uint64 provider, uint32 token) view returns(uint64, uint64, uint256)
func (_IFileSys *IFileSysSession) GetStoreInfo(user uint64, provider uint64, token uint32) (uint64, uint64, *big.Int, error) {
	return _IFileSys.Contract.GetStoreInfo(&_IFileSys.CallOpts, user, provider, token)
}

// GetStoreInfo is a free data retrieval call binding the contract method 0x3f5f363e.
//
// Solidity: function getStoreInfo(uint64 user, uint64 provider, uint32 token) view returns(uint64, uint64, uint256)
func (_IFileSys *IFileSysCallerSession) GetStoreInfo(user uint64, provider uint64, token uint32) (uint64, uint64, *big.Int, error) {
	return _IFileSys.Contract.GetStoreInfo(&_IFileSys.CallOpts, user, provider, token)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 kIndex) returns()
func (_IFileSys *IFileSysTransactor) AddKeeper(opts *bind.TransactOpts, kIndex uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addKeeper", kIndex)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 kIndex) returns()
func (_IFileSys *IFileSysSession) AddKeeper(kIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddKeeper(&_IFileSys.TransactOpts, kIndex)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 kIndex) returns()
func (_IFileSys *IFileSysTransactorSession) AddKeeper(kIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.AddKeeper(&_IFileSys.TransactOpts, kIndex)
}

// AddOrder is a paid mutator transaction binding the contract method 0xdee68733.
//
// Solidity: function addOrder(uint64 kIndex, uint64 uIndex, uint64 pIndex, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactor) AddOrder(opts *bind.TransactOpts, kIndex uint64, uIndex uint64, pIndex uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addOrder", kIndex, uIndex, pIndex, start, end, size, nonce, tokenIndex, sprice)
}

// AddOrder is a paid mutator transaction binding the contract method 0xdee68733.
//
// Solidity: function addOrder(uint64 kIndex, uint64 uIndex, uint64 pIndex, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysSession) AddOrder(kIndex uint64, uIndex uint64, pIndex uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, kIndex, uIndex, pIndex, start, end, size, nonce, tokenIndex, sprice)
}

// AddOrder is a paid mutator transaction binding the contract method 0xdee68733.
//
// Solidity: function addOrder(uint64 kIndex, uint64 uIndex, uint64 pIndex, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactorSession) AddOrder(kIndex uint64, uIndex uint64, pIndex uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, kIndex, uIndex, pIndex, start, end, size, nonce, tokenIndex, sprice)
}

// AddRepair is a paid mutator transaction binding the contract method 0x0f60c7b3.
//
// Solidity: function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactor) AddRepair(opts *bind.TransactOpts, kIndex uint64, pIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addRepair", kIndex, pIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// AddRepair is a paid mutator transaction binding the contract method 0x0f60c7b3.
//
// Solidity: function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysSession) AddRepair(kIndex uint64, pIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddRepair(&_IFileSys.TransactOpts, kIndex, pIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// AddRepair is a paid mutator transaction binding the contract method 0x0f60c7b3.
//
// Solidity: function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactorSession) AddRepair(kIndex uint64, pIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.AddRepair(&_IFileSys.TransactOpts, kIndex, pIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// CreateFs is a paid mutator transaction binding the contract method 0x8026d040.
//
// Solidity: function createFs(uint64 uIndex, uint32 payToken) returns()
func (_IFileSys *IFileSysTransactor) CreateFs(opts *bind.TransactOpts, uIndex uint64, payToken uint32) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "createFs", uIndex, payToken)
}

// CreateFs is a paid mutator transaction binding the contract method 0x8026d040.
//
// Solidity: function createFs(uint64 uIndex, uint32 payToken) returns()
func (_IFileSys *IFileSysSession) CreateFs(uIndex uint64, payToken uint32) (*types.Transaction, error) {
	return _IFileSys.Contract.CreateFs(&_IFileSys.TransactOpts, uIndex, payToken)
}

// CreateFs is a paid mutator transaction binding the contract method 0x8026d040.
//
// Solidity: function createFs(uint64 uIndex, uint32 payToken) returns()
func (_IFileSys *IFileSysTransactorSession) CreateFs(uIndex uint64, payToken uint32) (*types.Transaction, error) {
	return _IFileSys.Contract.CreateFs(&_IFileSys.TransactOpts, uIndex, payToken)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x69eafd02.
//
// Solidity: function proWithdraw(uint64 pIndex, uint32 tokenIndex, address pAddr, address tAddr, uint256 pay, uint256 lost) returns()
func (_IFileSys *IFileSysTransactor) ProWithdraw(opts *bind.TransactOpts, pIndex uint64, tokenIndex uint32, pAddr common.Address, tAddr common.Address, pay *big.Int, lost *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "proWithdraw", pIndex, tokenIndex, pAddr, tAddr, pay, lost)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x69eafd02.
//
// Solidity: function proWithdraw(uint64 pIndex, uint32 tokenIndex, address pAddr, address tAddr, uint256 pay, uint256 lost) returns()
func (_IFileSys *IFileSysSession) ProWithdraw(pIndex uint64, tokenIndex uint32, pAddr common.Address, tAddr common.Address, pay *big.Int, lost *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, pIndex, tokenIndex, pAddr, tAddr, pay, lost)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x69eafd02.
//
// Solidity: function proWithdraw(uint64 pIndex, uint32 tokenIndex, address pAddr, address tAddr, uint256 pay, uint256 lost) returns()
func (_IFileSys *IFileSysTransactorSession) ProWithdraw(pIndex uint64, tokenIndex uint32, pAddr common.Address, tAddr common.Address, pay *big.Int, lost *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, pIndex, tokenIndex, pAddr, tAddr, pay, lost)
}

// Recharge is a paid mutator transaction binding the contract method 0xe04f98ed.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, address uAddr, address tAddr, uint256 money) returns()
func (_IFileSys *IFileSysTransactor) Recharge(opts *bind.TransactOpts, uIndex uint64, tokenIndex uint32, uAddr common.Address, tAddr common.Address, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "recharge", uIndex, tokenIndex, uAddr, tAddr, money)
}

// Recharge is a paid mutator transaction binding the contract method 0xe04f98ed.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, address uAddr, address tAddr, uint256 money) returns()
func (_IFileSys *IFileSysSession) Recharge(uIndex uint64, tokenIndex uint32, uAddr common.Address, tAddr common.Address, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Recharge(&_IFileSys.TransactOpts, uIndex, tokenIndex, uAddr, tAddr, money)
}

// Recharge is a paid mutator transaction binding the contract method 0xe04f98ed.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, address uAddr, address tAddr, uint256 money) returns()
func (_IFileSys *IFileSysTransactorSession) Recharge(uIndex uint64, tokenIndex uint32, uAddr common.Address, tAddr common.Address, money *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Recharge(&_IFileSys.TransactOpts, uIndex, tokenIndex, uAddr, tAddr, money)
}

// SubOrder is a paid mutator transaction binding the contract method 0x8f6c3f03.
//
// Solidity: function subOrder(uint64 kIndex, uint64 uIndex, uint64 pIndex, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactor) SubOrder(opts *bind.TransactOpts, kIndex uint64, uIndex uint64, pIndex uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "subOrder", kIndex, uIndex, pIndex, start, end, size, nonce, tokenIndex, sprice)
}

// SubOrder is a paid mutator transaction binding the contract method 0x8f6c3f03.
//
// Solidity: function subOrder(uint64 kIndex, uint64 uIndex, uint64 pIndex, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysSession) SubOrder(kIndex uint64, uIndex uint64, pIndex uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, kIndex, uIndex, pIndex, start, end, size, nonce, tokenIndex, sprice)
}

// SubOrder is a paid mutator transaction binding the contract method 0x8f6c3f03.
//
// Solidity: function subOrder(uint64 kIndex, uint64 uIndex, uint64 pIndex, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactorSession) SubOrder(kIndex uint64, uIndex uint64, pIndex uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, kIndex, uIndex, pIndex, start, end, size, nonce, tokenIndex, sprice)
}

// SubRepair is a paid mutator transaction binding the contract method 0x75818519.
//
// Solidity: function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactor) SubRepair(opts *bind.TransactOpts, kIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "subRepair", kIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// SubRepair is a paid mutator transaction binding the contract method 0x75818519.
//
// Solidity: function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysSession) SubRepair(kIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.SubRepair(&_IFileSys.TransactOpts, kIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// SubRepair is a paid mutator transaction binding the contract method 0x75818519.
//
// Solidity: function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) returns()
func (_IFileSys *IFileSysTransactorSession) SubRepair(kIndex uint64, newPro uint64, start uint64, end uint64, size uint64, nonce uint64, tokenIndex uint32, sprice *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.SubRepair(&_IFileSys.TransactOpts, kIndex, newPro, start, end, size, nonce, tokenIndex, sprice)
}

// Withdraw is a paid mutator transaction binding the contract method 0xbab8d63a.
//
// Solidity: function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) returns()
func (_IFileSys *IFileSysTransactor) Withdraw(opts *bind.TransactOpts, index uint64, tokenIndex uint32, roleType uint8, tAddr common.Address, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "withdraw", index, tokenIndex, roleType, tAddr, addr, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xbab8d63a.
//
// Solidity: function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) returns()
func (_IFileSys *IFileSysSession) Withdraw(index uint64, tokenIndex uint32, roleType uint8, tAddr common.Address, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Withdraw(&_IFileSys.TransactOpts, index, tokenIndex, roleType, tAddr, addr, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xbab8d63a.
//
// Solidity: function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) returns()
func (_IFileSys *IFileSysTransactorSession) Withdraw(index uint64, tokenIndex uint32, roleType uint8, tAddr common.Address, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IFileSys.Contract.Withdraw(&_IFileSys.TransactOpts, index, tokenIndex, roleType, tAddr, addr, amount)
}

// IIssuanceMetaData contains all meta data concerning the IIssuance contract.
var IIssuanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setSP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"13df3828": "issu(uint64,uint64,uint64,uint256)",
		"6dd81f17": "setSP(uint256,uint256)",
		"11e65fc0": "setTP(uint256,uint256)",
	},
}

// IIssuanceABI is the input ABI used to generate the binding from.
// Deprecated: Use IIssuanceMetaData.ABI instead.
var IIssuanceABI = IIssuanceMetaData.ABI

// Deprecated: Use IIssuanceMetaData.Sigs instead.
// IIssuanceFuncSigs maps the 4-byte function signature to its string representation.
var IIssuanceFuncSigs = IIssuanceMetaData.Sigs

// IIssuance is an auto generated Go binding around an Ethereum contract.
type IIssuance struct {
	IIssuanceCaller     // Read-only binding to the contract
	IIssuanceTransactor // Write-only binding to the contract
	IIssuanceFilterer   // Log filterer for contract events
}

// IIssuanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IIssuanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IIssuanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IIssuanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IIssuanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IIssuanceSession struct {
	Contract     *IIssuance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IIssuanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IIssuanceCallerSession struct {
	Contract *IIssuanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IIssuanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IIssuanceTransactorSession struct {
	Contract     *IIssuanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IIssuanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IIssuanceRaw struct {
	Contract *IIssuance // Generic contract binding to access the raw methods on
}

// IIssuanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IIssuanceCallerRaw struct {
	Contract *IIssuanceCaller // Generic read-only contract binding to access the raw methods on
}

// IIssuanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IIssuanceTransactorRaw struct {
	Contract *IIssuanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIIssuance creates a new instance of IIssuance, bound to a specific deployed contract.
func NewIIssuance(address common.Address, backend bind.ContractBackend) (*IIssuance, error) {
	contract, err := bindIIssuance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IIssuance{IIssuanceCaller: IIssuanceCaller{contract: contract}, IIssuanceTransactor: IIssuanceTransactor{contract: contract}, IIssuanceFilterer: IIssuanceFilterer{contract: contract}}, nil
}

// NewIIssuanceCaller creates a new read-only instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceCaller(address common.Address, caller bind.ContractCaller) (*IIssuanceCaller, error) {
	contract, err := bindIIssuance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuanceCaller{contract: contract}, nil
}

// NewIIssuanceTransactor creates a new write-only instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IIssuanceTransactor, error) {
	contract, err := bindIIssuance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IIssuanceTransactor{contract: contract}, nil
}

// NewIIssuanceFilterer creates a new log filterer instance of IIssuance, bound to a specific deployed contract.
func NewIIssuanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IIssuanceFilterer, error) {
	contract, err := bindIIssuance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IIssuanceFilterer{contract: contract}, nil
}

// bindIIssuance binds a generic wrapper to an already deployed contract.
func bindIIssuance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IIssuanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuance *IIssuanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuance.Contract.IIssuanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuance *IIssuanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuance.Contract.IIssuanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuance *IIssuanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuance.Contract.IIssuanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IIssuance *IIssuanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IIssuance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IIssuance *IIssuanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IIssuance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IIssuance *IIssuanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IIssuance.Contract.contract.Transact(opts, method, params...)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 sPrice) returns(uint256)
func (_IIssuance *IIssuanceTransactor) Issu(opts *bind.TransactOpts, _start uint64, _end uint64, _size uint64, sPrice *big.Int) (*types.Transaction, error) {
	return _IIssuance.contract.Transact(opts, "issu", _start, _end, _size, sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 sPrice) returns(uint256)
func (_IIssuance *IIssuanceSession) Issu(_start uint64, _end uint64, _size uint64, sPrice *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, _start, _end, _size, sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 sPrice) returns(uint256)
func (_IIssuance *IIssuanceTransactorSession) Issu(_start uint64, _end uint64, _size uint64, sPrice *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, _start, _end, _size, sPrice)
}

// SetSP is a paid mutator transaction binding the contract method 0x6dd81f17.
//
// Solidity: function setSP(uint256 _price, uint256 _size) returns()
func (_IIssuance *IIssuanceTransactor) SetSP(opts *bind.TransactOpts, _price *big.Int, _size *big.Int) (*types.Transaction, error) {
	return _IIssuance.contract.Transact(opts, "setSP", _price, _size)
}

// SetSP is a paid mutator transaction binding the contract method 0x6dd81f17.
//
// Solidity: function setSP(uint256 _price, uint256 _size) returns()
func (_IIssuance *IIssuanceSession) SetSP(_price *big.Int, _size *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.SetSP(&_IIssuance.TransactOpts, _price, _size)
}

// SetSP is a paid mutator transaction binding the contract method 0x6dd81f17.
//
// Solidity: function setSP(uint256 _price, uint256 _size) returns()
func (_IIssuance *IIssuanceTransactorSession) SetSP(_price *big.Int, _size *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.SetSP(&_IIssuance.TransactOpts, _price, _size)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_IIssuance *IIssuanceTransactor) SetTP(opts *bind.TransactOpts, _add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _IIssuance.contract.Transact(opts, "setTP", _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_IIssuance *IIssuanceSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.SetTP(&_IIssuance.TransactOpts, _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_IIssuance *IIssuanceTransactorSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _IIssuance.Contract.SetTP(&_IIssuance.TransactOpts, _add, _sub)
}

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

// IRoleMetaData contains all meta data concerning the IRole contract.
var IRoleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"}],\"name\":\"addKeeperToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"addProviderToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"indexes\",\"type\":\"uint64[]\"},{\"internalType\":\"uint16\",\"name\":\"_level\",\"type\":\"uint16\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddrGindex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getFsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGKNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getRoleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taddr\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"payToken\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"name\":\"setGInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pPledge\",\"type\":\"uint256\"}],\"name\":\"setPledgeMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"withdrawFromFs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7977031c": "addKeeperToGroup(uint64,uint64)",
		"8ae21af3": "addProviderToGroup(uint64,uint64,bytes)",
		"cf8e99a8": "checkIR(uint64,uint8)",
		"de92e994": "checkT(uint32)",
		"f652391b": "createGroup(uint64[],uint16)",
		"9332aa6e": "getAddr(uint64)",
		"421795e5": "getAddrGindex(uint64)",
		"5f096376": "getFsAddr(uint64)",
		"429fb683": "getGKNum(uint64)",
		"07483499": "getRoleInfo(address)",
		"a6ed590b": "pledgeK()",
		"8ba61d28": "pledgeP()",
		"de909560": "pledgePool()",
		"517985b0": "recharge(uint64,uint32,uint256,bytes)",
		"24b8fbf6": "register(address,bytes)",
		"10e35bbe": "registerKeeper(uint64,bytes,bytes)",
		"d57e8a4e": "registerProvider(uint64,bytes)",
		"09824a80": "registerToken(address)",
		"e917e2c5": "registerUser(uint64,uint64,uint32,bytes,bytes)",
		"675841d4": "setGInfo(uint64,bool,uint256,uint256)",
		"97948fda": "setPledgeMoney(uint256,uint256)",
		"d30d0ce5": "withdrawFromFs(uint64,uint32,uint256,bytes)",
	},
}

// IRoleABI is the input ABI used to generate the binding from.
// Deprecated: Use IRoleMetaData.ABI instead.
var IRoleABI = IRoleMetaData.ABI

// Deprecated: Use IRoleMetaData.Sigs instead.
// IRoleFuncSigs maps the 4-byte function signature to its string representation.
var IRoleFuncSigs = IRoleMetaData.Sigs

// IRole is an auto generated Go binding around an Ethereum contract.
type IRole struct {
	IRoleCaller     // Read-only binding to the contract
	IRoleTransactor // Write-only binding to the contract
	IRoleFilterer   // Log filterer for contract events
}

// IRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoleSession struct {
	Contract     *IRole            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoleCallerSession struct {
	Contract *IRoleCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoleTransactorSession struct {
	Contract     *IRoleTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoleRaw struct {
	Contract *IRole // Generic contract binding to access the raw methods on
}

// IRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoleCallerRaw struct {
	Contract *IRoleCaller // Generic read-only contract binding to access the raw methods on
}

// IRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoleTransactorRaw struct {
	Contract *IRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRole creates a new instance of IRole, bound to a specific deployed contract.
func NewIRole(address common.Address, backend bind.ContractBackend) (*IRole, error) {
	contract, err := bindIRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRole{IRoleCaller: IRoleCaller{contract: contract}, IRoleTransactor: IRoleTransactor{contract: contract}, IRoleFilterer: IRoleFilterer{contract: contract}}, nil
}

// NewIRoleCaller creates a new read-only instance of IRole, bound to a specific deployed contract.
func NewIRoleCaller(address common.Address, caller bind.ContractCaller) (*IRoleCaller, error) {
	contract, err := bindIRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleCaller{contract: contract}, nil
}

// NewIRoleTransactor creates a new write-only instance of IRole, bound to a specific deployed contract.
func NewIRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoleTransactor, error) {
	contract, err := bindIRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoleTransactor{contract: contract}, nil
}

// NewIRoleFilterer creates a new log filterer instance of IRole, bound to a specific deployed contract.
func NewIRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoleFilterer, error) {
	contract, err := bindIRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoleFilterer{contract: contract}, nil
}

// bindIRole binds a generic wrapper to an already deployed contract.
func bindIRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRole *IRoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRole.Contract.IRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRole *IRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRole.Contract.IRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRole *IRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRole.Contract.IRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRole *IRoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRole *IRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRole *IRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRole.Contract.contract.Transact(opts, method, params...)
}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_IRole *IRoleCaller) CheckIR(opts *bind.CallOpts, _index uint64, _rType uint8) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "checkIR", _index, _rType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_IRole *IRoleSession) CheckIR(_index uint64, _rType uint8) (common.Address, error) {
	return _IRole.Contract.CheckIR(&_IRole.CallOpts, _index, _rType)
}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_IRole *IRoleCallerSession) CheckIR(_index uint64, _rType uint8) (common.Address, error) {
	return _IRole.Contract.CheckIR(&_IRole.CallOpts, _index, _rType)
}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_IRole *IRoleCaller) CheckT(opts *bind.CallOpts, tIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "checkT", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_IRole *IRoleSession) CheckT(tIndex uint32) (common.Address, error) {
	return _IRole.Contract.CheckT(&_IRole.CallOpts, tIndex)
}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_IRole *IRoleCallerSession) CheckT(tIndex uint32) (common.Address, error) {
	return _IRole.Contract.CheckT(&_IRole.CallOpts, tIndex)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_IRole *IRoleCaller) GetAddr(opts *bind.CallOpts, i uint64) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getAddr", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_IRole *IRoleSession) GetAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetAddr(&_IRole.CallOpts, i)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_IRole *IRoleCallerSession) GetAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetAddr(&_IRole.CallOpts, i)
}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_IRole *IRoleCaller) GetAddrGindex(opts *bind.CallOpts, i uint64) (common.Address, uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getAddrGindex", i)

	if err != nil {
		return *new(common.Address), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_IRole *IRoleSession) GetAddrGindex(i uint64) (common.Address, uint64, error) {
	return _IRole.Contract.GetAddrGindex(&_IRole.CallOpts, i)
}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_IRole *IRoleCallerSession) GetAddrGindex(i uint64) (common.Address, uint64, error) {
	return _IRole.Contract.GetAddrGindex(&_IRole.CallOpts, i)
}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_IRole *IRoleCaller) GetFsAddr(opts *bind.CallOpts, i uint64) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getFsAddr", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_IRole *IRoleSession) GetFsAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetFsAddr(&_IRole.CallOpts, i)
}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_IRole *IRoleCallerSession) GetFsAddr(i uint64) (common.Address, error) {
	return _IRole.Contract.GetFsAddr(&_IRole.CallOpts, i)
}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_IRole *IRoleCaller) GetGKNum(opts *bind.CallOpts, i uint64) (uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getGKNum", i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_IRole *IRoleSession) GetGKNum(i uint64) (uint64, error) {
	return _IRole.Contract.GetGKNum(&_IRole.CallOpts, i)
}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_IRole *IRoleCallerSession) GetGKNum(i uint64) (uint64, error) {
	return _IRole.Contract.GetGKNum(&_IRole.CallOpts, i)
}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_IRole *IRoleCaller) GetRoleInfo(opts *bind.CallOpts, acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getRoleInfo", acc)

	if err != nil {
		return *new(bool), *new(bool), *new(uint8), *new(uint64), *new(uint64), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)
	out3 := *abi.ConvertType(out[3], new(uint64)).(*uint64)
	out4 := *abi.ConvertType(out[4], new(uint64)).(*uint64)
	out5 := *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return out0, out1, out2, out3, out4, out5, err

}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_IRole *IRoleSession) GetRoleInfo(acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	return _IRole.Contract.GetRoleInfo(&_IRole.CallOpts, acc)
}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_IRole *IRoleCallerSession) GetRoleInfo(acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	return _IRole.Contract.GetRoleInfo(&_IRole.CallOpts, acc)
}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_IRole *IRoleCaller) PledgeK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "pledgeK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_IRole *IRoleSession) PledgeK() (*big.Int, error) {
	return _IRole.Contract.PledgeK(&_IRole.CallOpts)
}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_IRole *IRoleCallerSession) PledgeK() (*big.Int, error) {
	return _IRole.Contract.PledgeK(&_IRole.CallOpts)
}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_IRole *IRoleCaller) PledgeP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "pledgeP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_IRole *IRoleSession) PledgeP() (*big.Int, error) {
	return _IRole.Contract.PledgeP(&_IRole.CallOpts)
}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_IRole *IRoleCallerSession) PledgeP() (*big.Int, error) {
	return _IRole.Contract.PledgeP(&_IRole.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_IRole *IRoleCaller) PledgePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "pledgePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_IRole *IRoleSession) PledgePool() (common.Address, error) {
	return _IRole.Contract.PledgePool(&_IRole.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_IRole *IRoleCallerSession) PledgePool() (common.Address, error) {
	return _IRole.Contract.PledgePool(&_IRole.CallOpts)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_IRole *IRoleTransactor) AddKeeperToGroup(opts *bind.TransactOpts, _index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "addKeeperToGroup", _index, _gIndex)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_IRole *IRoleSession) AddKeeperToGroup(_index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _IRole.Contract.AddKeeperToGroup(&_IRole.TransactOpts, _index, _gIndex)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_IRole *IRoleTransactorSession) AddKeeperToGroup(_index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _IRole.Contract.AddKeeperToGroup(&_IRole.TransactOpts, _index, _gIndex)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_IRole *IRoleTransactor) AddProviderToGroup(opts *bind.TransactOpts, _index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "addProviderToGroup", _index, _gIndex, sign)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_IRole *IRoleSession) AddProviderToGroup(_index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.AddProviderToGroup(&_IRole.TransactOpts, _index, _gIndex, sign)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_IRole *IRoleTransactorSession) AddProviderToGroup(_index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.AddProviderToGroup(&_IRole.TransactOpts, _index, _gIndex, sign)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_IRole *IRoleTransactor) CreateGroup(opts *bind.TransactOpts, indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "createGroup", indexes, _level)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_IRole *IRoleSession) CreateGroup(indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _IRole.Contract.CreateGroup(&_IRole.TransactOpts, indexes, _level)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_IRole *IRoleTransactorSession) CreateGroup(indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _IRole.Contract.CreateGroup(&_IRole.TransactOpts, indexes, _level)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, uint256 money, bytes sign) payable returns()
func (_IRole *IRoleTransactor) Recharge(opts *bind.TransactOpts, uIndex uint64, tokenIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "recharge", uIndex, tokenIndex, money, sign)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, uint256 money, bytes sign) payable returns()
func (_IRole *IRoleSession) Recharge(uIndex uint64, tokenIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Recharge(&_IRole.TransactOpts, uIndex, tokenIndex, money, sign)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tokenIndex, uint256 money, bytes sign) payable returns()
func (_IRole *IRoleTransactorSession) Recharge(uIndex uint64, tokenIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Recharge(&_IRole.TransactOpts, uIndex, tokenIndex, money, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_IRole *IRoleTransactor) Register(opts *bind.TransactOpts, addr common.Address, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "register", addr, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_IRole *IRoleSession) Register(addr common.Address, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Register(&_IRole.TransactOpts, addr, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_IRole *IRoleTransactorSession) Register(addr common.Address, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.Register(&_IRole.TransactOpts, addr, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 index, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactor) RegisterKeeper(opts *bind.TransactOpts, index uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerKeeper", index, blsKey, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 index, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleSession) RegisterKeeper(index uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterKeeper(&_IRole.TransactOpts, index, blsKey, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 index, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactorSession) RegisterKeeper(index uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterKeeper(&_IRole.TransactOpts, index, blsKey, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 index, bytes sign) returns()
func (_IRole *IRoleTransactor) RegisterProvider(opts *bind.TransactOpts, index uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerProvider", index, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 index, bytes sign) returns()
func (_IRole *IRoleSession) RegisterProvider(index uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterProvider(&_IRole.TransactOpts, index, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 index, bytes sign) returns()
func (_IRole *IRoleTransactorSession) RegisterProvider(index uint64, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterProvider(&_IRole.TransactOpts, index, sign)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address taddr) returns()
func (_IRole *IRoleTransactor) RegisterToken(opts *bind.TransactOpts, taddr common.Address) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerToken", taddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address taddr) returns()
func (_IRole *IRoleSession) RegisterToken(taddr common.Address) (*types.Transaction, error) {
	return _IRole.Contract.RegisterToken(&_IRole.TransactOpts, taddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address taddr) returns()
func (_IRole *IRoleTransactorSession) RegisterToken(taddr common.Address) (*types.Transaction, error) {
	return _IRole.Contract.RegisterToken(&_IRole.TransactOpts, taddr)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xe917e2c5.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, uint32 payToken, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactor) RegisterUser(opts *bind.TransactOpts, index uint64, _gIndex uint64, payToken uint32, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerUser", index, _gIndex, payToken, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xe917e2c5.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, uint32 payToken, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleSession) RegisterUser(index uint64, _gIndex uint64, payToken uint32, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterUser(&_IRole.TransactOpts, index, _gIndex, payToken, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0xe917e2c5.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, uint32 payToken, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactorSession) RegisterUser(index uint64, _gIndex uint64, payToken uint32, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterUser(&_IRole.TransactOpts, index, _gIndex, payToken, blsKey, sign)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x675841d4.
//
// Solidity: function setGInfo(uint64 index, bool isAdd, uint256 _size, uint256 sPrice) returns()
func (_IRole *IRoleTransactor) SetGInfo(opts *bind.TransactOpts, index uint64, isAdd bool, _size *big.Int, sPrice *big.Int) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "setGInfo", index, isAdd, _size, sPrice)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x675841d4.
//
// Solidity: function setGInfo(uint64 index, bool isAdd, uint256 _size, uint256 sPrice) returns()
func (_IRole *IRoleSession) SetGInfo(index uint64, isAdd bool, _size *big.Int, sPrice *big.Int) (*types.Transaction, error) {
	return _IRole.Contract.SetGInfo(&_IRole.TransactOpts, index, isAdd, _size, sPrice)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x675841d4.
//
// Solidity: function setGInfo(uint64 index, bool isAdd, uint256 _size, uint256 sPrice) returns()
func (_IRole *IRoleTransactorSession) SetGInfo(index uint64, isAdd bool, _size *big.Int, sPrice *big.Int) (*types.Transaction, error) {
	return _IRole.Contract.SetGInfo(&_IRole.TransactOpts, index, isAdd, _size, sPrice)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_IRole *IRoleTransactor) SetPledgeMoney(opts *bind.TransactOpts, kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "setPledgeMoney", kPledge, pPledge)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_IRole *IRoleSession) SetPledgeMoney(kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _IRole.Contract.SetPledgeMoney(&_IRole.TransactOpts, kPledge, pPledge)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_IRole *IRoleTransactorSession) SetPledgeMoney(kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _IRole.Contract.SetPledgeMoney(&_IRole.TransactOpts, kPledge, pPledge)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_IRole *IRoleTransactor) WithdrawFromFs(opts *bind.TransactOpts, index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "withdrawFromFs", index, tIndex, amount, sign)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_IRole *IRoleSession) WithdrawFromFs(index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.WithdrawFromFs(&_IRole.TransactOpts, index, tIndex, amount, sign)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_IRole *IRoleTransactorSession) WithdrawFromFs(index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.WithdrawFromFs(&_IRole.TransactOpts, index, tIndex, amount, sign)
}

// RecoverMetaData contains all meta data concerning the Recover contract.
var RecoverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"19045a25": "recover(bytes32,bytes)",
	},
	Bin: "0x6102a861003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806319045a251461003a575b600080fd5b61004d610048366004610184565b610069565b6040516001600160a01b03909116815260200160405180910390f35b6000815160411461007c57506000610168565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156100c25760009350505050610168565b601b8160ff1610156100dc576100d981601b61023f565b90505b8060ff16601b141580156100f457508060ff16601c14155b156101055760009350505050610168565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610158573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561019757600080fd5b82359150602083013567ffffffffffffffff808211156101b657600080fd5b818501915085601f8301126101ca57600080fd5b8135818111156101dc576101dc61016e565b604051601f8201601f19908116603f011681019083821181831017156102045761020461016e565b8160405282815288602084870101111561021d57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600060ff821660ff84168060ff0382111561026a57634e487b7160e01b600052601160045260246000fd5b01939250505056fea2646970667358221220bb3ace1f08ca73947a397f9e9666b18a94fd4184ae056cd3e97c6316bd3c25ad64736f6c634300080a0033",
}

// RecoverABI is the input ABI used to generate the binding from.
// Deprecated: Use RecoverMetaData.ABI instead.
var RecoverABI = RecoverMetaData.ABI

// Deprecated: Use RecoverMetaData.Sigs instead.
// RecoverFuncSigs maps the 4-byte function signature to its string representation.
var RecoverFuncSigs = RecoverMetaData.Sigs

// RecoverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecoverMetaData.Bin instead.
var RecoverBin = RecoverMetaData.Bin

// DeployRecover deploys a new Ethereum contract, binding an instance of Recover to it.
func DeployRecover(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recover, error) {
	parsed, err := RecoverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecoverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// Recover is an auto generated Go binding around an Ethereum contract.
type Recover struct {
	RecoverCaller     // Read-only binding to the contract
	RecoverTransactor // Write-only binding to the contract
	RecoverFilterer   // Log filterer for contract events
}

// RecoverCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecoverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecoverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecoverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecoverSession struct {
	Contract     *Recover          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecoverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecoverCallerSession struct {
	Contract *RecoverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RecoverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecoverTransactorSession struct {
	Contract     *RecoverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RecoverRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecoverRaw struct {
	Contract *Recover // Generic contract binding to access the raw methods on
}

// RecoverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecoverCallerRaw struct {
	Contract *RecoverCaller // Generic read-only contract binding to access the raw methods on
}

// RecoverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecoverTransactorRaw struct {
	Contract *RecoverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecover creates a new instance of Recover, bound to a specific deployed contract.
func NewRecover(address common.Address, backend bind.ContractBackend) (*Recover, error) {
	contract, err := bindRecover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recover{RecoverCaller: RecoverCaller{contract: contract}, RecoverTransactor: RecoverTransactor{contract: contract}, RecoverFilterer: RecoverFilterer{contract: contract}}, nil
}

// NewRecoverCaller creates a new read-only instance of Recover, bound to a specific deployed contract.
func NewRecoverCaller(address common.Address, caller bind.ContractCaller) (*RecoverCaller, error) {
	contract, err := bindRecover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverCaller{contract: contract}, nil
}

// NewRecoverTransactor creates a new write-only instance of Recover, bound to a specific deployed contract.
func NewRecoverTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoverTransactor, error) {
	contract, err := bindRecover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoverTransactor{contract: contract}, nil
}

// NewRecoverFilterer creates a new log filterer instance of Recover, bound to a specific deployed contract.
func NewRecoverFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoverFilterer, error) {
	contract, err := bindRecover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoverFilterer{contract: contract}, nil
}

// bindRecover binds a generic wrapper to an already deployed contract.
func bindRecover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RecoverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.RecoverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.RecoverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recover *RecoverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recover *RecoverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recover *RecoverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recover.Contract.contract.Transact(opts, method, params...)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCaller) Recover(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Recover.contract.Call(opts, &out, "recover", hash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_Recover *RecoverCallerSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _Recover.Contract.Recover(&_Recover.CallOpts, hash, signature)
}

// RoleFSMetaData contains all meta data concerning the RoleFS contract.
var RoleFSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nPIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"kSigns\",\"type\":\"bytes[]\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"iss\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"r\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rt\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nPIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pSign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"kSigns\",\"type\":\"bytes[]\"}],\"name\":\"subRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"991d1195": "addOrder(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[])",
		"ffd0fd09": "addRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes[])",
		"b28ff3c9": "proWithdraw(uint64,uint32,uint256,uint256,bytes[])",
		"a7e7797d": "setAddr(address,address,address,address)",
		"08bf601e": "subOrder(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[])",
		"eb975779": "subRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50600380546001600160a01b031916331790556125c5806100326000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806308bf601e14610067578063991d11951461007c578063a7e7797d1461008f578063b28ff3c9146100a2578063eb975779146100b5578063ffd0fd09146100c8575b600080fd5b61007a610075366004611e86565b6100db565b005b61007a61008a366004611e86565b6102e9565b61007a61009d366004611fab565b610884565b61007a6100b0366004612007565b610912565b61007a6100c336600461207a565b610de5565b61007a6100d636600461207a565b6113b1565b6000546040516319f1d33560e31b81526001600160401b038d166004820152600160248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015610133573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610157919061215d565b506000546040516319f1d33560e31b81526001600160401b038c166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa1580156101b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d4919061215d565b506000806101e4338e8e8a61195a565b60005491935091506001600160a01b0316635f096376610205600184612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015610249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061026d919061215d565b6001600160a01b0316638f6c3f03838f8f8f8f8f8f8f8f6040518a63ffffffff1660e01b81526004016102a8999897969594939291906121bf565b600060405180830381600087803b1580156102c257600080fd5b505af11580156102d6573d6000803e3d6000fd5b5050505050505050505050505050505050565b6000876001600160401b03161161032c5760405162461bcd60e51b815260206004820152600260248201526139bd60f11b60448201526064015b60405180910390fd5b886001600160401b0316886001600160401b0316116103725760405162461bcd60e51b8152602060048201526002602482015261657360f01b6044820152606401610323565b6001600160401b038816610389620151808a612218565b610396906201518061224c565b6001600160401b0316146103d15760405162461bcd60e51b8152602060048201526002602482015261746560f01b6044820152606401610323565b6000546040516319f1d33560e31b81526001600160401b038d166004820152600160248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015610429573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044d919061215d565b506000546040516319f1d33560e31b81526001600160401b038c166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa1580156104a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ca919061215d565b506000806104da338e8e8a61195a565b60005491935091506001600160a01b0316635f0963766104fb600184612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa15801561053f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610563919061215d565b6001600160a01b031663dee68733838f8f8f8f8f8f8f8f6040518a63ffffffff1660e01b815260040161059e999897969594939291906121bf565b600060405180830381600087803b1580156105b857600080fd5b505af11580156105cc573d6000803e3d6000fd5b5050505063ffffffff87166108745760025460405163027be70560e31b81526001600160401b03808e166004830152808d1660248301528b166044820152606481018890526001600160a01b039091169060009082906313df3828906084016020604051808303816000875af115801561064a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066e919061227b565b6000549091506001600160a01b031663675841d461068d600186612197565b6040516001600160e01b031960e084901b1681526001600160401b03918216600482015260016024820152908e166044820152606481018b9052608401600060405180830381600087803b1580156106e457600080fd5b505af11580156106f8573d6000803e3d6000fd5b50505050806000141561070e5750505050610877565b60048054604051637bdbeb5b60e11b815260009281018390526001600160a01b039091169063f7b7d6b690602401602060405180830381865afa158015610759573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077d919061215d565b9050806001600160a01b03166379c6506860008054906101000a90046001600160a01b03166001600160a01b031663de9095606040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107df573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610803919061215d565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018590526044016020604051808303816000875af1158015610850573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d691906122a4565b50505b5050505050505050505050565b6003546001600160a01b031633146108c25760405162461bcd60e51b81526020600482015260016024820152602760f91b6044820152606401610323565b600280546001600160a01b039586166001600160a01b0319918216179091556000805494861694821694909417909355600180549285169284169290921790915560048054919093169116179055565b6000805481906001600160a01b031663421795e561093160018a612197565b6040516001600160e01b031960e084901b1681526001600160401b0390911660048201526024016040805180830381865afa158015610974573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099891906122bf565b60008054604051630748349960e01b81526001600160a01b03808616600483015294965092945090928392839283921690630748349990602401600060405180830381865afa1580156109ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610a179190810190612329565b505050925092509250828015610a2b575081155b8015610a3a57508060ff166002145b610a6b5760405162461bcd60e51b8152602060048201526002602482015261049560f41b6044820152606401610323565b6000546003906001600160a01b031663429fb683610a8a600189612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015610ace573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af29190612402565b610afd90600261224c565b610b079190612218565b6001600160401b031693508387511015610b485760405162461bcd60e51b8152602060048201526002602482015261534560f01b6044820152606401610323565b5050600080549091506001600160a01b0316635f096376610b6a600186612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015610bae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bd2919061215d565b6000546040516337a4ba6560e21b815263ffffffff8b1660048201529192506001600160a01b03808416926369eafd02928d928d928a92919091169063de92e99490602401602060405180830381865afa158015610c34573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c58919061215d565b6040516001600160e01b031960e087901b1681526001600160401b03909416600485015263ffffffff9290921660248401526001600160a01b039081166044840152166064820152608481018a905260a4810189905260c401600060405180830381600087803b158015610ccb57600080fd5b505af1158015610cdf573d6000803e3d6000fd5b5050505063ffffffff8816610dda57604051637b31a24d60e01b81526001600160401b038a16600482015263ffffffff891660248201526001600160a01b03821690637b31a24d9060440161016060405180830381865afa158015610d48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6c919061241f565b50506002546040516247997f60e61b8152939b506001600160a01b031698506311e65fc09750610dac96508e95508a945050600490910191506124b39050565b600060405180830381600087803b158015610dc657600080fd5b505af1158015610874573d6000803e3d6000fd5b505050505050505050565b6000546040516319f1d33560e31b81526001600160401b038c166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015610e3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e61919061215d565b506000546040516319f1d33560e31b81526001600160401b038b166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015610eba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ede919061215d565b50600080610eee338d8d8961195a565b9150915060008c8b8b8b8b8b8b604051602001610f71979695949392919060c097881b6001600160c01b0319908116825296881b8716600882015294871b8616601086015292861b85166018850152941b909216602082015260e09290921b6001600160e01b0319166028830152602c820152607360f81b604c820152604d0190565b60405160208183030381529060405280519060200120905061102a60008054906101000a90046001600160a01b03166001600160a01b0316639332aa6e60018f610fbb9190612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015610fff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611023919061215d565b8287611c24565b600080546003906001600160a01b031663429fb68361104a600187612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa15801561108e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b29190612402565b6110bd90600261224c565b6110c79190612218565b6001600160401b0316905080855110156110f35760405162461bcd60e51b8152600401610323906124c1565b60008060005b875181101561124f578473__$6230a6feddd2d01b0a9ffb5c5e188a1065$__6319045a2590918a8481518110611131576111316124de565b60200260200101516040518363ffffffff1660e01b81526004016111569291906124f4565b602060405180830381865af4158015611173573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611197919061215d565b60008054604051630748349960e01b81526001600160a01b03808516600483015293955091921690630748349990602401600060405180830381865afa1580156111e5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261120d9190810190612329565b50945050505050866001600160401b0316816001600160401b0316141561123c57836112388161252e565b9450505b508061124781612555565b9150506110f9565b5082826001600160401b031610156112795760405162461bcd60e51b8152600401610323906124c1565b50506000546001600160a01b03169150635f096376905061129b600184612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa1580156112df573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611303919061215d565b604051637581851960e01b81526001600160401b038085166004830152808e166024830152808d166044830152808c166064830152808b166084830152891660a482015263ffffffff881660c482015260e481018790526001600160a01b03919091169063758185199061010401600060405180830381600087803b15801561138b57600080fd5b505af115801561139f573d6000803e3d6000fd5b50505050505050505050505050505050565b6000546040516319f1d33560e31b81526001600160401b038c166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015611409573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142d919061215d565b506000546040516319f1d33560e31b81526001600160401b038b166004820152600260248201526001600160a01b039091169063cf8e99a890604401602060405180830381865afa158015611486573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114aa919061215d565b506000806114ba338d8d8961195a565b9150915060008c8b8b8b8b8b8b60405160200161153d979695949392919060c097881b6001600160c01b0319908116825296881b8716600882015294871b8616601086015292861b85166018850152941b909216602082015260e09290921b6001600160e01b0319166028830152602c820152606160f81b604c820152604d0190565b60405160208183030381529060405280519060200120905061158760008054906101000a90046001600160a01b03166001600160a01b0316639332aa6e60018f610fbb9190612197565b600080546003906001600160a01b031663429fb6836115a7600187612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa1580156115eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061160f9190612402565b61161a90600261224c565b6116249190612218565b6001600160401b0316905080855110156116505760405162461bcd60e51b8152600401610323906124c1565b60008060005b87518110156117ac578473__$6230a6feddd2d01b0a9ffb5c5e188a1065$__6319045a2590918a848151811061168e5761168e6124de565b60200260200101516040518363ffffffff1660e01b81526004016116b39291906124f4565b602060405180830381865af41580156116d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f4919061215d565b60008054604051630748349960e01b81526001600160a01b03808516600483015293955091921690630748349990602401600060405180830381865afa158015611742573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261176a9190810190612329565b50945050505050866001600160401b0316816001600160401b0316141561179957836117958161252e565b9450505b50806117a481612555565b915050611656565b5082826001600160401b031610156117d65760405162461bcd60e51b8152600401610323906124c1565b50506000546001600160a01b03169150635f09637690506117f8600184612197565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa15801561183c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611860919061215d565b6001600160a01b0316630f60c7b3838e8e8e8e8e8e8e8e6040518a63ffffffff1660e01b815260040161189b999897969594939291906121bf565b600060405180830381600087803b1580156118b557600080fd5b505af11580156118c9573d6000803e3d6000fd5b5050505063ffffffff861661194c57896001600160401b0316896001600160401b0316111561194c576002546001600160a01b0316806311e65fc060006119108e8e612197565b611923906001600160401b03168a612570565b6040516001600160e01b031960e085901b168152600481019290925260248201526044016102a8565b505050505050505050505050565b600080546040516337a4ba6560e21b815263ffffffff8416600482015282916001600160a01b031690819063de92e99490602401602060405180830381865afa1580156119ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119cf919061215d565b50604051630748349960e01b81526001600160a01b038881166004830152600091829182918291829190871690630748349990602401600060405180830381865afa158015611a22573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611a4a9190810190612329565b5094509450945094509450848015611a60575083155b8015611a6f57508260ff166003145b611aa05760405162461bcd60e51b8152602060048201526002602482015261434560f01b6044820152606401610323565b60006001600160a01b03871663421795e5611abc60018f612197565b6040516001600160e01b031960e084901b1681526001600160401b0390911660048201526024016040805180830381865afa158015611aff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b2391906122bf565b9150600090506001600160a01b03881663421795e5611b4360018f612197565b6040516001600160e01b031960e084901b1681526001600160401b0390911660048201526024016040805180830381865afa158015611b86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611baa91906122bf565b915050806001600160401b0316826001600160401b0316148015611bdf5750826001600160401b0316826001600160401b0316145b611c105760405162461bcd60e51b815260206004820152600260248201526111d160f21b6044820152606401610323565b50919c919b50909950505050505050505050565b6040516319045a2560e01b81526001600160a01b0384169073__$6230a6feddd2d01b0a9ffb5c5e188a1065$__906319045a2590611c6890869086906004016124f4565b602060405180830381865af4158015611c85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ca9919061215d565b6001600160a01b031614611ce45760405162461bcd60e51b8152602060048201526002602482015261494360f01b6044820152606401610323565b505050565b6001600160401b0381168114611cfe57600080fd5b50565b8035611d0c81611ce9565b919050565b803563ffffffff81168114611d0c57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715611d6357611d63611d25565b604052919050565b60006001600160401b03821115611d8457611d84611d25565b50601f01601f191660200190565b600082601f830112611da357600080fd5b8135611db6611db182611d6b565b611d3b565b818152846020838601011115611dcb57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112611df957600080fd5b813560206001600160401b0380831115611e1557611e15611d25565b8260051b611e24838201611d3b565b9384528581018301938381019088861115611e3e57600080fd5b84880192505b85831015611e7a57823584811115611e5c5760008081fd5b611e6a8a87838c0101611d92565b8352509184019190840190611e44565b98975050505050505050565b60008060008060008060008060008060006101608c8e031215611ea857600080fd5b611eb18c611d01565b9a50611ebf60208d01611d01565b9950611ecd60408d01611d01565b9850611edb60608d01611d01565b9750611ee960808d01611d01565b9650611ef760a08d01611d01565b9550611f0560c08d01611d11565b945060e08c013593506001600160401b03806101008e01351115611f2857600080fd5b611f398e6101008f01358f01611d92565b9350806101208e01351115611f4d57600080fd5b611f5e8e6101208f01358f01611d92565b9250806101408e01351115611f7257600080fd5b50611f848d6101408e01358e01611de8565b90509295989b509295989b9093969950565b6001600160a01b0381168114611cfe57600080fd5b60008060008060808587031215611fc157600080fd5b8435611fcc81611f96565b93506020850135611fdc81611f96565b92506040850135611fec81611f96565b91506060850135611ffc81611f96565b939692955090935050565b600080600080600060a0868803121561201f57600080fd5b853561202a81611ce9565b945061203860208701611d11565b9350604086013592506060860135915060808601356001600160401b0381111561206157600080fd5b61206d88828901611de8565b9150509295509295909350565b6000806000806000806000806000806101408b8d03121561209a57600080fd5b6120a38b611d01565b99506120b160208c01611d01565b98506120bf60408c01611d01565b97506120cd60608c01611d01565b96506120db60808c01611d01565b95506120e960a08c01611d01565b94506120f760c08c01611d11565b935060e08b013592506101008b01356001600160401b038082111561211b57600080fd5b6121278e838f01611d92565b93506101208d013591508082111561213e57600080fd5b5061214b8d828e01611de8565b9150509295989b9194979a5092959850565b60006020828403121561216f57600080fd5b815161217a81611f96565b9392505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03838116908316818110156121b7576121b7612181565b039392505050565b6001600160401b03998a1681529789166020890152958816604088015293871660608701529186166080860152851660a085015290931660c083015263ffffffff90921660e08201526101008101919091526101200190565b60006001600160401b038084168061224057634e487b7160e01b600052601260045260246000fd5b92169190910492915050565b60006001600160401b038083168185168183048111821515161561227257612272612181565b02949350505050565b60006020828403121561228d57600080fd5b5051919050565b80518015158114611d0c57600080fd5b6000602082840312156122b657600080fd5b61217a82612294565b600080604083850312156122d257600080fd5b82516122dd81611f96565b60208401519092506122ee81611ce9565b809150509250929050565b60005b838110156123145781810151838201526020016122fc565b83811115612323576000848401525b50505050565b60008060008060008060c0878903121561234257600080fd5b61234b87612294565b955061235960208801612294565b9450604087015160ff8116811461236f57600080fd5b606088015190945061238081611ce9565b608088015190935061239181611ce9565b60a08801519092506001600160401b038111156123ad57600080fd5b8701601f810189136123be57600080fd5b80516123cc611db182611d6b565b8181528a60208385010111156123e157600080fd5b6123f28260208301602086016122f9565b8093505050509295509295509295565b60006020828403121561241457600080fd5b815161217a81611ce9565b60008060008060008060008060008060006101608c8e03121561244157600080fd5b8b5161244c81611ce9565b60208d0151909b5061245d81611ce9565b809a505060408c0151985060608c0151975060808c0151965060a08c0151955060c08c0151945060e08c015193506101008c015192506101208c015191506101408c015190509295989b509295989b9093969950565b918252602082015260400190565b602080825260039082015262534e4560e81b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b82815260406020820152600082518060408401526125198160608501602087016122f9565b601f01601f1916919091016060019392505050565b60006001600160401b038083168181141561254b5761254b612181565b6001019392505050565b600060001982141561256957612569612181565b5060010190565b600081600019048311821515161561258a5761258a612181565b50029056fea26469706673582212200530fced51ad3bab80d5d91dc13d20f3d2615103efc8b779c037a0a8ad9065e864736f6c634300080a0033",
}

// RoleFSABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleFSMetaData.ABI instead.
var RoleFSABI = RoleFSMetaData.ABI

// Deprecated: Use RoleFSMetaData.Sigs instead.
// RoleFSFuncSigs maps the 4-byte function signature to its string representation.
var RoleFSFuncSigs = RoleFSMetaData.Sigs

// RoleFSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoleFSMetaData.Bin instead.
var RoleFSBin = RoleFSMetaData.Bin

// DeployRoleFS deploys a new Ethereum contract, binding an instance of RoleFS to it.
func DeployRoleFS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoleFS, error) {
	parsed, err := RoleFSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	recoverAddr, _, _, _ := DeployRecover(auth, backend)
	RoleFSBin = strings.Replace(RoleFSBin, "__$6230a6feddd2d01b0a9ffb5c5e188a1065$__", recoverAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoleFSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoleFS{RoleFSCaller: RoleFSCaller{contract: contract}, RoleFSTransactor: RoleFSTransactor{contract: contract}, RoleFSFilterer: RoleFSFilterer{contract: contract}}, nil
}

// RoleFS is an auto generated Go binding around an Ethereum contract.
type RoleFS struct {
	RoleFSCaller     // Read-only binding to the contract
	RoleFSTransactor // Write-only binding to the contract
	RoleFSFilterer   // Log filterer for contract events
}

// RoleFSCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleFSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleFSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleFSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleFSSession struct {
	Contract     *RoleFS           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleFSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleFSCallerSession struct {
	Contract *RoleFSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoleFSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleFSTransactorSession struct {
	Contract     *RoleFSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleFSRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleFSRaw struct {
	Contract *RoleFS // Generic contract binding to access the raw methods on
}

// RoleFSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleFSCallerRaw struct {
	Contract *RoleFSCaller // Generic read-only contract binding to access the raw methods on
}

// RoleFSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleFSTransactorRaw struct {
	Contract *RoleFSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoleFS creates a new instance of RoleFS, bound to a specific deployed contract.
func NewRoleFS(address common.Address, backend bind.ContractBackend) (*RoleFS, error) {
	contract, err := bindRoleFS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoleFS{RoleFSCaller: RoleFSCaller{contract: contract}, RoleFSTransactor: RoleFSTransactor{contract: contract}, RoleFSFilterer: RoleFSFilterer{contract: contract}}, nil
}

// NewRoleFSCaller creates a new read-only instance of RoleFS, bound to a specific deployed contract.
func NewRoleFSCaller(address common.Address, caller bind.ContractCaller) (*RoleFSCaller, error) {
	contract, err := bindRoleFS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleFSCaller{contract: contract}, nil
}

// NewRoleFSTransactor creates a new write-only instance of RoleFS, bound to a specific deployed contract.
func NewRoleFSTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleFSTransactor, error) {
	contract, err := bindRoleFS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleFSTransactor{contract: contract}, nil
}

// NewRoleFSFilterer creates a new log filterer instance of RoleFS, bound to a specific deployed contract.
func NewRoleFSFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleFSFilterer, error) {
	contract, err := bindRoleFS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleFSFilterer{contract: contract}, nil
}

// bindRoleFS binds a generic wrapper to an already deployed contract.
func bindRoleFS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleFSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleFS *RoleFSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleFS.Contract.RoleFSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleFS *RoleFSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleFS.Contract.RoleFSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleFS *RoleFSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleFS.Contract.RoleFSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoleFS *RoleFSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoleFS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoleFS *RoleFSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoleFS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoleFS *RoleFSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoleFS.Contract.contract.Transact(opts, method, params...)
}

// AddOrder is a paid mutator transaction binding the contract method 0x991d1195.
//
// Solidity: function addOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 _end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes usign, bytes psign, bytes[] ksigns) returns()
func (_RoleFS *RoleFSTransactor) AddOrder(opts *bind.TransactOpts, uIndex uint64, pIndex uint64, _start uint64, _end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, usign []byte, psign []byte, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "addOrder", uIndex, pIndex, _start, _end, _size, nonce, tIndex, sPrice, usign, psign, ksigns)
}

// AddOrder is a paid mutator transaction binding the contract method 0x991d1195.
//
// Solidity: function addOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 _end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes usign, bytes psign, bytes[] ksigns) returns()
func (_RoleFS *RoleFSSession) AddOrder(uIndex uint64, pIndex uint64, _start uint64, _end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, usign []byte, psign []byte, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.AddOrder(&_RoleFS.TransactOpts, uIndex, pIndex, _start, _end, _size, nonce, tIndex, sPrice, usign, psign, ksigns)
}

// AddOrder is a paid mutator transaction binding the contract method 0x991d1195.
//
// Solidity: function addOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 _end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes usign, bytes psign, bytes[] ksigns) returns()
func (_RoleFS *RoleFSTransactorSession) AddOrder(uIndex uint64, pIndex uint64, _start uint64, _end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, usign []byte, psign []byte, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.AddOrder(&_RoleFS.TransactOpts, uIndex, pIndex, _start, _end, _size, nonce, tIndex, sPrice, usign, psign, ksigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xffd0fd09.
//
// Solidity: function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactor) AddRepair(opts *bind.TransactOpts, pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sprice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "addRepair", pIndex, nPIndex, _start, end, _size, nonce, tIndex, sprice, pSign, kSigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xffd0fd09.
//
// Solidity: function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSSession) AddRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sprice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.AddRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sprice, pSign, kSigns)
}

// AddRepair is a paid mutator transaction binding the contract method 0xffd0fd09.
//
// Solidity: function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactorSession) AddRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sprice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.AddRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sprice, pSign, kSigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xb28ff3c9.
//
// Solidity: function proWithdraw(uint64 pIndex, uint32 tIndex, uint256 pay, uint256 lost, bytes[] ksigns) returns()
func (_RoleFS *RoleFSTransactor) ProWithdraw(opts *bind.TransactOpts, pIndex uint64, tIndex uint32, pay *big.Int, lost *big.Int, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "proWithdraw", pIndex, tIndex, pay, lost, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xb28ff3c9.
//
// Solidity: function proWithdraw(uint64 pIndex, uint32 tIndex, uint256 pay, uint256 lost, bytes[] ksigns) returns()
func (_RoleFS *RoleFSSession) ProWithdraw(pIndex uint64, tIndex uint32, pay *big.Int, lost *big.Int, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.ProWithdraw(&_RoleFS.TransactOpts, pIndex, tIndex, pay, lost, ksigns)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xb28ff3c9.
//
// Solidity: function proWithdraw(uint64 pIndex, uint32 tIndex, uint256 pay, uint256 lost, bytes[] ksigns) returns()
func (_RoleFS *RoleFSTransactorSession) ProWithdraw(pIndex uint64, tIndex uint32, pay *big.Int, lost *big.Int, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.ProWithdraw(&_RoleFS.TransactOpts, pIndex, tIndex, pay, lost, ksigns)
}

// SetAddr is a paid mutator transaction binding the contract method 0xa7e7797d.
//
// Solidity: function setAddr(address iss, address r, address f, address rt) returns()
func (_RoleFS *RoleFSTransactor) SetAddr(opts *bind.TransactOpts, iss common.Address, r common.Address, f common.Address, rt common.Address) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "setAddr", iss, r, f, rt)
}

// SetAddr is a paid mutator transaction binding the contract method 0xa7e7797d.
//
// Solidity: function setAddr(address iss, address r, address f, address rt) returns()
func (_RoleFS *RoleFSSession) SetAddr(iss common.Address, r common.Address, f common.Address, rt common.Address) (*types.Transaction, error) {
	return _RoleFS.Contract.SetAddr(&_RoleFS.TransactOpts, iss, r, f, rt)
}

// SetAddr is a paid mutator transaction binding the contract method 0xa7e7797d.
//
// Solidity: function setAddr(address iss, address r, address f, address rt) returns()
func (_RoleFS *RoleFSTransactorSession) SetAddr(iss common.Address, r common.Address, f common.Address, rt common.Address) (*types.Transaction, error) {
	return _RoleFS.Contract.SetAddr(&_RoleFS.TransactOpts, iss, r, f, rt)
}

// SubOrder is a paid mutator transaction binding the contract method 0x08bf601e.
//
// Solidity: function subOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes usign, bytes psign, bytes[] ksigns) returns()
func (_RoleFS *RoleFSTransactor) SubOrder(opts *bind.TransactOpts, uIndex uint64, pIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, usign []byte, psign []byte, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "subOrder", uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice, usign, psign, ksigns)
}

// SubOrder is a paid mutator transaction binding the contract method 0x08bf601e.
//
// Solidity: function subOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes usign, bytes psign, bytes[] ksigns) returns()
func (_RoleFS *RoleFSSession) SubOrder(uIndex uint64, pIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, usign []byte, psign []byte, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.SubOrder(&_RoleFS.TransactOpts, uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice, usign, psign, ksigns)
}

// SubOrder is a paid mutator transaction binding the contract method 0x08bf601e.
//
// Solidity: function subOrder(uint64 uIndex, uint64 pIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes usign, bytes psign, bytes[] ksigns) returns()
func (_RoleFS *RoleFSTransactorSession) SubOrder(uIndex uint64, pIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, usign []byte, psign []byte, ksigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.SubOrder(&_RoleFS.TransactOpts, uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice, usign, psign, ksigns)
}

// SubRepair is a paid mutator transaction binding the contract method 0xeb975779.
//
// Solidity: function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactor) SubRepair(opts *bind.TransactOpts, pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.contract.Transact(opts, "subRepair", pIndex, nPIndex, _start, end, _size, nonce, tIndex, sPrice, pSign, kSigns)
}

// SubRepair is a paid mutator transaction binding the contract method 0xeb975779.
//
// Solidity: function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSSession) SubRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.SubRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sPrice, pSign, kSigns)
}

// SubRepair is a paid mutator transaction binding the contract method 0xeb975779.
//
// Solidity: function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes pSign, bytes[] kSigns) returns()
func (_RoleFS *RoleFSTransactorSession) SubRepair(pIndex uint64, nPIndex uint64, _start uint64, end uint64, _size uint64, nonce uint64, tIndex uint32, sPrice *big.Int, pSign []byte, kSigns [][]byte) (*types.Transaction, error) {
	return _RoleFS.Contract.SubRepair(&_RoleFS.TransactOpts, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sPrice, pSign, kSigns)
}
