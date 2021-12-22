package main

func main() {}

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"math/big"
// 	callconts "memoContract/callcontracts"
// 	iface "memoContract/interfaces"
// 	"memoContract/test"
// 	"os"
// 	"strings"

// 	"github.com/ethereum/go-ethereum/common"
// )

// // package level vars
// var (
// 	ethEndPoint string
// 	//qethEndPoint string

// 	// tx options
// 	txopts *callconts.TxOpts

// 	// addrs and sks of test accounts
// 	addrs []common.Address
// 	sks   []string

// 	// admin account addresses
// 	adminAddr common.Address

// 	// contract addresses
// 	roleAddr       common.Address
// 	rtokenAddr     common.Address
// 	rolefsAddr     common.Address
// 	pledgePoolAddr common.Address
// 	issuAddr       common.Address

// 	// pledge values
// 	pledgeK *big.Int
// 	pledgeP *big.Int
// 	oneEth  *big.Int
// 	fiveEth *big.Int

// 	fast bool

// 	// Role caller for admin
// 	rAdmin iface.RoleInfo
// 	// Role caller for user
// 	rUser iface.RoleInfo
// 	// Role caller for normal account
// 	rAcc iface.RoleInfo

// 	// filesys address
// 	fsAddr1 common.Address
// 	fsAddr2 common.Address

// 	// group index
// 	gIndex1 uint64
// 	gIndex2 uint64

// 	config map[string]string

// 	err error
// )

// // init package level vars
// func init() {
// 	pledgeK = big.NewInt(1e18)
// 	pledgeP = big.NewInt(1e18)
// 	oneEth = big.NewInt(1e18)
// 	fiveEth = big.NewInt((5 * 1e18))

// 	txopts = &callconts.TxOpts{
// 		Nonce:    nil,
// 		GasPrice: big.NewInt(callconts.DefaultGasPrice),
// 		GasLimit: callconts.DefaultGasLimit,
// 	}

// 	addrs = []common.Address{
// 		common.HexToAddress(test.Acc1),
// 		common.HexToAddress(test.Acc2),
// 		common.HexToAddress(test.Acc3),
// 		common.HexToAddress(test.Acc4),
// 		common.HexToAddress(test.Acc5),
// 	}
// 	sks = []string{
// 		test.Sk1,
// 		test.Sk2,
// 		test.Sk3,
// 		test.Sk4,
// 		test.Sk5,
// 	}

// 	adminAddr = common.HexToAddress(test.AdminAddr)

// 	rAdmin = callconts.NewR(adminAddr, test.AdminSk, txopts)
// 	rUser = callconts.NewR(common.HexToAddress(test.Acc1), test.Sk1, txopts)
// }

// func main() {

// 	peth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;") //dev网
// 	//qeth = flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
// 	pfast := flag.Bool("f", false, "flag for fast test")

// 	flag.Parse()

// 	//qethEndPoint = *qeth
// 	fast = *pfast
// 	ethEndPoint = *peth
// 	callconts.EndPoint = *peth

// 	CheckEthBalance()
// 	CheckERC20Balance()

// 	fmt.Println("============ 1. prepair for testing ============")
// 	err = PrePaire()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("============ 2. test GetAddrGindex ============")

// 	// get acc address and gIndex by rIndex
// 	fmt.Println("test role 2")
// 	acc, gIndex, err := rAdmin.GetAddrGindex(roleAddr, 2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("acc: ", acc, " gIndex: ", gIndex)

// 	if acc.String() != test.Acc2 {
// 		log.Fatal("acc address error")
// 	}
// 	if gIndex != 1 {
// 		log.Fatal("gIndex error")
// 	}

// 	fmt.Println("test role 4")
// 	acc, gIndex, err = rAdmin.GetAddrGindex(roleAddr, 4)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("acc: ", acc, " gIndex: ", gIndex)

// 	if acc.String() != test.Acc4 {
// 		log.Fatal("acc address error")
// 	}
// 	if gIndex != 2 {
// 		log.Fatal("gIndex error")
// 	}

// 	fmt.Println("============ 3. test GetGroupsNum ============")
// 	// test GetGroupsNum
// 	gNum, err := rAdmin.GetGroupsNum(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("group number:", gNum)

// 	if gNum != 2 {
// 		log.Fatal("GetGroupsNum test failed, group number should be 2")
// 	}

