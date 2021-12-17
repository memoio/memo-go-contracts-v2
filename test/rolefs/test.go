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
	//acc1Addr := common.HexToAddress(test.Acc1) // 用于注册User
	acc2Addr := common.HexToAddress(test.Acc2) // 用于注册keeper
	acc3Addr := common.HexToAddress(test.Acc5) // 用于注册provider
	acc4Addr := common.HexToAddress(test.Acc4) // 用于注册provider
	//accs := []common.Address{acc1Addr, acc2Addr, acc3Addr, acc4Addr}
	pledgeK := big.NewInt(1e6)
	start := uint64(time.Now().Unix())
	end := start + 10
	size := uint64(10)
	nonce := uint64(0)
	sprice := big.NewInt(10)

	uIndex := uint64(2)
	gIndex := uint64(1)
	//kIndex := uint64(1)
	pIndex := uint64(3)

	// // 查看余额，支付交易Gas费，余额不足时，需充值（暂时手动）
	// bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	// fmt.Println("admin balance: ", bal, " in Ethereum")
	// for i, acc := range accs {
	// 	bal = callconts.QueryEthBalance(acc.Hex(), ethEndPoint)
	// 	fmt.Println("acc", i, " balance: ", bal, " in Ethereum")
	// }

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: big.NewInt(callconts.DefaultGasPrice),
		GasLimit: callconts.DefaultGasLimit,
	}

	// // 查询在erc20代币上的余额
	// erc20 := callconts.NewERC20(adminAddr, test.AdminSk, txopts)
	// bal, err := erc20.BalanceOf(test.PrimaryToken, adminAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("admin balance in primaryToken is ", bal)
	// // 余额不足，自动充值
	// if bal.Cmp(big.NewInt(test.MoneyTo)) < 0 {
	// 	// mintToken
	// 	err = erc20.MintToken(test.PrimaryToken, adminAddr, big.NewInt(test.MoneyTo))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bal, err = erc20.BalanceOf(test.PrimaryToken, adminAddr)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("after mint, admin balance in primaryToken is ", bal)
	// }
	// for i, acc := range accs {
	// 	bal, err = erc20.BalanceOf(test.PrimaryToken, acc)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("acc", i, " balance in primaryToken is ", bal)
	// 	if bal.Cmp(pledgeK) < 0 {
	// 		err = erc20.Transfer(test.PrimaryToken, acc, pledgeK)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		bal, err = erc20.BalanceOf(test.PrimaryToken, acc)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		fmt.Println("after transfer, acc", i, " balance in primaryToken is ", bal)
	// 	}
	// }

	rfs := callconts.NewRFS(adminAddr, test.AdminSk, txopts)

	// fmt.Println("============1. begin test deploy RoleFS contract============")
	// rolefsAddr, _, err := rfs.DeployRoleFS()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("RoleFS contract Address:", rolefsAddr.Hex())
	rolefsAddr := common.HexToAddress("0xC8b12827AB0562E2C7FFFa3D97Cfe89FcC109a45")

	// // 部署RoleFS合约之后，需要调用SetAddr函数赋予合约需要的状态变量信息
	// fmt.Println("============2. begin test SetAddr============")
	// // 部署Role合约
	// r := callconts.NewR(adminAddr, test.AdminSk, txopts)
	// roleAddr, _, err := r.DeployRole(test.Foundation, test.PrimaryToken, pledgeK, pledgeK)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("The Role contract address is ", roleAddr.Hex())
	roleAddr := common.HexToAddress("0x09635198900B3336F3400483d0f27fd8b71C96aE")
	// // 获得RToken合约地址
	// rTokenAddr, err := r.RToken(roleAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("The RToken contract address is ", rTokenAddr.Hex())
	rTokenAddr := common.HexToAddress("0x64f3a33353D6478c3cC4460291e366c93926a13f")
	// // 部署PledgePool合约
	// pp := callconts.NewPledgePool(adminAddr, test.AdminSk, txopts)
	// pledgePoolAddr, _, err := pp.DeployPledgePool(test.PrimaryToken, rTokenAddr, roleAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("The PledgePool contract address is ", pledgePoolAddr.Hex())
	pledgePoolAddr := common.HexToAddress("0x76d3E1374f4442e956b8B8520124eAC7eefF3513")
	// // 部署Issuance合约
	//issu := callconts.NewIssu(adminAddr, test.AdminSk, txopts)
	// issuanceAddr, _, err := issu.DeployIssuance(rolefsAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("The Issuance contract address is ", issuanceAddr.Hex())
	//issuanceAddr := common.HexToAddress("0xCcDf12eEFe3a42356D406E6a0D8D25edCCd548c0")
	// // setPI
	// err = r.SetPI(roleAddr, pledgePoolAddr, issuanceAddr, rolefsAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // keeper注册、质押、申请角色
	// r = callconts.NewR(acc2Addr, test.Sk2, txopts)
	// err = r.Register(roleAddr, acc2Addr, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// kIndex, err := r.GetRoleIndex(roleAddr, acc2Addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("kIndex ", kIndex)
	// pp = callconts.NewPledgePool(acc2Addr, test.Sk2, txopts)
	// err = pp.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, kIndex, pledgeK, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// callconts.PledgePoolAddr = pledgePoolAddr
	// err = r.RegisterKeeper(roleAddr, kIndex, []byte("test"), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // createGroup, and deploy FileSys contract
	// r = callconts.NewR(adminAddr, test.AdminSk, txopts)
	// gIndex, err := r.CreateGroup(roleAddr, rolefsAddr, 0, []uint64{kIndex}, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("gIndex ", gIndex)
	// // 获取FileSys合约地址
	// _, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(roleAddr, gIndex)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("The FileSys contract address is ", fsAddr.Hex())
	fsAddr := common.HexToAddress("0xbfB1120EB875Af9870B2B052abC18539efDe5603")
	// // 给RoleFS合约赋予状态变量值
	// err = rfs.SetAddr(rolefsAddr, issuanceAddr, roleAddr, fsAddr, rTokenAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("============3. begin test AddOrder============")
	// // user注册、充值
	// r = callconts.NewR(acc1Addr, test.Sk1, txopts)
	// err = r.Register(roleAddr, acc1Addr, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// uIndex, err := r.GetRoleIndex(roleAddr, acc1Addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("uIndex ", uIndex)
	// err = r.RegisterUser(roleAddr, rTokenAddr, uIndex, gIndex, 0, []byte("test"), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // user approve
	// erc20 = callconts.NewERC20(acc1Addr, test.Sk1, txopts)
	// err = erc20.Approve(test.PrimaryToken, fsAddr, pledgeK)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 查询该User对FileSys合约的allowance
	// allo, err := erc20.Allowance(test.PrimaryToken, acc1Addr, fsAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("The allowance of User to FileSys contract is ", allo)
	// // user往FileSys中充值pledgeK，用于存储服务付费
	// err = r.Recharge(roleAddr, rTokenAddr, uIndex, 0, pledgeK, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // user获取其在FileSys中的balance
	fs := callconts.NewFileSys(adminAddr, test.AdminSk, txopts)
	avail, _, err := fs.GetBalance(fsAddr, uIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	// if avail.Cmp(pledgeK) != 0 {
	// 	fmt.Println("user avail ", avail)
	// 	log.Fatal("user avail should be ", pledgeK)
	// }
	// // provider注册、质押、申请角色、加入group
	r := callconts.NewR(acc3Addr, test.Sk5, txopts)
	// err = r.Register(roleAddr, acc3Addr, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pIndex, err := r.GetRoleIndex(roleAddr, acc3Addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("pIndex ", pIndex)
	// pp = callconts.NewPledgePool(acc3Addr, test.Sk5, txopts)
	// err = pp.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, pIndex, pledgeK, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = r.RegisterProvider(roleAddr, pIndex, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = r.AddProviderToGroup(roleAddr, pIndex, gIndex, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// admin为RoleFS赋予Minter_Role权限
	erc20 := callconts.NewERC20(adminAddr, test.AdminSk, txopts)
	// err = erc20.SetUpRole(test.PrimaryToken, callconts.MinterRole, rolefsAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // keeper调用addOrder
	rfs = callconts.NewRFS(acc2Addr, test.Sk2, txopts)
	// err = rfs.AddOrder(rolefsAddr, roleAddr, rTokenAddr, uIndex, pIndex, start, end, size, nonce, 0, sprice, nil, nil, [][]byte{[]byte("test")})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // 获取addOrder后fs的信息并测试正确性
	// pSum, err := fs.GetFsProviderSum(fsAddr, uIndex)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if pSum != 1 {
	// 	fmt.Println("pSum ", pSum)
	// 	log.Fatal("The pSum should be 1")
	// }
	// pIndexGotten, err := fs.GetFsProvider(fsAddr, uIndex, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if pIndexGotten != pIndex {
	// 	fmt.Println("pIndexGotten ", pIndexGotten)
	// 	log.Fatal("The pIndex gotten should be ", pIndex)
	// }
	// // 获取addOrder后channel的信息并测试正确性
	// _, _, expire, err := fs.GetChannelInfo(fsAddr, uIndex, pIndex, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if expire != end {
	// 	fmt.Println("expire ", expire)
	// 	log.Fatal("expire should be ", end)
	// }
	// // 获取addOrder后store的信息并测试正确性
	// _, _size, _price, err := fs.GetStoreInfo(fsAddr, uIndex, pIndex, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if _price.Cmp(sprice) != 0 || _size != size {
	// 	fmt.Println("store info(after addOrder):price ", _price, " size ", _size)
	// 	log.Fatal("price in storeInfo after addOrder should be ", sprice, ", size in storeInfo should be ", size)
	// }
	// // 获取addOrder后proInfo中的Settlement的信息，并测试正确性
	// _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err := fs.GetSettleInfo(fsAddr, pIndex, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("after AddOrder: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid) // TODO:judge
	// // 1639654984 10 10 100 0 0 0 0 4 0 0
	// // if _time != start || _price != sprice || _size != size || _maxPay.Cmp(big.NewInt(100)) != 0 || _managePay.Cmp(big.NewInt(4)) != 0 {
	// // 	log.Fatal("time should be ", start, " price should be ", sprice, " size should be ", size, " maxPay should be ", 100, " managePay should be ", 4)
	// // }
	// // 获取addOrder后proInfo中的aggOrder的信息，并测试正确性
	// _nonce, _subNonce, err := fs.GetFsInfoAggOrder(fsAddr, uIndex, pIndex)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("aggOrder info=> nonce:", _nonce, " subNonce:", _subNonce)
	// // if _nonce != 1 {
	// // 	log.Fatal("nonce in aggOrder should be ", 1)
	// // }
	// // 获取addOrder后foundation的balance，并测试正确性
	// availF, tmpF, err := fs.GetBalance(fsAddr, 0, 0) // 查询基金会在FileSys中的balance
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if availF.Cmp(big.NewInt(1)) != 0 || tmpF.Cmp(big.NewInt(0)) != 0 {
	// 	fmt.Println("availF: ", availF, " tmpF: ", tmpF)
	// 	log.Fatal("availF should be ", 1, " tmpF should be ", tmpF)
	// }
	// // 获取addOrder后user的balance，并测试正确性
	// availNew, _, err := fs.GetBalance(fsAddr, uIndex, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if availNew.Sub(avail, availNew).Cmp(big.NewInt(105)) != 0 {
	// 	fmt.Println("user new avail: ", availNew)
	// 	log.Fatal("the availNew should be 105 less than avail")
	// }
	// // 获取addOrder后Role合约里的group信息，并测试正确性
	// _, _, _, _, sizeG, priceG, _, err := r.GetGroupInfo(roleAddr, gIndex)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if sizeG.Cmp(big.NewInt(int64(size))) != 0 || priceG.Cmp(sprice) != 0 {
	// 	fmt.Println("sizeG: ", sizeG, " priceG: ", priceG)
	// 	log.Fatal("the sizeG should be ", size, " priceG should be ", sprice)
	// }
	// // 获取addOrder后代币发行的相关情况，并测试正确性(初次addOrder，mintToken值为0)
	// mintLevel, err := issu.MintLevel(issuanceAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if mintLevel.Cmp(big.NewInt(0)) != 0 {
	// 	fmt.Println("MintLevel: ", mintLevel)
	// 	log.Fatal("mintLevel should be 0")
	// }
	// st, err := issu.SpaceTime(issuanceAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if st.Cmp(big.NewInt(100)) != 0 {
	// 	fmt.Println("spacetime in issuance:", st)   // 这里输出为0
	// 	log.Fatal("st should be 100")
	// }
	// size:10, price:10, totalPay:100,上面测试通过的话，这里就不需要再重复测试了

	fmt.Println("============4. begin test SubOrder============")
	//此处出错
	err = rfs.SubOrder(rolefsAddr, roleAddr, rTokenAddr, uIndex, pIndex, start+5, end, size-5, nonce, 0, sprice.Sub(sprice, big.NewInt(5)), nil, nil, [][]byte{[]byte("test")})
	if err != nil {
		log.Fatal(err)
	}
	// 这里不重复测试了，仅抽取部分信息进行测试
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err := fs.GetSettleInfo(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubOrder: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid) // TODO:judge
	// if _endPaid.Cmp(big.NewInt(1)) != 0 {
	// 	log.Fatal("endPaid should be 1")
	// }

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
	fmt.Println("after ProWithdraw: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid) // TODO:judge
	// if _lost.Cmp(lost) != 0 || _hasPaid.Cmp(pay) != 0 {
	// 	log.Fatal()
	// }
	// 获取balance信息
	avail, tmp, err := fs.GetBalance(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw, provider avail is ", avail, " tmp is ", tmp) // TODO:judge
	// 获取provider在tIndex上的代币余额
	bal, err := erc20.BalanceOf(test.PrimaryToken, acc3Addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after ProWithdraw, provider balance is ", bal) // TODO:judeg
	// if bal.Cmp(pay)!=0{
	// 	log.Fatal("balance should be ", pay)
	// }

	fmt.Println("============6. begin test AddRepair============")
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
	pp := callconts.NewPledgePool(acc4Addr, test.Sk4, txopts)
	err = pp.Pledge(pledgePoolAddr, test.PrimaryToken, roleAddr, p2Index, pledgeK, nil)
	if err != nil {
		log.Fatal(err)
	}
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
	npSig, err := callconts.SignForRepair(test.Sk4, p2Index, start+9, end, size, nonce, 0, sprice, "a") // new provider sign，此处需要确保sprice*(end-start)<=lost
	if err != nil {
		log.Fatal(err)
	}
	// 调用AddRepair
	err = rfs.AddRepair(rolefsAddr, roleAddr, rTokenAddr, pIndex, p2Index, start+9, end, size, nonce, 0, sprice, npSig, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid) // TODO:judge
	// if _lost.Cmp(lost) != 0 || _hasPaid.Cmp(pay) != 0 {
	// 	log.Fatal()
	// }
	// 获取balance信息
	avail, tmp, err = fs.GetBalance(fsAddr, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, provider avail is ", avail, " tmp is ", tmp) // TODO:judge
	// 获取repairFs信息
	pSum, err := fs.GetFsProviderSum(fsAddr, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs providerSum:", pSum)
	if pSum != 1 {
		log.Fatal("pSum should be 1")
	}
	nonce, subNonce, err := fs.GetFsInfoAggOrder(fsAddr, 0, p2Index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider nonce:", nonce, " subNonce:", subNonce) // TODO:judge
	_time, _size, _price, err = fs.GetStoreInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider storeInfo:", _time, _size, _price) // TODO:judge
	_amount, _nonce, _expire, err := fs.GetChannelInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after AddRepair, repairFs provider channelInfo:", _amount, _nonce, _expire)

	fmt.Println("============7. begin test SubRepair============")
	// 先构造签名信息
	//bytes32 h = keccak256(abi.encodePacked(pIndex, _start, end, _size, nonce, tIndex, sprice));
	npSig, err = callconts.SignForRepair(test.Sk4, p2Index, start+9, end, size, nonce, 0, sprice, "s") // new provider sign，此处需要确保sprice*(end-start)<=lost
	if err != nil {
		log.Fatal(err)
	}
	// 调用SubRepair
	err = rfs.SubRepair(rolefsAddr, roleAddr, rTokenAddr, pIndex, p2Index, start+9, end, size, nonce, 0, sprice, npSig, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 获取settlement信息，并判断正确性
	_time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid, err = fs.GetSettleInfo(fsAddr, pIndex, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair: ", _time, _size, _price, _maxPay, _hasPaid, _canPay, _lost, _lostPaid, _managePay, _endPaid, _linearPaid) // TODO:judge
	// if _lost.Cmp(lost) != 0 || _hasPaid.Cmp(pay) != 0 {
	// 	log.Fatal()
	// }
	// 获取balance信息
	avail, tmp, err = fs.GetBalance(fsAddr, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, provider avail is ", avail, " tmp is ", tmp) // TODO:judge
	// 获取repairFs信息
	pSum, err = fs.GetFsProviderSum(fsAddr, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs providerSum:", pSum)
	if pSum != 1 {
		log.Fatal("pSum should be 1")
	}
	nonce, subNonce, err = fs.GetFsInfoAggOrder(fsAddr, 0, p2Index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider nonce:", nonce, " subNonce:", subNonce) // TODO:judge
	_time, _size, _price, err = fs.GetStoreInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider storeInfo:", _time, _size, _price) // TODO:judge
	_amount, _nonce, _expire, err = fs.GetChannelInfo(fsAddr, 0, p2Index, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("after SubRepair, repairFs provider channelInfo:", _amount, _nonce, _expire)

	fmt.Println("============test success!============")
}
