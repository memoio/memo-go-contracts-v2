package callconts

import (
	"context"
	"errors"
	"log"
	"math/big"
	"memoContract/contracts/role"
	"strings"
	"time"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// EndPoint is rpc endpoint of geth node
	EndPoint string
	//ERC20 contract address
	ERC20 string
	// Role contract address
	Role string
	// RoleFS contract address
	RoleFS string
	// FileSys contract address
	FileSys string
	// PledgePool contract address
	PledgePool string
	// Issuance contract address
	Issuance string
)

const (
	//InvalidAddr implements invalid contracts-address
	InvalidAddr               = "0x0000000000000000000000000000000000000000"
	spaceTimePayGasLimit      = uint64(8000000)
	spaceTimePayGasPrice      = 2 * defaultGasPrice
	defaultGasPrice           = 200
	defaultGasLimit           = uint64(8000000)
	sendTransactionRetryCount = 5
	checkTxRetryCount         = 8
	checkTxSleepTime          = 5
	retryTxSleepTime          = time.Minute
	retryGetInfoSleepTime     = time.Minute
	waitTime                  = 3 * time.Second

	// AdminRole indicates the account has Admin right in ERC20 contract
	AdminRole = uint8(0)
	// MinterRole indicates the account has Minter right in ERC20 contract
	MinterRole = uint8(1)
	// PauserRole indicates the account has Pauser right in ERC20 contract
	PauserRole = uint8(2)

	// UserRoleType indicates user's roleType in Role contract
	UserRoleType = 1
	// ProviderRoleType indicates provider's roleType in Role contract
	ProviderRoleType = 2
	// KeeperRoleType indicates keeper's roleType in Role contract
	KeeperRoleType = 3
	register       = "role-register"
)

var (
	// ErrTxFail indicates that the transaction is not packaged or an error occurred during the packaging process
	ErrTxFail = errors.New("Transaction fails")
	// ErrTxExecu indicates that an error occurred during packaging
	ErrTxExecu = errors.New("Transaction mined but execution failed")
	// ErrBalNotE indicates that the account's balance is not enough to do something
	ErrBalNotE = errors.New("Balance is not enough")
	// ErrAlloNotE indicates that the account's allowance is not enough to transferfrom
	ErrAlloNotE = errors.New("Allowance is not enough")
	// ErrInValAddr indicates that the account address is invalid
	ErrInValAddr = errors.New("Invalid address")
	// ErrNoMintRight indicates that the account has not Mint right in erc20 contract
	ErrNoMintRight = errors.New("The account has not Mint right")
	// ErrNoPauseRight indicates that the account has not Pause right in erc20 contract
	ErrNoPauseRight = errors.New("The account has not Pause right")
	// ErrNoAdminRight indicates that the account has not Admin right in erc20 contract
	ErrNoAdminRight      = errors.New("The account has not Admin right")
	errAccessControlRole = errors.New("The role in accessControl is invalid")
	// ErrIndex indicates that the rindex does not meet the requirements
	ErrIndex = errors.New("The role index is invalid")
	// ErrIsBanned inidicates that the account is banned in Role contract, so some function about it cann't be called
	ErrIsBanned = errors.New("The account is banned in Role contract")
	// ErrTIndex tindex invalid
	ErrTIndex = errors.New("The token index is invalid")
	// ErrRoleReg has registered
	ErrRoleReg = errors.New("The account has already registered a role")
	// ErrInvalidG invalid gindex
	ErrInvalidG = errors.New("invalid group index")
	// ErrNotSetPP need set PledgePool address in Role contract
	ErrNotSetPP = errors.New("haven't set pledgePool address in Role contract before call RegisterToken")
	// ErrKSignsNE ksigns err
	ErrKSignsNE = errors.New("the account of kSigns is not enough")
	errAllowanceExc = errors.New("the account's allowance to other account excess balance")
)

// TxOpts contains some general parameters about sending ethereum transaction
type TxOpts struct {
	Nonce    *big.Int
	GasPrice *big.Int
	GasLimit uint64
}

// ContractModule  The basic information of node used for contract.
// Address and private key need to correspond.
type ContractModule struct {
	addr   common.Address //local address
	hexSk  string         //local privateKey
	txopts *TxOpts
}

// getClient get rpc-client based the endPoint
func getClient(endPoint string) *ethclient.Client {
	client, err := rpc.Dial(EndPoint)
	if err != nil {
		log.Println(err)
	}
	return ethclient.NewClient(client)
}

// makeAuth make the transactOpts to call contract
func makeAuth(hexSk string, moneyToContract *big.Int, txopts *TxOpts) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		log.Println("HexToECDSA err: ", err)
		return auth, err
	}

	auth = bind.NewKeyedTransactor(sk)
	auth.GasPrice = txopts.GasPrice
	auth.Value = moneyToContract //放进合约里的钱
	auth.Nonce = txopts.Nonce
	auth.GasLimit = txopts.GasLimit
	return auth, nil
}

//CheckTx check whether transaction is successful through receipt
func checkTx(tx *types.Transaction) error {
	log.Println("Check Tx hash:", tx.Hash().Hex(), "nonce:", tx.Nonce(), "gasPrice:", tx.GasPrice())

	var receipt *types.Receipt
	for i := 0; i < 20; i++ {
		receipt = getTransactionReceipt(tx.Hash())
		if receipt != nil {
			break
		}
		t := checkTxSleepTime * (i + 1)
		time.Sleep(time.Duration(t) * time.Second)
	}

	if receipt == nil { //245s获取不到交易信息，判定交易失败
		return ErrTxFail
	}

	if receipt.Status == 0 { //等于0表示交易失败，等于1表示成功
		log.Println("Transaction mined but execution failed")
		txReceipt, err := receipt.MarshalJSON()
		if err != nil {
			return err
		}
		log.Println("TxReceipt:", string(txReceipt))
		return ErrTxExecu
	}

	log.Println("GasUsed:", receipt.GasUsed, "CumulativeGasUsed:", receipt.CumulativeGasUsed)

	return nil
}

//GetTransactionReceipt 通过交易hash获得交易详情
func getTransactionReceipt(hash common.Hash) *types.Receipt {
	client, err := ethclient.Dial(EndPoint)
	if err != nil {
		log.Fatal("rpc.Dial err", err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	return receipt
}

func rebuild(err error, tx *types.Transaction, auth *bind.TransactOpts) {
	if err == ErrTxFail && tx != nil {
		auth.Nonce = big.NewInt(int64(tx.Nonce()))
		auth.GasPrice = new(big.Int).Add(tx.GasPrice(), big.NewInt(defaultGasPrice))
		log.Println("rebuild transaction... nonce is ", auth.Nonce, " gasPrice is ", auth.GasPrice)
	}
}

// get gIndex from logs in receipt
func getGIndexFromRLogs(hash common.Hash) (uint64, error) {
	receipt := getTransactionReceipt(hash)

	if len(receipt.Logs) != 1 {
		return 0, errors.New("length of logs in receipt is error")
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, err
	}

	intr, err := contractAbi.Events["CreateGroup"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, err
	}
	return intr[0].(uint64), nil
}

// SignForRegister Used to call Register on behalf of other accounts
func SignForRegister(caller common.Address, sk *ecdsa.PrivateKey) ([]byte, error) {
	//(caller, register)的哈希值
	label := common.LeftPadBytes([]byte(register), 32)
	hash := crypto.Keccak256(caller.Bytes(), label)

	//私钥对上述哈希值签名
	sig, err := crypto.Sign(hash, sk)
	if err != nil {
		return nil, err
	}

	return sig, nil
}
