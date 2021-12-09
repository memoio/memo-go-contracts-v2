package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	callconts "memoContract/callcontracts"
	"memoContract/test"
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
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 用于测试的一些参数
	adminAddr := common.HexToAddress(test.AdminAddr)
	rechargeMoney := big.NewInt(1e8)
	start := uint64(time.Now().Unix()) // 当前时间的时间戳
	end := start + 10
	size := uint64(10)
	sprice := big.NewInt(100)
	// 部署Role的参数
	pledgeK := big.NewInt(1e18)
	pledgeP := big.NewInt(1e18)
	var addrs []common.Address = []common.Address{common.HexToAddress(test.Addr), common.HexToAddress(test.Addr2), common.HexToAddress(test.Addr3), common.HexToAddress(test.Addr4), common.HexToAddress(test.Addr5)}
	var sks []string = []string{test.Sk, test.Sk2, test.Sk3, test.Sk4, test.Sk5}

	// 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("test admin-account has balance: ", bal, " in Ethereum")
	for _, addr := range addrs {
		bal = callconts.QueryEthBalance(addr.Hex(), ethEndPoint)
		fmt.Println("test common-account ", addr.Hex(), " has balance: ", bal, " in Ethereum")
	}

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	// 查看测试账户在ERC20代币上的余额，不足时，自动充值
	erc20 := callconts.NewERC20(adminAddr, test.AdminSk, txopts)
	bal, err := erc20.BalanceOf(test.PrimaryToken, adminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The balance of admin account in erc20 is ", bal)
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
		fmt.Println("After mint, the balance of admin account in erc20 is ", bal)
	}
	for _, addr := range addrs {
		bal, err = erc20.BalanceOf(test.PrimaryToken, addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The balance of ", addr.Hex(), " in erc20 is ", bal)
		if bal.Cmp(pledgeK) < 0 {
			err = erc20.Transfer(test.PrimaryToken, addr, pledgeK) // admin给测试账户转账，用于测试（充值或质押）
			if err != nil {
				log.Fatal(err)
			}
			bal, err = erc20.BalanceOf(test.PrimaryToken, addr)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("After transfer, the balance of ", addr.Hex(), " in erc20 is ", bal)
		}
	}

	fmt.Println("============1. begin test deploy FileSys contract============")
	// 部署FileSys合约前准备工作：部署Role合约、部署RoleFS合约、账户注册角色、调用Role合约中的CreateGroup函数
	// 顺序：Role、RoleFS、PledgePool、CreateGroup(FileSys)、
	r := callconts.NewR(adminAddr, test.AdminSk, txopts)
	roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeP)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Role contract address: ", roleAddr.Hex())
	//roleAddr := common.HexToAddress("0xe26F77e3268ae064514b18211Ed46DC0460197FE")

	rfs := callconts.NewRFS(adminAddr, test.AdminSk, txopts)
	rolefsAddr, _, err := rfs.DeployRoleFS()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The RoleFS contract address: ", rolefsAddr.Hex())
	//rolefsAddr := common.HexToAddress("0xAAaC6D27153BF52d66Eed127e0321372B2FFF67C")

	// 账户注册, addr:User addr2、addr3、addr4:Keeper addr5:Provider
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
	//rIndexes := []uint64{1, 2, 3, 4, 5}

	// 部署PledgePool合约用于账户质押，之后才可以申请Keeper、Provider角色
	rtokenAddr, err := r.RToken(roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The RToken contract address: ", rtokenAddr.Hex())
	//rtokenAddr := common.HexToAddress("0x081458b892fb2caEb3e4a6234F9183594531b715")
	p := callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
	pledgePoolAddr, _, err := p.DeployPledgePool(test.PrimaryToken, rtokenAddr, roleAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The PledgePool contract address: ", pledgePoolAddr.Hex())
	//pledgePoolAddr := common.HexToAddress("0xfB279b29E437cEE73ac2F2423a3Cfb22060Eb9d7")
	pledgek, err := r.PledgeK(roleAddr) // 申请Keeper最少需质押的金额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Apply Keeper need to pledge ", pledgek)
	pledgep, err := r.PledgeP(roleAddr) // 申请Provider最少需质押的金额
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Apply Provider need to pledge ", pledgep)
	// Keepers质押
	for i, rindex := range rIndexes[1:4] {
		err = toPledge(addrs[i+1], roleAddr, pledgePoolAddr, sks[i+1], rindex, pledgek, txopts)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Provider质押
	err = toPledge(addrs[4], roleAddr, pledgePoolAddr, sks[4], rIndexes[4], pledgep, txopts)
	if err != nil {
		log.Fatal(err)
	}

	issu := callconts.NewIssu(adminAddr, test.AdminSk, txopts)
	issuanceAddr, _, err := issu.DeployIssuance(rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The Issuance contract address is: ", issuanceAddr.Hex()) // 0xB15FEDB8017845331b460786fb5129C1Da06f6B1

	// 顺便给Role合约赋值
	r = callconts.NewR(adminAddr, test.AdminSk, txopts)
	err = r.SetPI(roleAddr, pledgePoolAddr, issuanceAddr, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}

	// 申请Keeper
	callconts.PledgePoolAddr = pledgePoolAddr
	for i, rindex := range rIndexes[1:4] {
		r := callconts.NewR(addrs[i+1], sks[i+1], txopts)
		fmt.Println(addrs[i+1].Hex(), " begin to register Keeper...")
		// 获取PledgePool合约中的balance
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

		err = r.RegisterKeeper(roleAddr, rindex, []byte("Hello, test"), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 申请Provider
	r = callconts.NewR(addrs[4], sks[4], txopts)
	err = r.RegisterProvider(roleAddr, rIndexes[4], nil)
	if err != nil {
		log.Fatal(err)
	}
	// 需要admin先调用CreateGroup,同时将部署FileSys合约
	r = callconts.NewR(adminAddr, test.AdminSk, txopts)
	gIndex, err := r.CreateGroup(roleAddr, rolefsAddr, uint64(0), rIndexes[1:3], 2)
	if err != nil {
		log.Fatal(err)
	}
	// 获取group信息
	isActive, isBanned, isReady, level, _size, price, fsAddr, err := r.GetGroupInfo(roleAddr, gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The group info: isActive:", isActive, " isBanned: ", isBanned, " isReady:", isReady, " level:", level, " size:", _size, " price:", price, " fsAddr:", fsAddr.Hex())

	fmt.Println("============2. begin test GetFsInfo============")
	// 先注册User，与group绑定
	r = callconts.NewR(addrs[0], sks[0], txopts)
	err = r.RegisterUser(roleAddr, rtokenAddr, rIndexes[0], gIndex, 0, []byte("Hello,test user"), nil)
	if err != nil {
		log.Fatal(err)
	}
	// 根据uIndex获取filesys信息
	fs := callconts.NewFileSys(addrs[0], sks[0], txopts)
	isActive, tokenIndex, err := fs.GetFsInfo(fsAddr, rIndexes[0])
	if err != nil {
		log.Fatal(err)
	}
	if !isActive || tokenIndex != 0 {
		log.Fatal("The filesys with uIndex: ", rIndexes[0], " isActive: ", isActive, " tokenIndex: ", tokenIndex)
	}

	fmt.Println("============3. begin test GetBalance============")
	// Provider申请加入group
	r = callconts.NewR(addrs[4], sks[4], txopts)
	err = r.AddProviderToGroup(roleAddr, rIndexes[4], gIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 获得FileSys中的AggOrder信息，则需要先提前调用RoleFS合约中的AddOrder函数，这样FileSys中才会有Order信息
	// 而调用RoleFS合约中的AddOrder函数前，需要先给RoleFS合约赋Issuance合约值
	// 给RoleFS合约赋值
	err = rfs.SetAddr(rolefsAddr, issuanceAddr, roleAddr, fsAddr, rtokenAddr)
	if err != nil {
		log.Fatal(err)
	}
	// User给FileSys充值，用于存储服务支付
	fmt.Println("User approve FileSys contract account")
	erc20 = callconts.NewERC20(addrs[0], sks[0], txopts)
	err = erc20.Approve(test.PrimaryToken, fsAddr, rechargeMoney)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User recharge to FileSys contract")
	// 查询该User对应的fsAddr,判断是否和上述相同
	r = callconts.NewR(addrs[0], sks[0], txopts)
	_, _, _, _, _gIndex, _, err := r.GetRoleInfo(roleAddr, addrs[0])
	if err != nil {
		log.Fatal(err)
	}
	_, _, _, _, _, _, _fsAddr, err := r.GetGroupInfo(roleAddr, _gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The fsAddr of user is ", _fsAddr.Hex())
	fmt.Println("The user address is ", addrs[0].Hex()) // 用来判断是否和rIndexes[0]指代的地址相同
	// 查询该User对FileSys合约的allowance
	allo, err := erc20.Allowance(test.PrimaryToken, addrs[0], fsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The allowance of User to FileSys contract is ", allo)
	// 查询fs是否isActive
	_isActive, _, err := fs.GetFsInfo(fsAddr, rIndexes[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The fs of user isActive: ", _isActive)
	err = r.Recharge(roleAddr, rtokenAddr, rIndexes[0], 0, rechargeMoney, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 查询User在FileSys合约中的可用balance
	avail, tmp, err := fs.GetBalance(fsAddr, rIndexes[0], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The User's avail balance in FileSys contract is ", avail, " tmp is ", tmp)

	fmt.Println("============4. begin test GetFsInfoAggOrder============")
	// 调用AddOrder函数前，需要为RoleFS赋予MINTER_ROLE权限
	erc20 = callconts.NewERC20(adminAddr, test.AdminSk, txopts)
	err = erc20.SetUpRole(test.PrimaryToken, callconts.MinterRole, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	// 调用RoleFS合约中的AddOrder函数,keeper调用
	rfs = callconts.NewRFS(addrs[1], sks[1], txopts)
	err = rfs.AddOrder(rolefsAddr, roleAddr, rtokenAddr, rIndexes[0], rIndexes[4], start, end, size, 0, 0, sprice, nil, nil, [][]byte{[]byte("Hello")})
	if err != nil {
		log.Fatal(err)
	}
	// 获取order信息
	_nonce, subNonce, err := fs.GetFsInfoAggOrder(fsAddr, rIndexes[0], rIndexes[4])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The order info nonce is ", _nonce, " subNonce is ", subNonce) // 应该都为0

	fmt.Println("============5. begin test GetStoreInfo============")
	_t, _s, _p, err := fs.GetStoreInfo(fsAddr, rIndexes[0], rIndexes[4], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The store information time is ", _t, " size is ", _s, " price is ", _p) // 应该和上述的time、size、price相同

	fmt.Println("============6. begin test GetChannelInfo============")
	amount, _n, expire, err := fs.GetChannelInfo(fsAddr, rIndexes[0], rIndexes[4], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The channel info amount is ", amount, " nonce is ", _n, " expire is ", expire)

	fmt.Println("============7. begin test GetSettleInfo============")
	setime, sesize, seprice, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err := fs.GetSettleInfo(fsAddr, rIndexes[4], 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The settlement info time is ", setime, " size is ", sesize, " price is ", seprice, " maxPay is ", maxPay, " hasPaid is ", hasPaid, " canPay is ", canPay, " lost is ", lost, " lostPaid is ", lostPaid, " managePay is ", managePay, " endPaid is ", endPaid, " linearPaid is ", linearPaid)

	fmt.Println("============test success!============")
}

func toPledge(addr, roleAddr, pledgePoolAddr common.Address, sk string, rindex uint64, pledgek *big.Int, txopts *callconts.TxOpts) error {
	p := callconts.NewPledgePool(addr, sk, txopts)
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
	err = p.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, rindex, pledgek, nil)
	if err != nil {
		return err
	}
	return nil
}
