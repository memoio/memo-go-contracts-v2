// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package role

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

// AOParams is an auto generated low-level Go binding around an user-defined struct.
type AOParams struct {
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint32
	SPrice *big.Int
	Usign  []byte
	Psign  []byte
	Ksigns [][]byte
}

// IssuParams is an auto generated low-level Go binding around an user-defined struct.
type IssuParams struct {
	Start  uint64
	End    uint64
	Size   uint64
	SPrice *big.Int
}

// PWParams is an auto generated low-level Go binding around an user-defined struct.
type PWParams struct {
	PIndex uint64
	TIndex uint32
	PAddr  common.Address
	TAddr  common.Address
	Pay    *big.Int
	Lost   *big.Int
	Ksigns [][]byte
}

// SGParams is an auto generated low-level Go binding around an user-defined struct.
type SGParams struct {
	Index  uint64
	IsAdd  bool
	Size   *big.Int
	SPrice *big.Int
}

// SOParams is an auto generated low-level Go binding around an user-defined struct.
type SOParams struct {
	KIndex uint64
	UIndex uint64
	PIndex uint64
	Start  uint64
	End    uint64
	Size   uint64
	Nonce  uint64
	TIndex uint32
	SPrice *big.Int
	Usign  []byte
	Psign  []byte
	Ksigns [][]byte
}

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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structAOParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newPro\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"addRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"}],\"name\":\"createFs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"token\",\"type\":\"uint32\"}],\"name\":\"getChannelInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"}],\"name\":\"getFsInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"}],\"name\":\"getFsInfoAggOrder\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"}],\"name\":\"getFsPSum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getFsPro\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getSettleInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"user\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provider\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"token\",\"type\":\"uint32\"}],\"name\":\"getStoreInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"pAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lost\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structPWParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"usign\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"psign\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"ksigns\",\"type\":\"bytes[]\"}],\"internalType\":\"structSOParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"subOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"kIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"newPro\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sprice\",\"type\":\"uint256\"}],\"name\":\"subRepair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"roleType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"50cbb46f": "addKeeper(uint64)",
		"2db099b9": "addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]))",
		"0f60c7b3": "addRepair(uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256)",
		"e49c0724": "createFs(uint64)",
		"761966b7": "getBalance(uint64,uint32)",
		"cc122253": "getChannelInfo(uint64,uint64,uint32)",
		"324ab551": "getFsInfo(uint64)",
		"03eb18aa": "getFsInfoAggOrder(uint64,uint64)",
		"650c6a89": "getFsPSum(uint64)",
		"25d931e8": "getFsPro(uint64,uint64)",
		"7b31a24d": "getSettleInfo(uint64,uint32)",
		"3f5f363e": "getStoreInfo(uint64,uint64,uint32)",
		"dc7d3249": "proWithdraw((uint64,uint32,address,address,uint256,uint256,bytes[]))",
		"e04f98ed": "recharge(uint64,uint32,address,address,uint256)",
		"b03c9e45": "subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]))",
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
// Solidity: function getFsInfo(uint64 user) view returns(bool)
func (_IFileSys *IFileSysCaller) GetFsInfo(opts *bind.CallOpts, user uint64) (bool, error) {
	var out []interface{}
	err := _IFileSys.contract.Call(opts, &out, "getFsInfo", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool)
func (_IFileSys *IFileSysSession) GetFsInfo(user uint64) (bool, error) {
	return _IFileSys.Contract.GetFsInfo(&_IFileSys.CallOpts, user)
}

// GetFsInfo is a free data retrieval call binding the contract method 0x324ab551.
//
// Solidity: function getFsInfo(uint64 user) view returns(bool)
func (_IFileSys *IFileSysCallerSession) GetFsInfo(user uint64) (bool, error) {
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

// AddOrder is a paid mutator transaction binding the contract method 0x2db099b9.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactor) AddOrder(opts *bind.TransactOpts, ps AOParams) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "addOrder", ps)
}

// AddOrder is a paid mutator transaction binding the contract method 0x2db099b9.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]) ps) returns()
func (_IFileSys *IFileSysSession) AddOrder(ps AOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, ps)
}

// AddOrder is a paid mutator transaction binding the contract method 0x2db099b9.
//
// Solidity: function addOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactorSession) AddOrder(ps AOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.AddOrder(&_IFileSys.TransactOpts, ps)
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

// CreateFs is a paid mutator transaction binding the contract method 0xe49c0724.
//
// Solidity: function createFs(uint64 uIndex) returns()
func (_IFileSys *IFileSysTransactor) CreateFs(opts *bind.TransactOpts, uIndex uint64) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "createFs", uIndex)
}

// CreateFs is a paid mutator transaction binding the contract method 0xe49c0724.
//
// Solidity: function createFs(uint64 uIndex) returns()
func (_IFileSys *IFileSysSession) CreateFs(uIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.CreateFs(&_IFileSys.TransactOpts, uIndex)
}

// CreateFs is a paid mutator transaction binding the contract method 0xe49c0724.
//
// Solidity: function createFs(uint64 uIndex) returns()
func (_IFileSys *IFileSysTransactorSession) CreateFs(uIndex uint64) (*types.Transaction, error) {
	return _IFileSys.Contract.CreateFs(&_IFileSys.TransactOpts, uIndex)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xdc7d3249.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactor) ProWithdraw(opts *bind.TransactOpts, ps PWParams) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "proWithdraw", ps)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xdc7d3249.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,bytes[]) ps) returns()
func (_IFileSys *IFileSysSession) ProWithdraw(ps PWParams) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, ps)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0xdc7d3249.
//
// Solidity: function proWithdraw((uint64,uint32,address,address,uint256,uint256,bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactorSession) ProWithdraw(ps PWParams) (*types.Transaction, error) {
	return _IFileSys.Contract.ProWithdraw(&_IFileSys.TransactOpts, ps)
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

// SubOrder is a paid mutator transaction binding the contract method 0xb03c9e45.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactor) SubOrder(opts *bind.TransactOpts, ps SOParams) (*types.Transaction, error) {
	return _IFileSys.contract.Transact(opts, "subOrder", ps)
}

// SubOrder is a paid mutator transaction binding the contract method 0xb03c9e45.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]) ps) returns()
func (_IFileSys *IFileSysSession) SubOrder(ps SOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, ps)
}

// SubOrder is a paid mutator transaction binding the contract method 0xb03c9e45.
//
// Solidity: function subOrder((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint32,uint256,bytes,bytes,bytes[]) ps) returns()
func (_IFileSys *IFileSysTransactorSession) SubOrder(ps SOParams) (*types.Transaction, error) {
	return _IFileSys.Contract.SubOrder(&_IFileSys.TransactOpts, ps)
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
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIssuParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c862ea50": "issu((uint64,uint64,uint64,uint256))",
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

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_IIssuance *IIssuanceTransactor) Issu(opts *bind.TransactOpts, ps IssuParams) (*types.Transaction, error) {
	return _IIssuance.contract.Transact(opts, "issu", ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_IIssuance *IIssuanceSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_IIssuance *IIssuanceTransactorSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _IIssuance.Contract.Issu(&_IIssuance.TransactOpts, ps)
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

// IPledgePoolMetaData contains all meta data concerning the IPledgePool contract.
var IPledgePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8859344c": "addToken(address,uint32)",
		"761966b7": "getBalance(uint64,uint32)",
		"c3b19ccc": "getPledge(uint32)",
		"364e4bf5": "pledge(uint64,uint256,bytes)",
		"32704298": "withdraw(uint64,uint32,uint256,bytes)",
	},
}

// IPledgePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use IPledgePoolMetaData.ABI instead.
var IPledgePoolABI = IPledgePoolMetaData.ABI

// Deprecated: Use IPledgePoolMetaData.Sigs instead.
// IPledgePoolFuncSigs maps the 4-byte function signature to its string representation.
var IPledgePoolFuncSigs = IPledgePoolMetaData.Sigs

// IPledgePool is an auto generated Go binding around an Ethereum contract.
type IPledgePool struct {
	IPledgePoolCaller     // Read-only binding to the contract
	IPledgePoolTransactor // Write-only binding to the contract
	IPledgePoolFilterer   // Log filterer for contract events
}

// IPledgePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPledgePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPledgePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPledgePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPledgePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPledgePoolSession struct {
	Contract     *IPledgePool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPledgePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPledgePoolCallerSession struct {
	Contract *IPledgePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IPledgePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPledgePoolTransactorSession struct {
	Contract     *IPledgePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IPledgePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPledgePoolRaw struct {
	Contract *IPledgePool // Generic contract binding to access the raw methods on
}

// IPledgePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPledgePoolCallerRaw struct {
	Contract *IPledgePoolCaller // Generic read-only contract binding to access the raw methods on
}

// IPledgePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPledgePoolTransactorRaw struct {
	Contract *IPledgePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPledgePool creates a new instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePool(address common.Address, backend bind.ContractBackend) (*IPledgePool, error) {
	contract, err := bindIPledgePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPledgePool{IPledgePoolCaller: IPledgePoolCaller{contract: contract}, IPledgePoolTransactor: IPledgePoolTransactor{contract: contract}, IPledgePoolFilterer: IPledgePoolFilterer{contract: contract}}, nil
}

// NewIPledgePoolCaller creates a new read-only instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePoolCaller(address common.Address, caller bind.ContractCaller) (*IPledgePoolCaller, error) {
	contract, err := bindIPledgePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgePoolCaller{contract: contract}, nil
}

// NewIPledgePoolTransactor creates a new write-only instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IPledgePoolTransactor, error) {
	contract, err := bindIPledgePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPledgePoolTransactor{contract: contract}, nil
}

// NewIPledgePoolFilterer creates a new log filterer instance of IPledgePool, bound to a specific deployed contract.
func NewIPledgePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IPledgePoolFilterer, error) {
	contract, err := bindIPledgePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPledgePoolFilterer{contract: contract}, nil
}