// 	fmt.Println("============ 4. test PledgePool ============")
// 	pp, err := rAdmin.PledgePool(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("pledge pool:", pp)
// 	if pp != pledgePoolAddr {
// 		log.Fatal("test PledgePool failed")
// 	}

// 	fmt.Println("============ 5. test Foundation ============")
// 	fd, err := rAdmin.Foundation(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("foundation:", fd)
// 	if fd != test.Foundation {
// 		log.Fatal("test Foundation failed")
// 	}

// 	fmt.Println("============ 6. test RToken ============")
// 	rt, err := rAdmin.RToken(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("rt:", rt)
// 	if rt != rtokenAddr {
// 		log.Fatal("test RToken failed")
// 	}

// 	fmt.Println("============ 7. test Issuance ============")
// 	is, err := rAdmin.Issuance(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("is:", is)
// 	if is != issuAddr {
// 		log.Fatal("test Issuance failed")
// 	}

// 	fmt.Println("============ 8. test Rolefs ============")
// 	rfs, err := rAdmin.Rolefs(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("rfs:", rfs)
// 	if rfs != rolefsAddr {
// 		log.Fatal("test Rolefs failed")
// 	}

// 	fmt.Println("============ 9. test GetAddrsNum ============")
// 	//
// 	if fast {
// 		fmt.Println("@@ skipped for fast testing")
// 	} else {
// 		an, err := rAdmin.GetAddrsNum(roleAddr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("an:", an)
// 		if an != 5 {
// 			log.Fatal("test GetAddrsNum failed")
// 		}
// 	}

// 	fmt.Println("============ 10. test GetAddr ============")

// 	ga, err := rAdmin.GetAddr(roleAddr, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("ga:", ga)
// 	if ga.String() != test.Acc1 {
// 		log.Fatal("test GetAddr failed")
// 	}

// 	fmt.Println("============ 11. test GetRoleInfo ============")
// 	isActive, isBanned, roleType, index, gIndex, extra, err := rAdmin.GetRoleInfo(roleAddr, common.HexToAddress(test.Acc1))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("role info: isActive, isBanned, roleType, index, gIndex, extra")
// 	fmt.Println("role info: ", isActive, isBanned, roleType, index, gIndex, extra)

// 	fmt.Println("============ 12. test GetGroupsNum ============")
// 	gn, err := rAdmin.GetGroupsNum(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("gn:", gn)
// 	if gn != 2 {
// 		log.Fatal("test GetGroupsNum failed")
// 	}

// 	fmt.Println("============ 13. test GetGKNum ============")
// 	gk, err := rAdmin.GetGKNum(roleAddr, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("gk:", gk)
// 	if gk != 2 {
// 		log.Fatal("test GetGKNum failed")
// 	}

// 	fmt.Println("============ 14. test GetGPNum ============")
// 	gu, gp, err := rAdmin.GetGUPNum(roleAddr, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("gu:", gu, "gp:", gp)

// 	// add a provider only when no provider in group
// 	// add role index 5 into group 1 as a provider
// 	if gp == 0 {
// 		fmt.Println("============ 15. test AddProviderToGroup ============")
// 		// caller must be provider himself to avoid sign
// 		rForProvider := callconts.NewR(common.HexToAddress(test.Acc5), test.Sk5, txopts)
// 		err := rForProvider.AddProviderToGroup(roleAddr, 5, 1, nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// output:
// 		// 2021/12/16 14:14:40 begin AddProviderToGroup in Role contract...
// 		// 2021/12/16 14:14:41 Check Tx hash: 0x6e5b639d01fc6432d8daa59d0d149fc0d62a714bbb23d99cf39e48c31b544461 nonce: 84 gasPrice: 200
// 		// 2021/12/16 14:14:41 waiting for miner...
// 		// 2021/12/16 14:14:56 GasUsed: 79781 CumulativeGasUsed: 79781
// 		// 2021/12/16 14:14:56 AddProviderToGroup in Role has been successful!
// 	} else {
// 		fmt.Println(">>>> role index 5 already exists in group 1, skip AddProviderToGroup")
// 	}

// 	fmt.Println("============ 16. test GetGroupK ============")
// 	// group 1 , keeper[1] == 3
// 	ggk, err := rAdmin.GetGroupK(roleAddr, 1, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("ggk:", ggk)
// 	if ggk != 3 {
// 		log.Fatal("test GetGroupK failed")
// 	}

