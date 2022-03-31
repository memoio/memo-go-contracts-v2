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

	
		// for test chain
		// EndPoint is rpc endpoint of geth node
		EndPoint  = "http://119.147.213.220:8191"
		QEndPoint = "http://119.147.213.220:8194"
		//ERC20 contract address
		ERC20Addr = common.HexToAddress("0x788E554e94f05e6b7862563d9D708d982ba96D5C")
		// Role contract address
		RoleAddr = common.HexToAddress("0xCa2C4103bd5679F43eC9E277C2bAf5598f94Fe6D")
		// RToken contract address
		RTokenAddr = common.HexToAddress("0x26ea029a4c778e74f2d64Ca8DFA7f36c6aF4FE60")
		// RoleFS contract address
		RoleFSAddr = common.HexToAddress("0xFd688f2D007019422E314EbA32C68a6F5D668fC0")
		// FileSys contract address
		FileSysAddr = common.HexToAddress("0xFB9FF16EB4093aa8fFf762F2dF4E61d3A7532Af9")
		// for group2, level 4: 0xE915C7cF05e8AeE4A001327c588D8837bD0CcF2C
		// for group3, level 4: 0x0E79b7F5aE172B39a9f3440561AB317011f3296B

		// PledgePool contract address
		PledgePoolAddr = common.HexToAddress("0x76CB45B03096Fd922aef3413cB04FCC6A97d93C4")
		// Issuance contract address
		IssuanceAddr = common.HexToAddress("0x16D62bA68Eda95967ACfFb9a03586A9965cd77d5")

		AdminAddr    = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")
		AdminSk      = "0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4"
		Foundation   = common.HexToAddress("0x98B0B2387f98206efbF6fbCe2462cE22916BAAa3")
		FoundationSk = "4d360550b16bcb24d8f89e92915679a488f60ddd70fb05d2c3ee84726a4983aa"

		GIndex = uint64(2)
	

	/*
		// for dev chain
		// EndPoint is rpc endpoint of geth node
		EndPoint  = "https://devchain.metamemo.one:8501"
		QEndPoint = "https://devchain.metamemo.one:8501"

		ERC20Addr = common.HexToAddress("0xb8b6256C5A6bFdf38717D16CcaB3C769F762Fbd7")
		RoleAddr = common.HexToAddress("0xa53Ad83C0c5b86bd93De00eD5233091680C0215B")
		RTokenAddr = common.HexToAddress("0xDf71cd95C72B45D51D364a53d63a552F583d507b")
		RoleFSAddr = common.HexToAddress("0xae63081f7F836FF6dd67C8440912f16111DC5980")
		PledgePoolAddr = common.HexToAddress("0xbbf742E1533732DAd71e4bd2B634dA23Fdc75792")
		IssuanceAddr = common.HexToAddress("0x1f7dEaC7aB4412fc1e9074026B2Cad8aB27c776B")
		FileSysAddr = common.HexToAddress("0xCeFf1C9a59193BaaB91f5Abe5e370dac5FDd4309")

		AdminAddr    = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")
		AdminSk      = "0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4"
		Foundation   = common.HexToAddress("0x98B0B2387f98206efbF6fbCe2462cE22916BAAa3")
		FoundationSk = "4d360550b16bcb24d8f89e92915679a488f60ddd70fb05d2c3ee84726a4983aa"

		GIndex = uint64(1)
	*/

	/*
	// for product chain
	// EndPoint is rpc endpoint of geth node
	EndPoint  = "https://chain.metamemo.one:8501"
	QEndPoint = "https://chain.metamemo.one:8501"
	//ERC20 contract address
	ERC20Addr = common.HexToAddress("0x7eF37CB66Bc9Ab2b88b0dC9a8461f99CEFA73D57")
	// Role contract address
	RoleAddr = common.HexToAddress("0x0fC98129261dadEc4576f4b052F64354F6f6C634")
	// RToken contract address
	RTokenAddr = common.HexToAddress("0xEa838C3198C031f5EBb2F079207aCF40cCB35A85")
	// RoleFS contract address
	RoleFSAddr = common.HexToAddress("0x95ac462Dec4Ffb4ce723FF35B7edd3Acb8F9d1d7")
	// PledgePool contract address
	PledgePoolAddr = common.HexToAddress("0x28Fd6aDA92F3C213a4439e7a8Ad2DaFDD9815291")
	// Issuance contract address
	IssuanceAddr = common.HexToAddress("0xEfA598A29b03A7c9176CA446F1d26e9E85ef0E9f")
	// FileSys contract address
	FileSysAddr = common.HexToAddress("0x9636865Dae73a72F88De1Afd15E582848765fefF")
	// for group 6, level=3: 0x7877BB55198D814521ad3daBa27e31fFb52d913D
	// for group 7, level=3: 0xA6122C0A0c1baD78A70B92537c8E374d01373041

	AdminAddr    = common.HexToAddress("0x76e391468b627c802d38ba1e56a1f5cfb05b98ee")
	AdminSk      = "474e530fc2aff814e9c4526fe06b4efd2043ffd8e52d67388b057919f0df2bdb"
	Foundation   = common.HexToAddress("0x98B0B2387f98206efbF6fbCe2462cE22916BAAa3")
	FoundationSk = "4d360550b16bcb24d8f89e92915679a488f60ddd70fb05d2c3ee84726a4983aa"

	// account/sk for deploy erc20：
	// 1.
	// 0x7ef756852946514924d7bdbe393096e13e895bc8
	// 511f471249ec5690d4ffd05af108b96f246a61f382b7da5cd1b587fa7bcacee3
	// 2.
	// 0x6820a0856c9c6fd13ba5b3e12cf0e77c88914817
	// b330076b4f869e41753196fd8c5a076877a8a6549646413b277985c2a7e37b9e
	// 3.
	// 0x50110b34244e8a88571e6c6c8270be6d8adb2058
	// 31c00148d644578ebebde6cbf9a75b4266821d2944c8af55276779c9e9edf7f1
	// 4.
	// 0x79fd9fbfeca6f3c01e79f86c28e35313db831404
	// bde31b12a25d4fe3696f567bb89273be7e0f425850a0972a438f7627eee7feba
	// 5.
	// 0x0a7c7f353c176474ab7da094b47f1189e976bc08
	// 3ea2511812b33f61b387f8a706378e8658a6f4a1bf9c20615e2af00985c59a6c

	GIndex = uint64(1)
	*/
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
	// ErrNotActive group not active
	ErrNotActive = xerrors.New("group not active")
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
	client, err := rpc.Dial(endPoint)
	if err != nil {
		log.Println(err)
	}
	return ethclient.NewClient(client)
}

// makeAuth make the transactOpts to call contract
func makeAuth(endPoint string, hexSk string, moneyToContract *big.Int, txopts *TxOpts) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		log.Println("HexToECDSA err: ", err)
		return auth, err
	}

	client := getClient(endPoint)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("client.NetworkID error,use the default chainID")
		chainID = big.NewInt(666)
	}
	fmt.Println("chainID: ", chainID)
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
func checkTx(endPoint string, tx *types.Transaction, status chan error, name string) {
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
		receipt = getTransactionReceipt(endPoint, tx.Hash())
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
		_, wd, err := getWithdrawInfoFromRLogs(endPoint, tx.Hash())
		if err != nil {
			log.Println("getWithdrawInfoFromRLogs error:", err)
		}
		log.Println("tx:", tx.Hash().Hex(), ", account withdraw money:", wd)
	}
}

//GetTransactionReceipt 通过交易hash获得交易详情
func getTransactionReceipt(endPoint string, hash common.Hash) *types.Receipt {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		log.Fatal("rpc.Dial err", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		log.Println("get transaction receipt: ", err)
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