// bindIPledgePool binds a generic wrapper to an already deployed contract.
func bindIPledgePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPledgePoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgePool *IPledgePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgePool.Contract.IPledgePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgePool *IPledgePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgePool.Contract.IPledgePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgePool *IPledgePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgePool.Contract.IPledgePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPledgePool *IPledgePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPledgePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPledgePool *IPledgePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPledgePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPledgePool *IPledgePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPledgePool.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCaller) GetBalance(opts *bind.CallOpts, index uint64, tIndex uint32) (*big.Int, error) {
	var out []interface{}
	err := _IPledgePool.contract.Call(opts, &out, "getBalance", index, tIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolSession) GetBalance(index uint64, tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetBalance(&_IPledgePool.CallOpts, index, tIndex)
}

// GetBalance is a free data retrieval call binding the contract method 0x761966b7.
//
// Solidity: function getBalance(uint64 index, uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCallerSession) GetBalance(index uint64, tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetBalance(&_IPledgePool.CallOpts, index, tIndex)
}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCaller) GetPledge(opts *bind.CallOpts, tIndex uint32) (*big.Int, error) {
	var out []interface{}
	err := _IPledgePool.contract.Call(opts, &out, "getPledge", tIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolSession) GetPledge(tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetPledge(&_IPledgePool.CallOpts, tIndex)
}

// GetPledge is a free data retrieval call binding the contract method 0xc3b19ccc.
//
// Solidity: function getPledge(uint32 tIndex) view returns(uint256)
func (_IPledgePool *IPledgePoolCallerSession) GetPledge(tIndex uint32) (*big.Int, error) {
	return _IPledgePool.Contract.GetPledge(&_IPledgePool.CallOpts, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_IPledgePool *IPledgePoolTransactor) AddToken(opts *bind.TransactOpts, tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _IPledgePool.contract.Transact(opts, "addToken", tAddr, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_IPledgePool *IPledgePoolSession) AddToken(tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _IPledgePool.Contract.AddToken(&_IPledgePool.TransactOpts, tAddr, tIndex)
}

// AddToken is a paid mutator transaction binding the contract method 0x8859344c.
//
// Solidity: function addToken(address tAddr, uint32 tIndex) returns()
func (_IPledgePool *IPledgePoolTransactorSession) AddToken(tAddr common.Address, tIndex uint32) (*types.Transaction, error) {
	return _IPledgePool.Contract.AddToken(&_IPledgePool.TransactOpts, tAddr, tIndex)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_IPledgePool *IPledgePoolTransactor) Pledge(opts *bind.TransactOpts, index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.contract.Transact(opts, "pledge", index, money, sign)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_IPledgePool *IPledgePoolSession) Pledge(index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Pledge(&_IPledgePool.TransactOpts, index, money, sign)
}

// Pledge is a paid mutator transaction binding the contract method 0x364e4bf5.
//
// Solidity: function pledge(uint64 index, uint256 money, bytes sign) payable returns()
func (_IPledgePool *IPledgePoolTransactorSession) Pledge(index uint64, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Pledge(&_IPledgePool.TransactOpts, index, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_IPledgePool *IPledgePoolTransactor) Withdraw(opts *bind.TransactOpts, index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.contract.Transact(opts, "withdraw", index, tIndex, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_IPledgePool *IPledgePoolSession) Withdraw(index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Withdraw(&_IPledgePool.TransactOpts, index, tIndex, money, sign)
}

// Withdraw is a paid mutator transaction binding the contract method 0x32704298.
//
// Solidity: function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes sign) returns()
func (_IPledgePool *IPledgePoolTransactorSession) Withdraw(index uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _IPledgePool.Contract.Withdraw(&_IPledgePool.TransactOpts, index, tIndex, money, sign)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"}],\"name\":\"addKeeperToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"addProviderToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"indexes\",\"type\":\"uint64[]\"},{\"internalType\":\"uint16\",\"name\":\"_level\",\"type\":\"uint16\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddrGindex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getFsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGKNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ig\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ik\",\"type\":\"uint64\"}],\"name\":\"getGroupK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getRoleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taddr\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structSGParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"setGInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pPledge\",\"type\":\"uint256\"}],\"name\":\"setPledgeMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"withdrawFromFs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
		"64fe6290": "getGroupK(uint64,uint64)",
		"07483499": "getRoleInfo(address)",
		"a6ed590b": "pledgeK()",
		"8ba61d28": "pledgeP()",
		"de909560": "pledgePool()",
		"517985b0": "recharge(uint64,uint32,uint256,bytes)",
		"24b8fbf6": "register(address,bytes)",
		"10e35bbe": "registerKeeper(uint64,bytes,bytes)",
		"d57e8a4e": "registerProvider(uint64,bytes)",
		"09824a80": "registerToken(address)",
		"488cee1c": "registerUser(uint64,uint64,bytes,bytes)",
		"121ed07f": "setGInfo((uint64,bool,uint256,uint256))",
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

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_IRole *IRoleCaller) GetGroupK(opts *bind.CallOpts, ig uint64, ik uint64) (uint64, error) {
	var out []interface{}
	err := _IRole.contract.Call(opts, &out, "getGroupK", ig, ik)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_IRole *IRoleSession) GetGroupK(ig uint64, ik uint64) (uint64, error) {
	return _IRole.Contract.GetGroupK(&_IRole.CallOpts, ig, ik)
}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_IRole *IRoleCallerSession) GetGroupK(ig uint64, ik uint64) (uint64, error) {
	return _IRole.Contract.GetGroupK(&_IRole.CallOpts, ig, ik)
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

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactor) RegisterUser(opts *bind.TransactOpts, index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "registerUser", index, _gIndex, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleSession) RegisterUser(index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterUser(&_IRole.TransactOpts, index, _gIndex, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_IRole *IRoleTransactorSession) RegisterUser(index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _IRole.Contract.RegisterUser(&_IRole.TransactOpts, index, _gIndex, blsKey, sign)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_IRole *IRoleTransactor) SetGInfo(opts *bind.TransactOpts, ps SGParams) (*types.Transaction, error) {
	return _IRole.contract.Transact(opts, "setGInfo", ps)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_IRole *IRoleSession) SetGInfo(ps SGParams) (*types.Transaction, error) {
	return _IRole.Contract.SetGInfo(&_IRole.TransactOpts, ps)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_IRole *IRoleTransactorSession) SetGInfo(ps SGParams) (*types.Transaction, error) {
	return _IRole.Contract.SetGInfo(&_IRole.TransactOpts, ps)
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

// OwnerMetaData contains all meta data concerning the Owner contract.
var OwnerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AlterOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"alterOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0ca05f9f": "alterOwner(address)",
		"893d20e8": "getOwner()",
	},
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163390811782556040805192835260208301919091527f8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90910160405180910390a16101808061006e6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630ca05f9f1461003b578063893d20e814610050575b600080fd5b61004e61004936600461011a565b61006f565b005b600054604080516001600160a01b039092168252519081900360200190f35b6000546001600160a01b031633146100b15760405162461bcd60e51b81526020600482015260016024820152602760f91b604482015260640160405180910390fd5b600054604080516001600160a01b03928316815291831660208301527f8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90910160405180910390a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006020828403121561012c57600080fd5b81356001600160a01b038116811461014357600080fd5b939250505056fea26469706673582212200c923b2149a502b9e95044154fa00b0a526e928fc9be89f1d6f937c357f50d7964736f6c634300080b0033",
}

// OwnerABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnerMetaData.ABI instead.
var OwnerABI = OwnerMetaData.ABI

// Deprecated: Use OwnerMetaData.Sigs instead.
// OwnerFuncSigs maps the 4-byte function signature to its string representation.
var OwnerFuncSigs = OwnerMetaData.Sigs

// OwnerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnerMetaData.Bin instead.
var OwnerBin = OwnerMetaData.Bin

// DeployOwner deploys a new Ethereum contract, binding an instance of Owner to it.
func DeployOwner(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owner, error) {
	parsed, err := OwnerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// Owner is an auto generated Go binding around an Ethereum contract.
type Owner struct {
	OwnerCaller     // Read-only binding to the contract
	OwnerTransactor // Write-only binding to the contract
	OwnerFilterer   // Log filterer for contract events
}

// OwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerSession struct {
	Contract     *Owner            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerCallerSession struct {
	Contract *OwnerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerTransactorSession struct {
	Contract     *OwnerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerRaw struct {
	Contract *Owner // Generic contract binding to access the raw methods on
}

// OwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerCallerRaw struct {
	Contract *OwnerCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerTransactorRaw struct {
	Contract *OwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwner creates a new instance of Owner, bound to a specific deployed contract.
func NewOwner(address common.Address, backend bind.ContractBackend) (*Owner, error) {
	contract, err := bindOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// NewOwnerCaller creates a new read-only instance of Owner, bound to a specific deployed contract.
func NewOwnerCaller(address common.Address, caller bind.ContractCaller) (*OwnerCaller, error) {
	contract, err := bindOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerCaller{contract: contract}, nil
}

// NewOwnerTransactor creates a new write-only instance of Owner, bound to a specific deployed contract.
func NewOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerTransactor, error) {
	contract, err := bindOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerTransactor{contract: contract}, nil
}

// NewOwnerFilterer creates a new log filterer instance of Owner, bound to a specific deployed contract.
func NewOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerFilterer, error) {
	contract, err := bindOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerFilterer{contract: contract}, nil
}

// bindOwner binds a generic wrapper to an already deployed contract.
func bindOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.OwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Owner *OwnerCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Owner *OwnerSession) GetOwner() (common.Address, error) {
	return _Owner.Contract.GetOwner(&_Owner.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Owner *OwnerCallerSession) GetOwner() (common.Address, error) {
	return _Owner.Contract.GetOwner(&_Owner.CallOpts)
}

// AlterOwner is a paid mutator transaction binding the contract method 0x0ca05f9f.
//
// Solidity: function alterOwner(address newOwner) returns()
func (_Owner *OwnerTransactor) AlterOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owner.contract.Transact(opts, "alterOwner", newOwner)
}

// AlterOwner is a paid mutator transaction binding the contract method 0x0ca05f9f.
//
// Solidity: function alterOwner(address newOwner) returns()
func (_Owner *OwnerSession) AlterOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owner.Contract.AlterOwner(&_Owner.TransactOpts, newOwner)
}

// AlterOwner is a paid mutator transaction binding the contract method 0x0ca05f9f.
//
// Solidity: function alterOwner(address newOwner) returns()
func (_Owner *OwnerTransactorSession) AlterOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owner.Contract.AlterOwner(&_Owner.TransactOpts, newOwner)
}

// OwnerAlterOwnerIterator is returned from FilterAlterOwner and is used to iterate over the raw logs and unpacked data for AlterOwner events raised by the Owner contract.
type OwnerAlterOwnerIterator struct {
	Event *OwnerAlterOwner // Event containing the contract specifics and raw log

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
func (it *OwnerAlterOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAlterOwner)
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
		it.Event = new(OwnerAlterOwner)
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
func (it *OwnerAlterOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAlterOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAlterOwner represents a AlterOwner event raised by the Owner contract.
type OwnerAlterOwner struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlterOwner is a free log retrieval operation binding the contract event 0x8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90.
//
// Solidity: event AlterOwner(address from, address to)
func (_Owner *OwnerFilterer) FilterAlterOwner(opts *bind.FilterOpts) (*OwnerAlterOwnerIterator, error) {

	logs, sub, err := _Owner.contract.FilterLogs(opts, "AlterOwner")
	if err != nil {
		return nil, err
	}
	return &OwnerAlterOwnerIterator{contract: _Owner.contract, event: "AlterOwner", logs: logs, sub: sub}, nil
}

// WatchAlterOwner is a free log subscription operation binding the contract event 0x8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90.
//
// Solidity: event AlterOwner(address from, address to)
func (_Owner *OwnerFilterer) WatchAlterOwner(opts *bind.WatchOpts, sink chan<- *OwnerAlterOwner) (event.Subscription, error) {

	logs, sub, err := _Owner.contract.WatchLogs(opts, "AlterOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAlterOwner)
				if err := _Owner.contract.UnpackLog(event, "AlterOwner", log); err != nil {
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

// ParseAlterOwner is a log parse operation binding the contract event 0x8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90.
//
// Solidity: event AlterOwner(address from, address to)
func (_Owner *OwnerFilterer) ParseAlterOwner(log types.Log) (*OwnerAlterOwner, error) {
	event := new(OwnerAlterOwner)
	if err := _Owner.contract.UnpackLog(event, "AlterOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Bin: "0x608060405234801561001057600080fd5b50600280546001600160a01b03191633179055610518806100326000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632df2685f1461005c5780633c7bdc19146100905780638ff8e15a146100b3578063b04c6c68146100db578063f7b7d6b6146100e3575b600080fd5b61006f61006a36600461042b565b61010e565b6040805163ffffffff90931683529015156020830152015b60405180910390f35b6100a361009e36600461045b565b6101e2565b6040519015158152602001610087565b6100c66100c136600461042b565b610259565b60405163ffffffff9091168152602001610087565b6000546100c6565b6100f66100f136600461045b565b61037c565b6040516001600160a01b039091168152602001610087565b6001600160a01b038116600090815260016020526040808220549051633c7bdc1960e01b815263ffffffff9091166004820181905282913090633c7bdc1990602401602060405180830381865afa15801561016d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101919190610481565b156101d85760008163ffffffff16815481106101af576101af6104a3565b6000918252602090912001546001600160a01b03858116911614156101d8579360019350915050565b9360009350915050565b6000805463ffffffff8316108015610244575060016000808463ffffffff1681548110610211576102116104a3565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff64010000000090910416155b1561025157506001919050565b506000919050565b6002546000906001600160a01b0316331461029f5760405162461bcd60e51b81526020600482015260026024820152614e4f60f01b604482015260640160405180910390fd5b6102a8826103c1565b156102cf57506001600160a01b031660009081526001602052604090205463ffffffff1690565b60008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563820180546001600160a01b0319166001600160a01b0387169081179091558084526020918252604093849020805463ffffffff191663ffffffff851690811790915584519182529181019190915290917f7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6910160405180910390a192915050565b6000805463ffffffff831610156102515760008263ffffffff16815481106103a6576103a66104a3565b6000918252602090912001546001600160a01b031692915050565b6000805b60005481101561042257826001600160a01b0316600082815481106103ec576103ec6104a3565b6000918252602090912001546001600160a01b031614156104105750600192915050565b8061041a816104b9565b9150506103c5565b50600092915050565b60006020828403121561043d57600080fd5b81356001600160a01b038116811461045457600080fd5b9392505050565b60006020828403121561046d57600080fd5b813563ffffffff8116811461045457600080fd5b60006020828403121561049357600080fd5b8151801515811461045457600080fd5b634e487b7160e01b600052603260045260246000fd5b60006000198214156104db57634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122080162e603925909c14e71489135a62d543ddc09b2963d9ccc556d03ae2890a0064736f6c634300080b0033",
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

// RecoverMetaData contains all meta data concerning the Recover contract.
var RecoverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"19045a25": "recover(bytes32,bytes)",
	},
	Bin: "0x6102a861003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c806319045a251461003a575b600080fd5b61004d610048366004610184565b610069565b6040516001600160a01b03909116815260200160405180910390f35b6000815160411461007c57506000610168565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156100c25760009350505050610168565b601b8160ff1610156100dc576100d981601b61023f565b90505b8060ff16601b141580156100f457508060ff16601c14155b156101055760009350505050610168565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610158573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561019757600080fd5b82359150602083013567ffffffffffffffff808211156101b657600080fd5b818501915085601f8301126101ca57600080fd5b8135818111156101dc576101dc61016e565b604051601f8201601f19908116603f011681019083821181831017156102045761020461016e565b8160405282815288602084870101111561021d57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600060ff821660ff84168060ff0382111561026a57634e487b7160e01b600052601160045260246000fd5b01939250505056fea2646970667358221220320f9df43c85c3b600193643aab8ef102a35d826bcca6389c8056f04e847353664736f6c634300080b0033",
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

// RoleMetaData contains all meta data concerning the Role contract.
var RoleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"f\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ppro\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AlterOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"}],\"name\":\"CreateGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RKeeper\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"RUser\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"}],\"name\":\"addKeeperToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"addProviderToGroup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"alterOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"}],\"name\":\"checkIR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"}],\"name\":\"checkT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"indexes\",\"type\":\"uint64[]\"},{\"internalType\":\"uint16\",\"name\":\"_level\",\"type\":\"uint16\"}],\"name\":\"createGroup\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"foundation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getAddrGindex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddrsNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getFsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGKNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ig\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"iu\",\"type\":\"uint64\"}],\"name\":\"getGU\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGUPNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGroupInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ig\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ik\",\"type\":\"uint64\"}],\"name\":\"getGroupK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"ig\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"ip\",\"type\":\"uint64\"}],\"name\":\"getGroupP\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGroupsNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getRoleIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"getRoleInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgeP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pledgePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"uIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_index\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rolefs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_fsAddr\",\"type\":\"address\"}],\"name\":\"setGF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isAdd\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structSGParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"setGInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_p\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"i\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rfs\",\"type\":\"address\"}],\"name\":\"setPI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pPledge\",\"type\":\"uint256\"}],\"name\":\"setPledgeMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"tIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sign\",\"type\":\"bytes\"}],\"name\":\"withdrawFromFs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"7977031c": "addKeeperToGroup(uint64,uint64)",
		"8ae21af3": "addProviderToGroup(uint64,uint64,bytes)",
		"0ca05f9f": "alterOwner(address)",
		"cf8e99a8": "checkIR(uint64,uint8)",
		"de92e994": "checkT(uint32)",
		"f652391b": "createGroup(uint64[],uint16)",
		"41fbb050": "foundation()",
		"9332aa6e": "getAddr(uint64)",
		"421795e5": "getAddrGindex(uint64)",
		"7ba783be": "getAddrsNum()",
		"5f096376": "getFsAddr(uint64)",
		"429fb683": "getGKNum(uint64)",
		"54b365a3": "getGU(uint64,uint64)",
		"0d0b065c": "getGUPNum(uint64)",
		"4496c991": "getGroupInfo(uint64)",
		"64fe6290": "getGroupK(uint64,uint64)",
		"3d180dd3": "getGroupP(uint64,uint64)",
		"5d39c56b": "getGroupsNum()",
		"893d20e8": "getOwner()",
		"ae0e4ffa": "getRoleIndex(address)",
		"07483499": "getRoleInfo(address)",
		"863623bb": "issuance()",
		"a6ed590b": "pledgeK()",
		"8ba61d28": "pledgeP()",
		"de909560": "pledgePool()",
		"40c65f72": "rToken()",
		"517985b0": "recharge(uint64,uint32,uint256,bytes)",
		"24b8fbf6": "register(address,bytes)",
		"10e35bbe": "registerKeeper(uint64,bytes,bytes)",
		"d57e8a4e": "registerProvider(uint64,bytes)",
		"09824a80": "registerToken(address)",
		"488cee1c": "registerUser(uint64,uint64,bytes,bytes)",
		"2d6d777e": "rolefs()",
		"a6773009": "setGF(uint64,address)",
		"121ed07f": "setGInfo((uint64,bool,uint256,uint256))",
		"eba091a6": "setPI(address,address,address)",
		"97948fda": "setPledgeMoney(uint256,uint256)",
		"d30d0ce5": "withdrawFromFs(uint64,uint32,uint256,bytes)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162003a8538038062003a858339810160408190526200003491620001de565b600080546001600160a01b0319163390811782556040805192835260208301919091527f8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90910160405180910390a1600280546001600160a01b0319166001600160a01b038616179055600582905560068190556040516001600160601b0319606086811b8216602084015285901b16603482015260009060480160405160208183030381529060405280519060200120905060008060405180602001620000fb90620001b3565b6020820181038252601f19601f820116604052509050828151602083016000f5600880546001600160a01b0319166001600160a01b038381169182179092556040516347fc70ad60e11b81529189166004830152919350839190638ff8e15a906024016020604051808303816000875af11580156200017e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001a4919062000226565b50505050505050505062000255565b61054a806200353b83390190565b80516001600160a01b0381168114620001d957600080fd5b919050565b60008060008060808587031215620001f557600080fd5b6200020085620001c1565b93506200021060208601620001c1565b6040860151606090960151949790965092505050565b6000602082840312156200023957600080fd5b815163ffffffff811681146200024e57600080fd5b9392505050565b6132d680620002656000396000f3fe6080604052600436106102295760003560e01c806364fe629011610123578063a6773009116100ab578063d57e8a4e1161006f578063d57e8a4e1461072d578063de9095601461074d578063de92e9941461076d578063eba091a61461078d578063f652391b146107ad57600080fd5b8063a677300914610671578063a6ed590b14610691578063ae0e4ffa146106a7578063cf8e99a8146106ed578063d30d0ce51461070d57600080fd5b8063893d20e8116100f2578063893d20e8146105cf5780638ae21af3146105ed5780638ba61d281461060d5780639332aa6e1461063157806397948fda1461065157600080fd5b806364fe62901461055a5780637977031c1461057a5780637ba783be1461059a578063863623bb146105af57600080fd5b806340c65f72116101b1578063488cee1c11610175578063488cee1c146104d2578063517985b0146104f257806354b365a3146105055780635d39c56b146105255780635f0963761461053a57600080fd5b806340c65f72146103c257806341fbb050146103e2578063421795e514610402578063429fb683146104495780634496c9911461046957600080fd5b806310e35bbe116101f857806310e35bbe146102f2578063121ed07f1461031257806324b8fbf6146103325780632d6d777e1461036a5780633d180dd3146103a257600080fd5b8063074834991461023557806309824a80146102705780630ca05f9f146102925780630d0b065c146102b257600080fd5b3661023057005b600080fd5b34801561024157600080fd5b50610255610250366004612a30565b6107cd565b60405161026796959493929190612aa9565b60405180910390f35b34801561027c57600080fd5b5061029061028b366004612a30565b6108cc565b005b34801561029e57600080fd5b506102906102ad366004612a30565b6109e6565b3480156102be57600080fd5b506102d26102cd366004612b16565b610a79565b604080516001600160401b03938416815292909116602083015201610267565b3480156102fe57600080fd5b5061029061030d366004612be6565b610ae7565b34801561031e57600080fd5b5061029061032d366004612c67565b610c90565b34801561033e57600080fd5b5061035261034d366004612cd9565b610df6565b6040516001600160401b039091168152602001610267565b34801561037657600080fd5b50600a5461038a906001600160a01b031681565b6040516001600160a01b039091168152602001610267565b3480156103ae57600080fd5b506103526103bd366004612d28565b610f44565b3480156103ce57600080fd5b5060085461038a906001600160a01b031681565b3480156103ee57600080fd5b5060025461038a906001600160a01b031681565b34801561040e57600080fd5b5061042261041d366004612b16565b610fbe565b604080516001600160a01b0390931683526001600160401b03909116602083015201610267565b34801561045557600080fd5b50610352610464366004612b16565b611020565b34801561047557600080fd5b50610489610484366004612b16565b611058565b60408051971515885295151560208801529315159486019490945261ffff9091166060850152608084015260a08301919091526001600160a01b031660c082015260e001610267565b3480156104de57600080fd5b506102906104ed366004612d5b565b6111f5565b610290610500366004612df1565b611436565b34801561051157600080fd5b50610352610520366004612d28565b611645565b34801561053157600080fd5b50600754610352565b34801561054657600080fd5b5061038a610555366004612b16565b61168e565b34801561056657600080fd5b50610352610575366004612d28565b6116cf565b34801561058657600080fd5b50610290610595366004612d28565b611718565b3480156105a657600080fd5b50600354610352565b3480156105bb57600080fd5b5060095461038a906001600160a01b031681565b3480156105db57600080fd5b506000546001600160a01b031661038a565b3480156105f957600080fd5b50610290610608366004612e4e565b6119e2565b34801561061957600080fd5b5061062360065481565b604051908152602001610267565b34801561063d57600080fd5b5061038a61064c366004612b16565b611b23565b34801561065d57600080fd5b5061029061066c366004612ea1565b611b5c565b34801561067d57600080fd5b5061029061068c366004612ec3565b611b91565b34801561069d57600080fd5b5061062360055481565b3480156106b357600080fd5b506103526106c2366004612a30565b6001600160a01b0316600090815260046020526040902054630100000090046001600160401b031690565b3480156106f957600080fd5b5061038a610708366004612efa565b611c1b565b34801561071957600080fd5b50610290610728366004612df1565b611cb3565b34801561073957600080fd5b50610290610748366004612f2c565b611f43565b34801561075957600080fd5b5060015461038a906001600160a01b031681565b34801561077957600080fd5b5061038a610788366004612f48565b6120e5565b34801561079957600080fd5b506102906107a8366004612f65565b6121f9565b3480156107b957600080fd5b506103526107c8366004612fc2565b612262565b6001600160a01b0381166000908152600460205260408120805460019091018054839283928392839260609260ff808316936101008404821693620100008104909216926001600160401b0363010000008404811693600160581b90041691819061083790613080565b80601f016020809104026020016040519081016040528092919081815260200182805461086390613080565b80156108b05780601f10610885576101008083540402835291602001916108b0565b820191906000526020600020905b81548152906001019060200180831161089357829003601f168201915b5050505050905095509550955095509550955091939550919395565b6000546001600160a01b031633146108ff5760405162461bcd60e51b81526004016108f6906130bb565b60405180910390fd5b6008546040516347fc70ad60e11b81526001600160a01b038381166004830152909116906000908290638ff8e15a906024016020604051808303816000875af1158015610950573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097491906130d6565b6001546040516322164d1360e21b81526001600160a01b03868116600483015263ffffffff84166024830152929350911690638859344c90604401600060405180830381600087803b1580156109c957600080fd5b505af11580156109dd573d6000803e3d6000fd5b50505050505050565b6000546001600160a01b03163314610a105760405162461bcd60e51b81526004016108f6906130bb565b600054604080516001600160a01b03928316815291831660208301527f8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90910160405180910390a1600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806007836001600160401b031681548110610a9857610a986130f3565b9060005260206000209060070201600301805490506007846001600160401b031681548110610ac957610ac96130f3565b90600052602060002090600702016002018054905091509150915091565b336000610af3856125f3565b9050806001600160a01b0316826001600160a01b031614610b485760008285604051602001610b23929190613109565b604051602081830303815290604052805190602001209050610b46828286612680565b505b60015460405163761966b760e01b81526001600160401b0387166004820152600060248201819052916001600160a01b03169063761966b790604401602060405180830381865afa158015610ba1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc59190613151565b9050600554811015610bfe5760405162461bcd60e51b81526020600482015260026024820152614e4560f01b60448201526064016108f6565b6001600160a01b0382166000908152600460209081526040909120805462ff00001916620300001781558651610c3c926001909201918801906128d3565b50604080516001600160401b03881681526001600160a01b03841660208201527fc50f17198ae53c50e1ad2f06d8348c7d6b31952e4bc9f52b15bb075e6f1bed0b91015b60405180910390a1505050505050565b600a546001600160a01b03163314610cba5760405162461bcd60e51b81526004016108f6906130bb565b806020015115610d5e578060400151600782600001516001600160401b031681548110610ce957610ce96130f3565b90600052602060002090600702016004016000828254610d099190613180565b9091555050606081015181516007805490916001600160401b0316908110610d3357610d336130f3565b90600052602060002090600702016005016000828254610d539190613180565b90915550610df39050565b8060400151600782600001516001600160401b031681548110610d8357610d836130f3565b90600052602060002090600702016004016000828254610da39190613198565b9091555050606081015181516007805490916001600160401b0316908110610dcd57610dcd6130f3565b90600052602060002090600702016005016000828254610ded9190613198565b90915550505b50565b6000336001600160a01b0384168114610e61576040516001600160601b0319606083901b1660208201526c3937b63296b932b3b4b9ba32b960991b6034820152600090604101604051602081830303815290604052805190602001209050610e5f858286612680565b505b6001600160a01b038416600090815260046020526040902054630100000090046001600160401b031615610ebf5750506001600160a01b038216600090815260046020526040902054630100000090046001600160401b0316610f3e565b5050600380546001810182557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b038516908117909155905460009182526004602052604090912080546affffffffffffffff000000191663010000006001600160401b038416021790555b92915050565b60006007836001600160401b031681548110610f6257610f626130f3565b9060005260206000209060070201600201826001600160401b031681548110610f8d57610f8d6130f3565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b0316905092915050565b60008060006003846001600160401b031681548110610fdf57610fdf6130f3565b6000918252602080832091909101546001600160a01b031680835260049091526040909120549095600160581b9091046001600160401b0316945092505050565b60006007826001600160401b03168154811061103e5761103e6130f3565b600091825260209091206001600790920201015492915050565b60008060008060008060006007886001600160401b03168154811061107f5761107f6130f3565b6000918252602090912060079182020154815460ff90911691906001600160401b038b169081106110b2576110b26130f3565b906000526020600020906007020160000160019054906101000a900460ff1660078a6001600160401b0316815481106110ed576110ed6130f3565b906000526020600020906007020160000160029054906101000a900460ff1660078b6001600160401b031681548110611128576111286130f3565b906000526020600020906007020160000160039054906101000a900461ffff1660078c6001600160401b031681548110611164576111646130f3565b90600052602060002090600702016004015460078d6001600160401b031681548110611192576111926130f3565b90600052602060002090600702016005015460078e6001600160401b0316815481106111c0576111c06130f3565b6000918252602090912060079091020160060154959e949d50929b50909950975095506001600160a01b039091169350915050565b6000611200856125f3565b90503361120c85612740565b816001600160a01b0316816001600160a01b03161461126157600081868660405160200161123c939291906131af565b60405160208183030381529060405280519060200120905061125f838286612680565b505b600761126e6001876131fd565b6001600160401b031681548110611287576112876130f3565b600091825260209091206007909102016006015460405163392701c960e21b81526001600160401b03881660048201526001600160a01b039091169063e49c072490602401600060405180830381600087803b1580156112e657600080fd5b505af11580156112fa573d6000803e3d6000fd5b50505050600760018661130d91906131fd565b6001600160401b031681548110611326576113266130f3565b6000918252602080832060036007909302018201805460018181018355918552828520600480830490910180546001600160401b03808f16600895909816949094026101000a9687029684021916959095179094556001600160a01b038716855292825260409093208054928916600160581b0272ffffffffffffffff0000000000000000ff000019909316929092176201000017825586516113d1939290920191908701906128d3565b506001600160a01b038216600081815260046020908152604091829020805460ff1916600117905581516001600160401b038a168152908101929092527f7defc6162296f3490e44c1787f4ae9852a8d6a8e67ba0a69c57bd7be5f8a0b1a9101610c80565b600060036114456001876131fd565b6001600160401b03168154811061145e5761145e6130f3565b60009182526020808320909101546001600160a01b0316808352600490915260409091205490915060ff62010000909104166001146114c45760405162461bcd60e51b81526020600482015260026024820152614e5560f01b60448201526064016108f6565b60006114cf856120e5565b9050336001600160a01b03831681146115515760408051606083811b6001600160601b031916602083015260c08a901b6001600160c01b031916603483015260e089901b6001600160e01b031916603c8301529181018790526000910160405160208183030381529060405280519060200120905061154f848287612680565b505b6001600160a01b03831660009081526004602052604090205460079061158990600190600160581b90046001600160401b03166131fd565b6001600160401b0316815481106115a2576115a26130f3565b600091825260209091206007909102016006015460405163e04f98ed60e01b81526001600160401b038916600482015263ffffffff881660248201526001600160a01b0385811660448301528481166064830152608482018890529091169063e04f98ed9060a401600060405180830381600087803b15801561162457600080fd5b505af1158015611638573d6000803e3d6000fd5b5050505050505050505050565b60006007836001600160401b031681548110611663576116636130f3565b9060005260206000209060070201600301826001600160401b031681548110610f8d57610f8d6130f3565b60006007826001600160401b0316815481106116ac576116ac6130f3565b60009182526020909120600660079092020101546001600160a01b031692915050565b60006007836001600160401b0316815481106116ed576116ed6130f3565b9060005260206000209060070201600101826001600160401b031681548110610f8d57610f8d6130f3565b6000546001600160a01b031633146117425760405162461bcd60e51b81526004016108f6906130bb565b600761174f6001836131fd565b6001600160401b031681548110611768576117686130f3565b6000918252602090912060079091020154610100900460ff16156117b35760405162461bcd60e51b815260206004820152600260248201526123a160f11b60448201526064016108f6565b60006117c08360036127f3565b905060076117cf6001846131fd565b6001600160401b0316815481106117e8576117e86130f3565b60009182526020909120600790910201600601546040516350cbb46f60e01b81526001600160401b03851660048201526001600160a01b03909116906350cbb46f90602401600060405180830381600087803b15801561184757600080fd5b505af115801561185b573d6000803e3d6000fd5b50505050600760018361186e91906131fd565b6001600160401b031681548110611887576118876130f3565b600091825260208083206007928302016001908101805480830182559085528285206004808304909101805460039093166008026101000a6001600160401b03818102199094168b8516919091021790556001600160a01b03871686529092526040909320805460ff67ffffffffffffffff60581b011916600160581b9287169290920260ff191691909117831790559061192290846131fd565b6001600160401b03168154811061193b5761193b6130f3565b60009182526020909120600791820201546301000000900461ffff16906119636001856131fd565b6001600160401b03168154811061197c5761197c6130f3565b906000526020600020906007020160010180549050106119dd57600160076119a482856131fd565b6001600160401b0316815481106119bd576119bd6130f3565b60009182526020909120600790910201805460ff19169115159190911790555b505050565b60006119ef8460026127f3565b90506119fa83612740565b336001600160a01b0382168114611a61576040516001600160601b0319606083901b1660208201526001600160c01b031960c086901b166034820152600090603c01604051602081830303815290604052805190602001209050611a5f838286612680565b505b6007611a6e6001866131fd565b6001600160401b031681548110611a8757611a876130f3565b600091825260208083206007929092029091016002018054600180820183559184528284206004808304909101805460039093166008026101000a6001600160401b03818102199094169b8416029a909a179099556001600160a01b0395909516835296905260409020805460ff67ffffffffffffffff60581b011916600160581b959093169490940260ff1916919091179093179091555050565b60006003826001600160401b031681548110611b4157611b416130f3565b6000918252602090912001546001600160a01b031692915050565b6000546001600160a01b03163314611b865760405162461bcd60e51b81526004016108f6906130bb565b600591909155600655565b6000546001600160a01b03163314611bbb5760405162461bcd60e51b81526004016108f6906130bb565b806007611bc96001856131fd565b6001600160401b031681548110611be257611be26130f3565b906000526020600020906007020160060160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505050565b6000806003611c2b6001866131fd565b6001600160401b031681548110611c4457611c446130f3565b60009182526020808320909101546001600160a01b0316808352600490915260409091205490915062010000900460ff90811690841614611cac5760405162461bcd60e51b8152602060048201526002602482015261524560f01b60448201526064016108f6565b9392505050565b6000611cbe846120e5565b905060008311611ccd57600080fd5b6000856001600160401b031611611ce357600080fd5b60006003611cf26001886131fd565b6001600160401b031681548110611d0b57611d0b6130f3565b60009182526020808320909101546001600160a01b031680835260049091526040822054909250600790611d52906001906001600160401b03600160581b909104166131fd565b6001600160401b031681548110611d6b57611d6b6130f3565b60009182526020808320600792909202909101600601546001600160a01b0385811684526004909252604090922054600254928216935062010000900460ff1691339116811415611dd057600254600099506001600160a01b03169350889150611eaa565b6001600160a01b038416600090815260046020526040902054610100900460ff1615611e235760405162461bcd60e51b8152602060048201526002602482015261494960f01b60448201526064016108f6565b8160ff1660021415611e3457600080fd5b836001600160a01b0316816001600160a01b031614611eaa576040516001600160601b0319606083901b1660208201526001600160e01b031960e08a901b16603482015260388101889052600090605801604051602081830303815290604052805190602001209050611ea8858289612680565b505b604051635d5c6b1d60e11b81526001600160401b038a16600482015263ffffffff8916602482015260ff831660448201526001600160a01b038681166064830152858116608483015260a4820189905284169063bab8d63a9060c401600060405180830381600087803b158015611f2057600080fd5b505af1158015611f34573d6000803e3d6000fd5b50505050505050505050505050565b336000611f4f846125f3565b9050806001600160a01b0316826001600160a01b031614611fbd576040516001600160601b0319606084901b16602082015267383937bb34b232b960c11b6034820152600090603c01604051602081830303815290604052805190602001209050611fbb828286612680565b505b60015460405163761966b760e01b81526001600160401b0386166004820152600060248201819052916001600160a01b03169063761966b790604401602060405180830381865afa158015612016573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203a9190613151565b90506006548110156120735760405162461bcd60e51b81526020600482015260026024820152614e4560f01b60448201526064016108f6565b6001600160a01b038216600081815260046020908152604091829020805462ff000019166202000017905581516001600160401b0389168152908101929092527ff06105db8b89019d932bb3ec85a22bbed723c3616043e02ca9857f3ba37005a5910160405180910390a15050505050565b600854604051633c7bdc1960e01b815263ffffffff831660048201526000916001600160a01b0316908190633c7bdc1990602401602060405180830381865afa158015612136573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061215a9190613225565b61218b5760405162461bcd60e51b8152602060048201526002602482015261544560f01b60448201526064016108f6565b604051637bdbeb5b60e11b815263ffffffff841660048201526001600160a01b0382169063f7b7d6b690602401602060405180830381865afa1580156121d5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cac9190613242565b6000546001600160a01b031633146122235760405162461bcd60e51b81526004016108f6906130bb565b600180546001600160a01b039485166001600160a01b0319918216179091556009805493851693821693909317909255600a8054919093169116179055565b600080546001600160a01b0316331461228d5760405162461bcd60e51b81526004016108f6906130bb565b604080516101408101825260008082526020820181905291810182905260608082018390526080820181905260a0820181905260c082015260e081018290526101008101829052610120810191909152600780546001810182556000829052825191027fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688810180546020808601516040870151606088015161ffff1663010000000264ffff0000001991151562010000029190911664ffffff0000199215156101000261ff00199815159890981661ffff199095169490941796909617169190911793909317815560808401518051859492936123b0937fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c689909101920190612957565b5060a082015180516123cc916002840191602090910190612957565b5060c082015180516123e8916003840191602090910190612957565b5060e08201516004820155610100820151600582015561012090910151600690910180546001600160a01b0319166001600160a01b0390921691909117905560075460005b85518160ff1610156124bb576000612462878360ff1681518110612453576124536130f3565b602002602001015160036127f3565b6001600160a01b03166000908152600460205260409020805460ff196001600160401b038616600160581b021660ff67ffffffffffffffff60581b011990911617600117905550806124b38161325f565b91505061242d565b508360076124ca6001846131fd565b6001600160401b0316815481106124e3576124e36130f3565b906000526020600020906007020160000160036101000a81548161ffff021916908361ffff16021790555084600760018361251e91906131fd565b6001600160401b031681548110612537576125376130f3565b9060005260206000209060070201600101908051906020019061255b929190612957565b508361ffff168551106125af576001600761257682846131fd565b6001600160401b03168154811061258f5761258f6130f3565b60009182526020909120600790910201805460ff19169115159190911790555b6040516001600160401b03821681527fc79ca32352cc5529f3c78b5cb44574fbc979555a04f5b6425ed2595417da2e649060200160405180910390a1949350505050565b60008060036126036001856131fd565b6001600160401b03168154811061261c5761261c6130f3565b60009182526020808320909101546001600160a01b0316808352600490915260409091205490915060ff620100009091041615610f3e5760405162461bcd60e51b8152602060048201526002602482015261272360f11b60448201526064016108f6565b6040516319045a2560e01b81526001600160a01b0384169073__$6230a6feddd2d01b0a9ffb5c5e188a1065$__906319045a25906126c4908690869060040161327f565b602060405180830381865af41580156126e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127059190613242565b6001600160a01b0316146119dd5760405162461bcd60e51b8152602060048201526002602482015261494360f01b60448201526064016108f6565b600761274d6001836131fd565b6001600160401b031681548110612766576127666130f3565b600091825260209091206007909102015460ff1680156127c25750600761278e6001836131fd565b6001600160401b0316815481106127a7576127a76130f3565b6000918252602090912060079091020154610100900460ff16155b610df35760405162461bcd60e51b815260206004820152600260248201526123a160f11b60448201526064016108f6565b60008060036128036001866131fd565b6001600160401b03168154811061281c5761281c6130f3565b60009182526020808320909101546001600160a01b0316808352600490915260409091205490915062010000900460ff90811690841614801561287857506001600160a01b03811660009081526004602052604090205460ff16155b80156128a257506001600160a01b038116600090815260046020526040902054610100900460ff16155b611cac5760405162461bcd60e51b8152602060048201526002602482015261414560f01b60448201526064016108f6565b8280546128df90613080565b90600052602060002090601f0160209004810192826129015760008555612947565b82601f1061291a57805160ff1916838001178555612947565b82800160010185558215612947579182015b8281111561294757825182559160200191906001019061292c565b50612953929150612a06565b5090565b828054828255906000526020600020906003016004900481019282156129475791602002820160005b838211156129ca57835183826101000a8154816001600160401b0302191690836001600160401b031602179055509260200192600801602081600701049283019260010302612980565b80156129fd5782816101000a8154906001600160401b0302191690556008016020816007010492830192600103026129ca565b50506129539291505b5b808211156129535760008155600101612a07565b6001600160a01b0381168114610df357600080fd5b600060208284031215612a4257600080fd5b8135611cac81612a1b565b60005b83811015612a68578181015183820152602001612a50565b83811115612a77576000848401525b50505050565b60008151808452612a95816020860160208601612a4d565b601f01601f19169290920160200192915050565b8615158152851515602082015260ff8516604082015260006001600160401b03808616606084015280851660808401525060c060a0830152612aee60c0830184612a7d565b98975050505050505050565b80356001600160401b0381168114612b1157600080fd5b919050565b600060208284031215612b2857600080fd5b611cac82612afa565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715612b6f57612b6f612b31565b604052919050565b600082601f830112612b8857600080fd5b81356001600160401b03811115612ba157612ba1612b31565b612bb4601f8201601f1916602001612b47565b818152846020838601011115612bc957600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215612bfb57600080fd5b612c0484612afa565b925060208401356001600160401b0380821115612c2057600080fd5b612c2c87838801612b77565b93506040860135915080821115612c4257600080fd5b50612c4f86828701612b77565b9150509250925092565b8015158114610df357600080fd5b600060808284031215612c7957600080fd5b604051608081018181106001600160401b0382111715612c9b57612c9b612b31565b604052612ca783612afa565b81526020830135612cb781612c59565b6020820152604083810135908201526060928301359281019290925250919050565b60008060408385031215612cec57600080fd5b8235612cf781612a1b565b915060208301356001600160401b03811115612d1257600080fd5b612d1e85828601612b77565b9150509250929050565b60008060408385031215612d3b57600080fd5b612d4483612afa565b9150612d5260208401612afa565b90509250929050565b60008060008060808587031215612d7157600080fd5b612d7a85612afa565b9350612d8860208601612afa565b925060408501356001600160401b0380821115612da457600080fd5b612db088838901612b77565b93506060870135915080821115612dc657600080fd5b50612dd387828801612b77565b91505092959194509250565b63ffffffff81168114610df357600080fd5b60008060008060808587031215612e0757600080fd5b612e1085612afa565b93506020850135612e2081612ddf565b92506040850135915060608501356001600160401b03811115612e4257600080fd5b612dd387828801612b77565b600080600060608486031215612e6357600080fd5b612e6c84612afa565b9250612e7a60208501612afa565b915060408401356001600160401b03811115612e9557600080fd5b612c4f86828701612b77565b60008060408385031215612eb457600080fd5b50508035926020909101359150565b60008060408385031215612ed657600080fd5b612edf83612afa565b91506020830135612eef81612a1b565b809150509250929050565b60008060408385031215612f0d57600080fd5b612f1683612afa565b9150602083013560ff81168114612eef57600080fd5b60008060408385031215612f3f57600080fd5b612cf783612afa565b600060208284031215612f5a57600080fd5b8135611cac81612ddf565b600080600060608486031215612f7a57600080fd5b8335612f8581612a1b565b92506020840135612f9581612a1b565b91506040840135612fa581612a1b565b809150509250925092565b803561ffff81168114612b1157600080fd5b60008060408385031215612fd557600080fd5b82356001600160401b0380821115612fec57600080fd5b818501915085601f83011261300057600080fd5b813560208282111561301457613014612b31565b8160051b9250613025818401612b47565b828152928401810192818101908985111561303f57600080fd5b948201945b848610156130645761305586612afa565b82529482019490820190613044565b96506130739050878201612fb0565b9450505050509250929050565b600181811c9082168061309457607f821691505b602082108114156130b557634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252600190820152602760f91b604082015260600190565b6000602082840312156130e857600080fd5b8151611cac81612ddf565b634e487b7160e01b600052603260045260246000fd5b6bffffffffffffffffffffffff198360601b16815260008251613133816014850160208701612a4d565b6535b2b2b832b960d11b6014939091019283015250601a0192915050565b60006020828403121561316357600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b600082198211156131935761319361316a565b500190565b6000828210156131aa576131aa61316a565b500390565b6bffffffffffffffffffffffff198460601b1681526001600160401b0360c01b8360c01b166014820152600082516131ee81601c850160208701612a4d565b91909101601c01949350505050565b60006001600160401b038381169083168181101561321d5761321d61316a565b039392505050565b60006020828403121561323757600080fd5b8151611cac81612c59565b60006020828403121561325457600080fd5b8151611cac81612a1b565b600060ff821660ff8114156132765761327661316a565b60010192915050565b8281526040602082015260006132986040830184612a7d565b94935050505056fea2646970667358221220df1e5cac7efd13e3c4867d7fccd18c3c28832ec94eed23423adda6968d1d8fc864736f6c634300080b0033608060405234801561001057600080fd5b50600280546001600160a01b03191633179055610518806100326000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632df2685f1461005c5780633c7bdc19146100905780638ff8e15a146100b3578063b04c6c68146100db578063f7b7d6b6146100e3575b600080fd5b61006f61006a36600461042b565b61010e565b6040805163ffffffff90931683529015156020830152015b60405180910390f35b6100a361009e36600461045b565b6101e2565b6040519015158152602001610087565b6100c66100c136600461042b565b610259565b60405163ffffffff9091168152602001610087565b6000546100c6565b6100f66100f136600461045b565b61037c565b6040516001600160a01b039091168152602001610087565b6001600160a01b038116600090815260016020526040808220549051633c7bdc1960e01b815263ffffffff9091166004820181905282913090633c7bdc1990602401602060405180830381865afa15801561016d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101919190610481565b156101d85760008163ffffffff16815481106101af576101af6104a3565b6000918252602090912001546001600160a01b03858116911614156101d8579360019350915050565b9360009350915050565b6000805463ffffffff8316108015610244575060016000808463ffffffff1681548110610211576102116104a3565b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff64010000000090910416155b1561025157506001919050565b506000919050565b6002546000906001600160a01b0316331461029f5760405162461bcd60e51b81526020600482015260026024820152614e4f60f01b604482015260640160405180910390fd5b6102a8826103c1565b156102cf57506001600160a01b031660009081526001602052604090205463ffffffff1690565b60008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563820180546001600160a01b0319166001600160a01b0387169081179091558084526020918252604093849020805463ffffffff191663ffffffff851690811790915584519182529181019190915290917f7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6910160405180910390a192915050565b6000805463ffffffff831610156102515760008263ffffffff16815481106103a6576103a66104a3565b6000918252602090912001546001600160a01b031692915050565b6000805b60005481101561042257826001600160a01b0316600082815481106103ec576103ec6104a3565b6000918252602090912001546001600160a01b031614156104105750600192915050565b8061041a816104b9565b9150506103c5565b50600092915050565b60006020828403121561043d57600080fd5b81356001600160a01b038116811461045457600080fd5b9392505050565b60006020828403121561046d57600080fd5b813563ffffffff8116811461045457600080fd5b60006020828403121561049357600080fd5b8151801515811461045457600080fd5b634e487b7160e01b600052603260045260246000fd5b60006000198214156104db57634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122080162e603925909c14e71489135a62d543ddc09b2963d9ccc556d03ae2890a0064736f6c634300080b0033",
}

// RoleABI is the input ABI used to generate the binding from.
// Deprecated: Use RoleMetaData.ABI instead.
var RoleABI = RoleMetaData.ABI

// Deprecated: Use RoleMetaData.Sigs instead.
// RoleFuncSigs maps the 4-byte function signature to its string representation.
var RoleFuncSigs = RoleMetaData.Sigs

// RoleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoleMetaData.Bin instead.
var RoleBin = RoleMetaData.Bin

// DeployRole deploys a new Ethereum contract, binding an instance of Role to it.
func DeployRole(auth *bind.TransactOpts, backend bind.ContractBackend, f common.Address, t common.Address, pk *big.Int, ppro *big.Int) (common.Address, *types.Transaction, *Role, error) {
	parsed, err := RoleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	recoverAddr, _, _, _ := DeployRecover(auth, backend)
	RoleBin = strings.Replace(RoleBin, "__$6230a6feddd2d01b0a9ffb5c5e188a1065$__", recoverAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoleBin), backend, f, t, pk, ppro)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Role{RoleCaller: RoleCaller{contract: contract}, RoleTransactor: RoleTransactor{contract: contract}, RoleFilterer: RoleFilterer{contract: contract}}, nil
}

// Role is an auto generated Go binding around an Ethereum contract.
type Role struct {
	RoleCaller     // Read-only binding to the contract
	RoleTransactor // Write-only binding to the contract
	RoleFilterer   // Log filterer for contract events
}

// RoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoleSession struct {
	Contract     *Role             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoleCallerSession struct {
	Contract *RoleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoleTransactorSession struct {
	Contract     *RoleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoleRaw struct {
	Contract *Role // Generic contract binding to access the raw methods on
}

// RoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoleCallerRaw struct {
	Contract *RoleCaller // Generic read-only contract binding to access the raw methods on
}

// RoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoleTransactorRaw struct {
	Contract *RoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRole creates a new instance of Role, bound to a specific deployed contract.
func NewRole(address common.Address, backend bind.ContractBackend) (*Role, error) {
	contract, err := bindRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Role{RoleCaller: RoleCaller{contract: contract}, RoleTransactor: RoleTransactor{contract: contract}, RoleFilterer: RoleFilterer{contract: contract}}, nil
}

// NewRoleCaller creates a new read-only instance of Role, bound to a specific deployed contract.
func NewRoleCaller(address common.Address, caller bind.ContractCaller) (*RoleCaller, error) {
	contract, err := bindRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoleCaller{contract: contract}, nil
}

// NewRoleTransactor creates a new write-only instance of Role, bound to a specific deployed contract.
func NewRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*RoleTransactor, error) {
	contract, err := bindRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoleTransactor{contract: contract}, nil
}

// NewRoleFilterer creates a new log filterer instance of Role, bound to a specific deployed contract.
func NewRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*RoleFilterer, error) {
	contract, err := bindRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoleFilterer{contract: contract}, nil
}

// bindRole binds a generic wrapper to an already deployed contract.
func bindRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Role *RoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Role.Contract.RoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Role *RoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Role.Contract.RoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Role *RoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Role.Contract.RoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Role *RoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Role.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Role *RoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Role.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Role *RoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Role.Contract.contract.Transact(opts, method, params...)
}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_Role *RoleCaller) CheckIR(opts *bind.CallOpts, _index uint64, _rType uint8) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "checkIR", _index, _rType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_Role *RoleSession) CheckIR(_index uint64, _rType uint8) (common.Address, error) {
	return _Role.Contract.CheckIR(&_Role.CallOpts, _index, _rType)
}

