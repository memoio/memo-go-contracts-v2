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

// IssuanceMetaData contains all meta data concerning the Issuance contract.
var IssuanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rfs\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"name\":\"setSP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spaceTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"13df3828": "issu(uint64,uint64,uint64,uint256)",
		"586fc5b5": "lastMint()",
		"988bda70": "mintLevel()",
		"a035b1fe": "price()",
		"6dd81f17": "setSP(uint256,uint256)",
		"11e65fc0": "setTP(uint256,uint256)",
		"949d225d": "size()",
		"43a2755c": "spaceTime()",
		"e7b0f666": "totalPaid()",
		"e154adf5": "totalPay()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610a5a380380610a5a83398101604081905261002f916103e8565b60004290508060028190555060016040518060600160405280606461ffff16815260200160016001600160401b0316815260200160016001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280607861ffff168152602001656400000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280609661ffff16815260200166040000000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b0316021790555050506001604051806060016040528060c861ffff16815260200166280000000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505081600860006101000a8154816001600160a01b0302191690836001600160a01b031602179055505050610418565b6000602082840312156103fa57600080fd5b81516001600160a01b038116811461041157600080fd5b9392505050565b610633806104276000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063949d225d11610066578063949d225d14610102578063988bda701461010b578063a035b1fe14610114578063e154adf51461011d578063e7b0f6661461012657600080fd5b806311e65fc0146100a357806313df3828146100b857806343a2755c146100dd578063586fc5b5146100e65780636dd81f17146100ef575b600080fd5b6100b66100b1366004610478565b61012f565b005b6100cb6100c63660046104b7565b610196565b60405190815260200160405180910390f35b6100cb60055481565b6100cb60025481565b6100b66100fd366004610478565b610423565b6100cb60045481565b6100cb60005481565b6100cb60035481565b6100cb60065481565b6100cb60075481565b6008546001600160a01b031633146101625760405162461bcd60e51b815260040161015990610502565b60405180910390fd5b81600760008282546101749190610533565b92505081905550806007600082825461018d919061054b565b90915550505050565b6008546000906001600160a01b031633146101c35760405162461bcd60e51b815260040161015990610502565b60025442906000906101d5908361054b565b90506000816003546101e79190610562565b905060006007546006546101fb919061054b565b60035490915015610232576000600354826102169190610581565b90506102228183610581565b915061022e8483610562565b9150505b60008054610241906001610533565b90505b60015481101561030157600060018281548110610263576102636105a3565b6000918252602090912001546004546201000090910467ffffffffffffffff16915081101561029157506004545b6000600183815481106102a6576102a66105a3565b600091825260209091200154600554600160501b90910467ffffffffffffffff16915081906102d6908490610581565b106102e55760008390556102ec565b50506102ef565b50505b806102f9816105b9565b915050610244565b50600061030e8a8a6105d4565b67ffffffffffffffff166004546103259190610562565b9050806005546103359190610533565b6005819055508767ffffffffffffffff16600460008282546103579190610533565b9250508190555086600360008282546103709190610533565b90915550600090506103828b8b6105d4565b6103969067ffffffffffffffff1689610562565b905080600660008282546103aa9190610533565b925050819055506001600054815481106103c6576103c66105a3565b6000918252602090912001546103e09061ffff1684610562565b92506103ed606484610581565b925083600760008282546104019190610533565b90915550505067ffffffffffffffff9094166002559350505050949350505050565b6008546001600160a01b0316331461044d5760405162461bcd60e51b815260040161015990610502565b816003600082825461045f919061054b565b92505081905550806004600082825461018d919061054b565b6000806040838503121561048b57600080fd5b50508035926020909101359150565b803567ffffffffffffffff811681146104b257600080fd5b919050565b600080600080608085870312156104cd57600080fd5b6104d68561049a565b93506104e46020860161049a565b92506104f26040860161049a565b9396929550929360600135925050565b6020808252600190820152602760f91b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600082198211156105465761054661051d565b500190565b60008282101561055d5761055d61051d565b500390565b600081600019048311821515161561057c5761057c61051d565b500290565b60008261059e57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052603260045260246000fd5b60006000198214156105cd576105cd61051d565b5060010190565b600067ffffffffffffffff838116908316818110156105f5576105f561051d565b03939250505056fea2646970667358221220aa6690d33e40887abde265ae66b025a2a9a6a300633816e719a0b09e13a7ce1d64736f6c634300080a0033",
}

// IssuanceABI is the input ABI used to generate the binding from.
// Deprecated: Use IssuanceMetaData.ABI instead.
var IssuanceABI = IssuanceMetaData.ABI

// Deprecated: Use IssuanceMetaData.Sigs instead.
// IssuanceFuncSigs maps the 4-byte function signature to its string representation.
var IssuanceFuncSigs = IssuanceMetaData.Sigs

// IssuanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IssuanceMetaData.Bin instead.
var IssuanceBin = IssuanceMetaData.Bin

