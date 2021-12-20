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

// 该测试需花费约14分钟
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
	acc3Addr := common.HexToAddress(test.Acc3) // 用于注册provider
	acc4Addr := common.HexToAddress(test.Acc4) // 用于注册provider
	accs := []common.Address{acc1Addr, acc2Addr, acc3Addr, acc4Addr}
	pledgeK := big.NewInt(1e6)
	start := uint64(time.Now().Unix())
	fmt.Println("start:", start)
	end := start + 10
	size := uint64(10)
	nonce := uint64(0)
	sprice := big.NewInt(10)
	zero := big.NewInt(0)
	hundred := big.NewInt(100)

	//start := 1639911209
	//uIndex := uint64(2)
	//gIndex := uint64(1)
	//kIndex := uint64(1)
	//pIndex := uint64(3)
	//p2Index := uint64(4)
	//rolefsAddr := common.HexToAddress("0x506703d5e2126003944B707bDb10c1030F891e15")
	//roleAddr := common.HexToAddress("0x05B2A72D0045aF7A5dD574c138a28528f46A6E33")
	//rTokenAddr := common.HexToAddress("0xAEAFe15CBd5e6f7788779330Ad4A152E8ddB06eF")
	//pledgePoolAddr := common.HexToAddress("0xcfcBCac22D55e1FFeD9887158E642Fde09924071")
	//issuanceAddr := common.HexToAddress("0x5Be65871560ee3038aE99c9D9d2b0c23AEe381ab")
	//fsAddr := common.HexToAddress("0x44f813AFaAA5e65c8D06Be806C2DeB2E2335Bd76")

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
	callconts.PledgePoolAddr = pledgePoolAddr
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
	// user approve
	erc20 = callconts.NewERC20(acc1Addr, test.Sk1, txopts)
	err = erc20.Approve(test.PrimaryToken, fsAddr, pledgeK)
	if err != nil {
		log.Fatal(err)
	}
	// 查询该User对FileSys合约的allowance
	allo, err := erc20.Allowance(test.PrimaryToken, acc1Addr, fsAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The allowance of User to FileSys contract is ", allo)
	if allo.Cmp(pledgeK) != 0 {
		log.Fatal("allowance should be ", pledgeK)
	}
	// user往FileSys中充值pledgeK，用于存储服务付费
	err = r.Recharge(roleAddr, rTokenAddr, uIndex, 0, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	// user获取其在FileSys中的balance
	fs := callconts.NewFileSys(adminAddr, test.AdminSk, txopts)
	avail, _, err := fs.GetBalance(fsAddr, uIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user avail in fs:", avail)
	if avail.Cmp(pledgeK) < 0 {
		log.Fatal("user avail should not less than ", pledgeK)
	}
	// provider注册、质押、申请角色、加入group
	r = callconts.NewR(acc3Addr, test.Sk3, txopts)
	err = r.Register(roleAddr, acc3Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	pIndex, err := r.GetRoleIndex(roleAddr, acc3Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pIndex ", pIndex)
	pp = callconts.NewPledgePool(acc3Addr, test.Sk3, txopts)
	err = pp.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, pIndex, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = r.RegisterProvider(roleAddr, pIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = r.AddProviderToGroup(roleAddr, pIndex, gIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	//admin为RoleFS赋予Minter_Role权限
	erc20 = callconts.NewERC20(adminAddr, test.AdminSk, txopts)
	err = erc20.SetUpRole(test.PrimaryToken, callconts.MinterRole, rolefsAddr)
	if err != nil {
		log.Fatal(err)
	}
	// keeper调用addOrder
	rfs = callconts.NewRFS(acc2Addr, test.Sk2, txopts)
	err = rfs.AddOrder(rolefsAddr, roleAddr, rTokenAddr, uIndex, pIndex, start, end, size, nonce, 0, sprice, nil, nil, [][]byte{[]byte("test")})
	if err != nil {
		log.Fatal(err)
	}
	// 获取addOrder后fs的信息并测试正确性
	pSum, err := fs.GetFsProviderSum(fsAddr, uIndex)
	if err != nil {
		log.Fatal(err)
	}
	if pSum != 1 {
		fmt.Println("pSum ", pSum)
		log.Fatal("The pSum should be 1")
	}
	pIndexGotten, err := fs.GetFsProvider(fsAddr, uIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	if pIndexGotten != pIndex {
		fmt.Println("pIndexGotten ", pIndexGotten)
		log.Fatal("The pIndex gotten should be ", pIndex)
	}
	// 获取addOrder后channel的信息并测试正确性
	_amount, _nonce, _expire, err := fs.GetChannelInfo(fsAddr, uIndex, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: channel amount", _amount, ",nonce", _nonce, ",expire", _expire)
	if _amount.Cmp(zero) != 0 || _nonce != 0 || _expire != end {
		log.Fatal("result error")
	}
	// 获取addOrder后store的信息并测试正确性
	_time, _size, _price, err := fs.GetStoreInfo(fsAddr, uIndex, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: storeinfo time", _time, ",size", _size, ",price", _price)
	if _price.Cmp(sprice) != 0 || _size != size || _time != 0 {
		log.Fatal("price in storeInfo after addOrder should be ", sprice, ", size in storeInfo should be ", size)
	}
	// 获取addOrder后proInfo中的Settlement的信息，并测试正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err := fs.GetSettleInfo(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: settleinfo time", _time, ",size", _size, ",price", _price, ",maxPay", _maxPay, ",hasPaid", _hasPaid, ",canPay", _canPay, ",lost", _lost, ",lostPaid", _lostPaid, ",managePay", _managePay, ",endPaid", _endPaid, ",linearPaid", _linearPaid)
	if _time != start || _price != sprice || _size != size || _maxPay.Cmp(hundred) != 0 || _managePay.Cmp(big.NewInt(4)) != 0 {
		log.Fatal("time should be ", start, " price should be ", sprice, " size should be ", size, " maxPay should be ", 100, " managePay should be ", 4)
	}
	// 获取addOrder后proInfo中的aggOrder的信息，并测试正确性
	_nonce, _subNonce, err := fs.GetFsInfoAggOrder(fsAddr, uIndex, pIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: aggOrderInfo nonce ", _nonce, ",subNonce:", _subNonce)
	if _nonce != 1 || _subNonce != 0 {
		log.Fatal("nonce in aggOrder should be ", 1)
	}
	// 获取addOrder后foundation的balance，并测试正确性
	_availF, _tmpF, err := fs.GetBalance(fsAddr, 0, 0) // 查询基金会在FileSys中的balance
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: foundation avail ", _availF, ",tmp ", _tmpF)
	if _availF.Cmp(big.NewInt(1)) != 0 || _tmpF.Cmp(zero) != 0 {
		log.Fatal("availF should be ", 1, " tmpF should be ", _tmpF)
	}
	// 获取addOrder后user的balance，并测试正确性
	_avail, _tmp, err := fs.GetBalance(fsAddr, uIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: user avail ", _avail, ",tmp ", _tmp)
	if _avail.Sub(avail, _avail).Cmp(big.NewInt(105)) != 0 {
		log.Fatal("the new avail should be 105 less than avail")
	}
	// 获取addOrder后Role合约里的group信息，并测试正确性
	_isActive, _isBanned, _isReady, _level, sizeG, priceG, _fsAddr, err := r.GetGroupInfo(roleAddr, gIndex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After addOrder: groupInfo isActive ", _isActive, ",isBanned ", _isBanned, ",isReady ", _isReady, ",level ", _level, ",size ", sizeG, ",price ", priceG, ",fsAddr ", _fsAddr.Hex())
	if sizeG.Cmp(big.NewInt(int64(size))) != 0 || priceG.Cmp(sprice) != 0 {
		log.Fatal("the sizeG should be ", size, " priceG should be ", sprice)
	}
	// 获取addOrder后代币发行的相关情况，并测试正确性(初次addOrder，mintToken值为0)
	mintLevel, err := issu.MintLevel(issuanceAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mintLevel:", mintLevel)
	if mintLevel.Cmp(zero) != 0 {
		log.Fatal("mintLevel should be 0")
	}
	st, err := issu.SpaceTime(issuanceAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("spaceTime:", st)
	if st.Cmp(zero) != 0 {
		log.Fatal("st should be 0")
	}
	//size:10, price:10, totalPay:100,上面测试通过的话，这里就不需要再重复测试了

	fmt.Println("============4. begin test SubOrder============")
	err = rfs.SubOrder(rolefsAddr, roleAddr, rTokenAddr, uIndex, pIndex, start+5, end, size-5, nonce, 0, sprice.Sub(sprice, big.NewInt(5)), nil, nil, [][]byte{[]byte("test")})
	if err != nil {
		log.Fatal(err)
	}
	// 这里不重复测试了，仅抽取部分信息进行测试
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubOrder: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid)
	// after SubOrder:  end 5 5 100 0 100 0 0 4 0 0
	if _time != end || _size != (size-5) || _maxPay.Cmp(hundred) != 0 || _hasPaid.Cmp(zero) != 0 || _canPay.Cmp(hundred) != 0 || _lost.Cmp(zero) != 0 || _managePay.Cmp(big.NewInt(4)) != 0 {
		log.Fatal("result wrong")
	}

	fmt.Println("============5. begin test ProWithdraw============")
	pay := big.NewInt(20)
	lost := big.NewInt(10)
	err = rfs.ProWithdraw(rolefsAddr, roleAddr, rTokenAddr, pIndex, 0, pay, lost, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid)
	// end+1010 5 5 100 20 5150 10 0 4 0 0
	if _size != (size-5) || _hasPaid.Cmp(pay) != 0 || _lost.Cmp(lost) != 0 {
		log.Fatal("result wrong")
	}
	// 获取balance信息
	_avail, _tmp, err = fs.GetBalance(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw, provider avail", _avail, " tmp", _tmp)
	if _avail.Cmp(zero) != 0 || _tmp.Cmp(big.NewInt(70)) != 0 {
		log.Fatal("result wrong")
	}
	// 获取provider在tIndex上的代币余额
	bal, err = erc20.BalanceOf(test.PrimaryToken, acc3Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw, provider balance is ", bal)

	fmt.Println("============6. begin test AddRepair============")
	// 获取Issuance中的totalPaid
	_totalPaid, err := issu.TotalPaid(issuanceAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issuance totalPaid:", _totalPaid)
	if _totalPaid.Cmp(zero) != 0 {
		log.Fatal("result wrong")
	}
	// 调用AddRepair前，需要先调用ProWithdraw、指定lost值
	// 新注册一个provider
	r = callconts.NewR(acc4Addr, test.Sk4, txopts)
	err = r.Register(roleAddr, acc4Addr, nil)
	if err != nil {
		log.Fatal(err)
	}
	p2Index, err := r.GetRoleIndex(roleAddr, acc4Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("p2Index ", p2Index)
	pp = callconts.NewPledgePool(acc4Addr, test.Sk4, txopts)
	err = pp.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, p2Index, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	err = r.RegisterProvider(roleAddr, p2Index, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = r.AddProviderToGroup(roleAddr, p2Index, gIndex, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 先构造签名信息
	//bytes32 h = keccak256(abi.encodePacked(pIndex, _start, end, _size, nonce, tIndex, sprice));
	npSig, err := callconts.SignForRepair(test.Sk4, pIndex, start+9, end, size, nonce, 0, big.NewInt(5), "a") // new provider sign，此处需要确保sprice*(end-start)<=lost
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("params:", pIndex, start+9, end, size, nonce, 0, big.NewInt(5), "a")
	fmt.Println("npSign:", npSig)
	// 调用AddRepair
	time.Sleep(2 * time.Second)
	err = rfs.AddRepair(rolefsAddr, roleAddr, rTokenAddr, pIndex, p2Index, start+9, end, size, nonce, 0, big.NewInt(5), npSig, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(fsAddr, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid)
	if _time != (start+9) || _size != size || _price.Cmp(big.NewInt(5)) != 0 || _maxPay.Cmp(big.NewInt(5)) != 0 || _lost.Cmp(zero) != 0 {
		log.Fatal("result wrong")
	}
	// 获取balance信息
	_avail, _tmp, err = fs.GetBalance(fsAddr, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, provider avail", _avail, " tmp", _tmp)
	if _avail.Cmp(zero) != 0 || _tmp.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("result wrong")
	}
	// 获取repairFs信息
	pSum, err = fs.GetFsProviderSum(fsAddr, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs providerSum:", pSum)
	if pSum != 1 {
		log.Fatal("pSum should be 1")
	}
	_nonce, _subNonce, err = fs.GetFsInfoAggOrder(fsAddr, 0, p2Index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider nonce:", _nonce, " subNonce:", _subNonce)
	if _nonce != 1 || _subNonce != 0 {
		log.Fatal("result wrong")
	}
	_time, _size, _price, err = fs.GetStoreInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider storeInfo:", _time, _size, _price)
	if _time != 0 || _size != size || _price.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("result wrong")
	}
	_amount, _nonce, _expire, err = fs.GetChannelInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider channelInfo:", _amount, _nonce, _expire)
	if _amount.Cmp(zero) != 0 || _nonce != 0 || _expire != 0 {
		log.Fatal("result wrong")
	}

	fmt.Println("============7. begin test SubRepair============")
	// 先构造签名信息
	//bytes32 h = keccak256(abi.encodePacked(pIndex, _start, end, _size, nonce, tIndex, sprice));
	npSig, err = callconts.SignForRepair(test.Sk4, pIndex, start+9, end, size, nonce, 0, big.NewInt(5), "s") // new provider sign
	if err != nil {
		log.Fatal(err)
	}
	// 调用SubRepair
	err = rfs.SubRepair(rolefsAddr, roleAddr, rTokenAddr, pIndex, p2Index, start+9, end, size, nonce, 0, big.NewInt(5), npSig, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(fsAddr, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid)
	// 获取balance信息
	_avail, _tmp, err = fs.GetBalance(fsAddr, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, provider avail", _avail, " tmp", _tmp)
	if _avail.Cmp(zero) != 0 || _tmp.Cmp(big.NewInt(5)) != 0 {
		log.Fatal("result wrong")
	}
	// 获取repairFs信息
	pSum, err = fs.GetFsProviderSum(fsAddr, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs providerSum:", pSum)
	if pSum != 1 {
		log.Fatal("pSum should be 1")
	}
	_nonce, _subNonce, err = fs.GetFsInfoAggOrder(fsAddr, 0, p2Index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider nonce:", _nonce, " subNonce:", _subNonce)
	if _nonce != 1 || _subNonce != 1 {
		log.Fatal("result wrong")
	}
	_time, _size, _price, err = fs.GetStoreInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider storeInfo:", _time, _size, _price)
	if _time != 0 || _size != 0 || _price.Cmp(zero) != 0 {
		log.Fatal("result wrong")
	}
	_amount, _nonce, _expire, err = fs.GetChannelInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider channelInfo:", _amount, _nonce, _expire)
	if _amount.Cmp(zero) != 0 || _nonce != 0 || _expire != 0 {
		log.Fatal("result wrong")
	}

	fmt.Println("============test success!============")
}