// 	fmt.Println("============ 17. test GetGroupP ============")
// 	// group 1 , provider[0] == 5
// 	ggp, err := rAdmin.GetGroupP(roleAddr, 1, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("ggp:", ggp)
// 	if ggp != 5 {
// 		log.Fatal("test GetGroupP failed")
// 	}

// 	fmt.Println("============ 18. test PledgeK,PledgeP,SetPledgeMoney ============")

// 	// show old
// 	oldPK, err := rAdmin.PledgeK(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("pledgeK before set: ", oldPK)
// 	oldPP, err := rAdmin.PledgeP(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("pledgeP before set: ", oldPP)

// 	// plus pledge by 1 eth
// 	err = rAdmin.SetPledgeMoney(roleAddr, new(big.Int).Add(oldPK, oneEth), new(big.Int).Add(oldPP, oneEth))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// show new
// 	newPK, err := rAdmin.PledgeK(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("pledgeK after set: ", newPK)
// 	newPP, err := rAdmin.PledgeP(roleAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("pledgeP after set: ", newPP)

// 	if new(big.Int).Sub(newPK, oldPK).Cmp(oneEth) != 0 {
// 		log.Fatal("test set PledgeK failed")
// 	}

// 	if new(big.Int).Sub(newPP, oldPP).Cmp(oneEth) != 0 {
// 		log.Fatal("test set PledgeP failed")
// 	}

// 	fmt.Println("============ 19. test Recharge ============")

// 	// to assure balance is enough
// 	CheckERC20Balance()

// 	// check old balance of user in erc20
// 	userAddr := common.HexToAddress(test.Acc1)
// 	erc20 := callconts.NewERC20(userAddr, test.Sk1, txopts)

// 	oldBalErc20, err := erc20.BalanceOf(test.PrimaryToken, userAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> old balance in erc20 is: ", oldBalErc20)

// 	// check old balance of user in fs
// 	fs := callconts.NewFileSys(adminAddr, test.AdminSk, txopts)
// 	oldBalFs, tmp, err := fs.GetBalance(fsAddr1, 1, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> old balance in fs is: ", oldBalFs, " tmp:", tmp)

// 	// user approve before recharge
// 	err = erc20.Approve(test.PrimaryToken, fsAddr1, oneEth)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// 查询user allowance
// 	allo, err := erc20.Allowance(test.PrimaryToken, userAddr, fsAddr1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The allowance of ", userAddr, " to ", fsAddr1, " is ", allo)

// 	// user recharge 1 eth into fileSys
// 	usrIndex, err := rUser.GetRoleIndex(roleAddr, userAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("user index:", usrIndex)

// 	// query token index
// 	rtCaller := callconts.NewRT(userAddr, test.Sk1, txopts)
// 	tIndex, ok, err := rtCaller.GetTI(rtokenAddr, test.PrimaryToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if !ok {
// 		log.Fatal("GetTI return invalid")
// 	}
// 	fmt.Println("token index:", tIndex)

// 	// user(acc1) recharge 1 eth into fs, with primary token
// 	err = rUser.Recharge(roleAddr, rtokenAddr, 1, 0, oneEth, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("ggp:", ggp)
// 	if ggp != 5 {
// 		log.Fatal("test Recharge failed")
// 	}

// 	// check new balance of user in erc20
// 	newBalErc20, err := erc20.BalanceOf(test.PrimaryToken, userAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> new balance in erc20 is: ", newBalErc20)

// 	// check new balance of user in fs
// 	newBalFs, tmp, err := fs.GetBalance(fsAddr1, 1, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> new balance in fs is: ", newBalFs, " tmp:", tmp)

// 	// test
// 	if new(big.Int).Sub(oldBalErc20, newBalErc20).Cmp(oneEth) != 0 {
// 		log.Fatal("test Recharge failed, erc20 balance error")
// 	}
// 	if new(big.Int).Sub(newBalFs, oldBalFs).Cmp(oneEth) != 0 {
// 		log.Fatal("test Recharge failed, fs balance error")
// 	}

// 	fmt.Println("============ 20. test WithdrawFromFs ============")

// 	// to assure balance is enough
// 	CheckERC20Balance()