// CheckIR is a free data retrieval call binding the contract method 0xcf8e99a8.
//
// Solidity: function checkIR(uint64 _index, uint8 _rType) view returns(address)
func (_Role *RoleCallerSession) CheckIR(_index uint64, _rType uint8) (common.Address, error) {
	return _Role.Contract.CheckIR(&_Role.CallOpts, _index, _rType)
}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_Role *RoleCaller) CheckT(opts *bind.CallOpts, tIndex uint32) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "checkT", tIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_Role *RoleSession) CheckT(tIndex uint32) (common.Address, error) {
	return _Role.Contract.CheckT(&_Role.CallOpts, tIndex)
}

// CheckT is a free data retrieval call binding the contract method 0xde92e994.
//
// Solidity: function checkT(uint32 tIndex) view returns(address)
func (_Role *RoleCallerSession) CheckT(tIndex uint32) (common.Address, error) {
	return _Role.Contract.CheckT(&_Role.CallOpts, tIndex)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() view returns(address)
func (_Role *RoleCaller) Foundation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "foundation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() view returns(address)
func (_Role *RoleSession) Foundation() (common.Address, error) {
	return _Role.Contract.Foundation(&_Role.CallOpts)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() view returns(address)
func (_Role *RoleCallerSession) Foundation() (common.Address, error) {
	return _Role.Contract.Foundation(&_Role.CallOpts)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_Role *RoleCaller) GetAddr(opts *bind.CallOpts, i uint64) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getAddr", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_Role *RoleSession) GetAddr(i uint64) (common.Address, error) {
	return _Role.Contract.GetAddr(&_Role.CallOpts, i)
}

// GetAddr is a free data retrieval call binding the contract method 0x9332aa6e.
//
// Solidity: function getAddr(uint64 i) view returns(address)
func (_Role *RoleCallerSession) GetAddr(i uint64) (common.Address, error) {
	return _Role.Contract.GetAddr(&_Role.CallOpts, i)
}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_Role *RoleCaller) GetAddrGindex(opts *bind.CallOpts, i uint64) (common.Address, uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getAddrGindex", i)

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
func (_Role *RoleSession) GetAddrGindex(i uint64) (common.Address, uint64, error) {
	return _Role.Contract.GetAddrGindex(&_Role.CallOpts, i)
}

// GetAddrGindex is a free data retrieval call binding the contract method 0x421795e5.
//
// Solidity: function getAddrGindex(uint64 i) view returns(address, uint64)
func (_Role *RoleCallerSession) GetAddrGindex(i uint64) (common.Address, uint64, error) {
	return _Role.Contract.GetAddrGindex(&_Role.CallOpts, i)
}

// GetAddrsNum is a free data retrieval call binding the contract method 0x7ba783be.
//
// Solidity: function getAddrsNum() view returns(uint64)
func (_Role *RoleCaller) GetAddrsNum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getAddrsNum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAddrsNum is a free data retrieval call binding the contract method 0x7ba783be.
//
// Solidity: function getAddrsNum() view returns(uint64)
func (_Role *RoleSession) GetAddrsNum() (uint64, error) {
	return _Role.Contract.GetAddrsNum(&_Role.CallOpts)
}

// GetAddrsNum is a free data retrieval call binding the contract method 0x7ba783be.
//
// Solidity: function getAddrsNum() view returns(uint64)
func (_Role *RoleCallerSession) GetAddrsNum() (uint64, error) {
	return _Role.Contract.GetAddrsNum(&_Role.CallOpts)
}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_Role *RoleCaller) GetFsAddr(opts *bind.CallOpts, i uint64) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getFsAddr", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_Role *RoleSession) GetFsAddr(i uint64) (common.Address, error) {
	return _Role.Contract.GetFsAddr(&_Role.CallOpts, i)
}

