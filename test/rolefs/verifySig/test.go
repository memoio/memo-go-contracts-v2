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

// 该测试需花费约14分钟
func main() {
	eth := flag.String("eth", "http://119.147.213.220:8193", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: nil,
		GasLimit: callconts.DefaultGasLimit,
	}

	status := make(chan error)

	// rfs caller for deploy
	rfs := callconts.NewRFS(common.Address{}, adminAddr, test.AdminSk, txopts, ethEndPoint, status)

	fmt.Println("============1. begin test deploy RoleFS contract============")
	rolefsAddr, _, err := rfs.DeployRoleFS()
	if err != nil {
		log.Fatal(err)
	}
	if err = <-status; err != nil {
		log.Fatal(err)
	}
	fmt.Println("RoleFS contract Address:", rolefsAddr.Hex())

	fmt.Println("============2. begin call addOrder ============")

	// rfs caller for call method
	rfs = callconts.NewRFS(rolefsAddr, adminAddr, test.AdminSk, txopts, ethEndPoint, status)
	err = rfs.AddOrder(
		common.Address{},
		common.Address{},
		1,
		1,
		1,
		86400,
		1,
		1,
		1,
		big.NewInt(1),
		nil,
		nil,
		nil)

	if err != nil {
		log.Print("call add order err:", err)
	}

}