// DeployIssuance deploys a new Ethereum contract, binding an instance of Issuance to it.
func DeployIssuance(auth *bind.TransactOpts, backend bind.ContractBackend, rfs common.Address) (common.Address, *types.Transaction, *Issuance, error) {
	parsed, err := IssuanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IssuanceBin), backend, rfs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Issuance{IssuanceCaller: IssuanceCaller{contract: contract}, IssuanceTransactor: IssuanceTransactor{contract: contract}, IssuanceFilterer: IssuanceFilterer{contract: contract}}, nil
}

// Issuance is an auto generated Go binding around an Ethereum contract.
type Issuance struct {
	IssuanceCaller     // Read-only binding to the contract
	IssuanceTransactor // Write-only binding to the contract
	IssuanceFilterer   // Log filterer for contract events
}

// IssuanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IssuanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IssuanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IssuanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IssuanceSession struct {
	Contract     *Issuance         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IssuanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IssuanceCallerSession struct {
	Contract *IssuanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IssuanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IssuanceTransactorSession struct {
	Contract     *IssuanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IssuanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IssuanceRaw struct {
	Contract *Issuance // Generic contract binding to access the raw methods on
}

// IssuanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IssuanceCallerRaw struct {
	Contract *IssuanceCaller // Generic read-only contract binding to access the raw methods on
}

// IssuanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IssuanceTransactorRaw struct {
	Contract *IssuanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIssuance creates a new instance of Issuance, bound to a specific deployed contract.
func NewIssuance(address common.Address, backend bind.ContractBackend) (*Issuance, error) {
	contract, err := bindIssuance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Issuance{IssuanceCaller: IssuanceCaller{contract: contract}, IssuanceTransactor: IssuanceTransactor{contract: contract}, IssuanceFilterer: IssuanceFilterer{contract: contract}}, nil
}

// NewIssuanceCaller creates a new read-only instance of Issuance, bound to a specific deployed contract.
func NewIssuanceCaller(address common.Address, caller bind.ContractCaller) (*IssuanceCaller, error) {
	contract, err := bindIssuance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IssuanceCaller{contract: contract}, nil
}

// NewIssuanceTransactor creates a new write-only instance of Issuance, bound to a specific deployed contract.
func NewIssuanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IssuanceTransactor, error) {
	contract, err := bindIssuance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IssuanceTransactor{contract: contract}, nil
}

// NewIssuanceFilterer creates a new log filterer instance of Issuance, bound to a specific deployed contract.
func NewIssuanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IssuanceFilterer, error) {
	contract, err := bindIssuance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IssuanceFilterer{contract: contract}, nil
}

// bindIssuance binds a generic wrapper to an already deployed contract.
func bindIssuance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IssuanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuance *IssuanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuance.Contract.IssuanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuance *IssuanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuance.Contract.IssuanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuance *IssuanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuance.Contract.IssuanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuance *IssuanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuance *IssuanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuance *IssuanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuance.Contract.contract.Transact(opts, method, params...)
}

// LastMint is a free data retrieval call binding the contract method 0x586fc5b5.
//
// Solidity: function lastMint() view returns(uint256)
func (_Issuance *IssuanceCaller) LastMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "lastMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastMint is a free data retrieval call binding the contract method 0x586fc5b5.
//
// Solidity: function lastMint() view returns(uint256)
func (_Issuance *IssuanceSession) LastMint() (*big.Int, error) {
	return _Issuance.Contract.LastMint(&_Issuance.CallOpts)
}

// LastMint is a free data retrieval call binding the contract method 0x586fc5b5.
//
// Solidity: function lastMint() view returns(uint256)
func (_Issuance *IssuanceCallerSession) LastMint() (*big.Int, error) {
	return _Issuance.Contract.LastMint(&_Issuance.CallOpts)
}

// MintLevel is a free data retrieval call binding the contract method 0x988bda70.
//
// Solidity: function mintLevel() view returns(uint256)
func (_Issuance *IssuanceCaller) MintLevel(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "mintLevel")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintLevel is a free data retrieval call binding the contract method 0x988bda70.
//
// Solidity: function mintLevel() view returns(uint256)
func (_Issuance *IssuanceSession) MintLevel() (*big.Int, error) {
	return _Issuance.Contract.MintLevel(&_Issuance.CallOpts)
}