// GetFsAddr is a free data retrieval call binding the contract method 0x5f096376.
//
// Solidity: function getFsAddr(uint64 i) view returns(address)
func (_Role *RoleCallerSession) GetFsAddr(i uint64) (common.Address, error) {
	return _Role.Contract.GetFsAddr(&_Role.CallOpts, i)
}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_Role *RoleCaller) GetGKNum(opts *bind.CallOpts, i uint64) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getGKNum", i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_Role *RoleSession) GetGKNum(i uint64) (uint64, error) {
	return _Role.Contract.GetGKNum(&_Role.CallOpts, i)
}

// GetGKNum is a free data retrieval call binding the contract method 0x429fb683.
//
// Solidity: function getGKNum(uint64 i) view returns(uint64)
func (_Role *RoleCallerSession) GetGKNum(i uint64) (uint64, error) {
	return _Role.Contract.GetGKNum(&_Role.CallOpts, i)
}

// GetGU is a free data retrieval call binding the contract method 0x54b365a3.
//
// Solidity: function getGU(uint64 ig, uint64 iu) view returns(uint64)
func (_Role *RoleCaller) GetGU(opts *bind.CallOpts, ig uint64, iu uint64) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getGU", ig, iu)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGU is a free data retrieval call binding the contract method 0x54b365a3.
//
// Solidity: function getGU(uint64 ig, uint64 iu) view returns(uint64)
func (_Role *RoleSession) GetGU(ig uint64, iu uint64) (uint64, error) {
	return _Role.Contract.GetGU(&_Role.CallOpts, ig, iu)
}

