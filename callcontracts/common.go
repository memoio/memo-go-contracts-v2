package callconts

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"memoContract/contracts/pledgepool"
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
	EndPoint  = "http://119.147.213.220:8191"
	QEndPoint = "http://119.147.213.220:8194"
	//ERC20 contract address
	ERC20Addr = common.HexToAddress("0xa96303D074eF892F39BCF5E19CD25Eeff7A73BAA")
	// Role contract address
	RoleAddr = common.HexToAddress("0xFbC6db9ae0f847F99907DA27068254e3482578d9")
	// RoleFS contract address
	RoleFSAddr = common.HexToAddress("0xCA3Ad5b308f1d00f4Dbb0eB7ff4EB8FDE079d138")
	// RToken contract address
	RTokenAddr = common.HexToAddress("0xEA490b38Fb5F872169D917e052D799942440a916")
	// FileSys contract address
	FileSysAddr = common.HexToAddress("0x139743cD6F0B5CA710F35552411971787D53d5FD")
	// PledgePool contract address
	PledgePoolAddr = common.HexToAddress("0xEF42eC9Cd7c140ec37D4470A62e95b57aDf24371")
	// Issuance contract address
	IssuanceAddr = common.HexToAddress("0xe8924Ed18C4270270696175F90b0C9D84b774A26")

	AdminAddr  = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")
	AdminSk    = "0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4"
	Foundation = common.HexToAddress("0x30F6551c2F970b21C1A9426aeb289c4ED6d570Fd")

	GIndex = uint64(1)
)

const (
	//InvalidAddr implements invalid contracts-address
	InvalidAddr = "0x0000000000000000000000000000000000000000"
	// DefaultGasPrice default gas price in sending transaction
	DefaultGasPrice = 200
	// DefaultGasLimit default gas limit in sending transaction
	DefaultGasLimit           = uint64(5000000) // as small as possible
	sendTransactionRetryCount = 5
	checkTxRetryCount         = 8
	checkTxSleepTime          = 6 // 先等待6s（出块时间加1）
	nextBlockTime             = 5 // 出块时间5s
	retryTxSleepTime          = time.Minute
	retryGetInfoSleepTime     = 30 * time.Second
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
	labelKeeper    = "keeper"
	labelProvider  = "provider"
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
	ErrTxFail = errors.New("transaction not packaged")
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
	// ErrIndexZero shouldn't be zero
	ErrIndexZero = errors.New("the roleIndex or groupIndex should not be 0")
	// ErrOARange index out of range
	ErrOARange = errors.New("the index out of array range")
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
	EndPoint string
}

// ContractModule  The basic information of node used for contract.
// Address and private key need to correspond.
type ContractModule struct {
	addr            common.Address //local address
	hexSk           string         //local privateKey
	txopts          *TxOpts
	contractAddress common.Address
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

	chainID := new(big.Int).SetUint64(35896)
	auth, err = bind.NewKeyedTransactorWithChainID(sk, chainID)
	if err != nil {
		return nil, errors.New("new keyed transaction failed")
	}
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
	t := checkTxSleepTime
	for i := 0; i < 10; i++ {
		if i != 0 {
			t = nextBlockTime * i
		}
		log.Printf("waiting %v sec.\n", t)
		time.Sleep(time.Duration(t) * time.Second)
		log.Println("getting txReceipt..")
		receipt = getTransactionReceipt(tx.Hash())
		if receipt != nil {
			break
		}
	}

	if receipt == nil { //231s获取不到交易信息，判定交易失败
		log.Println("get tx receipt nil, tx not packaged")
		return ErrTxFail
	}

	log.Println("GasUsed:", receipt.GasUsed, "CumulativeGasUsed:", receipt.CumulativeGasUsed)

	if receipt.Status == 0 { //等于0表示交易失败，等于1表示成功
		log.Println("Transaction mined but execution failed, please check your tx input")
		// txReceipt, err := receipt.MarshalJSON()
		// if err != nil {
		// 	return err
		// }
		// log.Println("TxReceipt:", string(txReceipt))
		if receipt.GasUsed != receipt.CumulativeGasUsed {
			log.Println("Err: tx exceed gas limit")
		}
		return ErrTxExecu
	}

	return nil
}