// 	// check old balance of user in erc20
// 	oldBalErc20, err = erc20.BalanceOf(test.PrimaryToken, userAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> old balance in erc20 is: ", oldBalErc20)

// 	// query user balance before withdraw
// 	fs = callconts.NewFileSys(adminAddr, test.AdminSk, txopts)
// 	oldBalFs, tmp, err = fs.GetBalance(fsAddr1, 1, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> old balance in erc20 is: ", oldBalFs, " tmp:", tmp)

// 	// withdraw
// 	userAddr = common.HexToAddress(test.Acc1)
// 	rUser = callconts.NewR(userAddr, test.Sk1, txopts)
// 	err = rUser.WithdrawFromFs(roleAddr, rtokenAddr, 1, 0, oneEth, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// check new balance of user in erc20
// 	newBalErc20, err = erc20.BalanceOf(test.PrimaryToken, userAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> new balance in erc20 is: ", newBalErc20)

// 	// query new balance in fs after withdraw
// 	newBalFs, tmp, err = fs.GetBalance(fsAddr1, 1, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(">>>> old balance in erc20 is: ", newBalFs, " tmp:", tmp)

// 	// test
// 	if new(big.Int).Sub(newBalErc20, oldBalErc20).Cmp(oneEth) != 0 {
// 		log.Fatal("test Withdraw failed, erc20 balance error")
// 	}
// 	if new(big.Int).Sub(oldBalFs, newBalFs).Cmp(oneEth) != 0 {
// 		log.Fatal("test WithdrawFromFs failed")
// 	}

// 	fmt.Println("============ 21. test SignForRegister ============")

// 	fmt.Println(">>>> test call register by other account")

// 	if fast {
// 		fmt.Println("@@ skipped for fast testing")
// 	} else {
// 		// signed by acc's sk
// 		sig, err := callconts.SignForRegister(adminAddr, test.Sk6)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// other account call register for regist acc6
// 		rOther := callconts.NewR(adminAddr, test.AdminSk, txopts)
// 		err = rOther.Register(roleAddr, common.HexToAddress(test.Acc6), sig)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// get role info
// 		fmt.Println("call GetRoleIndex")
// 		rAcc = callconts.NewR(common.HexToAddress(test.Acc6), test.Sk6, txopts)
// 		rIndex, err := rAcc.GetRoleIndex(roleAddr, common.HexToAddress(test.Acc6))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("rIndex:", rIndex)
// 		// this is the 6th registered account
// 		if rIndex != 6 {
// 			log.Fatal("test SignForRegister failed, rIndex error")
// 		}
// 	}

// 	fmt.Println("============test success!============")

// 	// if success and not fast test, save new config into config file
// 	if !fast {
// 		cfg := fmt.Sprintf("role=%s\nrtoken=%s\nrolefs=%s\npledgePool=%s\nissuance=%s", roleAddr, rtokenAddr, rolefsAddr, pledgePoolAddr, issuAddr)
// 		SaveConfig(cfg)
// 	}
// }

// // approve before pledge
// func toPledge(addr, roleAddr, pledgePoolAddr common.Address, sk string, rindex uint64, pledgek *big.Int, txopts *callconts.TxOpts) error {

// 	// 调用pledge前需要先approve
// 	erc20 := callconts.NewERC20(addr, sk, txopts)
// 	err := erc20.Approve(test.PrimaryToken, pledgePoolAddr, pledgek)
// 	if err != nil {
// 		return err
// 	}
// 	// 查询allowance
// 	allo, err := erc20.Allowance(test.PrimaryToken, addr, pledgePoolAddr)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("The allowance of ", addr, " to ", pledgePoolAddr, " is ", allo)

// 	// 质押
// 	p := callconts.NewPledgePool(addr, sk, txopts)
// 	err = p.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, rindex, pledgek, nil)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // read contract addresses from config file
// func ReadConfig(path string) map[string]string {
// 	config := make(map[string]string)

// 	f, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	r := bufio.NewReader(f)
// 	for {
// 		b, _, err := r.ReadLine()
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			panic(err)
// 		}
// 		s := strings.TrimSpace(string(b))
// 		index := strings.Index(s, "=")
// 		if index < 0 {
// 			continue
// 		}
// 		key := strings.TrimSpace(s[:index])
// 		if len(key) == 0 {
// 			continue
// 		}
// 		value := strings.TrimSpace(s[index+1:])
// 		if len(value) == 0 {
// 			continue
// 		}
// 		config[key] = value
// 	}
// 	return config
// }