// GetGU is a free data retrieval call binding the contract method 0x54b365a3.
//
// Solidity: function getGU(uint64 ig, uint64 iu) view returns(uint64)
func (_Role *RoleCallerSession) GetGU(ig uint64, iu uint64) (uint64, error) {
	return _Role.Contract.GetGU(&_Role.CallOpts, ig, iu)
}

// GetGUPNum is a free data retrieval call binding the contract method 0x0d0b065c.
//
// Solidity: function getGUPNum(uint64 i) view returns(uint64, uint64)
func (_Role *RoleCaller) GetGUPNum(opts *bind.CallOpts, i uint64) (uint64, uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getGUPNum", i)

	if err != nil {
		return *new(uint64), *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return out0, out1, err

}

// GetGUPNum is a free data retrieval call binding the contract method 0x0d0b065c.
//
// Solidity: function getGUPNum(uint64 i) view returns(uint64, uint64)
func (_Role *RoleSession) GetGUPNum(i uint64) (uint64, uint64, error) {
	return _Role.Contract.GetGUPNum(&_Role.CallOpts, i)
}

// GetGUPNum is a free data retrieval call binding the contract method 0x0d0b065c.
//
// Solidity: function getGUPNum(uint64 i) view returns(uint64, uint64)
func (_Role *RoleCallerSession) GetGUPNum(i uint64) (uint64, uint64, error) {
	return _Role.Contract.GetGUPNum(&_Role.CallOpts, i)
}

// GetGroupInfo is a free data retrieval call binding the contract method 0x4496c991.
//
// Solidity: function getGroupInfo(uint64 i) view returns(bool, bool, bool, uint16, uint256, uint256, address)
func (_Role *RoleCaller) GetGroupInfo(opts *bind.CallOpts, i uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getGroupInfo", i)

	if err != nil {
		return *new(bool), *new(bool), *new(bool), *new(uint16), *new(*big.Int), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(uint16)).(*uint16)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetGroupInfo is a free data retrieval call binding the contract method 0x4496c991.
//
// Solidity: function getGroupInfo(uint64 i) view returns(bool, bool, bool, uint16, uint256, uint256, address)
func (_Role *RoleSession) GetGroupInfo(i uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	return _Role.Contract.GetGroupInfo(&_Role.CallOpts, i)
}

// GetGroupInfo is a free data retrieval call binding the contract method 0x4496c991.
//
// Solidity: function getGroupInfo(uint64 i) view returns(bool, bool, bool, uint16, uint256, uint256, address)
func (_Role *RoleCallerSession) GetGroupInfo(i uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	return _Role.Contract.GetGroupInfo(&_Role.CallOpts, i)
}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_Role *RoleCaller) GetGroupK(opts *bind.CallOpts, ig uint64, ik uint64) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getGroupK", ig, ik)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_Role *RoleSession) GetGroupK(ig uint64, ik uint64) (uint64, error) {
	return _Role.Contract.GetGroupK(&_Role.CallOpts, ig, ik)
}

