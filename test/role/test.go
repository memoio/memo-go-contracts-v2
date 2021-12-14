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
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	roleAddr := common.HexToAddress("0x0aE0e81C338E9128Aa630Ece410ee73F0CEEc88c")
	// rtokenAddr := common.HexToAddress("0x1c40028BB314D4Ecd3Ead67104fb5D790Fe52F9F")
	// rolefsAddr := common.HexToAddress("0x5c3C6b6aEfaA577548B07871688e109787d7F225")
	// pledgePoolAddr := common.HexToAddress("0x4FD381B8b7f96B8BE2B6044c795A380EE47CF640")
	// issuanceAddr := common.HexToAddress("0x9b2D613F811895fb67BF743A745caD256F6Ce2e0")
	var gIndex uint64 = 1

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	// rechargeMoney := big.NewInt(1e8)
	// start := uint64(time.Now().Unix()) // 当前时间的时间戳
	// end := start + 10
	// size := uint64(10)
	// sprice := big.NewInt(100)
	// 部署Role的参数

	pledgeK := big.NewInt(1e18)
	// pledgeP := big.NewInt(1e18)
	var addrs []common.Address = []common.Address{common.HexToAddress(test.Acc1), common.HexToAddress(test.Acc2), common.HexToAddress(test.Acc3), common.HexToAddress(test.Acc4), common.HexToAddress(test.Acc5)}
	//var sks []string = []string{test.Sk1, test.Sk2, test.Sk3, test.Sk4, test.Sk5}

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("admin balance: ", bal, " in Ethereum")
	for i, addr := range addrs {
		bal = callconts.QueryEthBalance(addr.Hex(), ethEndPoint)
		fmt.Println("acc", i, " balance: ", bal, " in Ethereum")
	}

	// tx options
	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	fmt.Println(">>>> Prepairing balances")
	// 查看测试账户的ERC20代币余额，不足时自动充值
	erc20 := callconts.NewERC20(adminAddr, test.AdminSk, txopts)
	// 确保admin账户的ERC20代币余额充足（至少5 eth）
	bal, err := erc20.BalanceOf(test.PrimaryToken, adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("admin balance in primaryToken is ", bal)
	if bal.Cmp(big.NewInt(test.MoneyTo*5)) < 0 {
		// mintToken
		err = erc20.MintToken(test.PrimaryToken, adminAddr, big.NewInt(test.MoneyTo*5))
		if err != nil {
			log.Fatal(err)
		}
		bal, err = erc20.BalanceOf(test.PrimaryToken, adminAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("after mint, admin balance in primaryToken is ", bal)
	}
	// 确保每个测试账户的ERC20代币余额充足（不少于Pledge值，默认1 eth）
	for i, addr := range addrs {
		bal, err = erc20.BalanceOf(test.PrimaryToken, addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("acc", i, " balance in primaryToken is ", bal)
		if bal.Cmp(pledgeK) < 0 {
			err = erc20.Transfer(test.PrimaryToken, addr, pledgeK) // admin给测试账户转账，用于测试（充值或质押）
			if err != nil {
				log.Fatal(err)
			}
			bal, err = erc20.BalanceOf(test.PrimaryToken, addr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("after transfer, acc", i, " balance in primaryToken is ", bal)
		}
	}

	fmt.Println("============ 1. prepair for testing ============")

	// 准备工作：
	// 部署<Role合约>，获取<RToken合约>地址
	// 部署<RoleFS合约>
	// 注册所有帐号
	// 部署<PledgePool合约>
	// Keeper和Provider质押代币
	// 部署<Issuance合约>
	// 注册Keeper和Provider角色
	// 调用<Role合约>中的CreateGroup函数,自动部署<FileSys合约>
	// 顺序：Role、RoleFS、PledgePool、CreateGroup(FileSys)

	// deploy Role
	r := callconts.NewR(adminAddr, test.AdminSk, txopts)
	/*
		// fmt.Println(">>>> begin deploy Role")

		// 部署Role合约的时候指定使用的ERC20 Token,以及最小的Pledge值
		roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The Role contract address: ", roleAddr.Hex())

		// get RToken
		rtokenAddr, err := r.RToken(roleAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The RToken contract address: ", rtokenAddr.Hex())
		//rtokenAddr := common.HexToAddress("0x081458b892fb2caEb3e4a6234F9183594531b715")

		// deploy RoleFS
		rfs := callconts.NewRFS(adminAddr, test.AdminSk, txopts)
		rolefsAddr, _, err := rfs.DeployRoleFS()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The RoleFS contract address: ", rolefsAddr.Hex())

	*/

	// fmt.Println(">>>> begin register accountss")
	// register accounts
	// acc1: User
	// acc2, acc3, acc4: Keeper
	// acc5: Provider
	/*
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
		fmt.Println("Role indexes is ", rIndexes)
	*/
	rIndexes := []uint64{1, 2, 3, 4, 5}

	// fmt.Println(">>>> begin deploy PledgePool")

	/*
		p := callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
		pledgePoolAddr, _, err := p.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The PledgePool contract address: ", pledgePoolAddr.Hex())
		//pledgePoolAddr := common.HexToAddress("0xfB279b29E437cEE73ac2F2423a3Cfb22060Eb9d7")
	*/

	// get min Pledge for keeper and provider
	pledgek, err := r.PledgeK(roleAddr) // 申请Keeper最少需质押的金额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("min pledge for applying keeper: ", pledgek)
	pledgep, err := r.PledgeP(roleAddr) // 申请Provider最少需质押的金额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("min pledge for applying provider: ", pledgep)

	/*
		fmt.Println(">>>> begin pledge")
		// Keepers pledge minPledgeK for register
		for i, rindex := range rIndexes[1:4] {
			err = toPledge(addrs[i+1], roleAddr, pledgePoolAddr, sks[i+1], rindex, pledgek, txopts)
			if err != nil {
				log.Fatal(err)
			}
		}
		// Provider pledge minPledgeP for register
		err = toPledge(addrs[4], roleAddr, pledgePoolAddr, sks[4], rIndexes[4], pledgep, txopts)
		if err != nil {
			log.Fatal(err)
		}
	*/

	// fmt.Println(">>>> begin deploy Issuance")
	/*
		issu := callconts.NewIssu(adminAddr, test.AdminSk, txopts)

			issuanceAddr, _, err := issu.DeployIssuance(rolefsAddr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("----> The Issuance contract address is: ", issuanceAddr.Hex()) // 0xB15FEDB8017845331b460786fb5129C1Da06f6B1
	*/

	/*
		// 给Role合约指定所有相关合约地址
		r = callconts.NewR(adminAddr, test.AdminSk, txopts)
		err = r.SetPI(roleAddr, pledgePoolAddr, issuanceAddr, rolefsAddr)
		if err != nil {
			log.Fatal(err)
		}
	*/

	/*
		fmt.Println(">>>> begin register keepers")
		// Register Keepers
		callconts.PledgePoolAddr = pledgePoolAddr
		for i, rindex := range rIndexes[1:4] {
			r := callconts.NewR(addrs[i+1], sks[i+1], txopts)
			fmt.Println(addrs[i+1].Hex(), " begin to register Keeper...")
			// 查询admin帐号在PledgePool合约中的ERC20 balance
			bal, err := p.GetBalanceInPPool(pledgePoolAddr, rindex, 0)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("rindex ", rindex, " balance in PledgePool is ", bal)
			// 获取账户信息
			_, _, roleType, index, _, _, err := r.GetRoleInfo(roleAddr, addrs[i+1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("The rindex", rindex, " information in Role contract, roleType:", roleType, " index:", index)

			if roleType == 0 {
				err = r.RegisterKeeper(roleAddr, rindex, []byte("Hello, test"), nil)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		fmt.Println(">>>> begin register provider")
		// Register Provider
		r = callconts.NewR(addrs[4], sks[4], txopts)
		// isActive, isBanned, roleType, index, gIndex, extra
		isActive, _, roleType, index, gIndex, _, err := r.GetRoleInfo(roleAddr, addrs[4])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("provider info: ")
		fmt.Println("isActive:", isActive, " roleType:", roleType, " index:", index, " gIndex:", gIndex)
		if roleType == 0 {
			err = r.RegisterProvider(roleAddr, rIndexes[4], nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	*/

	fmt.Println("INFOS:")
	fmt.Println("rIndexes: ", rIndexes)
	fmt.Println("addrs: ", addrs)
	fmt.Println("gIndex: ", gIndex)

	/*
		fmt.Println(">>>> begin create group")
		// 需要admin先调用CreateGroup,同时将部署FileSys合约
		r = callconts.NewR(adminAddr, test.AdminSk, txopts)
		gIndex, err = r.CreateGroup(roleAddr, rolefsAddr, uint64(0), rIndexes[1:3], 2)
		if err != nil {
			log.Fatal(err)
		}
	*/

	// 获取group信息
	fmt.Println(">>>> getting group info")
	isActive, isBanned, isReady, level, _size, price, fsAddr, err := r.GetGroupInfo(roleAddr, gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The group info- isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr.Hex())

	// ----> The Role contract address:  0x0aE0e81C338E9128Aa630Ece410ee73F0CEEc88c
	// ----> The RToken contract address:  0x1c40028BB314D4Ecd3Ead67104fb5D790Fe52F9F
	// ----> The RoleFS contract address:  0x5c3C6b6aEfaA577548B07871688e109787d7F225
	// ----> The PledgePool contract address:  0x4FD381B8b7f96B8BE2B6044c795A380EE47CF640
	// ----> The Issuance contract address is:  0x9b2D613F811895fb67BF743A745caD256F6Ce2e0

	// roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP)
	// pledgePoolAddr, _, err := p.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
	// issuanceAddr, _, err := issu.DeployIssuance(rolefsAddr)

	fmt.Println("============ 2. test GetAddrGindex ============")

	// get acc address and gIndex by rIndex
	acc, gi, err := r.GetAddrGindex(roleAddr, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("acc: ", acc, " gi: ", gi)
	if acc != addrs[1] {
		log.Fatal("test GetAddrGindex failed: acc address error")
	}
	if gi != gIndex {
		log.Fatal("test GetAddrGindex failed: gIndex error")
	}

	fmt.Println("============test success!============")

}

// approve before pledge
func toPledge(addr, roleAddr, pledgePoolAddr common.Address, sk string, rindex uint64, pledgek *big.Int, txopts *callconts.TxOpts) error {

	// 调用pledge前需要先approve
	erc20 := callconts.NewERC20(addr, sk, txopts)
	err := erc20.Approve(test.PrimaryToken, pledgePoolAddr, pledgek)
	if err != nil {
		return err
	}
	// 查询allowance
	allo, err := erc20.Allowance(test.PrimaryToken, addr, pledgePoolAddr)
	if err != nil {
		return err
	}
	fmt.Println("The allowance of ", addr, " to ", pledgePoolAddr, " is ", allo)

	// 质押
	p := callconts.NewPledgePool(addr, sk, txopts)
	err = p.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, rindex, pledgek, nil)
	if err != nil {
		return err
	}
	return nil
}
