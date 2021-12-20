package callconts

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
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

// the following variables need to be assigned according to the running results in actual applications
var (
	// EndPoint is rpc endpoint of geth node
	EndPoint string
	//ERC20 contract address
	ERC20Addr common.Address
	// Role contract address
	RoleAddr common.Address
	// RoleFS contract address
	RoleFSAddr common.Address
	// FileSys contract address
	FileSysAddr common.Address
	// PledgePool contract address
	PledgePoolAddr common.Address
	// Issuance contract address
	IssuanceAddr common.Address
)

const (
	//InvalidAddr implements invalid contracts-address
	InvalidAddr          = "0x0000000000000000000000000000000000000000"
	spaceTimePayGasLimit = uint64(8000000)
	spaceTimePayGasPrice = 2 * DefaultGasPrice
	// DefaultGasPrice default gas price in sending transaction
	DefaultGasPrice = 200
	// DefaultGasLimit default gas limit in sending transaction
	DefaultGasLimit           = uint64(8000000)
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
	// topic of contract log
	transferTopic    = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	alterOwnerTopic  = "0x8c153ecee6895f15da72e646b4029e0ef7cbf971986d8d9cfe48c5563d368e90"
	addTTopic        = "0x7a5e6bb234636aa6f5c8428d056a65a5c9ec9d25638a69ad3bd3e362e64a8de6"
	approvalTopic    = "0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"
	pledgeTopic      = "0x5e91ea8ea1c46300eb761859be01d7b16d44389ef91e03a163a87413cbf55b95"
	rKeeperTopic     = "0xc50f17198ae53c50e1ad2f06d8348c7d6b31952e4bc9f52b15bb075e6f1bed0b"
	rProviderTopic   = "0xf06105db8b89019d932bb3ec85a22bbed723c3616043e02ca9857f3ba37005a5"
	rUserTopic       = "0x7defc6162296f3490e44c1787f4ae9852a8d6a8e67ba0a69c57bd7be5f8a0b1a"
	createGroupTopic = "0xc79ca32352cc5529f3c78b5cb44574fbc979555a04f5b6425ed2595417da2e64"
	// EthSkLength sk length in ethereum
	EthSkLength = 66 // 0x
)