// GetGroupK is a free data retrieval call binding the contract method 0x64fe6290.
//
// Solidity: function getGroupK(uint64 ig, uint64 ik) view returns(uint64)
func (_Role *RoleCallerSession) GetGroupK(ig uint64, ik uint64) (uint64, error) {
	return _Role.Contract.GetGroupK(&_Role.CallOpts, ig, ik)
}

// GetGroupP is a free data retrieval call binding the contract method 0x3d180dd3.
//
// Solidity: function getGroupP(uint64 ig, uint64 ip) view returns(uint64)
func (_Role *RoleCaller) GetGroupP(opts *bind.CallOpts, ig uint64, ip uint64) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getGroupP", ig, ip)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGroupP is a free data retrieval call binding the contract method 0x3d180dd3.
//
// Solidity: function getGroupP(uint64 ig, uint64 ip) view returns(uint64)
func (_Role *RoleSession) GetGroupP(ig uint64, ip uint64) (uint64, error) {
	return _Role.Contract.GetGroupP(&_Role.CallOpts, ig, ip)
}

// GetGroupP is a free data retrieval call binding the contract method 0x3d180dd3.
//
// Solidity: function getGroupP(uint64 ig, uint64 ip) view returns(uint64)
func (_Role *RoleCallerSession) GetGroupP(ig uint64, ip uint64) (uint64, error) {
	return _Role.Contract.GetGroupP(&_Role.CallOpts, ig, ip)
}

// GetGroupsNum is a free data retrieval call binding the contract method 0x5d39c56b.
//
// Solidity: function getGroupsNum() view returns(uint64)
func (_Role *RoleCaller) GetGroupsNum(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getGroupsNum")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGroupsNum is a free data retrieval call binding the contract method 0x5d39c56b.
//
// Solidity: function getGroupsNum() view returns(uint64)
func (_Role *RoleSession) GetGroupsNum() (uint64, error) {
	return _Role.Contract.GetGroupsNum(&_Role.CallOpts)
}

// GetGroupsNum is a free data retrieval call binding the contract method 0x5d39c56b.
//
// Solidity: function getGroupsNum() view returns(uint64)
func (_Role *RoleCallerSession) GetGroupsNum() (uint64, error) {
	return _Role.Contract.GetGroupsNum(&_Role.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Role *RoleCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Role *RoleSession) GetOwner() (common.Address, error) {
	return _Role.Contract.GetOwner(&_Role.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Role *RoleCallerSession) GetOwner() (common.Address, error) {
	return _Role.Contract.GetOwner(&_Role.CallOpts)
}

// GetRoleIndex is a free data retrieval call binding the contract method 0xae0e4ffa.
//
// Solidity: function getRoleIndex(address acc) view returns(uint64)
func (_Role *RoleCaller) GetRoleIndex(opts *bind.CallOpts, acc common.Address) (uint64, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getRoleIndex", acc)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRoleIndex is a free data retrieval call binding the contract method 0xae0e4ffa.
//
// Solidity: function getRoleIndex(address acc) view returns(uint64)
func (_Role *RoleSession) GetRoleIndex(acc common.Address) (uint64, error) {
	return _Role.Contract.GetRoleIndex(&_Role.CallOpts, acc)
}

// GetRoleIndex is a free data retrieval call binding the contract method 0xae0e4ffa.
//
// Solidity: function getRoleIndex(address acc) view returns(uint64)
func (_Role *RoleCallerSession) GetRoleIndex(acc common.Address) (uint64, error) {
	return _Role.Contract.GetRoleIndex(&_Role.CallOpts, acc)
}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_Role *RoleCaller) GetRoleInfo(opts *bind.CallOpts, acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "getRoleInfo", acc)

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
func (_Role *RoleSession) GetRoleInfo(acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	return _Role.Contract.GetRoleInfo(&_Role.CallOpts, acc)
}

// GetRoleInfo is a free data retrieval call binding the contract method 0x07483499.
//
// Solidity: function getRoleInfo(address acc) view returns(bool, bool, uint8, uint64, uint64, bytes)
func (_Role *RoleCallerSession) GetRoleInfo(acc common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	return _Role.Contract.GetRoleInfo(&_Role.CallOpts, acc)
}

// Issuance is a free data retrieval call binding the contract method 0x863623bb.
//
// Solidity: function issuance() view returns(address)
func (_Role *RoleCaller) Issuance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "issuance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Issuance is a free data retrieval call binding the contract method 0x863623bb.
//
// Solidity: function issuance() view returns(address)
func (_Role *RoleSession) Issuance() (common.Address, error) {
	return _Role.Contract.Issuance(&_Role.CallOpts)
}

// Issuance is a free data retrieval call binding the contract method 0x863623bb.
//
// Solidity: function issuance() view returns(address)
func (_Role *RoleCallerSession) Issuance() (common.Address, error) {
	return _Role.Contract.Issuance(&_Role.CallOpts)
}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_Role *RoleCaller) PledgeK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "pledgeK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_Role *RoleSession) PledgeK() (*big.Int, error) {
	return _Role.Contract.PledgeK(&_Role.CallOpts)
}

// PledgeK is a free data retrieval call binding the contract method 0xa6ed590b.
//
// Solidity: function pledgeK() view returns(uint256)
func (_Role *RoleCallerSession) PledgeK() (*big.Int, error) {
	return _Role.Contract.PledgeK(&_Role.CallOpts)
}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_Role *RoleCaller) PledgeP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "pledgeP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_Role *RoleSession) PledgeP() (*big.Int, error) {
	return _Role.Contract.PledgeP(&_Role.CallOpts)
}

