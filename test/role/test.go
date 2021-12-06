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
	pledgeK := big.NewInt(1e18)
	pledgeP := big.NewInt(1e18)

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}
	e := callconts.NewR(adminAddr, test.AdminSk, txopts)

	fmt.Println("============1. begin test deploy Role contract============")
	roleAddr, _, err := e.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Role contract Address:", roleAddr.Hex())
}
