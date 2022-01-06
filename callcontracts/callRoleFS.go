package callconts

import (
	"errors"
	"log"
	"math/big"
	rolefs "memoContract/contracts/rolefs"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewRFS new a instance of ContractModule
func NewRFS(roleFSAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string) iface.RoleFSInfo {
	rfs := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: roleFSAddr,
		endPoint:        endPoint,
	}

	return rfs
}

// DeployRoleFS deploy an RoleFS contract, called by admin
func (rfs *ContractModule) DeployRoleFS() (common.Address, *rolefs.RoleFS, error) {
	var roleFSAddr, roleFSAddress common.Address
	var roleFSInstance, roleFSIns *rolefs.RoleFS
	var err error

	log.Println("begin deploy RoleFS contract...")
	client := getClient(rfs.endPoint)
	defer client.Close()
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0
	for {
		auth, errMA := makeAuth(rfs.hexSk, nil, rfs.txopts)
		if errMA != nil {
			return roleFSAddr, nil, errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		roleFSAddress, tx, roleFSInstance, err = rolefs.DeployRoleFS(auth, client)
		if roleFSAddress.String() != InvalidAddr {
			roleFSAddr = roleFSAddress
			roleFSIns = roleFSInstance
		}
		if err != nil {
			retryCount++
			log.Println("deploy RoleFS Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return roleFSAddr, roleFSIns, err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("deploy RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return roleFSAddr, roleFSIns, err
			}
			continue
		}
		if err != nil {
			return roleFSAddr, roleFSIns, err
		}
		break
	}
	log.Println("RoleFS has been successfully deployed! The address is ", roleFSAddr.Hex())
	return roleFSAddr, roleFSIns, nil
}

func newRoleFS(roleFSAddr common.Address, client *ethclient.Client) (*rolefs.RoleFS, error) {
	roleFSIns, err := rolefs.NewRoleFS(roleFSAddr, client)
	if err != nil {
		return nil, err
	}
	return roleFSIns, nil
}

func (rfs *ContractModule) checkParam(uIndex, pIndex uint64, uRoleType, pRoleType uint8, tIndex uint32, roleAddr, rTokenAddr common.Address) (uint64, error) {
	// check whether uIndex is user
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	uaddr, err := r.GetAddr(uIndex)
	if err != nil {
		return 0, err
	}
	isActive, isBanned, roleType, _, ugIndex, _, err := r.GetRoleInfo(uaddr)
	if err != nil {
		return 0, err
	}
	if roleType != uRoleType || !isActive || isBanned {
		log.Println("uIndex ", uIndex, " roleType:", roleType, " isBanned:", isBanned, " isActive:", isActive)
		return 0, ErrIndex
	}
	// check whether pIndex is provider
	paddr, err := r.GetAddr(pIndex)
	if err != nil {
		return 0, err
	}
	isActive, isBanned, roleType, _, pgIndex, _, err := r.GetRoleInfo(paddr)
	if err != nil {
		return 0, err
	}
	if roleType != pRoleType || !isActive || isBanned {
		log.Println("pIndex ", pIndex, " roleType:", roleType, " isBanned:", isBanned, " isActive:", isActive)
		return 0, ErrIndex
	}
	// check whether their gIndex is same
	isActive, isBanned, roleType, _, cgIndex, _, err := r.GetRoleInfo(rfs.addr)
	if err != nil {
		return 0, err
	}
	if roleType != KeeperRoleType || !isActive || isBanned {
		log.Println("caller ", rfs.addr.Hex(), " roleType:", roleType, " isBanned:", isBanned, " isActive:", isActive)
		return 0, ErrIndex
	}
	if ugIndex != pgIndex || ugIndex != cgIndex {
		log.Println("uIndex's gIndex:", ugIndex, " pIndex's gIndex:", pgIndex, " caller's gIndex:", cgIndex)
		return 0, ErrIndex
	}
	// check tindex
	rt := NewRT(rTokenAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	isValid, err := rt.IsValid(tIndex)
	if err != nil {
		return 0, err
	}
	if !isValid {
		return 0, ErrTIndex
	}
	return ugIndex, nil
}

// SetAddr called by admin, which is the deployer. Set some address type variables.
func (rfs *ContractModule) SetAddr(issuan, role, fileSys, rtoken common.Address) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check caller
	r := NewOwn(role, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	owner, err := r.GetOwner()
	if err != nil {
		return err
	}
	if owner.Hex() != rfs.addr.Hex() {
		log.Println("owner is", owner.Hex(), " but caller is", rfs.addr.Hex())
		return errNotOwner
	}

	log.Println("begin SetAddr in RoleFS contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(rfs.hexSk, nil, rfs.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleFSIns.SetAddr(auth, issuan, role, fileSys, rtoken)
		if err != nil {
			retryCount++
			log.Println("setAddr Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("SetAddr in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	log.Println("SetAddr in RoleFS has been successful!")
	return nil
}

// AddOrder called by keeper? Add the storage order in the FileSys.
// hash(uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice)?
// 目前合约中还未对签名进行判断处理
// nonce需要从0开始依次累加
// 调用该函数前，需要admin为RoleFS合约账户赋予MINTER_ROLE权限
func (rfs *ContractModule) AddOrder(roleAddr, rTokenAddr common.Address, uIndex, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, usign, psign []byte, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}
	// check start,end,size
	if size == 0 {
		return errSize
	}
	if end <= start {
		log.Println("start:", start, " end:", end)
		return errEnd
	}
	if (end/86400)*86400 != end {
		log.Println("end:", end)
		return errors.New("end should be divisible by 86400(one day)")
	}
	// check uIndex,pIndex,gIndex,tIndex
	gIndex, err := rfs.checkParam(uIndex, pIndex, UserRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr)
	if err != nil {
		return err
	}
	// check ksigns's length
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	if len(ksigns) < int(gkNum*2/3) {
		return ErrKSignsNE
	}
	// check balance
	pay := big.NewInt(0).Mul(sprice, new(big.Int).SetUint64(end-start))
	manageAndTax := big.NewInt(0).Div(pay, big.NewInt(20)) // pay/100*4 + pay/100*1
	payAndTax := big.NewInt(0).Add(pay, manageAndTax)
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	avail, _, err := fs.GetBalance(uIndex, tIndex)
	if err != nil {
		return err
	}
	if avail.Cmp(payAndTax) < 0 {
		log.Println("payAndTax is", payAndTax, " but avail is", avail)
		return ErrBalNotE
	}
	// check nonce
	_nonce, _, err := fs.GetFsInfoAggOrder(uIndex, pIndex)
	if err != nil {
		return err
	}
	if _nonce != nonce {
		log.Println("nonce:", nonce, " should be", _nonce)
		return errNonce
	}
	// check start
	_time, _, _, err := fs.GetStoreInfo(uIndex, pIndex, tIndex)
	if err != nil {
		return err
	}
	if start < _time {
		log.Println("start:", start, " should be less than time:", _time)
		return errors.New("start error")
	}
	// check whether rolefsAddr has Minter-Role
	if tIndex == 0 {
		erc20 := NewERC20(ERC20Addr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
		has, err := erc20.HasRole(MinterRole, rfs.contractAddress)
		if err != nil {
			return err
		}
		if !has {
			log.Println("rolefsAddr:", rfs.contractAddress.Hex(), " hasn't MinterRole, please setUpRole first")
			return errors.New("rolefsAddr has not MinterRole")
		}
	}

	log.Println("begin AddOrder in RoleFS contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(rfs.hexSk, nil, rfs.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleFSIns.AddOrder(auth, uIndex, pIndex, start, end, size, nonce, tIndex, sprice, usign, psign, ksigns)
		if err != nil {
			retryCount++
			log.Println("AddOrder Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("AddOrder in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	log.Println("AddOrder in RoleFS has been successful!")
	return nil
}

// SubOrder called by keeper? Reduce the storage order in the FileSys.
// hash(uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice)?
// 目前合约中还未对签名信息做判断处理
func (rfs *ContractModule) SubOrder(roleAddr, rTokenAddr common.Address, uIndex, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, usign, psign []byte, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check size,start.end
	if size <= 0 {
		return errSize
	}
	now := uint64(time.Now().Unix())
	if end <= start || end > now {
		log.Println("end:", end, " start:", start, " now:", now)
		return errEndNow
	}
	// check uIndex,pIndex,gIndex,tIndex
	gIndex, err := rfs.checkParam(uIndex, pIndex, UserRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr)
	if err != nil {
		return err
	}
	// check ksigns's length
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	if len(ksigns) < int(gkNum*2/3) {
		return ErrKSignsNE
	}
	// check nonce
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	_, _subNonce, err := fs.GetFsInfoAggOrder(uIndex, pIndex)
	if err != nil {
		return err
	}
	if _subNonce != nonce {
		log.Println("nonce:", nonce, " should be", _subNonce)
		return errNonce
	}
	// check size
	_, _size, _price, err := fs.GetStoreInfo(uIndex, pIndex, tIndex)
	if err != nil {
		return err
	}
	if size > _size {
		log.Println("size:", size, " shouldn't be more than store.size:", _size)
		return errSize
	}
	if sprice.Cmp(_price) > 0 {
		log.Println("sprice:", sprice, " shouldn't be more than store.price:", _price)
		return errSprice
	}

	log.Println("begin SubOrder in RoleFS contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(rfs.hexSk, nil, rfs.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleFSIns.SubOrder(auth, uIndex, pIndex, start, end, size, nonce, tIndex, sprice, usign, psign, ksigns)
		if err != nil {
			retryCount++
			log.Println("SubOrder Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("SubOrder in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	log.Println("SubOrder in RoleFS has been successful!")
	return nil
}

// AddRepair called by keeper. Add the repair order in the FileSys.
// hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "a"), signed by newProvider
func (rfs *ContractModule) AddRepair(roleAddr, rTokenAddr common.Address, pIndex, nPIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, psign []byte, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check pIndex, nPIndex,tIndex,gIndex
	gIndex, err := rfs.checkParam(pIndex, nPIndex, ProviderRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr)
	if err != nil {
		return err
	}
	// check ksigns's length
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	if len(ksigns) < int(gkNum*2/3) {
		return ErrKSignsNE
	}
	// check start,end,size
	if size <= 0 {
		return errSize
	}
	if end <= start {
		log.Println("end:", end, " start:", start)
		return errEnd
	}
	// check lost,lostPaid
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	_, _, _, _, _, _, _lost, _lostPaid, _, _, _, err := fs.GetSettleInfo(pIndex, tIndex)
	if err != nil {
		return err
	}
	if _lost.Cmp(_lostPaid) < 0 {
		log.Println("pIndex:", pIndex, " lost:", _lost, " lostPaid:", _lostPaid, ", lost shouldn't less than lostPaid")
		return errors.New("lost error")
	}
	bal := big.NewInt(0).Sub(_lost, _lostPaid)
	pay := big.NewInt(0).Mul(sprice, new(big.Int).SetUint64(end-start))
	if bal.Cmp(pay) < 0 {
		log.Println("pIndex:", pIndex, " bal:", bal, " pay:", pay, ", bal shouldn't be less than pay")
		return errors.New("pay error")
	}
	// check nonce
	_nonce, _, err := fs.GetFsInfoAggOrder(0, nPIndex)
	if err != nil {
		return err
	}
	if _nonce != nonce {
		log.Println("newPro:", nPIndex, " repairFs.nonce:", _nonce, " nonce:", nonce, ", they should be same")
		return errNonce
	}

	log.Println("begin AddRepair in RoleFS contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(rfs.hexSk, nil, rfs.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleFSIns.AddRepair(auth, pIndex, nPIndex, start, end, size, nonce, tIndex, sprice, psign, ksigns)
		if err != nil {
			retryCount++
			log.Println("AddRepair Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("AddRepair in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	log.Println("AddRepair in RoleFS has been successful!")
	return nil
}

// SubRepair called by keeper. Reduce the repair order in the FileSys.
// hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "s"), signed by newProvider
func (rfs *ContractModule) SubRepair(roleAddr, rTokenAddr common.Address, pIndex, nPIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, psign []byte, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check pIndex,npIndex,gIndex,tIndex
	gIndex, err := rfs.checkParam(pIndex, nPIndex, ProviderRoleType, ProviderRoleType, tIndex, roleAddr, rTokenAddr)
	if err != nil {
		return err
	}
	// check ksigns's length
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	if len(ksigns) < int(gkNum*2/3) {
		return ErrKSignsNE
	}
	// check start,end,size
	if size <= 0 {
		return errSize
	}
	now := uint64(time.Now().Unix())
	if end <= start || end > now {
		log.Println("end:", end, " start:", start, " now:", now)
		return errEndNow
	}
	// check nonce
	_, _, _, _, _, _, fsAddr, err := r.GetGroupInfo(gIndex)
	if err != nil {
		return err
	}
	fs := NewFileSys(fsAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	_, _subNonce, err := fs.GetFsInfoAggOrder(0, nPIndex)
	if err != nil {
		return err
	}
	if _subNonce != nonce {
		log.Println("nonce:", nonce, " should be", _subNonce)
		return errNonce
	}
	// check size
	_, _size, _price, err := fs.GetStoreInfo(0, nPIndex, tIndex)
	if err != nil {
		return err
	}
	if size > _size {
		log.Println("size:", size, " shouldn't be more than store.size:", _size)
		return errSize
	}
	if sprice.Cmp(_price) > 0 {
		log.Println("sprice:", sprice, " shouldn't be more than store.price:", _price)
		return errSprice
	}

	log.Println("begin SubRepair in RoleFS contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(rfs.hexSk, nil, rfs.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleFSIns.SubRepair(auth, pIndex, nPIndex, start, end, size, nonce, tIndex, sprice, psign, ksigns)
		if err != nil {
			retryCount++
			log.Println("SubRepair Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("SubRepair in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	log.Println("SubRepair in RoleFS has been successful!")
	return nil
}

// ProWithdraw called by keeper? Retrieve the Provider's balance in FileSys.
// hash(pIndex, tIndex, pay, lost)?
func (rfs *ContractModule) ProWithdraw(roleAddr, rTokenAddr common.Address, pIndex uint64, tIndex uint32, pay, lost *big.Int, ksigns [][]byte) error {
	client := getClient(rfs.endPoint)
	defer client.Close()
	roleFSIns, err := newRoleFS(rfs.contractAddress, client)
	if err != nil {
		return err
	}

	// check pIndex
	r := NewR(roleAddr, rfs.addr, rfs.hexSk, rfs.txopts, rfs.endPoint)
	addr, err := r.GetAddr(pIndex)
	if err != nil {
		return err
	}
	isActive, isBanned, roleType, _, gIndex, _, err := r.GetRoleInfo(addr)
	if err != nil {
		return err
	}
	if !isActive || isBanned || roleType != ProviderRoleType {
		log.Println("pIndex isActive:", isActive, " isBanned:", isBanned, " roleType:", roleType, ", should be active,not be banned,roleType should be 2")
		return ErrIndex
	}

	// check ksigns's length
	gkNum, err := r.GetGKNum(gIndex)
	if err != nil {
		return err
	}
	l := int(gkNum * 2 / 3)
	le := len(ksigns)
	if le < l {
		log.Println("ksigns length", le, " shouldn't be less than", l)
		return ErrKSignsNE
	}

	log.Println("begin call ProWithdraw in RoleFS contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(rfs.hexSk, nil, rfs.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleFSIns.ProWithdraw(auth, pIndex, tIndex, pay, lost, ksigns)
		if err != nil {
			retryCount++
			log.Println("ProWithdraw Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("ProWithdraw in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		if err != nil {
			return err
		}
		break
	}
	log.Println("ProWithdraw in RoleFS has been successful!")
	return nil
}
