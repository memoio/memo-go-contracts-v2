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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rfs\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"_start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_end\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_size\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"sPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIssuParams\",\"name\":\"ps\",\"type\":\"tuple\"}],\"name\":\"issu\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minRatio\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_add\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sub\",\"type\":\"uint256\"}],\"name\":\"setTP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spaceTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subPMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"subSMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c862ea50": "issu((uint64,uint64,uint64,uint256))",
		"44178c25": "issuRatio()",
		"586fc5b5": "lastMint()",
		"86d8745b": "minRatio()",
		"988bda70": "mintLevel()",
		"030bb079": "periodTarget()",
		"517463e6": "periodTotalReward()",
		"a035b1fe": "price()",
		"11e65fc0": "setTP(uint256,uint256)",
		"949d225d": "size()",
		"43a2755c": "spaceTime()",
		"f0b44ab1": "subPMap(uint64)",
		"45fe29f6": "subSMap(uint64)",
		"e7b0f666": "totalPaid()",
		"e154adf5": "totalPay()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610dcb380380610dcb83398101604081905261002f91610448565b60004290508060028190555060016040518060600160405280603261ffff1681526020016564000000000081526020016283d6006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280605061ffff1681526020016604000000000000815260200162c5c1006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280606461ffff16815260200166c80000000000008152602001630107ac006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280605061ffff16815260200167100000000000000081526020016283d6006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505060016040518060600160405280603261ffff1681526020016803200000000000000081526020016303c267006001600160401b0316815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548161ffff021916908361ffff1602179055506020820151816001015560408201518160020160006101000a8154816001600160401b0302191690836001600160401b03160217905550505081600e60006101000a8154816001600160a01b0302191690836001600160a01b031602179055506003600b60006101000a81548161ffff021916908361ffff1602179055506ba18f07d736b90be5500000006008819055506032600a819055505050610478565b60006020828403121561045a57600080fd5b81516001600160a01b038116811461047157600080fd5b9392505050565b610944806104876000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806386d8745b11610097578063c862ea5011610066578063c862ea50146101ab578063e154adf5146101be578063e7b0f666146101c7578063f0b44ab1146101d057600080fd5b806386d8745b1461016f578063949d225d14610190578063988bda7014610199578063a035b1fe146101a257600080fd5b806344178c25116100d357806344178c251461013457806345fe29f61461013d578063517463e61461015d578063586fc5b51461016657600080fd5b8063030bb079146100fa57806311e65fc01461011657806343a2755c1461012b575b600080fd5b61010360085481565b6040519081526020015b60405180910390f35b61012961012436600461074e565b6101f0565b005b61010360055481565b610103600a5481565b61010361014b366004610787565b600d6020526000908152604090205481565b61010360095481565b61010360025481565b600b5461017d9061ffff1681565b60405161ffff909116815260200161010d565b61010360045481565b61010360005481565b61010360035481565b6101036101b93660046107a9565b61026f565b61010360065481565b61010360075481565b6101036101de366004610787565b600c6020526000908152604090205481565b600e546001600160a01b031633146102335760405162461bcd60e51b81526020600482015260016024820152602760f91b60448201526064015b60405180910390fd5b81600760008282546102459190610845565b9091555050600754811161026b578060076000828254610265919061085d565b90915550505b5050565b600e546000906001600160a01b031633146102b05760405162461bcd60e51b81526020600482015260016024820152602760f91b604482015260640161022a565b60608201516020808401516001600160401b03166000908152600c9091526040812080549091906102e2908490610845565b9250508190555081604001516001600160401b0316600d600084602001516001600160401b03166001600160401b03168152602001908152602001600020600082825461032f9190610845565b909155505060025442906201518090610348908361085d565b11156103615760025461035e9062015180610845565b90505b600060025482610371919061085d565b90506000816003546103839190610874565b90506103926201518084610893565b620151806002546103a39190610893565b101561045f5760006103b86201518085610893565b6103c59062015180610874565b6001600160401b0381166000908152600c602052604090205490915080156104245760006103f3838761085d565b6103fd9083610874565b9050610409818561085d565b9350816003600082825461041d919061085d565b9091555050505b6001600160401b0382166000908152600d6020526040902054801561045b578060046000828254610455919061085d565b90915550505b5050505b80600760008282546104719190610845565b909155505084516020860151600091610489916108b5565b6001600160401b031686604001516001600160401b03166104aa9190610874565b905080600560008282546104be9190610845565b9091555050855160208701516000916104d6916108b5565b6001600160401b031687606001516104ee9190610874565b905080600660008282546105029190610845565b9250508190555086604001516001600160401b0316600460008282546105289190610845565b9091555050606087015160038054600090610544908490610845565b9091555050826105605750505060029190915550600092915050565b6000545b60015481101561066757600060018281548110610583576105836108dd565b90600052602060002090600302016001015490506004548110156105a657506004545b6000600183815481106105bb576105bb6108dd565b60009182526020909120600260039092020101546005546001600160401b03909116915081906105ec908490610893565b1061064b576105fc836001610845565b600081905550600160005481548110610617576106176108dd565b600091825260209091206003909102015461ffff16600a5560085461063e90600290610893565b6008556000600955610652565b5050610667565b5050808061065f906108f3565b915050610564565b506000600a54846106789190610874565b9050610685606482610893565b90508560028190555060008160095461069e9190610845565b9050600854811180156106ba5750600b54600a5461ffff909116105b156107405760006009546008546106d1919061085d565b905060006106df828561085d565b90506106ec600282610893565b90506106f8818561085d565b9350836009600082825461070c9190610845565b9091555050600a5461072090600290610893565b600a819055600b5461ffff16111561073d57600b5461ffff16600a555b50505b50955050505050505b919050565b6000806040838503121561076157600080fd5b50508035926020909101359150565b80356001600160401b038116811461074957600080fd5b60006020828403121561079957600080fd5b6107a282610770565b9392505050565b6000608082840312156107bb57600080fd5b604051608081018181106001600160401b03821117156107eb57634e487b7160e01b600052604160045260246000fd5b6040526107f783610770565b815261080560208401610770565b602082015261081660408401610770565b6040820152606083013560608201528091505092915050565b634e487b7160e01b600052601160045260246000fd5b600082198211156108585761085861082f565b500190565b60008282101561086f5761086f61082f565b500390565b600081600019048311821515161561088e5761088e61082f565b500290565b6000826108b057634e487b7160e01b600052601260045260246000fd5b500490565b60006001600160401b03838116908316818110156108d5576108d561082f565b039392505050565b634e487b7160e01b600052603260045260246000fd5b60006000198214156109075761090761082f565b506001019056fea2646970667358221220080a7c6b7b8587c93fbfe4ffe108298ca66ee619c872b2070e9620e654e8a73964736f6c634300080b0033",
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

