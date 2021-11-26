// only prepare to complete calling basic functions of contract now, after that, will judge the input parameters of functions
// And also need to add Getter functions.

package callconts

import (
	"log"
	"math/big"
	"memoContract/contracts/role"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewR new a instance of ContractModule
func NewR(addr common.Address, hexSk string, txopts *TxOpts) iface.RoleInfo {
	r := &ContractModule{
		addr:   addr,
		hexSk:  hexSk,
		txopts: txopts,
	}

	return r
}

// DeployRole deploy a Role contract, called by admin, specify foundation、primaryToken、pledgeK、pledgeP
func (r *ContractModule) DeployRole(foundation, primaryToken common.Address, pledgeKeeper, pledgeProvider *big.Int) (common.Address, *role.Role, error) {
	var roleAddr, roleAddress common.Address
	var roleInstance, roleIns *role.Role
	var err error

	log.Println("begin deploy Role contract...")
	client := getClient(EndPoint)
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return roleAddr, nil, errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		roleAddress, tx, roleInstance, err = role.DeployRole(auth, client, foundation, primaryToken, pledgeKeeper, pledgeProvider)
		if roleAddress.String() != InvalidAddr {
			roleAddr = roleAddress
			roleIns = roleInstance
		}
		if err != nil {
			retryCount++
			log.Println("deploy Role Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return roleAddr, roleIns, err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("deploy Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return roleAddr, roleIns, err
			}
			continue
		}
		break
	}
	log.Println("Role has been successfully deployed! The address is ", roleAddr.Hex())
	return roleAddr, roleIns, nil
}

// newRole new an instance of Role contract, 'roleAddr' indicates Role contract address
func newRole(roleAddr common.Address) (*role.Role, error) {
	roleIns, err := role.NewRole(roleAddr, getClient(EndPoint))
	if err != nil {
		return nil, err
	}
	return roleIns, nil
}

