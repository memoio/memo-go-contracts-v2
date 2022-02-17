// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package issuance

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

// IssuParams is an auto generated low-level Go binding around an user-defined struct.
type IssuParams struct {
	Start  uint64
	End    uint64
	Size   uint64
	SPrice *big.Int
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

// IssuanceMetaData contains all meta data concerning the Issuance contract.
var IssuanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rfs\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIssuParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spaceTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subPMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subSMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c862ea50": "issu((uint64,uint64,uint64,uint256))",
		"586fc5b5": "lastMint()",
		"988bda70": "mintLevel()",
		"a035b1fe": "price()",
		"11e65fc0": "setTP(uint256,uint256)",
		"949d225d": "size()",
		"43a2755c": "spaceTime()",
		"f0b44ab1": "subPMap(uint64)",
		"45fe29f6": "subSMap(uint64)",
		"e7b0f666": "totalPaid()",
		"e154adf5": "totalPay()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610be2380380610be283398101604081905261002f916103e8565b60004290508060028190555060016040518060600160405280606461ffff16815260200160016001600160401b0316815260200160016001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280607861ffff168152602001656400000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280609661ffff16815260200166040000000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b0316021790555050506001604051806060016040528060c861ffff16815260200166280000000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505081600a60006101000a8154816001600160a01b0302191690836001600160a01b031602179055505050610418565b6000602082840312156103fa57600080fd5b81516001600160a01b038116811461041157600080fd5b9392505050565b6107bb806104276000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063988bda7011610071578063988bda7014610110578063a035b1fe14610119578063c862ea5014610122578063e154adf514610135578063e7b0f6661461013e578063f0b44ab11461014757600080fd5b806311e65fc0146100ae57806343a2755c146100c357806345fe29f6146100de578063586fc5b5146100fe578063949d225d14610107575b600080fd5b6100c16100bc3660046105c0565b610167565b005b6100cc60055481565b60405190815260200160405180910390f35b6100cc6100ec3660046105fe565b60096020526000908152604090205481565b6100cc60025481565b6100cc60045481565b6100cc60005481565b6100cc60035481565b6100cc610130366004610620565b6101e5565b6100cc60065481565b6100cc60075481565b6100cc6101553660046105fe565b60086020526000908152604090205481565b600a546001600160a01b031633146101a95760405162461bcd60e51b81526020600482015260016024820152602760f91b604482015260640160405180910390fd5b81600760008282546101bb91906106bc565b909155505060075481116101e15780600760008282546101db91906106d4565b90915550505b5050565b600081606001516008600084602001516001600160401b03166001600160401b03168152602001908152602001600020600082825461022491906106bc565b9250508190555081604001516001600160401b03166009600084602001516001600160401b03166001600160401b03168152602001908152602001600020600082825461027191906106bc565b90915550506002544290620151809061028a90836106d4565b11156102a3576002546102a090620151806106bc565b90505b6000600254826102b391906106d4565b90506000816003546102c591906106eb565b90506102d4620151808461070a565b620151806002546102e5919061070a565b10156103a15760006102fa620151808561070a565b61030790620151806106eb565b6001600160401b038116600090815260086020526040902054909150801561036657600061033583876106d4565b61033f90836106eb565b905061034b81856106d4565b9350816003600082825461035f91906106d4565b9091555050505b6001600160401b038216600090815260096020526040902054801561039d57806004600082825461039791906106d4565b90915550505b5050505b80600760008282546103b391906106bc565b9091555050845160208601516000916103cb9161072c565b6001600160401b031686604001516001600160401b03166103ec91906106eb565b9050806005600082825461040091906106bc565b9091555050855160208701516000916104189161072c565b6001600160401b0316876060015161043091906106eb565b9050806006600082825461044491906106bc565b9250508190555086604001516001600160401b03166004600082825461046a91906106bc565b90915550506060870151600380546000906104869084906106bc565b9091555050826104a25750505060029190915550600092915050565b600080546104b19060016106bc565b90505b60015481101561056f576000600182815481106104d3576104d3610754565b600091825260209091200154600454620100009091046001600160401b0316915081101561050057506004545b60006001838154811061051557610515610754565b600091825260209091200154600554600160501b9091046001600160401b03169150819061054490849061070a565b1061055357600083905561055a565b505061056f565b505080806105679061076a565b9150506104b4565b50600060016000548154811061058757610587610754565b6000918252602090912001546105a19061ffff16856106eb565b90506105ae60648261070a565b60029690965550939695505050505050565b600080604083850312156105d357600080fd5b50508035926020909101359150565b80356001600160401b03811681146105f957600080fd5b919050565b60006020828403121561061057600080fd5b610619826105e2565b9392505050565b60006080828403121561063257600080fd5b604051608081018181106001600160401b038211171561066257634e487b7160e01b600052604160045260246000fd5b60405261066e836105e2565b815261067c602084016105e2565b602082015261068d604084016105e2565b6040820152606083013560608201528091505092915050565b634e487b7160e01b600052601160045260246000fd5b600082198211156106cf576106cf6106a6565b500190565b6000828210156106e6576106e66106a6565b500390565b6000816000190483118215151615610705576107056106a6565b500290565b60008261072757634e487b7160e01b600052601260045260246000fd5b500490565b60006001600160401b038381169083168181101561074c5761074c6106a6565b039392505050565b634e487b7160e01b600052603260045260246000fd5b600060001982141561077e5761077e6106a6565b506001019056fea264697066735822122076b4f0291edfe13f98bda48eaf338a7425026e4e41f98620418f7c69cf7085c464736f6c634300080b0033",
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

// SubPMap is a free data retrieval call binding the contract method 0xf0b44ab1.
//
// Solidity: function subPMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCaller) SubPMap(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "subPMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubPMap is a free data retrieval call binding the contract method 0xf0b44ab1.
//
// Solidity: function subPMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceSession) SubPMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubPMap(&_Issuance.CallOpts, arg0)
}

// SubPMap is a free data retrieval call binding the contract method 0xf0b44ab1.
//
// Solidity: function subPMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCallerSession) SubPMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubPMap(&_Issuance.CallOpts, arg0)
}

// SubSMap is a free data retrieval call binding the contract method 0x45fe29f6.
//
// Solidity: function subSMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCaller) SubSMap(opts *bind.CallOpts, arg0 uint64) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "subSMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubSMap is a free data retrieval call binding the contract method 0x45fe29f6.
//
// Solidity: function subSMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceSession) SubSMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubSMap(&_Issuance.CallOpts, arg0)
}

// SubSMap is a free data retrieval call binding the contract method 0x45fe29f6.
//
// Solidity: function subSMap(uint64 ) view returns(uint256)
func (_Issuance *IssuanceCallerSession) SubSMap(arg0 uint64) (*big.Int, error) {
	return _Issuance.Contract.SubSMap(&_Issuance.CallOpts, arg0)
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

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_Issuance *IssuanceTransactor) Issu(opts *bind.TransactOpts, ps IssuParams) (*types.Transaction, error) {
	return _Issuance.contract.Transact(opts, "issu", ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_Issuance *IssuanceSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, ps)
}

// Issu is a paid mutator transaction binding the contract method 0xc862ea50.
//
// Solidity: function issu((uint64,uint64,uint64,uint256) ps) returns(uint256)
func (_Issuance *IssuanceTransactorSession) Issu(ps IssuParams) (*types.Transaction, error) {
	return _Issuance.Contract.Issu(&_Issuance.TransactOpts, ps)
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
