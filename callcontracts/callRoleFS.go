package callconts

import (
	"log"
	"math/big"
	rolefs "memoContract/contracts/rolefs"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewRFS new a instance of ContractModule
func NewRFS(addr common.Address, hexSk string, txopts *TxOpts) iface.RoleFSInfo {
	rfs := &ContractModule{
		addr:   addr,
		hexSk:  hexSk,
		txopts: txopts,
	}

	return rfs
}

// DeployRoleFS deploy an RoleFS contract, called by admin
func (rfs *ContractModule) DeployRoleFS() (common.Address, *rolefs.RoleFS, error) {
	var roleFSAddr, roleFSAddress common.Address
	var roleFSInstance, roleFSIns *rolefs.RoleFS
	var err error

	log.Println("begin deploy RoleFS contract...")
	client := getClient(EndPoint)
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
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
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
		if err != nil {
			checkRetryCount++
			log.Println("deploy RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return roleFSAddr, roleFSIns, err
			}
			continue
		}
		break
	}
	log.Println("RoleFS has been successfully deployed! The address is ", roleFSAddr.Hex())
	return roleFSAddr, roleFSIns, nil
}

func newRoleFS(roleFSAddr common.Address) (*rolefs.RoleFS, error) {
	roleFSIns, err := rolefs.NewRoleFS(roleFSAddr, getClient(EndPoint))
	if err != nil {
		return nil, err
	}
	return roleFSIns, nil
}

// SetAddr called by admin, which is the deployer. Set some address type variables.
func (rfs *ContractModule) SetAddr(roleFS, issuan, role, fileSys, rtoken common.Address) error {
	roleFSIns, err := newRoleFS(roleFS)
	if err != nil {
		return err
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
			log.Println("SetAddr in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("SetAddr in RoleFS has been successful!")
	return nil
}

// AddOrder called by keeper? Add the storage order in the FileSys.
// hash(uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice)?
func (rfs *ContractModule) AddOrder(roleFS common.Address, uIndex, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, usign, psign []byte, ksigns [][]byte) error {
	roleFSIns, err := newRoleFS(roleFS)
	if err != nil {
		return err
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
			log.Println("AddOrder in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("AddOrder in RoleFS has been successful!")
	return nil
}

// SubOrder called by keeper? Reduce the storage order in the FileSys.
// hash(uIndex, pIndex, _start, end, _size, nonce, tIndex, sPrice)?
func (rfs *ContractModule) SubOrder(roleFS common.Address, uIndex, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, usign, psign []byte, ksigns [][]byte) error {
	roleFSIns, err := newRoleFS(roleFS)
	if err != nil {
		return err
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

		tx, err = roleFSIns.AddOrder(auth, uIndex, pIndex, start, end, size, nonce, tIndex, sprice, usign, psign, ksigns)
		if err != nil {
			retryCount++
			log.Println("SubOrder Err:", err)
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
			log.Println("SubOrder in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("SubOrder in RoleFS has been successful!")
	return nil
}

// AddRepair called by keeper. Add the repair order in the FileSys.
// hash(pIndex, _start, end, _size, nonce, tIndex, sPrice), signed by newProvider
func (rfs *ContractModule) AddRepair(roleFS common.Address, pIndex, nPIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, psign []byte, ksigns [][]byte) error {
	roleFSIns, err := newRoleFS(roleFS)
	if err != nil {
		return err
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
			log.Println("AddRepair in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("AddRepair in RoleFS has been successful!")
	return nil
}

// SubRepair called by keeper. Reduce the repair order in the FileSys.
// hash(pIndex, _start, end, _size, nonce, tIndex, sPrice), signed by newProvider
func (rfs *ContractModule) SubRepair(roleFS common.Address, pIndex, nPIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, psign []byte, ksigns [][]byte) error {
	roleFSIns, err := newRoleFS(roleFS)
	if err != nil {
		return err
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
			log.Println("SubRepair in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("SubRepair in RoleFS has been successful!")
	return nil
}

// ProWithdraw called by keeper? Retrieve the Provider's balance in FileSys.
// hash(pIndex, tIndex, pay, lost)?
func (rfs *ContractModule) ProWithdraw(roleFS common.Address, pIndex uint64, tIndex uint32, pay, lost *big.Int, ksigns [][]byte) error {
	roleFSIns, err := newRoleFS(roleFS)
	if err != nil {
		return err
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
			log.Println("ProWithdraw in RoleFS transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("ProWithdraw in RoleFS has been successful!")
	return nil
}