// PledgeP is a free data retrieval call binding the contract method 0x8ba61d28.
//
// Solidity: function pledgeP() view returns(uint256)
func (_Role *RoleCallerSession) PledgeP() (*big.Int, error) {
	return _Role.Contract.PledgeP(&_Role.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_Role *RoleCaller) PledgePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "pledgePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_Role *RoleSession) PledgePool() (common.Address, error) {
	return _Role.Contract.PledgePool(&_Role.CallOpts)
}

// PledgePool is a free data retrieval call binding the contract method 0xde909560.
//
// Solidity: function pledgePool() view returns(address)
func (_Role *RoleCallerSession) PledgePool() (common.Address, error) {
	return _Role.Contract.PledgePool(&_Role.CallOpts)
}

// RToken is a free data retrieval call binding the contract method 0x40c65f72.
//
// Solidity: function rToken() view returns(address)
func (_Role *RoleCaller) RToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "rToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RToken is a free data retrieval call binding the contract method 0x40c65f72.
//
// Solidity: function rToken() view returns(address)
func (_Role *RoleSession) RToken() (common.Address, error) {
	return _Role.Contract.RToken(&_Role.CallOpts)
}

// RToken is a free data retrieval call binding the contract method 0x40c65f72.
//
// Solidity: function rToken() view returns(address)
func (_Role *RoleCallerSession) RToken() (common.Address, error) {
	return _Role.Contract.RToken(&_Role.CallOpts)
}

// Rolefs is a free data retrieval call binding the contract method 0x2d6d777e.
//
// Solidity: function rolefs() view returns(address)
func (_Role *RoleCaller) Rolefs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Role.contract.Call(opts, &out, "rolefs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rolefs is a free data retrieval call binding the contract method 0x2d6d777e.
//
// Solidity: function rolefs() view returns(address)
func (_Role *RoleSession) Rolefs() (common.Address, error) {
	return _Role.Contract.Rolefs(&_Role.CallOpts)
}

// Rolefs is a free data retrieval call binding the contract method 0x2d6d777e.
//
// Solidity: function rolefs() view returns(address)
func (_Role *RoleCallerSession) Rolefs() (common.Address, error) {
	return _Role.Contract.Rolefs(&_Role.CallOpts)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_Role *RoleTransactor) AddKeeperToGroup(opts *bind.TransactOpts, _index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "addKeeperToGroup", _index, _gIndex)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_Role *RoleSession) AddKeeperToGroup(_index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _Role.Contract.AddKeeperToGroup(&_Role.TransactOpts, _index, _gIndex)
}

// AddKeeperToGroup is a paid mutator transaction binding the contract method 0x7977031c.
//
// Solidity: function addKeeperToGroup(uint64 _index, uint64 _gIndex) returns()
func (_Role *RoleTransactorSession) AddKeeperToGroup(_index uint64, _gIndex uint64) (*types.Transaction, error) {
	return _Role.Contract.AddKeeperToGroup(&_Role.TransactOpts, _index, _gIndex)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_Role *RoleTransactor) AddProviderToGroup(opts *bind.TransactOpts, _index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "addProviderToGroup", _index, _gIndex, sign)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_Role *RoleSession) AddProviderToGroup(_index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.AddProviderToGroup(&_Role.TransactOpts, _index, _gIndex, sign)
}

// AddProviderToGroup is a paid mutator transaction binding the contract method 0x8ae21af3.
//
// Solidity: function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes sign) returns()
func (_Role *RoleTransactorSession) AddProviderToGroup(_index uint64, _gIndex uint64, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.AddProviderToGroup(&_Role.TransactOpts, _index, _gIndex, sign)
}

// AlterOwner is a paid mutator transaction binding the contract method 0x0ca05f9f.
//
// Solidity: function alterOwner(address newOwner) returns()
func (_Role *RoleTransactor) AlterOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "alterOwner", newOwner)
}

// AlterOwner is a paid mutator transaction binding the contract method 0x0ca05f9f.
//
// Solidity: function alterOwner(address newOwner) returns()
func (_Role *RoleSession) AlterOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Role.Contract.AlterOwner(&_Role.TransactOpts, newOwner)
}

// AlterOwner is a paid mutator transaction binding the contract method 0x0ca05f9f.
//
// Solidity: function alterOwner(address newOwner) returns()
func (_Role *RoleTransactorSession) AlterOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Role.Contract.AlterOwner(&_Role.TransactOpts, newOwner)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_Role *RoleTransactor) CreateGroup(opts *bind.TransactOpts, indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "createGroup", indexes, _level)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_Role *RoleSession) CreateGroup(indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _Role.Contract.CreateGroup(&_Role.TransactOpts, indexes, _level)
}

// CreateGroup is a paid mutator transaction binding the contract method 0xf652391b.
//
// Solidity: function createGroup(uint64[] indexes, uint16 _level) returns(uint64)
func (_Role *RoleTransactorSession) CreateGroup(indexes []uint64, _level uint16) (*types.Transaction, error) {
	return _Role.Contract.CreateGroup(&_Role.TransactOpts, indexes, _level)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tIndex, uint256 money, bytes sign) payable returns()
func (_Role *RoleTransactor) Recharge(opts *bind.TransactOpts, uIndex uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "recharge", uIndex, tIndex, money, sign)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tIndex, uint256 money, bytes sign) payable returns()
func (_Role *RoleSession) Recharge(uIndex uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.Recharge(&_Role.TransactOpts, uIndex, tIndex, money, sign)
}

// Recharge is a paid mutator transaction binding the contract method 0x517985b0.
//
// Solidity: function recharge(uint64 uIndex, uint32 tIndex, uint256 money, bytes sign) payable returns()
func (_Role *RoleTransactorSession) Recharge(uIndex uint64, tIndex uint32, money *big.Int, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.Recharge(&_Role.TransactOpts, uIndex, tIndex, money, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_Role *RoleTransactor) Register(opts *bind.TransactOpts, addr common.Address, sign []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "register", addr, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_Role *RoleSession) Register(addr common.Address, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.Register(&_Role.TransactOpts, addr, sign)
}

// Register is a paid mutator transaction binding the contract method 0x24b8fbf6.
//
// Solidity: function register(address addr, bytes sign) returns(uint64)
func (_Role *RoleTransactorSession) Register(addr common.Address, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.Register(&_Role.TransactOpts, addr, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 _index, bytes _blsKey, bytes sign) returns()
func (_Role *RoleTransactor) RegisterKeeper(opts *bind.TransactOpts, _index uint64, _blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "registerKeeper", _index, _blsKey, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 _index, bytes _blsKey, bytes sign) returns()
func (_Role *RoleSession) RegisterKeeper(_index uint64, _blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.RegisterKeeper(&_Role.TransactOpts, _index, _blsKey, sign)
}

// RegisterKeeper is a paid mutator transaction binding the contract method 0x10e35bbe.
//
// Solidity: function registerKeeper(uint64 _index, bytes _blsKey, bytes sign) returns()
func (_Role *RoleTransactorSession) RegisterKeeper(_index uint64, _blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.RegisterKeeper(&_Role.TransactOpts, _index, _blsKey, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 _index, bytes sign) returns()
func (_Role *RoleTransactor) RegisterProvider(opts *bind.TransactOpts, _index uint64, sign []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "registerProvider", _index, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 _index, bytes sign) returns()
func (_Role *RoleSession) RegisterProvider(_index uint64, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.RegisterProvider(&_Role.TransactOpts, _index, sign)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0xd57e8a4e.
//
// Solidity: function registerProvider(uint64 _index, bytes sign) returns()
func (_Role *RoleTransactorSession) RegisterProvider(_index uint64, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.RegisterProvider(&_Role.TransactOpts, _index, sign)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tAddr) returns()
func (_Role *RoleTransactor) RegisterToken(opts *bind.TransactOpts, tAddr common.Address) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "registerToken", tAddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tAddr) returns()
func (_Role *RoleSession) RegisterToken(tAddr common.Address) (*types.Transaction, error) {
	return _Role.Contract.RegisterToken(&_Role.TransactOpts, tAddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address tAddr) returns()
func (_Role *RoleTransactorSession) RegisterToken(tAddr common.Address) (*types.Transaction, error) {
	return _Role.Contract.RegisterToken(&_Role.TransactOpts, tAddr)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 _index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_Role *RoleTransactor) RegisterUser(opts *bind.TransactOpts, _index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "registerUser", _index, _gIndex, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 _index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_Role *RoleSession) RegisterUser(_index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.RegisterUser(&_Role.TransactOpts, _index, _gIndex, blsKey, sign)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x488cee1c.
//
// Solidity: function registerUser(uint64 _index, uint64 _gIndex, bytes blsKey, bytes sign) returns()
func (_Role *RoleTransactorSession) RegisterUser(_index uint64, _gIndex uint64, blsKey []byte, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.RegisterUser(&_Role.TransactOpts, _index, _gIndex, blsKey, sign)
}

// SetGF is a paid mutator transaction binding the contract method 0xa6773009.
//
// Solidity: function setGF(uint64 _gIndex, address _fsAddr) returns()
func (_Role *RoleTransactor) SetGF(opts *bind.TransactOpts, _gIndex uint64, _fsAddr common.Address) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "setGF", _gIndex, _fsAddr)
}

// SetGF is a paid mutator transaction binding the contract method 0xa6773009.
//
// Solidity: function setGF(uint64 _gIndex, address _fsAddr) returns()
func (_Role *RoleSession) SetGF(_gIndex uint64, _fsAddr common.Address) (*types.Transaction, error) {
	return _Role.Contract.SetGF(&_Role.TransactOpts, _gIndex, _fsAddr)
}

// SetGF is a paid mutator transaction binding the contract method 0xa6773009.
//
// Solidity: function setGF(uint64 _gIndex, address _fsAddr) returns()
func (_Role *RoleTransactorSession) SetGF(_gIndex uint64, _fsAddr common.Address) (*types.Transaction, error) {
	return _Role.Contract.SetGF(&_Role.TransactOpts, _gIndex, _fsAddr)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_Role *RoleTransactor) SetGInfo(opts *bind.TransactOpts, ps SGParams) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "setGInfo", ps)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_Role *RoleSession) SetGInfo(ps SGParams) (*types.Transaction, error) {
	return _Role.Contract.SetGInfo(&_Role.TransactOpts, ps)
}

// SetGInfo is a paid mutator transaction binding the contract method 0x121ed07f.
//
// Solidity: function setGInfo((uint64,bool,uint256,uint256) ps) returns()
func (_Role *RoleTransactorSession) SetGInfo(ps SGParams) (*types.Transaction, error) {
	return _Role.Contract.SetGInfo(&_Role.TransactOpts, ps)
}

// SetPI is a paid mutator transaction binding the contract method 0xeba091a6.
//
// Solidity: function setPI(address _p, address i, address rfs) returns()
func (_Role *RoleTransactor) SetPI(opts *bind.TransactOpts, _p common.Address, i common.Address, rfs common.Address) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "setPI", _p, i, rfs)
}

// SetPI is a paid mutator transaction binding the contract method 0xeba091a6.
//
// Solidity: function setPI(address _p, address i, address rfs) returns()
func (_Role *RoleSession) SetPI(_p common.Address, i common.Address, rfs common.Address) (*types.Transaction, error) {
	return _Role.Contract.SetPI(&_Role.TransactOpts, _p, i, rfs)
}