// MintLevel is a free data retrieval call binding the contract method 0x988bda70.
//
// Solidity: function mintLevel() view returns(uint256)
func (_Issuance *IssuanceCallerSession) MintLevel() (*big.Int, error) {
	return _Issuance.Contract.MintLevel(&_Issuance.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Issuance *IssuanceCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Issuance *IssuanceSession) Price() (*big.Int, error) {
	return _Issuance.Contract.Price(&_Issuance.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Issuance *IssuanceCallerSession) Price() (*big.Int, error) {
	return _Issuance.Contract.Price(&_Issuance.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_Issuance *IssuanceCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_Issuance *IssuanceSession) Size() (*big.Int, error) {
	return _Issuance.Contract.Size(&_Issuance.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_Issuance *IssuanceCallerSession) Size() (*big.Int, error) {
	return _Issuance.Contract.Size(&_Issuance.CallOpts)
}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceCaller) SpaceTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "spaceTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceSession) SpaceTime() (*big.Int, error) {
	return _Issuance.Contract.SpaceTime(&_Issuance.CallOpts)
}

// SpaceTime is a free data retrieval call binding the contract method 0x43a2755c.
//
// Solidity: function spaceTime() view returns(uint256)
func (_Issuance *IssuanceCallerSession) SpaceTime() (*big.Int, error) {
	return _Issuance.Contract.SpaceTime(&_Issuance.CallOpts)
}

// TotalPaid is a free data retrieval call binding the contract method 0xe7b0f666.
//
// Solidity: function totalPaid() view returns(uint256)
func (_Issuance *IssuanceCaller) TotalPaid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "totalPaid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPaid is a free data retrieval call binding the contract method 0xe7b0f666.
//
// Solidity: function totalPaid() view returns(uint256)
func (_Issuance *IssuanceSession) TotalPaid() (*big.Int, error) {
	return _Issuance.Contract.TotalPaid(&_Issuance.CallOpts)
}

// TotalPaid is a free data retrieval call binding the contract method 0xe7b0f666.
//
// Solidity: function totalPaid() view returns(uint256)
func (_Issuance *IssuanceCallerSession) TotalPaid() (*big.Int, error) {
	return _Issuance.Contract.TotalPaid(&_Issuance.CallOpts)
}

// TotalPay is a free data retrieval call binding the contract method 0xe154adf5.
//
// Solidity: function totalPay() view returns(uint256)
func (_Issuance *IssuanceCaller) TotalPay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "totalPay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPay is a free data retrieval call binding the contract method 0xe154adf5.
//
// Solidity: function totalPay() view returns(uint256)
func (_Issuance *IssuanceSession) TotalPay() (*big.Int, error) {
	return _Issuance.Contract.TotalPay(&_Issuance.CallOpts)
}

// TotalPay is a free data retrieval call binding the contract method 0xe154adf5.
//
// Solidity: function totalPay() view returns(uint256)
func (_Issuance *IssuanceCallerSession) TotalPay() (*big.Int, error) {
	return _Issuance.Contract.TotalPay(&_Issuance.CallOpts)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 sPrice) returns(uint256)
func (_Issuance *IssuanceTransactor) Issu(opts *bind.TransactOpts, _start uint64, _end uint64, _size uint64, sPrice *big.Int) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "issu", _start, _end, _size, sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 sPrice) returns(uint256)
func (_Issuance *IssuanceSession) Issu(_start uint64, _end uint64, _size uint64, sPrice *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, _start, _end, _size, sPrice)
}

// Issu is a paid mutator transaction binding the contract method 0x13df3828.
//
// Solidity: function issu(uint64 _start, uint64 _end, uint64 _size, uint256 sPrice) returns(uint256)
func (_Issuance *IssuanceTransactorSession) Issu(_start uint64, _end uint64, _size uint64, sPrice *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, _start, _end, _size, sPrice)
}

// SetSP is a paid mutator transaction binding the contract method 0x6dd81f17.
//
// Solidity: function setSP(uint256 _price, uint256 _size) returns()
func (_Issuance *IssuanceTransactor) SetSP(opts *bind.TransactOpts, _price *big.Int, _size *big.Int) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "setSP", _price, _size)
}

// SetSP is a paid mutator transaction binding the contract method 0x6dd81f17.
//
// Solidity: function setSP(uint256 _price, uint256 _size) returns()
func (_Issuance *IssuanceSession) SetSP(_price *big.Int, _size *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.SetSP(&_Issuance.TransactOpts, _price, _size)
}

// SetSP is a paid mutator transaction binding the contract method 0x6dd81f17.
//
// Solidity: function setSP(uint256 _price, uint256 _size) returns()
func (_Issuance *IssuanceTransactorSession) SetSP(_price *big.Int, _size *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.SetSP(&_Issuance.TransactOpts, _price, _size)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_Issuance *IssuanceTransactor) SetTP(opts *bind.TransactOpts, _add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "setTP", _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_Issuance *IssuanceSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.SetTP(&_Issuance.TransactOpts, _add, _sub)
}

// SetTP is a paid mutator transaction binding the contract method 0x11e65fc0.
//
// Solidity: function setTP(uint256 _add, uint256 _sub) returns()
func (_Issuance *IssuanceTransactorSession) SetTP(_add *big.Int, _sub *big.Int) (*types.Transaction, error) {
	return _Issuance.Contract.SetTP(&_Issuance.TransactOpts, _add, _sub)
}