// // save contract addresses into config file if test success
// func SaveConfig(cfg string) error {
// 	err := ioutil.WriteFile("contracts.ini", []byte(cfg), 0666)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return nil
// }

// // check accounts' eth balance, used for sending tx, recharge if not enough
// func CheckEthBalance() {

// 	fmt.Println(">>>> checking eth balance")

// 	ethBal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
// 	fmt.Printf("admin balance: %x in Ethereum\n", ethBal)
// 	for i, acc := range addrs {
// 		ethBal = callconts.QueryEthBalance(acc.Hex(), ethEndPoint)
// 		fmt.Println("acc", i, " balance: ", ethBal, " in Ethereum")
// 	}
// }

// // check each account's balance, increase it if not enough
// func CheckERC20Balance() {

// 	fmt.Println(">>>> checking balance of admin, 5 eth at least")

// 	// 查看测试账户的ERC20代币余额，不足时自动充值
// 	erc20 := callconts.NewERC20(adminAddr, test.AdminSk, txopts)

// 	// 确保admin账户的ERC20代币余额充足（至少5 eth）
// 	bal, err := erc20.BalanceOf(test.PrimaryToken, adminAddr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("admin balance in primaryToken is ", bal)

// 	if bal.Cmp(fiveEth) < 0 {
// 		// mintToken
// 		err = erc20.MintToken(test.PrimaryToken, adminAddr, fiveEth)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		bal, err = erc20.BalanceOf(test.PrimaryToken, adminAddr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("after mint, admin balance in primaryToken is ", bal)
// 	}

// 	fmt.Println(">>>> checking balance of accounts, 1 eth each")
// 	// 确保每个测试账户的ERC20代币余额充足（不少于Pledge值，默认1 eth）
// 	for i, addr := range addrs {
// 		bal, err = erc20.BalanceOf(test.PrimaryToken, addr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("acc", i, " balance in primaryToken is ", bal)
// 		if bal.Cmp(oneEth) < 0 {
// 			err = erc20.Transfer(test.PrimaryToken, addr, oneEth) // admin给测试账户转账，用于测试（充值或质押）
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			bal, err = erc20.BalanceOf(test.PrimaryToken, addr)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			fmt.Println("after transfer, acc", i, " balance in primaryToken is ", bal)
// 		}
// 	}
// }

// // prepare everything before starting test
// func PrePaire() (err error) {

// 	// read contract addresses from config file for fast test
// 	if fast {
// 		fmt.Println(">>>> Fast test read contract addresses from config file")

// 		config = ReadConfig("./contracts.ini")

// 		roleAddr = common.HexToAddress(config["role"])
// 		rtokenAddr = common.HexToAddress(config["rtoken"])
// 		rolefsAddr = common.HexToAddress(config["rolefs"])
// 		pledgePoolAddr = common.HexToAddress(config["pledgePool"])
// 		issuAddr = common.HexToAddress(config["issuance"])
// 	}

// 	// deploy: Role、RoleFS、PledgePool、CreateGroup(FileSys)

// 	// deploy Role
// 	if fast {
// 		fmt.Println("@@ fast test skip deploy Role, use existing address")
// 	} else {
// 		fmt.Println(">>>> begin deploy Role")

// 		// deploy Role
// 		roleAddr, _, err = rAdmin.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("----> The Role contract address: ", roleAddr.Hex())

// 		// get RToken
// 		rtokenAddr, err = rAdmin.RToken(roleAddr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("----> The RToken contract address: ", rtokenAddr.Hex())

// 		// deploy RoleFS
// 		rfsAdmin := callconts.NewRFS(adminAddr, test.AdminSk, txopts)
// 		rolefsAddr, _, err = rfsAdmin.DeployRoleFS()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("----> The RoleFS contract address: ", rolefsAddr.Hex())
// 	}