// IssuRatio is a free data retrieval call binding the contract method 0x44178c25.
//
// Solidity: function issuRatio() view returns(uint256)
func (_Issuance *IssuanceCaller) IssuRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "issuRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IssuRatio is a free data retrieval call binding the contract method 0x44178c25.
//
// Solidity: function issuRatio() view returns(uint256)
func (_Issuance *IssuanceSession) IssuRatio() (*big.Int, error) {
	return _Issuance.Contract.IssuRatio(&_Issuance.CallOpts)
}

// IssuRatio is a free data retrieval call binding the contract method 0x44178c25.
//
// Solidity: function issuRatio() view returns(uint256)
func (_Issuance *IssuanceCallerSession) IssuRatio() (*big.Int, error) {
	return _Issuance.Contract.IssuRatio(&_Issuance.CallOpts)
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

// MinRatio is a free data retrieval call binding the contract method 0x86d8745b.
//
// Solidity: function minRatio() view returns(uint16)
func (_Issuance *IssuanceCaller) MinRatio(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "minRatio")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MinRatio is a free data retrieval call binding the contract method 0x86d8745b.
//
// Solidity: function minRatio() view returns(uint16)
func (_Issuance *IssuanceSession) MinRatio() (uint16, error) {
	return _Issuance.Contract.MinRatio(&_Issuance.CallOpts)
}

// MinRatio is a free data retrieval call binding the contract method 0x86d8745b.
//
// Solidity: function minRatio() view returns(uint16)
func (_Issuance *IssuanceCallerSession) MinRatio() (uint16, error) {
	return _Issuance.Contract.MinRatio(&_Issuance.CallOpts)
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

// PeriodTarget is a free data retrieval call binding the contract method 0x030bb079.
//
// Solidity: function periodTarget() view returns(uint256)
func (_Issuance *IssuanceCaller) PeriodTarget(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "periodTarget")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTarget is a free data retrieval call binding the contract method 0x030bb079.
//
// Solidity: function periodTarget() view returns(uint256)
func (_Issuance *IssuanceSession) PeriodTarget() (*big.Int, error) {
	return _Issuance.Contract.PeriodTarget(&_Issuance.CallOpts)
}

// PeriodTarget is a free data retrieval call binding the contract method 0x030bb079.
//
// Solidity: function periodTarget() view returns(uint256)
func (_Issuance *IssuanceCallerSession) PeriodTarget() (*big.Int, error) {
	return _Issuance.Contract.PeriodTarget(&_Issuance.CallOpts)
}

// PeriodTotalReward is a free data retrieval call binding the contract method 0x517463e6.
//
// Solidity: function periodTotalReward() view returns(uint256)
func (_Issuance *IssuanceCaller) PeriodTotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Issuance.contract.Call(opts, &out, "periodTotalReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTotalReward is a free data retrieval call binding the contract method 0x517463e6.
//
// Solidity: function periodTotalReward() view returns(uint256)
func (_Issuance *IssuanceSession) PeriodTotalReward() (*big.Int, error) {
	return _Issuance.Contract.PeriodTotalReward(&_Issuance.CallOpts)
}

// PeriodTotalReward is a free data retrieval call binding the contract method 0x517463e6.
//
// Solidity: function periodTotalReward() view returns(uint256)
func (_Issuance *IssuanceCallerSession) PeriodTotalReward() (*big.Int, error) {
	return _Issuance.Contract.PeriodTotalReward(&_Issuance.CallOpts)
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
