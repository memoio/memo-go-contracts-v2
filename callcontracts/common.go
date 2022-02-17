package callconts

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"golang.org/x/xerrors"

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
	RoleAddr = common.HexToAddress("0xaAAe8E331Fa0dEF6636c655C9b66388E9eE95Cc1")
	// RoleFS contract address
	RoleFSAddr = common.HexToAddress("0x75e8d452Ac0E8A9a1031B54AdeF9d78c7f60fc62")
	// RToken contract address
	RTokenAddr = common.HexToAddress("0x0E9d1F6a17614EFF6Ee15Bf8B1F2A2CBF47Bfb95")
	// FileSys contract address
	FileSysAddr = common.HexToAddress("0x66d1697f8263f4a264Ff4D8110Ac32de17004450")
	// PledgePool contract address
	PledgePoolAddr = common.HexToAddress("0xCb523235d0FeC5103bB34593407135380DA28E93")
	// Issuance contract address
	IssuanceAddr = common.HexToAddress("0x064B3D22AF5A25C87Df41Bf0dF61c5E271026267")
	AdminAddr    = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")
	AdminSk      = "0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4"
	Foundation   = common.HexToAddress("0x98B0B2387f98206efbF6fbCe2462cE22916BAAa3")
	FoundationSk = "4d360550b16bcb24d8f89e92915679a488f60ddd70fb05d2c3ee84726a4983aa"

	GIndex = uint64(2)
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
	waitTime                  = 1 * time.Second // after transaction, get info from chain

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
	EthSkLength = 66 // with prefix 0x
)

var (
	// ErrTxFail indicates that the transaction is not packaged or an error occurred during the packaging process
	ErrTxFail = xerrors.New("transaction not packaged")
	// ErrTxExecu indicates that an error occurred during packaging
	ErrTxExecu = xerrors.New("transaction mined but execution failed")
	// ErrBalNotE indicates that the account's balance is not enough to do something
	ErrBalNotE = xerrors.New("balance is not enough")
	// ErrAlloNotE indicates that the account's allowance is not enough to transferfrom
	ErrAlloNotE = xerrors.New("allowance is not enough")
	// ErrInValAddr indicates that the account address is invalid
	ErrInValAddr = xerrors.New("invalid address")
	// ErrNoMintRight indicates that the account has not Mint right in erc20 contract
	ErrNoMintRight = xerrors.New("the account has not Mint right")
	// ErrNoPauseRight indicates that the account has not Pause right in erc20 contract
	ErrNoPauseRight = xerrors.New("the account has not Pause right")
	// ErrNoAdminRight indicates that the account has not Admin right in erc20 contract
	ErrNoAdminRight      = xerrors.New("the account has not Admin right")
	errAccessControlRole = xerrors.New("the role in accessControl is invalid")
	// ErrIndex indicates that the rindex does not meet the requirements
	ErrIndex = xerrors.New("the role index is invalid")
	// ErrIndexZero shouldn't be zero
	ErrIndexZero = xerrors.New("the roleIndex or groupIndex should not be 0")
	// ErrOARange index out of range
	ErrOARange = xerrors.New("the index out of array range")
	// ErrIsBanned inidicates that the account is banned in Role contract, so some function about it cann't be called
	ErrIsBanned = xerrors.New("the account is banned in Role contract")
	// ErrTIndex tindex invalid
	ErrTIndex = xerrors.New("the token index is invalid")
	// ErrRoleReg has registered
	ErrRoleReg = xerrors.New("the account has already registered a role")
	// ErrInvalidG invalid gindex
	ErrInvalidG = xerrors.New("invalid group index")
	// ErrNotSetPP need set PledgePool address in Role contract
	ErrNotSetPP = xerrors.New("haven't set pledgePool address in Role contract before call RegisterToken")
	// ErrKSignsNE ksigns err
	ErrKSignsNE     = xerrors.New("the account of kSigns is not enough")
	errAllowanceExc = xerrors.New("the account's allowance to other account excess balance")
	errPledgeNE     = xerrors.New("the pledge money is not enough to pledgeKeeper or pledgeProvider")
	errHexskFormat  = xerrors.New("the hexsk'format is wrong")
	errPaused       = xerrors.New("transfer is paused now")
	errNotOwner     = xerrors.New("the caller is not owner")
	errSize         = xerrors.New("size should be greater than 0, or size err")
	errEnd          = xerrors.New("end should be greater than start")
	errEndNow       = xerrors.New("end should be more than start,and shouldn't be more than now")
	errNonce        = xerrors.New("nonce error")
	errSprice       = xerrors.New("sprice error")
	errCaller       = xerrors.New("caller error")
)

// TxOpts contains some general parameters about sending ethereum transaction
type TxOpts struct {
	Nonce    *big.Int // 赋值nil
	GasPrice *big.Int // 赋值nil，从而使用推荐值
	GasLimit uint64   // 赋值默认值
}

// ContractModule  The basic information of node used for contract.
// Address and private key need to correspond.
type ContractModule struct {
	addr            common.Address //local address
	hexSk           string         //local privateKey
	txopts          *TxOpts
	contractAddress common.Address
	endPoint        string     // ethClient endPoint
	Status          chan error // 上层调用函数需要监听该通道，持续从该通道中读取数据
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
		return nil, xerrors.New("new keyed transaction failed")
	}
	auth.GasPrice = txopts.GasPrice
	auth.Value = moneyToContract //放进合约里的钱
	auth.Nonce = txopts.Nonce
	auth.GasLimit = txopts.GasLimit
	return auth, nil
}

//CheckTx check whether transaction is successful through receipt
func checkTx(tx *types.Transaction, status chan error, name string) {
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

	// 矿工挂掉等情况导致交易无法被打包
	if receipt == nil { //231s获取不到交易信息，判定交易失败
		log.Println(name, ": cann't get tx receipt, tx not packaged")
		status <- ErrTxFail
		return
	}

	log.Println("GasUsed:", receipt.GasUsed, "CumulativeGasUsed:", receipt.CumulativeGasUsed)

	if receipt.Status == 0 { //等于0表示交易失败，等于1表示成功
		log.Println(name, ": transaction mined but execution failed, please check your tx input")
		if receipt.GasUsed != receipt.CumulativeGasUsed {
			log.Println(name, ": tx exceed gas limit")
		}
		status <- ErrTxExecu
		return
	}

	// 交易成功
	status <- nil
	log.Println(name, "has been successful!")

	// 获取withdraw money,输出到日志中
	if name == "Withdraw" {
		_, wd, err := getWithdrawInfoFromRLogs(tx.Hash())
		if err != nil {
			log.Println("getWithdrawInfoFromRLogs error:", err)
		}
		log.Println("tx:", tx.Hash().Hex(), ", account withdraw money:", wd)
	}
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
