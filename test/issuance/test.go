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

// 仍然需要通过调用rolefs合约中的addOrder等函数，从而触发代币发行、触发Issuance合约中的参数被更改，再次测试getter类函数
func main() {
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	rolefsAddr := common.HexToAddress("0xAAaC6D27153BF52d66Eed127e0321372B2FFF67C") // 测试期间部署的RoleFS合约地址

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	fmt.Println("============1. begin test deploy Issuance contract============")
	issu := callconts.NewIssu(adminAddr, test.AdminSk, txopts)
	issuAddr, _, err := issu.DeployIssuance(rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Issuance contract address is ", issuAddr.Hex())

	fmt.Println("============2. begin test MintLevel============")
	mintLevel, err := issu.MintLevel(issuAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The mintLevel is ", mintLevel)

	fmt.Println("============3. begin test LastMint============")
	lastMint, err := issu.LastMint(issuAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The lastMint is ", lastMint)

	fmt.Println("============4. begin test Price============")
	price, err := issu.Price(issuAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The price is ", price)

	fmt.Println("============5. begin test Size============")
	size, err := issu.Size(issuAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The size is ", size)

	fmt.Println("============6. begin test SpaceTime============")
	spaceTime, err := issu.SpaceTime(issuAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The spaceTime is ", spaceTime)

	fmt.Println("============7. begin test TotalPay============")
	totalPay, err := issu.TotalPay(issuAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The totalPay is ", totalPay)

	fmt.Println("============8. begin test TotalPaid============")
	totalPaid, err := issu.TotalPaid(issuAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The totalPaid is ", totalPaid)

	fmt.Println("============test success!============")
}
