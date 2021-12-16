package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	callconts "memoContract/callcontracts"
	iface "memoContract/interfaces"
	"memoContract/test"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ethEndPoint  string
	qethEndPoint string
)

func main() {
	eth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	pFast := flag.Bool("f", false, "flag for fast test")
	flag.Parse()

	Fast := *pFast
	ethEndPoint = *eth
	qethEndPoint = *qeth
	_ = qethEndPoint
	callconts.EndPoint = ethEndPoint

	// vars for test
	var (
		roleAddr       common.Address
		rtokenAddr     common.Address
		rolefsAddr     common.Address
		pledgePoolAddr common.Address
		issuanceAddr   common.Address
		gIndex1        uint64
		gIndex2        uint64
	)

	if Fast {
		// read contracts from ini
		config := InitConfig("./contracts.ini")
		// fast test read contract addresses from config file
		roleAddr = common.HexToAddress(config["role"])
		rtokenAddr = common.HexToAddress(config["rtoken"])
		rolefsAddr = common.HexToAddress(config["rolefs"])
		pledgePoolAddr = common.HexToAddress(config["pledgePool"])
		issuanceAddr = common.HexToAddress(config["issuance"])

		// group index start from 1
		gIndex1 = 1
		gIndex2 = 2
	} else {
		_ = roleAddr
		_ = rtokenAddr
		_ = rolefsAddr
		_ = pledgePoolAddr
		_ = issuanceAddr
	}

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)

	// 部署Role的参数
	var (
		rechargeMoney *big.Int
		start         uint64
		end           uint64
		size          uint64
		sprice        *big.Int

		err error
	)
	// fast test don't use them
	if !Fast {
		rechargeMoney = big.NewInt(1e8)
		start = uint64(time.Now().Unix()) // 当前时间的时间戳
		end = start + 10
		size = uint64(10)
		sprice = big.NewInt(100)
	}
	_ = rechargeMoney
	_ = start
	_ = end
	_ = size
	_ = sprice

	pledgeK := big.NewInt(1e18)
	pledgeP := big.NewInt(1e18)
	_ = pledgeP
	var addrs []common.Address = []common.Address{common.HexToAddress(test.Acc1), common.HexToAddress(test.Acc2), common.HexToAddress(test.Acc3), common.HexToAddress(test.Acc4), common.HexToAddress(test.Acc5)}
	var sks []string = []string{test.Sk1, test.Sk2, test.Sk3, test.Sk4, test.Sk5}

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

	if !Fast {
		fmt.Println(">>>> checking balance of admin, 5 eth at least")
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

		fmt.Println(">>>> checking balance of accounts, 1 eth each")
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
	}

	fmt.Println("============ 1. prepair for testing ============")

	// 合约部署顺序：Role、RoleFS、PledgePool、CreateGroup(FileSys)

	r := callconts.NewR(adminAddr, test.AdminSk, txopts)
	// deploy Role
	if Fast {
		fmt.Println("@@ fast test skip deploy Role, use existing address")
	} else {
		fmt.Println(">>>> begin deploy Role")

		// deploy Role
		roleAddr, _, err = r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The Role contract address: ", roleAddr.Hex())

		// get RToken
		rtokenAddr, err = r.RToken(roleAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The RToken contract address: ", rtokenAddr.Hex())

		// deploy RoleFS
		rfs := callconts.NewRFS(adminAddr, test.AdminSk, txopts)
		rolefsAddr, _, err = rfs.DeployRoleFS()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The RoleFS contract address: ", rolefsAddr.Hex())
	}

	var rIndexes = make([]uint64, 5)
	// register accounts： User=[acc1] Keeper=[acc2, acc3, acc4] Provider=[acc5]
	if Fast {
		fmt.Println("@@ fast test skip register accounts, use default role indexes")
		rIndexes = []uint64{1, 2, 3, 4, 5}
	} else {
		fmt.Println(">>>> begin register accounts to get role indexes")
		for i, addr := range addrs {
			fmt.Printf("Role address: %s, register acc: %s\n", roleAddr.Hex(), addr)
			r = callconts.NewR(addr, sks[i], txopts)
			err = r.Register(roleAddr, addr, nil)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("call GetRoleIndex")
			rIndexes[i], err = r.GetRoleIndex(roleAddr, addr)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Role indexes is ", rIndexes)

	var p iface.PledgePoolInfo
	// deploy pledge pool
	if Fast {
		fmt.Println("@@ fast test skip deloy pledge pool, use existing contract")
	} else {
		fmt.Println(">>>> begin deploy PledgePool")
		p = callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
		pledgePoolAddr, _, err = p.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The PledgePool contract address: ", pledgePoolAddr.Hex())
	}

	// get pledge
	var pledgek *big.Int
	var pledgep *big.Int
	if Fast {
		fmt.Println("@@ fast test skip get Pledge, use default: 1 eth")
	} else {
		fmt.Println(">>>> Getting min pledge from Role")
		// get min Pledge for keeper and provider
		pledgek, err = r.PledgeK(roleAddr) // 申请Keeper最少需质押的金额
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("min pledge for applying keeper: ", pledgek)
		pledgep, err = r.PledgeP(roleAddr) // 申请Provider最少需质押的金额
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("min pledge for applying provider: ", pledgep)
	}

	// pledge
	if Fast {
		fmt.Println("@@ fast test skip Pledge, use existing")
	} else {
		fmt.Println(">>>> begin keepers pledge")

		for i, rindex := range rIndexes[1:4] {
			fmt.Println("params:")
			fmt.Println(
				" add:", addrs[i+1],
				" roleAddr:", roleAddr,
				" pledgePoolAddr:", pledgePoolAddr,
				" sk:", sks[i+1],
				" rindex:", rindex,
				" pledgek:", pledgek,
				" txopts:", txopts,
			)
			err = toPledge(addrs[i+1], roleAddr, pledgePoolAddr, sks[i+1], rindex, pledgek, txopts)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(">>>> begin provider pledge")
		err = toPledge(addrs[4], roleAddr, pledgePoolAddr, sks[4], rIndexes[4], pledgep, txopts)
		if err != nil {
			log.Fatal(err)
		}
	}

	// deploy Issuance
	if Fast {
		fmt.Println("@@ fast test skip deploy Issuance")
	} else {
		fmt.Println(">>>> begin deploy Issuance")

		issu := callconts.NewIssu(adminAddr, test.AdminSk, txopts)

		issuanceAddr, _, err = issu.DeployIssuance(rolefsAddr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----> The Issuance contract address is: ", issuanceAddr.Hex()) // 0xB15FEDB8017845331b460786fb5129C1Da06f6B1

		// 给Role合约指定所有相关合约地址
		r = callconts.NewR(adminAddr, test.AdminSk, txopts)
		err = r.SetPI(roleAddr, pledgePoolAddr, issuanceAddr, rolefsAddr)
		if err != nil {
			log.Fatal(err)
		}
	}

	// register keepers and provider role
	if Fast {
		fmt.Println("@@ fast test skip register roles, use existing roles")
	} else {
		fmt.Println(">>>> begin register keepers")
		callconts.PledgePoolAddr = pledgePoolAddr
		for i, rindex := range rIndexes[1:4] {
			r := callconts.NewR(addrs[i+1], sks[i+1], txopts)
			fmt.Println(addrs[i+1].Hex(), " begin to register Keeper...")

			// 查询admin帐号在PledgePool合约中的ERC20 balance
			fmt.Println("pledgePoolAddr: ", pledgePoolAddr, " rindex:", rindex)
			p = callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
			bal, err := p.GetBalanceInPPool(pledgePoolAddr, rindex, 0)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("rindex ", rindex, " balance in PledgePool is ", bal)

			// 获取角色信息
			_, _, roleType, index, _, _, err := r.GetRoleInfo(roleAddr, addrs[i+1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("The rindex", rindex, " information in Role contract, roleType:", roleType, " index:", index)

			// register keeper role
			if roleType == 0 {
				err = r.RegisterKeeper(roleAddr, rindex, []byte("Hello, test"), nil)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		fmt.Println(">>>> begin register provider ")
		r = callconts.NewR(addrs[4], sks[4], txopts)
		// isActive, isBanned, roleType, index, gIndex, extra
		isActive, _, roleType, index, gIndex, _, err := r.GetRoleInfo(roleAddr, addrs[4])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("provider info: ")
		fmt.Println("isActive:", isActive, " roleType:", roleType, " index:", index, " gIndex:", gIndex)

		// register provider
		if roleType == 0 {
			err = r.RegisterProvider(roleAddr, rIndexes[4], nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// create groups
	if Fast {
		fmt.Println("@@ fast test with existing group index 1,2")
		gIndex1 = 1
		gIndex2 = 2
	} else {
		// create group, keeper1 keeper2 in group1, keeper3 in group2
		fmt.Println(">>>> begin create group 1")
		// 需要admin先调用CreateGroup,同时将部署FileSys合约
		r = callconts.NewR(adminAddr, test.AdminSk, txopts)
		gIndex1, err = r.CreateGroup(roleAddr, rolefsAddr, uint64(0), rIndexes[1:3], 2)
		if err != nil {
			log.Fatal(err)
		}

		if gIndex1 != 1 {
			log.Fatal("gIndex1 should be 1")
		}

		// test create group 2 with inActive, level=2, but only 1 keeper involved
		fmt.Println(">>>> begin create group 2")
		r = callconts.NewR(adminAddr, test.AdminSk, txopts)
		gIndex2, err = r.CreateGroup(roleAddr, rolefsAddr, uint64(0), rIndexes[3:4], 2)
		if err != nil {
			log.Fatal(err)
		}

		if gIndex2 != 2 {
			log.Fatal("gIndex2 should be 2")
		}
	}

	// show info
	{
		fmt.Println(">>>> Show Info:")
		fmt.Println("rIndexes: ", rIndexes)
		fmt.Println("addrs: ", addrs)
		fmt.Println("gIndex1: ", gIndex1)
		fmt.Println("gIndex2: ", gIndex2)
		fmt.Println("role: ", roleAddr)
		fmt.Println("rtoken: ", rtokenAddr)
		fmt.Println("rolefs: ", rolefsAddr)
		fmt.Println("pledgePool: ", pledgePoolAddr)
		fmt.Println("issuance: ", issuanceAddr)
	}

	// test group1
	fmt.Println(">>>> getting group1 info")
	isActive, isBanned, isReady, level, _size, price, fsAddr, err := r.GetGroupInfo(roleAddr, gIndex1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The group info- isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr.Hex())

	if !isActive {
		log.Fatal("group1 should be active")
	}

	// test group2
	fmt.Println(">>>> getting group2 info")
	isActive, isBanned, isReady, level, _size, price, fsAddr, err = r.GetGroupInfo(roleAddr, gIndex2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The group info- isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr.Hex())

	if isActive {
		log.Fatal("group2 should be inactive")
	}

	fmt.Println("============ 2. test GetAddrGindex ============")

	// get acc address and gIndex by rIndex

	fmt.Println("test role 2")
	acc, gIndex, err := r.GetAddrGindex(roleAddr, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("acc: ", acc, " gIndex: ", gIndex)

	if acc.String() != test.Acc2 {
		log.Fatal("acc address error")
	}
	if gIndex != 1 {
		log.Fatal("gIndex error")
	}

	fmt.Println("test role 4")
	acc, gIndex, err = r.GetAddrGindex(roleAddr, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("acc: ", acc, " gIndex: ", gIndex)

	if acc.String() != test.Acc4 {
		log.Fatal("acc address error")
	}
	if gIndex != 2 {
		log.Fatal("gIndex error")
	}

	fmt.Println("============ 3. test GetGroupsNum ============")
	// test GetGroupsNum
	gNum, err := r.GetGroupsNum(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("group number:", gNum)

	if gNum != 2 {
		log.Fatal("GetGroupsNum test failed, group number should be 2")
	}

	fmt.Println("============ 4. test PledgePool ============")
	pp, err := r.PledgePool(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pledge pool:", pp)
	if pp != pledgePoolAddr {
		log.Fatal("test PledgePool failed")
	}

	fmt.Println("============ 5. test Foundation ============")
	fd, err := r.Foundation(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("foundation:", fd)
	if fd != test.Foundation {
		log.Fatal("test Foundation failed")
	}

	fmt.Println("============ 6. test PledgeK ============")
	plk, err := r.PledgeK(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("plk:", plk)
	if plk.Cmp(pledgeK) != 0 {
		log.Fatal("test PledgeK failed")
	}

	fmt.Println("============ 7. test PledgeP ============")
	plp, err := r.PledgeP(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pk:", plp)
	if plp.Cmp(pledgeP) != 0 {
		log.Fatal("test PledgeK failed")
	}

	fmt.Println("============ 8. test RToken ============")
	rt, err := r.RToken(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rt:", rt)
	if rt != rtokenAddr {
		log.Fatal("test RToken failed")
	}

	fmt.Println("============ 9. test Issuance ============")
	is, err := r.Issuance(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is:", is)
	if is != issuanceAddr {
		log.Fatal("test Issuance failed")
	}

	fmt.Println("============ 10. test Rolefs ============")
	rfs, err := r.Rolefs(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rfs:", rfs)
	if rfs != rolefsAddr {
		log.Fatal("test Rolefs failed")
	}

	fmt.Println("============ 11. test GetAddrsNum ============")
	an, err := r.GetAddrsNum(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("an:", an)
	if an != 5 {
		log.Fatal("test GetAddrsNum failed")
	}

	fmt.Println("============ 12. test GetAddr ============")
	ga, err := r.GetAddr(roleAddr, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ga:", ga)
	if ga.String() != test.Acc1 {
		log.Fatal("test GetAddr failed")
	}

	fmt.Println("============ 13. test GetRoleInfo ============")
	isActive, isBanned, roleType, index, gIndex, extra, err := r.GetRoleInfo(roleAddr, common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("role info: isActive, isBanned, roleType, index, gIndex, extra")
	fmt.Println("role info: ", isActive, isBanned, roleType, index, gIndex, extra)

	fmt.Println("============ 14. test GetGroupsNum ============")
	gn, err := r.GetGroupsNum(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gn:", gn)
	if gn != 2 {
		log.Fatal("test GetGroupsNum failed")
	}

	fmt.Println("============ 15. test GetGKNum ============")
	gk, err := r.GetGKNum(roleAddr, 0) // bug
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gk:", gk)
	if gk != 2 {
		log.Fatal("test GetGKNum failed")
	}

	fmt.Println("============ 16. test GetGPNum ============")
	gp, err := r.GetGPNum(roleAddr, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gp:", gp)
	if gp != 2 {
		log.Fatal("test GetGPNum failed")
	}

	fmt.Println("============test success!============")

	// if success and not fast test, save new config into config file
	if !Fast {
		cfg := fmt.Sprintf("role=%s\nrtoken=%s\nrolefs=%s\npledgePool=%s\nissuance=%s", roleAddr, rtokenAddr, rolefsAddr, pledgePoolAddr, issuanceAddr)
		SaveConfig(cfg)
	}
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

//读取key=value类型的配置文件
func InitConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

func SaveConfig(cfg string) error {
	err := ioutil.WriteFile("contracts.ini", []byte(cfg), 0666)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