// SetPI callled by admin, set pledgePool-address、 issuance-address and rolefs-address
func (r *ContractModule) SetPI(roleAddr, pledgePoolAddr, issuAddr, rolefsAddr common.Address) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin SetPI in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.SetPI(auth, pledgePoolAddr, issuAddr, rolefsAddr)
		if err != nil {
			retryCount++
			log.Println("setPI Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("SetPI in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("SetPI in Role has been successful!")
	return nil
}

// Register called by anyone to get index
func (r *ContractModule) Register(roleAddr, addr common.Address, sign []byte) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin Register in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.Register(auth, addr, sign)
		if err != nil {
			retryCount++
			log.Println("register Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("Register in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Register in Role has been successful!")
	return nil
}

// RegisterKeeper called by anyone to register Keeper role
func (r *ContractModule) RegisterKeeper(roleAddr common.Address, index uint64, blskey []byte, sign []byte) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin RegisterKeeper in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.RegisterKeeper(auth, index, blskey, sign)
		if err != nil {
			retryCount++
			log.Println("registerKeeper Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("RegisterKeeper in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("RegisterKeeper in Role has been successful!")
	return nil
}

// RegisterProvider called by anyone to register Provider role
func (r *ContractModule) RegisterProvider(roleAddr common.Address, index uint64, sign []byte) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin RegisterProvider in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.RegisterProvider(auth, index, sign)
		if err != nil {
			retryCount++
			log.Println("registerProvider Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("RegisterProvider in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("RegisterProvider in Role has been successful!")
	return nil
}

// RegisterUser called by anyone to register User role
func (r *ContractModule) RegisterUser(roleAddr common.Address, index uint64, gindex uint64, tindex uint32, blskey []byte, sign []byte) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin RegisterUser in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.RegisterUser(auth, index, gindex, tindex, blskey, sign)
		if err != nil {
			retryCount++
			log.Println("registerUser Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("RegisterUser in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("RegisterUser in Role has been successful!")
	return nil
}

// RegisterToken called by admin to register token. Once token is registered, it is supported by memo.
func (r *ContractModule) RegisterToken(roleAddr common.Address, tokenAddr common.Address) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin RegisterToken in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.RegisterToken(auth, tokenAddr)
		if err != nil {
			retryCount++
			log.Println("registerToken Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("RegisterToken in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("RegisterToken in Role has been successful!")
	return nil
}

// CreateGroup called by admin to create a group.
func (r *ContractModule) CreateGroup(roleAddr common.Address, kindexes []uint64, level uint16) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin CreateGroup in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.CreateGroup(auth, kindexes, level)
		if err != nil {
			retryCount++
			log.Println("CreateGroup Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("RegisterToken in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("RegisterToken in Role has been successful!")
	return nil
}

// SetGF called by admin to set fsAddress for group after CreateGroup and deployFileSys.
func (r *ContractModule) SetGF(roleAddr, fsAddr common.Address, gIndex uint64) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin SetGF in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.SetGF(auth, gIndex, fsAddr)
		if err != nil {
			retryCount++
			log.Println("SetGF Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("SetGF in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("SetGF in Role has been successful!")
	return nil
}

// AddKeeperToGroup called by admin.
func (r *ContractModule) AddKeeperToGroup(roleAddr common.Address, kIndex uint64, gIndex uint64) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin AddKeeperToGroup in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.AddKeeperToGroup(auth, kIndex, gIndex)
		if err != nil {
			retryCount++
			log.Println("AddKeeperToGroup Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("AddKeeperToGroup in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("AddKeeperToGroup in Role has been successful!")
	return nil
}

// AddProviderToGroup called by provider or called by others.
func (r *ContractModule) AddProviderToGroup(roleAddr common.Address, pIndex uint64, gIndex uint64, sign []byte) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin AddProviderToGroup in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.AddProviderToGroup(auth, pIndex, gIndex, sign)
		if err != nil {
			retryCount++
			log.Println("AddProviderToGroup Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("AddProviderToGroup in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("AddProviderToGroup in Role has been successful!")
	return nil
}

// SetPledgeMoney called by admin to set the amount that the keeper and provider needs to pledge.
func (r *ContractModule) SetPledgeMoney(roleAddr common.Address, kpledge *big.Int, ppledge *big.Int) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin SetPledgeMoney in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.SetPledgeMoney(auth, kpledge, ppledge)
		if err != nil {
			retryCount++
			log.Println("SetPledgeMoney Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("SetPledgeMoney in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("SetPledgeMoney in Role has been successful!")
	return nil
}

// Recharge called by user or called by others.
func (r *ContractModule) Recharge(roleAddr common.Address, uIndex uint64, tIndex uint32, money *big.Int, sign []byte) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin Recharge in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.Recharge(auth, uIndex, tIndex, money, sign)
		if err != nil {
			retryCount++
			log.Println("Recharge Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("Recharge in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Recharge in Role has been successful!")
	return nil
}

// WithdrawFromFs called by memo-role or called by others.
func (r *ContractModule) WithdrawFromFs(roleAddr common.Address, rIndex uint64, tIndex uint32, amount *big.Int, sign []byte) error {
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return err
	}

	log.Println("begin WithdrawFromFs in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = roleIns.WithdrawFromFs(auth, rIndex, tIndex, amount, sign)
		if err != nil {
			retryCount++
			log.Println("WithdrawFromFs Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("WithdrawFromFs in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("WithdrawFromFs in Role has been successful!")
	return nil
}

// GetAddrsNum get the number of registered addresses.
func (r *ContractModule) GetAddrsNum(roleAddr common.Address) (uint64, error) {
	var anum uint64
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return anum, err
	}

	retryCount := 0
	for {
		retryCount++
		anum, err = roleIns.GetAddrsNum(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return anum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return anum, nil
	}
}

// GetAddr get address by role index.
func (r *ContractModule) GetAddr(roleAddr common.Address, rIndex uint64) (common.Address, error) {
	var addr common.Address
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return addr, err
	}

	retryCount := 0
	rIndex-- // get address by array index actually in contract, which is rIndex minus 1
	for {
		retryCount++
		addr, err = roleIns.GetAddr(&bind.CallOpts{
			From: r.addr,
		}, rIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return addr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return addr, nil
	}
}

// GetRoleIndex get the account role index by address.
func (r *ContractModule) GetRoleIndex(roleAddr common.Address, addr common.Address) (uint64, error) {
	var rIndex uint64
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return rIndex, err
	}

	retryCount := 0
	for {
		retryCount++
		rIndex, err = roleIns.GetRoleIndex(&bind.CallOpts{
			From: r.addr,
		}, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return rIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return rIndex, nil
	}
}

// GetRoleInfo get account information by address.
func (r *ContractModule) GetRoleInfo(roleAddr common.Address, addr common.Address) (bool, bool, uint8, uint64, uint64, []byte, error) {
	var isActive, isBanned bool
	var roleType uint8
	var index, gIndex uint64
	var extra []byte

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return isActive, isBanned, roleType, index, gIndex, extra, err
	}

	retryCount := 0
	for {
		retryCount++
		isActive, isBanned, roleType, index, gIndex, extra, err = roleIns.GetRoleInfo(&bind.CallOpts{
			From: r.addr,
		}, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return isActive, isBanned, roleType, index, gIndex, extra, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isActive, isBanned, roleType, index, gIndex, extra, nil
	}
}

// GetGroupsNum get the number of groups.
func (r *ContractModule) GetGroupsNum(roleAddr common.Address) (uint64, error) {
	var gnum uint64

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return gnum, err
	}

	retryCount := 0
	for {
		retryCount++
		gnum, err = roleIns.GetGroupsNum(&bind.CallOpts{
			From: r.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return gnum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return gnum, nil
	}
}

// GetGroupInfo get group information by gIndex.
func (r *ContractModule) GetGroupInfo(roleAddr common.Address, gIndex uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error) {
	var isActive, isBanned, isReady bool
	var level uint16
	var size, price *big.Int
	var fsAddr common.Address

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return isActive, isBanned, isReady, level, size, price, fsAddr, err
	}

	retryCount := 0
	gIndex-- // get group info by array index actually in contract, which is gIndex minus 1
	for {
		retryCount++
		isActive, isBanned, isReady, level, size, price, fsAddr, err = roleIns.GetGroupInfo(&bind.CallOpts{
			From: r.addr,
		}, gIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return isActive, isBanned, isReady, level, size, price, fsAddr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isActive, isBanned, isReady, level, size, price, fsAddr, nil
	}
}

// GetAddrGindex get account's address and its gIndex by rIndex.
func (r *ContractModule) GetAddrGindex(roleAddr common.Address, rIndex uint64) (common.Address, uint64, error) {
	var addr common.Address
	var gIndex uint64

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return addr, gIndex, err
	}

	retryCount := 0
	for {
		retryCount++
		addr, gIndex, err = roleIns.GetAddrGindex(&bind.CallOpts{
			From: r.addr,
		}, rIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return addr, gIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return addr, gIndex, nil
	}
}

// GetGKNum get the number of keepers in the group.
func (r *ContractModule) GetGKNum(roleAddr common.Address, gIndex uint64) (uint64, error) {
	var gkNum uint64

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return gkNum, err
	}

	retryCount := 0
	for {
		retryCount++
		gkNum, err = roleIns.GetGKNum(&bind.CallOpts{
			From: r.addr,
		}, gIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return gkNum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return gkNum, nil
	}
}

// GetGPNum get the number of providers in the group.
func (r *ContractModule) GetGPNum(roleAddr common.Address, gIndex uint64) (uint64, error) {
	var gpNum uint64

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return gpNum, err
	}

	retryCount := 0
	for {
		retryCount++
		gpNum, err = roleIns.GetGPNum(&bind.CallOpts{
			From: r.addr,
		}, gIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return gpNum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return gpNum, nil
	}
}

// GetGroupK get keeper role index by gIndex and keeper array index.
func (r *ContractModule) GetGroupK(roleAddr common.Address, gIndex uint64, index uint64) (uint64, error) {
	var kIndex uint64

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return kIndex, err
	}

	retryCount := 0
	gIndex-- // get group info by array index actually in contract, which is gIndex minus 1
	for {
		retryCount++
		kIndex, err = roleIns.GetGroupK(&bind.CallOpts{
			From: r.addr,
		}, gIndex, index)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return kIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return kIndex, nil
	}
}

// GetGroupP get provider role index by gIndex and provider array index.
func (r *ContractModule) GetGroupP(roleAddr common.Address, gIndex uint64, index uint64) (uint64, error) {
	var pIndex uint64

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return pIndex, err
	}

	retryCount := 0
	gIndex-- // get group info by array index actually in contract, which is gIndex minus 1
	for {
		retryCount++
		pIndex, err = roleIns.GetGroupP(&bind.CallOpts{
			From: r.addr,
		}, gIndex, index)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return pIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return pIndex, nil
	}
}