var (
	// ErrTxFail indicates that the transaction is not packaged or an error occurred during the packaging process
	ErrTxFail = errors.New("transaction fails")
	// ErrTxExecu indicates that an error occurred during packaging
	ErrTxExecu = errors.New("transaction mined but execution failed")
	// ErrBalNotE indicates that the account's balance is not enough to do something
	ErrBalNotE = errors.New("balance is not enough")
	// ErrAlloNotE indicates that the account's allowance is not enough to transferfrom
	ErrAlloNotE = errors.New("allowance is not enough")
	// ErrInValAddr indicates that the account address is invalid
	ErrInValAddr = errors.New("invalid address")
	// ErrNoMintRight indicates that the account has not Mint right in erc20 contract
	ErrNoMintRight = errors.New("the account has not Mint right")
	// ErrNoPauseRight indicates that the account has not Pause right in erc20 contract
	ErrNoPauseRight = errors.New("the account has not Pause right")
	// ErrNoAdminRight indicates that the account has not Admin right in erc20 contract
	ErrNoAdminRight      = errors.New("the account has not Admin right")
	errAccessControlRole = errors.New("the role in accessControl is invalid")
	// ErrIndex indicates that the rindex does not meet the requirements
	ErrIndex = errors.New("the role index is invalid")
	// ErrIsBanned inidicates that the account is banned in Role contract, so some function about it cann't be called
	ErrIsBanned = errors.New("the account is banned in Role contract")
	// ErrTIndex tindex invalid
	ErrTIndex = errors.New("the token index is invalid")
	// ErrRoleReg has registered
	ErrRoleReg = errors.New("the account has already registered a role")
	// ErrInvalidG invalid gindex
	ErrInvalidG = errors.New("invalid group index")
	// ErrNotSetPP need set PledgePool address in Role contract
	ErrNotSetPP = errors.New("haven't set pledgePool address in Role contract before call RegisterToken")
	// ErrKSignsNE ksigns err
	ErrKSignsNE     = errors.New("the account of kSigns is not enough")
	errAllowanceExc = errors.New("the account's allowance to other account excess balance")
	errPledgeNE     = errors.New("the pledge money is not enough to pledgeKeeper or pledgeProvider")
	errHexskFormat  = errors.New("the hexsk'format is wrong")
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
	log.Println("waiting for miner...")

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
		auth.GasPrice = new(big.Int).Add(tx.GasPrice(), big.NewInt(DefaultGasPrice))
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
func SignForRegister(caller common.Address, sk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(sk)
	if err != nil {
		log.Fatal(err)
	}
	//(caller, register)的哈希值
	//label := common.LeftPadBytes([]byte(register), 32)
	label := []byte(register)
	hash := crypto.Keccak256(caller.Bytes(), label)

	//私钥对上述哈希值签名
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// todo
func SignForRegisterKeeper(caller common.Address, sk string) ([]byte, error) {
	return nil, nil
}

// todo
func SignForRegisterProvider(caller common.Address, sk string) ([]byte, error) {
	return nil, nil
}

// todo
func SignForRegisterUser(caller common.Address, sk string) ([]byte, error) {
	return nil, nil
}

// todo
func SignForAddProviderToGroup(caller common.Address, sk string) ([]byte, error) {
	return nil, nil
}

// todo
func SignForWithdrawFromFs(caller common.Address, sk string) ([]byte, error) {
	return nil, nil
}

// SignForRepair used to call AddRepair or SubRepair, when subRepair, label is "s"; when addRepair, label is "a"
func SignForRepair(sk string, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, label string) ([]byte, error) {
	ecdsaSk, err := HexSkToEcdsa(sk)
	if err != nil {
		log.Fatal(err)
	}

	// (npIndex, _start, end, _size, nonce, tIndex, sprice, label)的哈希值
	by := make([]byte, 77)
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, pIndex)
	for i, b := range tmp {
		by[i] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, start)
	for i, b := range tmp {
		by[i+8] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, end)
	for i, b := range tmp {
		by[i+16] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, size)
	for i, b := range tmp {
		by[i+24] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, nonce)
	for i, b := range tmp {
		by[i+32] = byte(b)
	}
	t := make([]byte, 4)
	binary.BigEndian.PutUint32(t, tIndex)
	for i, b := range tmp {
		by[i+40] = byte(b)
	}
	spriceNew := common.LeftPadBytes(sprice.Bytes(), 32)
	for i, b := range spriceNew {
		by[i+44] = byte(b)
	}
	labelNew := []byte(label)
	by[76] = byte(labelNew[0])

	hash := crypto.Keccak256(by)

	fmt.Println("hash:", hash)

	// 私钥签名
	sig, err := crypto.Sign(hash, ecdsaSk)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// QueryEthBalance gets eth balance
func QueryEthBalance(addr, ethEndPoint string) *big.Int {
	var result string
	client, err := rpc.Dial(ethEndPoint)
	if err != nil {
		log.Fatal("rpc.dial err:", err)
	}
	err = client.Call(&result, "eth_getBalance", addr, "latest")
	if err != nil {
		log.Fatal("client.call err:", err)
	}
	fmt.Println("balance:", result)
	a := big.NewInt(0)
	a, success := a.SetString(result[2:], 16)
	if !success {
		log.Fatal("hex to bigInt fails")
	}
	return a
}

// HexSkToByte transfer hex string to byte
func HexSkToByte(hexsk string) ([]byte, error) {
	var src []byte
	skLengthNoPrefix := EthSkLength - 2
	skByteEthLength := skLengthNoPrefix / 2

	switch len(hexsk) {
	case EthSkLength:
		if hexsk[:2] == "0x" {
			src = []byte(hexsk[2:])
		} else {
			return nil, errHexskFormat
		}
	case skLengthNoPrefix:
		src = []byte(hexsk)
	default:
		return nil, errHexskFormat
	}

	skByteEth := make([]byte, skByteEthLength)

	_, err := hex.Decode(skByteEth, src)
	if err != nil {
		return nil, err
	}

	return skByteEth, nil
}

// HexSkToEcdsa transfer hex string to ecdsa type
func HexSkToEcdsa(hexsk string) (*ecdsa.PrivateKey, error) {
	skByteEth, err := HexSkToByte(hexsk)
	if err != nil {
		return nil, err
	}
	skECDSA, err := crypto.ToECDSA(skByteEth)
	if err != nil {
		return nil, err
	}
	return skECDSA, nil
}