// SetPI is a paid mutator transaction binding the contract method 0xeba091a6.
//
// Solidity: function setPI(address _p, address i, address rfs) returns()
func (_Role *RoleTransactorSession) SetPI(_p common.Address, i common.Address, rfs common.Address) (*types.Transaction, error) {
	return _Role.Contract.SetPI(&_Role.TransactOpts, _p, i, rfs)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_Role *RoleTransactor) SetPledgeMoney(opts *bind.TransactOpts, kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "setPledgeMoney", kPledge, pPledge)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_Role *RoleSession) SetPledgeMoney(kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _Role.Contract.SetPledgeMoney(&_Role.TransactOpts, kPledge, pPledge)
}

// SetPledgeMoney is a paid mutator transaction binding the contract method 0x97948fda.
//
// Solidity: function setPledgeMoney(uint256 kPledge, uint256 pPledge) returns()
func (_Role *RoleTransactorSession) SetPledgeMoney(kPledge *big.Int, pPledge *big.Int) (*types.Transaction, error) {
	return _Role.Contract.SetPledgeMoney(&_Role.TransactOpts, kPledge, pPledge)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_Role *RoleTransactor) WithdrawFromFs(opts *bind.TransactOpts, index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Role.contract.Transact(opts, "withdrawFromFs", index, tIndex, amount, sign)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_Role *RoleSession) WithdrawFromFs(index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.WithdrawFromFs(&_Role.TransactOpts, index, tIndex, amount, sign)
}

// WithdrawFromFs is a paid mutator transaction binding the contract method 0xd30d0ce5.
//
// Solidity: function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes sign) returns()
func (_Role *RoleTransactorSession) WithdrawFromFs(index uint64, tIndex uint32, amount *big.Int, sign []byte) (*types.Transaction, error) {
	return _Role.Contract.WithdrawFromFs(&_Role.TransactOpts, index, tIndex, amount, sign)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Role *RoleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Role.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Role *RoleSession) Receive() (*types.Transaction, error) {
	return _Role.Contract.Receive(&_Role.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Role *RoleTransactorSession) Receive() (*types.Transaction, error) {
	return _Role.Contract.Receive(&_Role.TransactOpts)
}

// RoleAlterOwnerIterator is returned from FilterAlterOwner and is used to iterate over the raw logs and unpacked data for AlterOwner events raised by the Role contract.
type RoleAlterOwnerIterator struct {
	Event *RoleAlterOwner // Event containing the contract specifics and raw log

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
func (it *RoleAlterOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleAlterOwner)
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
		it.Event = new(RoleAlterOwner)
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
func (it *RoleAlterOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleAlterOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleAlterOwner represents a AlterOwner event raised by the Role contract.
type RoleAlterOwner struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlterOwner is a free log retrieval operation binding the contract event 0x8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90.
//
// Solidity: event AlterOwner(address from, address to)
func (_Role *RoleFilterer) FilterAlterOwner(opts *bind.FilterOpts) (*RoleAlterOwnerIterator, error) {

	logs, sub, err := _Role.contract.FilterLogs(opts, "AlterOwner")
	if err != nil {
		return nil, err
	}
	return &RoleAlterOwnerIterator{contract: _Role.contract, event: "AlterOwner", logs: logs, sub: sub}, nil
}

// WatchAlterOwner is a free log subscription operation binding the contract event 0x8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90.
//
// Solidity: event AlterOwner(address from, address to)
func (_Role *RoleFilterer) WatchAlterOwner(opts *bind.WatchOpts, sink chan<- *RoleAlterOwner) (event.Subscription, error) {

	logs, sub, err := _Role.contract.WatchLogs(opts, "AlterOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleAlterOwner)
				if err := _Role.contract.UnpackLog(event, "AlterOwner", log); err != nil {
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

// ParseAlterOwner is a log parse operation binding the contract event 0x8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90.
//
// Solidity: event AlterOwner(address from, address to)
func (_Role *RoleFilterer) ParseAlterOwner(log types.Log) (*RoleAlterOwner, error) {
	event := new(RoleAlterOwner)
	if err := _Role.contract.UnpackLog(event, "AlterOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleCreateGroupIterator is returned from FilterCreateGroup and is used to iterate over the raw logs and unpacked data for CreateGroup events raised by the Role contract.
type RoleCreateGroupIterator struct {
	Event *RoleCreateGroup // Event containing the contract specifics and raw log

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
func (it *RoleCreateGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleCreateGroup)
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
		it.Event = new(RoleCreateGroup)
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
func (it *RoleCreateGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleCreateGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleCreateGroup represents a CreateGroup event raised by the Role contract.
type RoleCreateGroup struct {
	GIndex uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreateGroup is a free log retrieval operation binding the contract event 0xc79ca32352cc5529f3c78b5cb44574fbc979555a04f5b6425ed2595417da2e64.
//
// Solidity: event CreateGroup(uint64 gIndex)
func (_Role *RoleFilterer) FilterCreateGroup(opts *bind.FilterOpts) (*RoleCreateGroupIterator, error) {

	logs, sub, err := _Role.contract.FilterLogs(opts, "CreateGroup")
	if err != nil {
		return nil, err
	}
	return &RoleCreateGroupIterator{contract: _Role.contract, event: "CreateGroup", logs: logs, sub: sub}, nil
}

// WatchCreateGroup is a free log subscription operation binding the contract event 0xc79ca32352cc5529f3c78b5cb44574fbc979555a04f5b6425ed2595417da2e64.
//
// Solidity: event CreateGroup(uint64 gIndex)
func (_Role *RoleFilterer) WatchCreateGroup(opts *bind.WatchOpts, sink chan<- *RoleCreateGroup) (event.Subscription, error) {

	logs, sub, err := _Role.contract.WatchLogs(opts, "CreateGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleCreateGroup)
				if err := _Role.contract.UnpackLog(event, "CreateGroup", log); err != nil {
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

// ParseCreateGroup is a log parse operation binding the contract event 0xc79ca32352cc5529f3c78b5cb44574fbc979555a04f5b6425ed2595417da2e64.
//
// Solidity: event CreateGroup(uint64 gIndex)
func (_Role *RoleFilterer) ParseCreateGroup(log types.Log) (*RoleCreateGroup, error) {
	event := new(RoleCreateGroup)
	if err := _Role.contract.UnpackLog(event, "CreateGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleRKeeperIterator is returned from FilterRKeeper and is used to iterate over the raw logs and unpacked data for RKeeper events raised by the Role contract.
type RoleRKeeperIterator struct {
	Event *RoleRKeeper // Event containing the contract specifics and raw log

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
func (it *RoleRKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleRKeeper)
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
		it.Event = new(RoleRKeeper)
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
func (it *RoleRKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleRKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleRKeeper represents a RKeeper event raised by the Role contract.
type RoleRKeeper struct {
	Index uint64
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRKeeper is a free log retrieval operation binding the contract event 0xc50f17198ae53c50e1ad2f06d8348c7d6b31952e4bc9f52b15bb075e6f1bed0b.
//
// Solidity: event RKeeper(uint64 index, address addr)
func (_Role *RoleFilterer) FilterRKeeper(opts *bind.FilterOpts) (*RoleRKeeperIterator, error) {

	logs, sub, err := _Role.contract.FilterLogs(opts, "RKeeper")
	if err != nil {
		return nil, err
	}
	return &RoleRKeeperIterator{contract: _Role.contract, event: "RKeeper", logs: logs, sub: sub}, nil
}

// WatchRKeeper is a free log subscription operation binding the contract event 0xc50f17198ae53c50e1ad2f06d8348c7d6b31952e4bc9f52b15bb075e6f1bed0b.
//
// Solidity: event RKeeper(uint64 index, address addr)
func (_Role *RoleFilterer) WatchRKeeper(opts *bind.WatchOpts, sink chan<- *RoleRKeeper) (event.Subscription, error) {

	logs, sub, err := _Role.contract.WatchLogs(opts, "RKeeper")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleRKeeper)
				if err := _Role.contract.UnpackLog(event, "RKeeper", log); err != nil {
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

// ParseRKeeper is a log parse operation binding the contract event 0xc50f17198ae53c50e1ad2f06d8348c7d6b31952e4bc9f52b15bb075e6f1bed0b.
//
// Solidity: event RKeeper(uint64 index, address addr)
func (_Role *RoleFilterer) ParseRKeeper(log types.Log) (*RoleRKeeper, error) {
	event := new(RoleRKeeper)
	if err := _Role.contract.UnpackLog(event, "RKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleRProviderIterator is returned from FilterRProvider and is used to iterate over the raw logs and unpacked data for RProvider events raised by the Role contract.
type RoleRProviderIterator struct {
	Event *RoleRProvider // Event containing the contract specifics and raw log

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
func (it *RoleRProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleRProvider)
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
		it.Event = new(RoleRProvider)
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
func (it *RoleRProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleRProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleRProvider represents a RProvider event raised by the Role contract.
type RoleRProvider struct {
	Index uint64
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRProvider is a free log retrieval operation binding the contract event 0xf06105db8b89019d932bb3ec85a22bbed723c3616043e02ca9857f3ba37005a5.
//
// Solidity: event RProvider(uint64 index, address addr)
func (_Role *RoleFilterer) FilterRProvider(opts *bind.FilterOpts) (*RoleRProviderIterator, error) {

	logs, sub, err := _Role.contract.FilterLogs(opts, "RProvider")
	if err != nil {
		return nil, err
	}
	return &RoleRProviderIterator{contract: _Role.contract, event: "RProvider", logs: logs, sub: sub}, nil
}

// WatchRProvider is a free log subscription operation binding the contract event 0xf06105db8b89019d932bb3ec85a22bbed723c3616043e02ca9857f3ba37005a5.
//
// Solidity: event RProvider(uint64 index, address addr)
func (_Role *RoleFilterer) WatchRProvider(opts *bind.WatchOpts, sink chan<- *RoleRProvider) (event.Subscription, error) {

	logs, sub, err := _Role.contract.WatchLogs(opts, "RProvider")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleRProvider)
				if err := _Role.contract.UnpackLog(event, "RProvider", log); err != nil {
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

// ParseRProvider is a log parse operation binding the contract event 0xf06105db8b89019d932bb3ec85a22bbed723c3616043e02ca9857f3ba37005a5.
//
// Solidity: event RProvider(uint64 index, address addr)
func (_Role *RoleFilterer) ParseRProvider(log types.Log) (*RoleRProvider, error) {
	event := new(RoleRProvider)
	if err := _Role.contract.UnpackLog(event, "RProvider", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoleRUserIterator is returned from FilterRUser and is used to iterate over the raw logs and unpacked data for RUser events raised by the Role contract.
type RoleRUserIterator struct {
	Event *RoleRUser // Event containing the contract specifics and raw log

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
func (it *RoleRUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoleRUser)
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
		it.Event = new(RoleRUser)
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
func (it *RoleRUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoleRUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoleRUser represents a RUser event raised by the Role contract.
type RoleRUser struct {
	Index uint64
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRUser is a free log retrieval operation binding the contract event 0x7defc6162296f3490e44c1787f4ae9852a8d6a8e67ba0a69c57bd7be5f8a0b1a.
//
// Solidity: event RUser(uint64 index, address addr)
func (_Role *RoleFilterer) FilterRUser(opts *bind.FilterOpts) (*RoleRUserIterator, error) {

	logs, sub, err := _Role.contract.FilterLogs(opts, "RUser")
	if err != nil {
		return nil, err
	}
	return &RoleRUserIterator{contract: _Role.contract, event: "RUser", logs: logs, sub: sub}, nil
}

// WatchRUser is a free log subscription operation binding the contract event 0x7defc6162296f3490e44c1787f4ae9852a8d6a8e67ba0a69c57bd7be5f8a0b1a.
//
// Solidity: event RUser(uint64 index, address addr)
func (_Role *RoleFilterer) WatchRUser(opts *bind.WatchOpts, sink chan<- *RoleRUser) (event.Subscription, error) {

	logs, sub, err := _Role.contract.WatchLogs(opts, "RUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoleRUser)
				if err := _Role.contract.UnpackLog(event, "RUser", log); err != nil {
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

// ParseRUser is a log parse operation binding the contract event 0x7defc6162296f3490e44c1787f4ae9852a8d6a8e67ba0a69c57bd7be5f8a0b1a.
//
// Solidity: event RUser(uint64 index, address addr)
func (_Role *RoleFilterer) ParseRUser(log types.Log) (*RoleRUser, error) {
	event := new(RoleRUser)
	if err := _Role.contract.UnpackLog(event, "RUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
