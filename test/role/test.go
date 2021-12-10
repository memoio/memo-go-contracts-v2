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

// test process
// 1. deploy Role
// 2. deploy PledgePool
// 3. deploy RoleFS
// 4. deploy Issuance
// 5. call setPI
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
	bal = callconts.QueryEthBalance(test.Acc1, ethEndPoint)
	fmt.Println("test common-account has balance: ", bal, " in Ethereum")

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	pledgeK := big.NewInt(1e18)
	pledgeP := big.NewInt(1e18)

	// tx options
	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	fmt.Println("========== 1. deploy Role contract ==========")
	// deploy Role with admin
	roleCaller := callconts.NewR(adminAddr, test.AdminSk, txopts)
	roleAddr, _, err := roleCaller.DeployRole(
		test.Foundation,
		test.PrimaryToken,
		pledgeK,
		pledgeP,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Role Address:", roleAddr.Hex())

	fmt.Println("========== 2. deploy PledgePool contract ==========")
	// get RToken address
	rtAddr, err := roleCaller.RToken(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RToken Address:", rtAddr.Hex())

	// deploy PledgePool with admin
	pledgeCaller := callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
	pledgeAddr, _, err := pledgeCaller.DeployPledgePool(
		test.PrimaryToken,
		rtAddr,
		roleAddr,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PledgePool Address:", pledgeAddr.Hex())

	fmt.Println("========== 3. deploy RoleFS contract ==========")
	// caller by admin
	rfsCaller := callconts.NewRFS(adminAddr, test.AdminSk, txopts)
	rfsAddr, _, err := rfsCaller.DeployRoleFS()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RoleFS address: ", rfsAddr.Hex())

	fmt.Println("========== 4. deploy Issuance contract ==========")
	// caller by admin
	issuCaller := callconts.NewIssu(adminAddr, test.AdminSk, txopts)
	issuAddr, _, err := issuCaller.DeployIssuance(rfsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issuance address: ", issuAddr.Hex()) // 0xB15FEDB8017845331b460786fb5129C1Da06f6B1

	fmt.Println("========== 5. call setPI() ==========")
	// call with addrs
	roleCaller.SetPI(
		roleAddr,
		pledgeAddr,
		issuAddr,
		rfsAddr,
	)

}
