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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setUpRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060016000808060001b815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000807f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000807f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506108e08061019b6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806391d148541161006657806391d148541461015d578063a217fddf1461018d578063d5391393146101ab578063d547741f146101c9578063e63ab1e9146101f95761009e565b80632d958471146100a35780633f4ba83a146100d35780636805b84b146100f15780638456cb591461010f5780638bb9c5bf1461012d575b600080fd5b6100bd60048036038101906100b89190610688565b610217565b6040516100ca91906106e3565b60405180910390f35b6100db6102d7565b6040516100e891906106e3565b60405180910390f35b6100f9610364565b60405161010691906106e3565b60405180910390f35b61011761037b565b60405161012491906106e3565b60405180910390f35b610147600480360381019061014291906106fe565b610407565b60405161015491906106e3565b60405180910390f35b61017760048036038101906101729190610688565b610479565b60405161018491906106e3565b60405180910390f35b6101956104e0565b6040516101a2919061073a565b60405180910390f35b6101b36104e7565b6040516101c0919061073a565b60405180910390f35b6101e360048036038101906101de9190610688565b61050b565b6040516101f091906106e3565b60405180910390f35b6102016105cb565b60405161020e919061073a565b60405180910390f35b60006102266000801b33610479565b610265576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025c906107b2565b60405180910390fd5b600160008085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905092915050565b60006103037f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a33610479565b610342576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103399061081e565b60405180910390fd5b6000600160006101000a81548160ff0219169083151502179055506001905090565b6000600160009054906101000a900460ff16905090565b60006103a77f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a33610479565b6103e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103dd9061088a565b60405180910390fd5b60018060006101000a81548160ff0219169083151502179055506001905090565b60008060008084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b600080600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b600061051a6000801b33610479565b610559576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610550906107b2565b60405180910390fd5b600080600085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905092915050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b600080fd5b6000819050919050565b610607816105f4565b811461061257600080fd5b50565b600081359050610624816105fe565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106558261062a565b9050919050565b6106658161064a565b811461067057600080fd5b50565b6000813590506106828161065c565b92915050565b6000806040838503121561069f5761069e6105ef565b5b60006106ad85828601610615565b92505060206106be85828601610673565b9150509250929050565b60008115159050919050565b6106dd816106c8565b82525050565b60006020820190506106f860008301846106d4565b92915050565b600060208284031215610714576107136105ef565b5b600061072284828501610615565b91505092915050565b610734816105f4565b82525050565b600060208201905061074f600083018461072b565b92915050565b600082825260208201905092915050565b7f4e41520000000000000000000000000000000000000000000000000000000000600082015250565b600061079c600383610755565b91506107a782610766565b602082019050919050565b600060208201905081810360008301526107cb8161078f565b9050919050565b7f20434e5550000000000000000000000000000000000000000000000000000000600082015250565b6000610808600583610755565b9150610813826107d2565b602082019050919050565b60006020820190508181036000830152610837816107fb565b9050919050565b7f434e500000000000000000000000000000000000000000000000000000000000600082015250565b6000610874600383610755565b915061087f8261083e565b602082019050919050565b600060208201905081810360008301526108a381610867565b905091905056fea2646970667358221220322c214655df83bbf83cbac6aad7ae0aadcf2a4cb3558fb4a477fecba89f4ad964736f6c634300080a0033",
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

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
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) MINTERROLE() ([32]byte, error) {
	return _AccessControl.Contract.MINTERROLE(&_AccessControl.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) MINTERROLE() ([32]byte, error) {
	return _AccessControl.Contract.MINTERROLE(&_AccessControl.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) PAUSERROLE() ([32]byte, error) {
	return _AccessControl.Contract.PAUSERROLE(&_AccessControl.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) PAUSERROLE() ([32]byte, error) {
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

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
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

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns(bool)
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns(bool)
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns(bool)
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns(bool)
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x2d958471.
//
// Solidity: function setUpRole(bytes32 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactor) SetUpRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "setUpRole", role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x2d958471.
//
// Solidity: function setUpRole(bytes32 role, address account) returns(bool)
func (_AccessControl *AccessControlSession) SetUpRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.SetUpRole(&_AccessControl.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x2d958471.
//
// Solidity: function setUpRole(bytes32 role, address account) returns(bool)
func (_AccessControl *AccessControlTransactorSession) SetUpRole(role [32]byte, account common.Address) (*types.Transaction, error) {
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
