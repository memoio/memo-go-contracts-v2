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
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	acc1Addr := common.HexToAddress(test.Acc1) // 用于注册User
	acc2Addr := common.HexToAddress(test.Acc2) // 用于注册keeper
	acc3Addr := common.HexToAddress(test.Acc5) // 用于注册provider
	accs := []common.Address{acc1Addr, acc2Addr, acc3Addr}
	pledgeK := big.NewInt(1e6)

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")
	for i, acc := range accs {
		bal = callconts.QueryEthBalance(acc.Hex(), ethEndPoint)
		fmt.Println("acc", i, " balance: ", bal, " in Ethereum")
	}

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	// 查询在erc20代币上的余额
	erc20 := callconts.NewERC20(adminAddr, test.AdminSk, txopts)
	bal, err := erc20.BalanceOf(test.PrimaryToken, adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("admin balance in primaryToken is ", bal)
	// 余额不足，自动充值
	if bal.Cmp(big.NewInt(test.MoneyTo)) < 0 {
		// mintToken
		err = erc20.MintToken(test.PrimaryToken, adminAddr, big.NewInt(test.MoneyTo))
		if err != nil {
			log.Fatal(err)
		}
		bal, err = erc20.BalanceOf(test.PrimaryToken, adminAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("after mint, admin balance in primaryToken is ", bal)
	}
	for i, acc := range accs {
		bal, err = erc20.BalanceOf(test.PrimaryToken, acc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("acc", i, " balance in primaryToken is ", bal)
		if bal.Cmp(pledgeK) < 0 {
			err = erc20.Transfer(test.PrimaryToken, acc, pledgeK)
			if err != nil {
				log.Fatal(err)
			}
			bal, err = erc20.BalanceOf(test.PrimaryToken, acc)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("after transfer, acc", i, " balance in primaryToken is ", bal)
		}
	}

	rfs := callconts.NewRFS(adminAddr, test.AdminSk, txopts)

	fmt.Println("============1. begin test deploy RoleFS contract============")
	rolefsAddr, _, err := rfs.DeployRoleFS()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RoleFS contract Address:", rolefsAddr.Hex())

	// 部署RoleFS合约之后，需要调用SetAddr函数赋予合约需要的状态变量信息
	fmt.Println("============2. begin test SetAddr============")
	// 部署Role合约
	r := callconts.NewR(adminAddr, test.AdminSk, txopts)
	roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Role contract address is ", roleAddr.Hex())
	// 获得RToken合约地址
	rTokenAddr, err := r.RToken(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The RToken contract address is ", rTokenAddr.Hex())
	// 部署PledgePool合约
	pp := callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
	pledgePoolAddr, _, err := pp.DeployPledgePool(test.PrimaryToken, rTokenAddr, roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The PledgePool contract address is ", pledgePoolAddr.Hex())
	// 部署Issuance合约
	issu := callconts.NewIssu(adminAddr, test.AdminSk, txopts)
	issuanceAddr, _, err := issu.DeployIssuance(rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Issuance contract address is ", issuanceAddr.Hex())
	// setPI
	err = r.SetPI(roleAddr, pledgePoolAddr, issuanceAddr, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	// keeper注册、质押、申请角色
	r = callconts.NewR(acc2Addr, test.Sk2, txopts)
	err = r.Register(roleAddr, acc2Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	kIndex, err := r.GetRoleIndex(roleAddr, acc2Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("kIndex ", kIndex)
	pp = callconts.NewPledgePool(acc2Addr, test.Sk2, txopts)
	err = pp.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, kIndex, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = r.RegisterKeeper(roleAddr, kIndex, []byte("test"), nil)
	if err != nil {
		log.Fatal(err)
	}
	// createGroup, and deploy FileSys contract
	r = callconts.NewR(adminAddr, test.AdminSk, txopts)
	gIndex, err := r.CreateGroup(roleAddr, rolefsAddr, 0, []uint64{kIndex}, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gIndex ", gIndex)
	// 获取FileSys合约地址
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(roleAddr, gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The FileSys contract address is ", fsAddr.Hex())
	// 给RoleFS合约赋予状态变量值
	err = rfs.SetAddr(rolefsAddr, issuanceAddr, roleAddr, fsAddr, rTokenAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("============3. begin test AddOrder============")
	// user注册、充值
	r = callconts.NewR(acc1Addr, test.Sk1, txopts)
	err = r.Register(roleAddr, acc1Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	uIndex, err := r.GetRoleIndex(roleAddr, acc1Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uIndex ", uIndex)
	err = r.RegisterUser(roleAddr, rTokenAddr, uIndex, gIndex, 0, []byte("test"), nil)
	if err != nil {
		log.Fatal(err)
	}
	// user往FileSys中充值pledgeK，用于存储服务付费
	err = r.Recharge(roleAddr, rTokenAddr, uIndex, 0, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	// provider注册、质押、申请角色、加入group
	r = callconts.NewR(acc3Addr, test.Sk5, txopts)
	err = r.Register(roleAddr, acc3Addr, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("============4. begin test SubOrder============")

	fmt.Println("============5. begin test AddRepair============")

	fmt.Println("============6. begin test SubRepair============")

	fmt.Println("============7. begin test ProWithdraw============")

	fmt.Println("============test success!============")
}
