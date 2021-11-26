// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"renounceRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setUpRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"d5391393": "MINTER_ROLE()",
		"e63ab1e9": "PAUSER_ROLE()",
		"6805b84b": "getPaused()",
		"9e97b8f6": "hasRole(uint8,address)",
		"8456cb59": "pause()",
		"057c9cb8": "renounceRole(uint8)",
		"4cbb87d3": "revokeRole(uint8,address)",
		"6d6aa720": "setUpRole(uint8,address)",
		"3f4ba83a": "unpause()",
	},
	Bin: "0x608060405234801561001057600080fd5b503360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602090815260408083208054600160ff1991821681179092557fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d845282852080548216831790557fabbb5caa7dda850e60932de0934eb1f9d0f59695050f761dc64e443e5030a56990935290832080549092161790556103d69081906100be90396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638456cb59116100665780638456cb59146101265780639e97b8f61461012e578063a217fddf14610141578063d53913931461015b578063e63ab1e91461016357600080fd5b8063057c9cb8146100a35780633f4ba83a146100ed5780634cbb87d3146100f55780636805b84b146101085780636d6aa72014610113575b600080fd5b6100d86100b136600461033b565b60ff166000908152602081815260408083203384529091529020805460ff19169055600190565b60405190151581526020015b60405180910390f35b6100d861016b565b6100d861010336600461035d565b6101bf565b60015460ff166100d8565b6100d861012136600461035d565b610233565b6100d86102ab565b6100d861013c36600461035d565b6102fa565b610149600081565b60405160ff90911681526020016100e4565b610149600181565b610149600281565b60006101786002336102fa565b6101b15760405162461bcd60e51b8152602060048201526005602482015264020434e55560dc1b60448201526064015b60405180910390fd5b506001805460ff1916815590565b60006101cc6000336102fa565b6101fe5760405162461bcd60e51b81526020600482015260036024820152622720a960e91b60448201526064016101a8565b5060ff82166000908152602081815260408083206001600160a01b03851684529091529020805460ff19169055600192915050565b60006102406000336102fa565b6102725760405162461bcd60e51b81526020600482015260036024820152622720a960e91b60448201526064016101a8565b5060ff82166000908152602081815260408083206001600160a01b03851684529091529020805460ff1916600190811790915592915050565b60006102b86002336102fa565b6102ea5760405162461bcd60e51b81526020600482015260036024820152620434e560ec1b60448201526064016101a8565b506001805460ff19168117815590565b60ff9182166000908152602081815260408083206001600160a01b0394909416835292905220541690565b803560ff8116811461033657600080fd5b919050565b60006020828403121561034d57600080fd5b61035682610325565b9392505050565b6000806040838503121561037057600080fd5b61037983610325565b915060208301356001600160a01b038116811461039557600080fd5b80915050925092905056fea2646970667358221220f8437ea7131ec2cd6a18076346e4707e21c616c48e7fdd2107e3c9a4e53fd16264736f6c634300080a0033",
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// Deprecated: Use AccessControlMetaData.Sigs instead.
// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = AccessControlMetaData.Sigs

// AccessControlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessControlMetaData.Bin instead.
var AccessControlBin = AccessControlMetaData.Bin