// 	var rIndexes = make([]uint64, 5)
// 	// register accounts： User=[acc1] Keeper=[acc2, acc3, acc4] Provider=[acc5]
// 	if fast {
// 		fmt.Println("@@ fast test skip register accounts, use default role indexes")
// 		rIndexes = []uint64{1, 2, 3, 4, 5}
// 	} else {
// 		fmt.Println(">>>> begin register accounts to get role indexes")
// 		for i, addr := range addrs {
// 			fmt.Printf("Role address: %s, register acc: %s\n", roleAddr.Hex(), addr)
// 			rAcc = callconts.NewR(addr, sks[i], txopts)
// 			err = rAcc.Register(roleAddr, addr, nil)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			fmt.Println("call GetRoleIndex")
// 			rIndexes[i], err = rAcc.GetRoleIndex(roleAddr, addr)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 	}

// 	fmt.Println("Role indexes is ", rIndexes)

// 	var p iface.PledgePoolInfo
// 	// deploy pledge pool
// 	if fast {
// 		fmt.Println("@@ fast test skip deloy pledge pool, use existing contract")
// 	} else {
// 		fmt.Println(">>>> begin deploy PledgePool")
// 		p = callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
// 		pledgePoolAddr, _, err = p.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("----> The PledgePool contract address: ", pledgePoolAddr.Hex())
// 	}

// 	// get pledge
// 	var pledgek *big.Int
// 	var pledgep *big.Int
// 	if fast {
// 		fmt.Println("@@ fast test skip get Pledge, use default: 1 eth")
// 	} else {
// 		fmt.Println(">>>> Getting min pledge from Role")
// 		// get min Pledge for keeper and provider
// 		pledgek, err = rAdmin.PledgeK(roleAddr) // 申请Keeper最少需质押的金额
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("min pledge for applying keeper: ", pledgek)
// 		pledgep, err = rAdmin.PledgeP(roleAddr) // 申请Provider最少需质押的金额
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("min pledge for applying provider: ", pledgep)
// 	}

// 	// deploy Issuance
// 	if fast {
// 		fmt.Println("@@ fast test skip deploy Issuance")
// 	} else {
// 		fmt.Println(">>>> begin deploy Issuance")

// 		issu := callconts.NewIssu(adminAddr, test.AdminSk, txopts)

// 		issuAddr, _, err = issu.DeployIssuance(rolefsAddr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("----> The Issuance contract address is: ", issuAddr.Hex()) // 0xB15FEDB8017845331b460786fb5129C1Da06f6B1

// 		// 给Role合约指定所有相关合约地址
// 		err = rAdmin.SetPI(roleAddr, pledgePoolAddr, issuAddr, rolefsAddr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// keeper and provider pledge
// 	if fast {
// 		fmt.Println("@@ fast test skip Pledge, use existing")
// 	} else {
// 		fmt.Println(">>>> begin keepers pledge")

// 		for i, rindex := range rIndexes[1:4] {
// 			fmt.Println("params:")
// 			fmt.Println(
// 				" add:", addrs[i+1],
// 				" roleAddr:", roleAddr,
// 				" pledgePoolAddr:", pledgePoolAddr,
// 				" sk:", sks[i+1],
// 				" rindex:", rindex,
// 				" pledgek:", pledgek,
// 				" txopts:", txopts,
// 			)
// 			err = toPledge(addrs[i+1], roleAddr, pledgePoolAddr, sks[i+1], rindex, pledgek, txopts)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 		fmt.Println(">>>> begin provider pledge")
// 		err = toPledge(addrs[4], roleAddr, pledgePoolAddr, sks[4], rIndexes[4], pledgep, txopts)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// keepers and provider register role
// 	if fast {
// 		fmt.Println("@@ fast test skip register roles, use existing roles")
// 	} else {
// 		fmt.Println(">>>> begin register keepers")

// 		callconts.PledgePoolAddr = pledgePoolAddr
// 		for i, rindex := range rIndexes[1:4] {
// 			rKeeper := callconts.NewR(addrs[i+1], sks[i+1], txopts)
// 			fmt.Println(addrs[i+1].Hex(), " begin to register Keeper...")

// 			// query admin's ERC20 balance in pledgepool
// 			fmt.Println("pledgePoolAddr: ", pledgePoolAddr, " rindex:", rindex)
// 			p = callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
// 			bal, err := p.GetBalanceInPPool(pledgePoolAddr, rindex, 0)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			fmt.Println("rindex ", rindex, " balance in PledgePool is ", bal)

// 			// get keeper role info
// 			_, _, roleType, index, _, _, err := rKeeper.GetRoleInfo(roleAddr, addrs[i+1])
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			fmt.Println("The rindex", rindex, " information in Role contract, roleType:", roleType, " index:", index)