//GetTransactionReceipt 通过交易hash获得交易详情
func getTransactionReceipt(hash common.Hash) *types.Receipt {
	client, err := ethclient.Dial(EndPoint)
	if err != nil {
		log.Fatal("rpc.Dial err", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		log.Println("get transaction receipt err:", err)
	}
	return receipt
}

// increase tx's gasprice when error occur
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

// get RUser event info from tx hash
func getRUserInfoFromRLogs(hash common.Hash) (uint64, common.Address, error) {
	receipt := getTransactionReceipt(hash)

	if len(receipt.Logs) != 1 {
		return 0, common.Address{}, errors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, common.Address{}, err
	}

	// only 1 log in this tx
	intr, err := contractAbi.Events["RUser"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, common.Address{}, err
	}

	// index, address
	return intr[0].(uint64), intr[1].(common.Address), nil
}

// get RKeeper event info from tx hash
func getRKeeperInfoFromRLogs(hash common.Hash) (uint64, common.Address, error) {
	receipt := getTransactionReceipt(hash)

	if len(receipt.Logs) != 1 {
		return 0, common.Address{}, errors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, common.Address{}, err
	}

	// only 1 log in this tx
	intr, err := contractAbi.Events["RKeeper"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, common.Address{}, err
	}

	// index, address
	return intr[0].(uint64), intr[1].(common.Address), nil
}

// get RProvider event info from tx hash
func getRProviderInfoFromRLogs(hash common.Hash) (uint64, common.Address, error) {
	receipt := getTransactionReceipt(hash)

	if len(receipt.Logs) != 1 {
		return 0, common.Address{}, errors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(role.RoleABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return 0, common.Address{}, err
	}

	// only 1 log in this tx
	intr, err := contractAbi.Events["RProvider"].Inputs.UnpackValues(receipt.Logs[0].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return 0, common.Address{}, err
	}

	// index, address
	return intr[0].(uint64), intr[1].(common.Address), nil
}

// get Pledge event info in PledgePool contract from tx hash
func getPledgeInfoFromRLogs(hash common.Hash) (common.Address, *big.Int, error) {
	receipt := getTransactionReceipt(hash)

	log.Println("tx logs count: ", len(receipt.Logs))

	if len(receipt.Logs) != 3 {
		return common.Address{}, nil, errors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(pledgepool.PledgePoolABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return common.Address{}, nil, err
	}

	// 3 logs in this tx, event Pledge is the last one
	// get data
	d, err := contractAbi.Events["Pledge"].Inputs.UnpackValues(receipt.Logs[2].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return common.Address{}, nil, err
	}
	log.Println("data:", d)

	// actual topic start from index 1, and 0 is keccack256 of event signature: Pledge(address,uint256)
	t := receipt.Logs[2].Topics[1].String()
	log.Println("topic:", t)

	m, ok := d[0].(*big.Int)
	if !ok {
		log.Println("money type assertion failed")
	}

	// address, money
	return common.HexToAddress(t), m, nil
}

// get Withdraw event info in PledgePool contract from tx hash
func getWithdrawInfoFromRLogs(hash common.Hash) (common.Address, *big.Int, error) {
	receipt := getTransactionReceipt(hash)

	log.Println("tx logs count: ", len(receipt.Logs))

	if len(receipt.Logs) != 2 {
		return common.Address{}, nil, errors.New("length of logs in receipt is error")
	}

	// get role contract abi
	contractAbi, err := abi.JSON(strings.NewReader(string(pledgepool.PledgePoolABI)))
	if err != nil {
		log.Println("abi json err:", err)
		return common.Address{}, nil, err
	}

	// 2 logs in this tx, event Pledge is the last one
	// get data
	d, err := contractAbi.Events["Pledge"].Inputs.UnpackValues(receipt.Logs[1].Data)
	if err != nil {
		log.Println("unpack log err: ", err)
		return common.Address{}, nil, err
	}
	log.Println("data:", d)

	// actual topic start from index 1, and 0 is keccack256 of event signature: Pledge(address,uint256)
	t := receipt.Logs[1].Topics[1].String()
	log.Println("topic:", t)

	m, ok := d[0].(*big.Int)
	if !ok {
		log.Println("money type assertion failed")
	}

	// address, money
	return common.HexToAddress(t), m, nil
}

// SignForRegister Used to call Register on behalf of other accounts
// hash(caller, register)
func SignForRegister(caller common.Address, regAccSK string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSK)
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

// hash(caller,_blsKey,"keeper")
func SignForRegisterKeeper(caller common.Address, _blsKey []byte, regAccSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,_blsKey,"keeper")
	label := []byte(labelKeeper)
	hash := crypto.Keccak256(caller.Bytes(), _blsKey, label)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,"provider")
func SignForRegisterProvider(caller common.Address, regAccSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,"provider")
	label := []byte(labelProvider)
	hash := crypto.Keccak256(caller.Bytes(), label)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,gIndex,payToken,blsKey)
func SignForRegisterUser(caller common.Address, gIndex uint64, payToken uint32, _blsKey []byte, regAccSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,gIndex,payToken,blsKey)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	tmp4 := make([]byte, 4)
	// append caller
	b = append(b, caller.Bytes()...)
	// append gIndex
	binary.BigEndian.PutUint64(tmp8, gIndex)
	b = append(b, tmp8...)
	// append payToken
	binary.BigEndian.PutUint32(tmp4, payToken)
	b = append(b, tmp4...)
	// append blsKey
	b = append(b, _blsKey...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,gIndex)
func SignForAddProviderToGroup(caller common.Address, gIndex uint64, accSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(accSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,gIndex,payToken,blsKey)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	// append caller
	b = append(b, caller.Bytes()...)
	// append gIndex
	binary.BigEndian.PutUint64(tmp8, gIndex)
	b = append(b, tmp8...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller, uIndex, tIndex, money)
func SignForRecharge(caller common.Address, uIndex uint64, tIndex uint32, money *big.Int, accSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(caller, uIndex, tIndex, money)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	tmp4 := make([]byte, 4)

	// append caller
	b = append(b, caller.Bytes()...)
	// append uIndex
	binary.BigEndian.PutUint64(tmp8, uIndex)
	b = append(b, tmp8...)
	// append tIndex
	binary.BigEndian.PutUint32(tmp4, tIndex)
	b = append(b, tmp4...)
	// append money
	m := common.LeftPadBytes(money.Bytes(), 32)
	b = append(b, m...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller, tIndex, amount)
func SignForWithdrawFromFs(caller common.Address, tIndex uint32, amount *big.Int, accSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(caller, tIndex, amount)
	b := make([]byte, 0)
	tmp4 := make([]byte, 4)

	// append caller
	b = append(b, caller.Bytes()...)
	// append tIndex
	binary.BigEndian.PutUint32(tmp4, tIndex)
	b = append(b, tmp4...)
	// append amount
	m := common.LeftPadBytes(amount.Bytes(), 32)
	b = append(b, m...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
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