// DeployAccessControl deploys a new Ethereum contract, binding an instance of AccessControl to it.
func DeployAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControl, error) {
	parsed, err := AccessControlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(uint8)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(uint8)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() (uint8, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(uint8)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() (uint8, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(uint8)
func (_AccessControl *AccessControlCaller) MINTERROLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(uint8)
func (_AccessControl *AccessControlSession) MINTERROLE() (uint8, error) {
	return _AccessControl.Contract.MINTERROLE(&_AccessControl.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(uint8)
func (_AccessControl *AccessControlCallerSession) MINTERROLE() (uint8, error) {
	return _AccessControl.Contract.MINTERROLE(&_AccessControl.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(uint8)
func (_AccessControl *AccessControlCaller) PAUSERROLE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(uint8)
func (_AccessControl *AccessControlSession) PAUSERROLE() (uint8, error) {
	return _AccessControl.Contract.PAUSERROLE(&_AccessControl.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(uint8)
func (_AccessControl *AccessControlCallerSession) PAUSERROLE() (uint8, error) {
	return _AccessControl.Contract.PAUSERROLE(&_AccessControl.CallOpts)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_AccessControl *AccessControlCaller) GetPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_AccessControl *AccessControlSession) GetPaused() (bool, error) {
	return _AccessControl.Contract.GetPaused(&_AccessControl.CallOpts)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_AccessControl *AccessControlCallerSession) GetPaused() (bool, error) {
	return _AccessControl.Contract.GetPaused(&_AccessControl.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role uint8, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role uint8, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role uint8, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_AccessControl *AccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_AccessControl *AccessControlSession) Pause() (*types.Transaction, error) {
	return _AccessControl.Contract.Pause(&_AccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_AccessControl *AccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _AccessControl.Contract.Pause(&_AccessControl.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x057c9cb8.
//
// Solidity: function renounceRole(uint8 role) returns(bool)
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role uint8) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x057c9cb8.
//
// Solidity: function renounceRole(uint8 role) returns(bool)
func (_AccessControl *AccessControlSession) RenounceRole(role uint8) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x057c9cb8.
//
// Solidity: function renounceRole(uint8 role) returns(bool)
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role uint8) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x4cbb87d3.
//
// Solidity: function revokeRole(uint8 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role uint8, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x4cbb87d3.
//
// Solidity: function revokeRole(uint8 role, address account) returns(bool)
func (_AccessControl *AccessControlSession) RevokeRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x4cbb87d3.
//
// Solidity: function revokeRole(uint8 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x6d6aa720.
//
// Solidity: function setUpRole(uint8 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactor) SetUpRole(opts *bind.TransactOpts, role uint8, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "setUpRole", role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x6d6aa720.
//
// Solidity: function setUpRole(uint8 role, address account) returns(bool)
func (_AccessControl *AccessControlSession) SetUpRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetUpRole(&_AccessControl.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x6d6aa720.
//
// Solidity: function setUpRole(uint8 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactorSession) SetUpRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetUpRole(&_AccessControl.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_AccessControl *AccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_AccessControl *AccessControlSession) Unpause() (*types.Transaction, error) {
	return _AccessControl.Contract.Unpause(&_AccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_AccessControl *AccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _AccessControl.Contract.Unpause(&_AccessControl.TransactOpts)
}

// IAccessControlMetaData contains all meta data concerning the IAccessControl contract.
var IAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"renounceRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setUpRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6805b84b": "getPaused()",
		"9e97b8f6": "hasRole(uint8,address)",
		"8456cb59": "pause()",
		"057c9cb8": "renounceRole(uint8)",
		"4cbb87d3": "revokeRole(uint8,address)",
		"6d6aa720": "setUpRole(uint8,address)",
		"3f4ba83a": "unpause()",
	},
}

// IAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlMetaData.ABI instead.
var IAccessControlABI = IAccessControlMetaData.ABI

// Deprecated: Use IAccessControlMetaData.Sigs instead.
// IAccessControlFuncSigs maps the 4-byte function signature to its string representation.
var IAccessControlFuncSigs = IAccessControlMetaData.Sigs

// IAccessControl is an auto generated Go binding around an Ethereum contract.
type IAccessControl struct {
	IAccessControlCaller     // Read-only binding to the contract
	IAccessControlTransactor // Write-only binding to the contract
	IAccessControlFilterer   // Log filterer for contract events
}

// IAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlSession struct {
	Contract     *IAccessControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlCallerSession struct {
	Contract *IAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlTransactorSession struct {
	Contract     *IAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlRaw struct {
	Contract *IAccessControl // Generic contract binding to access the raw methods on
}

// IAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlCallerRaw struct {
	Contract *IAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlTransactorRaw struct {
	Contract *IAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControl creates a new instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControl(address common.Address, backend bind.ContractBackend) (*IAccessControl, error) {
	contract, err := bindIAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControl{IAccessControlCaller: IAccessControlCaller{contract: contract}, IAccessControlTransactor: IAccessControlTransactor{contract: contract}, IAccessControlFilterer: IAccessControlFilterer{contract: contract}}, nil
}

// NewIAccessControlCaller creates a new read-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlCaller, error) {
	contract, err := bindIAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlCaller{contract: contract}, nil
}

// NewIAccessControlTransactor creates a new write-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlTransactor, error) {
	contract, err := bindIAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlTransactor{contract: contract}, nil
}

// NewIAccessControlFilterer creates a new log filterer instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlFilterer, error) {
	contract, err := bindIAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlFilterer{contract: contract}, nil
}

// bindIAccessControl binds a generic wrapper to an already deployed contract.
func bindIAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.IAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transact(opts, method, params...)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_IAccessControl *IAccessControlCaller) GetPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "getPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_IAccessControl *IAccessControlSession) GetPaused() (bool, error) {
	return _IAccessControl.Contract.GetPaused(&_IAccessControl.CallOpts)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) GetPaused() (bool, error) {
	return _IAccessControl.Contract.GetPaused(&_IAccessControl.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCaller) HasRole(opts *bind.CallOpts, role uint8, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlSession) HasRole(role uint8, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x9e97b8f6.
//
// Solidity: function hasRole(uint8 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) HasRole(role uint8, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IAccessControl *IAccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IAccessControl *IAccessControlSession) Pause() (*types.Transaction, error) {
	return _IAccessControl.Contract.Pause(&_IAccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_IAccessControl *IAccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _IAccessControl.Contract.Pause(&_IAccessControl.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x057c9cb8.
//
// Solidity: function renounceRole(uint8 role) returns(bool)
func (_IAccessControl *IAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role uint8) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x057c9cb8.
//
// Solidity: function renounceRole(uint8 role) returns(bool)
func (_IAccessControl *IAccessControlSession) RenounceRole(role uint8) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x057c9cb8.
//
// Solidity: function renounceRole(uint8 role) returns(bool)
func (_IAccessControl *IAccessControlTransactorSession) RenounceRole(role uint8) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x4cbb87d3.
//
// Solidity: function revokeRole(uint8 role, address account) returns(bool)
func (_IAccessControl *IAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role uint8, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x4cbb87d3.
//
// Solidity: function revokeRole(uint8 role, address account) returns(bool)
func (_IAccessControl *IAccessControlSession) RevokeRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x4cbb87d3.
//
// Solidity: function revokeRole(uint8 role, address account) returns(bool)
func (_IAccessControl *IAccessControlTransactorSession) RevokeRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x6d6aa720.
//
// Solidity: function setUpRole(uint8 role, address account) returns(bool)
func (_IAccessControl *IAccessControlTransactor) SetUpRole(opts *bind.TransactOpts, role uint8, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "setUpRole", role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x6d6aa720.
//
// Solidity: function setUpRole(uint8 role, address account) returns(bool)
func (_IAccessControl *IAccessControlSession) SetUpRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.SetUpRole(&_IAccessControl.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x6d6aa720.
//
// Solidity: function setUpRole(uint8 role, address account) returns(bool)
func (_IAccessControl *IAccessControlTransactorSession) SetUpRole(role uint8, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.SetUpRole(&_IAccessControl.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IAccessControl *IAccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IAccessControl *IAccessControlSession) Unpause() (*types.Transaction, error) {
	return _IAccessControl.Contract.Unpause(&_IAccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_IAccessControl *IAccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _IAccessControl.Contract.Unpause(&_IAccessControl.TransactOpts)
}