// 			// register keeper role
// 			if roleType == 0 {
// 				err = rKeeper.RegisterKeeper(roleAddr, rindex, []byte("Hello, test"), nil)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 			}
// 		}

// 		fmt.Println(">>>> begin register provider ")

// 		rProvider := callconts.NewR(addrs[4], sks[4], txopts)
// 		// isActive, isBanned, roleType, index, gIndex, extra
// 		isActive, _, roleType, index, gIndex, _, err := rProvider.GetRoleInfo(roleAddr, addrs[4])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("provider info: ")
// 		fmt.Println("isActive:", isActive, " roleType:", roleType, " index:", index, " gIndex:", gIndex)

// 		// register provider
// 		if roleType == 0 {
// 			err = rProvider.RegisterProvider(roleAddr, rIndexes[4], nil)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 	}

// 	// create groups
// 	if fast {
// 		fmt.Println("@@ fast test with existing group index 1,2")
// 		gIndex1 = 1
// 		gIndex2 = 2
// 	} else {
// 		// create group, keeper1 keeper2 in group1, keeper3 in group2
// 		fmt.Println(">>>> begin create group 1")
// 		// 需要admin先调用CreateGroup,同时将部署FileSys合约
// 		gIndex1, err = rAdmin.CreateGroup(roleAddr, rolefsAddr, uint64(0), rIndexes[1:3], 2)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if gIndex1 != 1 {
// 			log.Fatal("gIndex1 should be 1")
// 		}

// 		// test create group 2 with inActive, level=2, but only 1 keeper involved
// 		fmt.Println(">>>> begin create group 2")
// 		gIndex2, err = rAdmin.CreateGroup(roleAddr, rolefsAddr, uint64(0), rIndexes[3:4], 2)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if gIndex2 != 2 {
// 			log.Fatal("gIndex2 should be 2")
// 		}
// 	}

// 	// register user role
// 	if fast {
// 		fmt.Println("@@ fast test skip register user, use existing user")
// 	} else {
// 		fmt.Println(">>>> begin register user")
// 		// role caller for user
// 		rUser := callconts.NewR(addrs[0], sks[0], txopts)
// 		// isActive, isBanned, roleType, index, gIndex, extra
// 		isActive, _, roleType, index, gIndex, _, err := rUser.GetRoleInfo(roleAddr, addrs[0])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("user info: ")
// 		fmt.Println("isActive:", isActive, " roleType:", roleType, " index:", index, " gIndex:", gIndex)

// 		// register user
// 		if roleType == 0 {
// 			err = rUser.RegisterUser(roleAddr, rtokenAddr, rIndexes[0], 1, 0, nil, nil)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}
// 	}

// 	// show info
// 	{
// 		fmt.Println(">>>> Show Info:")
// 		fmt.Println("rIndexes: ", rIndexes)
// 		fmt.Println("addrs: ", addrs)
// 		fmt.Println("gIndex1: ", gIndex1)
// 		fmt.Println("gIndex2: ", gIndex2)
// 		fmt.Println("role: ", roleAddr)
// 		fmt.Println("rtoken: ", rtokenAddr)
// 		fmt.Println("rolefs: ", rolefsAddr)
// 		fmt.Println("pledgePool: ", pledgePoolAddr)
// 		fmt.Println("issuance: ", issuAddr)
// 	}

// 	var (
// 		//isActive, isBanned, isReady, level, _size, price, fsAddr1, err
// 		isActive bool
// 		isBanned bool
// 		isReady  bool
// 		level    uint16
// 		_size    *big.Int
// 		price    *big.Int
// 	)
// 	// get group1 info
// 	fmt.Println(">>>> getting group1 info")
// 	isActive, isBanned, isReady, level, _size, price, fsAddr1, err = rAdmin.GetGroupInfo(roleAddr, gIndex1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The group info- isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr1.Hex())

// 	if !isActive {
// 		log.Fatal("group1 should be active")
// 	}

// 	// get group2 info
// 	fmt.Println(">>>> getting group2 info")
// 	isActive, isBanned, isReady, level, _size, price, fsAddr2, err = rAdmin.GetGroupInfo(roleAddr, gIndex2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("The group info- isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr2.Hex())

// 	if isActive {
// 		log.Fatal("group2 should be inactive")
// 	}

// 	return nil
// }
