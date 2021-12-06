package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	callconts "memoContract/callcontracts"
	"memoContract/test"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ethEndPoint  string
	qethEndPoint string
)

func main() {
	//--eth=http://119.147.213.219:8101 --qeth=http://119.147.213.219:8101      testnet网
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 查看余额，支付交易Gas费
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("test admin-account has balance: ", bal, " in Ethereum")
	bal = callconts.QueryEthBalance(test.Addr, ethEndPoint)
	fmt.Println("test common-account has balance: ", bal, " in Ethereum")

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	addr := common.HexToAddress(test.Addr)
	founder := uint64(0)           // 基金会地址的索引，固定为0
	gIndex := uint64(1)            // Role合约中CreateGroup函数的返回值
	rfs := common.HexToAddress("") // RoleFS合约地址
	keepers := []uint64{}          // keeper角色索引，等同于CreateGroup函数中的keepers参数
	uIndex := uint64(1)
	// 部署Role的参数
	pledgeK := big.NewInt(1e18)
	pledgeP := big.NewInt(1e18)

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	fmt.Println("============1. begin test deploy FileSys contract============")
	// 部署FileSys合约前准备工作：部署Role合约、账户注册角色、调用Role合约中的CreateGroup函数、部署RoleFS合约
	r := callconts.NewR(adminAddr, test.AdminSk, txopts)
	roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Role contract address: ", roleAddr.Hex())

	// 账户注册, addr:User addr2、addr3、addr4:Keeper addr5:Provider
	var addrs []common.Address = []common.Address{common.HexToAddress(test.Addr), common.HexToAddress(test.Addr2), common.HexToAddress(test.Addr3), common.HexToAddress(test.Addr4), common.HexToAddress(test.Addr5)}
	var sks []string = []string{test.Sk, test.Sk2, test.Sk3, test.Sk4, test.Sk5}
	rIndexes := make([]uint64, 5)
	for i, addr := range addrs {
		r = callconts.NewR(addr, sks[i], txopts)
		err = r.Register(roleAddr, addr, nil)
		if err != nil {
			log.Fatal(err)
		}
		rIndexes[i], err = r.GetRoleIndex(roleAddr, addr)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 部署PledgePool合约用于账户质押，之后才可以申请Keeper、Provider角色
	rtokenAddr, err := r.RToken(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The RToken contract address: ", rtokenAddr.Hex())
	p := callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
	pledgePoolAddr, _, err := p.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The PledgePool contract address: ", pledgePoolAddr.Hex())
	pledgek, err := r.PledgeK(roleAddr) // 申请Keeper最少需质押的金额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Apply Keeper need to pledge ", pledgek)
	for i, rindex := range rIndexes[1:4] {
		p = callconts.NewPledgePool(addrs[i+1], sks[i+1], txopts)
		// 调用pledge前需要先approve

		err = p.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, rindex, pledgek, nil)
	}

	e := callconts.NewFileSys(adminAddr, test.AdminSk, txopts)
	filesysAddr, _, err := e.DeployFileSys(founder, gIndex, roleAddr, rfs, keepers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FileSys contract Address:", filesysAddr.Hex())

	fmt.Println("============2. begin test GetFsInfo============")
	e = callconts.NewFileSys(addr, test.Sk, txopts)
	isActive, tokenIndex, err := e.GetFsInfo(filesysAddr, uIndex)
	if err != nil {
		log.Fatal(err)
	}
	if !isActive || tokenIndex != 0 {
		log.Fatal("The filesys with uIndex: ", uIndex, " isActive: ", isActive, " tokenIndex: ", tokenIndex)
	}
}
