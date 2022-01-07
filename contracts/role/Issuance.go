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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rfs\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spaceTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subPMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subSMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"13df3828": "issu(uint64,uint64,uint64,uint256)",
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
	Bin: "0x608060405234801561001057600080fd5b50604051610bad380380610bad83398101604081905261002f916103e8565b60004290508060028190555060016040518060600160405280606461ffff16815260200160016001600160401b0316815260200160016001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280607861ffff168152602001656400000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280609661ffff16815260200166040000000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b0316021790555050506001604051806060016040528060c861ffff16815260200166280000000000006001600160401b031681526020016283d6006001600160401b03168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff16021790555060208201518160000160026101000a8154816001600160401b0302191690836001600160401b03160217905550604082015181600001600a6101000a8154816001600160401b0302191690836001600160401b03160217905550505081600a60006101000a8154816001600160a01b0302191690836001600160a01b031602179055505050610418565b6000602082840312156103fa57600080fd5b81516001600160a01b038116811461041157600080fd5b9392505050565b610786806104276000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063949d225d11610071578063949d225d1461011a578063988bda7014610123578063a035b1fe1461012c578063e154adf514610135578063e7b0f6661461013e578063f0b44ab11461014757600080fd5b806311e65fc0146100ae57806313df3828146100c357806343a2755c146100e857806345fe29f6146100f1578063586fc5b514610111575b600080fd5b6100c16100bc3660046105c4565b610167565b005b6100d66100d1366004610603565b6101e6565b60405190815260200160405180910390f35b6100d660055481565b6100d66100ff36600461064e565b60096020526000908152604090205481565b6100d660025481565b6100d660045481565b6100d660005481565b6100d660035481565b6100d660065481565b6100d660075481565b6100d661015536600461064e565b60086020526000908152604090205481565b600a546001600160a01b031633146101aa5760405162461bcd60e51b81526020600482015260016024820152602760f91b60448201526064015b60405180910390fd5b81600760008282546101bc9190610686565b909155505060075481116101e25780600760008282546101dc919061069e565b90915550505b5050565b600a546000906001600160a01b031633146102275760405162461bcd60e51b81526020600482015260016024820152602760f91b60448201526064016101a1565b67ffffffffffffffff841660009081526008602052604081208054849290610250908490610686565b909155505067ffffffffffffffff8481166000908152600960205260408120805492861692909190610283908490610686565b90915550506002544290620151809061029c908361069e565b11156102b5576002546102b29062015180610686565b90505b6000600254826102c5919061069e565b90506000816003546102d791906106b5565b90506102e662015180846106d4565b620151806002546102f791906106d4565b10156103b557600061030c62015180856106d4565b61031990620151806106b5565b67ffffffffffffffff81166000908152600860205260409020549091508015610379576000610348838761069e565b61035290836106b5565b905061035e818561069e565b93508160036000828254610372919061069e565b9091555050505b67ffffffffffffffff821660009081526009602052604090205480156103b15780600460008282546103ab919061069e565b90915550505b5050505b80600760008282546103c79190610686565b90915550600090506103d989896106f6565b67ffffffffffffffff168767ffffffffffffffff166103f891906106b5565b9050806005600082825461040c9190610686565b909155506000905061041e8a8a6106f6565b6104329067ffffffffffffffff16886106b5565b905080600660008282546104469190610686565b925050819055508767ffffffffffffffff16600460008282546104699190610686565b9250508190555086600360008282546104829190610686565b90915550508261049f5750505060029190915550600090506105bc565b600080546104ae906001610686565b90505b60015481101561056e576000600182815481106104d0576104d061071f565b6000918252602090912001546004546201000090910467ffffffffffffffff1691508110156104fe57506004545b6000600183815481106105135761051361071f565b600091825260209091200154600554600160501b90910467ffffffffffffffff16915081906105439084906106d4565b10610552576000839055610559565b505061056e565b5050808061056690610735565b9150506104b1565b5060006001600054815481106105865761058661071f565b6000918252602090912001546105a09061ffff16856106b5565b90506105ad6064826106d4565b60029690965550939450505050505b949350505050565b600080604083850312156105d757600080fd5b50508035926020909101359150565b803567ffffffffffffffff811681146105fe57600080fd5b919050565b6000806000806080858703121561061957600080fd5b610622856105e6565b9350610630602086016105e6565b925061063e604086016105e6565b9396929550929360600135925050565b60006020828403121561066057600080fd5b610669826105e6565b9392505050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561069957610699610670565b500190565b6000828210156106b0576106b0610670565b500390565b60008160001904831182151516156106cf576106cf610670565b500290565b6000826106f157634e487b7160e01b600052601260045260246000fd5b500490565b600067ffffffffffffffff8381169083168181101561071757610717610670565b039392505050565b634e487b7160e01b600052603260045260246000fd5b600060001982141561074957610749610670565b506001019056fea26469706673582212205f32fee57aa65dd26294c6de9b8af01249ba323c78ac3ebbffec3b6798a2641064736f6c634300080a0033",
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
