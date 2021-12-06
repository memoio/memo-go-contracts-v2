package callconts

import (
	"log"
	"math/big"
	filesys "memoContract/contracts/filesystem"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewFileSys new a instance of ContractModule
func NewFileSys(addr common.Address, hexSk string, txopts *TxOpts) iface.FileSysInfo {
	fs := &ContractModule{
		addr:   addr,
		hexSk:  hexSk,
		txopts: txopts,
	}

	return fs
}

// newFileSys new an instance of FileSys contract, 'fsAddr' indicates FileSys contract address.
func newFileSys(fsAddr common.Address) (*filesys.FileSys, error) {
	fsIns, err := filesys.NewFileSys(fsAddr, getClient(EndPoint))
	if err != nil {
		return nil, err
	}
	return fsIns, nil
}

// DeployFileSys deploy a FileSys contract, called by admin.
// Called after the admin calls the CreateGroup function in the Role Contract.
// 'r' indicates role-contract address, 'rfs' indicates RoleFS-contract address.
func (fs *ContractModule) DeployFileSys(founder, gIndex uint64, r, rfs common.Address, keepers []uint64) (common.Address, *filesys.FileSys, error) {
	var fsAddr, fsAddress common.Address
	var fsInstance, fsIns *filesys.FileSys
	var err error

	log.Println("begin deploy FileSys contract...")
	client := getClient(EndPoint)
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0
	for {
		auth, errMA := makeAuth(fs.hexSk, nil, fs.txopts)
		if errMA != nil {
			return fsAddr, nil, errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		fsAddress, tx, fsInstance, err = filesys.DeployFileSys(auth, client, founder, gIndex, r, rfs, keepers)
		if fsAddress.String() != InvalidAddr {
			fsAddr = fsAddress
			fsIns = fsInstance
		}
		if err != nil {
			retryCount++
			log.Println("deploy FileSys Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return fsAddr, fsIns, err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("deploy FileSys transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return fsAddr, fsIns, err
			}
			continue
		}
		break
	}
	log.Println("FileSys has been successfully deployed! The address is ", fsAddr.Hex())
	return fsAddr, fsIns, nil
}

// GetFsInfo Get information of filesystem.
func (fs *ContractModule) GetFsInfo(fsAddr common.Address, uIndex uint64) (bool, uint32, error) {
	var result struct {
		IsActive   bool
		TokenIndex uint32
	}
	result.IsActive = false
	result.TokenIndex = 0

	fsIns, err := newFileSys(fsAddr)
	if err != nil {
		return result.IsActive, result.TokenIndex, err
	}

	retryCount := 0
	for {
		retryCount++
		// multiple return values are returned as a structure because named return values are used in the contract function
		result, err = fsIns.GetFsInfo(&bind.CallOpts{
			From: fs.addr,
		}, uIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return result.IsActive, result.TokenIndex, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return result.IsActive, result.TokenIndex, nil
	}
}

// GetFsInfoAggOrder Get information of aggOrder.
func (fs *ContractModule) GetFsInfoAggOrder(fsAddr common.Address, uIndex uint64, pIndex uint64) (uint64, uint64, error) {
	var nonce uint64
	var subNonce uint64

	fsIns, err := newFileSys(fsAddr)
	if err != nil {
		return nonce, subNonce, err
	}

	retryCount := 0
	for {
		retryCount++
		nonce, subNonce, err = fsIns.GetFsInfoAggOrder(&bind.CallOpts{
			From: fs.addr,
		}, uIndex, pIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return nonce, subNonce, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return nonce, subNonce, nil
	}
}

// GetStoreInfo Get information of storage order.
func (fs *ContractModule) GetStoreInfo(fsAddr common.Address, uIndex uint64, pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, error) {
	var _time, size uint64
	var price *big.Int

	fsIns, err := newFileSys(fsAddr)
	if err != nil {
		return _time, size, price, err
	}

	retryCount := 0
	for {
		retryCount++
		_time, size, price, err = fsIns.GetStoreInfo(&bind.CallOpts{
			From: fs.addr,
		}, uIndex, pIndex, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return _time, size, price, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return _time, size, price, nil
	}
}

// GetChannelInfo Get information of channel.
func (fs *ContractModule) GetChannelInfo(fsAddr common.Address, uIndex uint64, pIndex uint64, tIndex uint32) (*big.Int, uint64, uint64, error) {
	var nonce, expire uint64
	var amount *big.Int

	fsIns, err := newFileSys(fsAddr)
	if err != nil {
		return amount, nonce, expire, err
	}

	retryCount := 0
	for {
		retryCount++
		amount, nonce, expire, err = fsIns.GetChannelInfo(&bind.CallOpts{
			From: fs.addr,
		}, uIndex, pIndex, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return amount, nonce, expire, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return amount, nonce, expire, nil
	}
}

// GetSettleInfo Get information of settlement.
func (fs *ContractModule) GetSettleInfo(fsAddr common.Address, pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var _time, size uint64
	var price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid *big.Int

	fsIns, err := newFileSys(fsAddr)
	if err != nil {
		return _time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err
	}

	retryCount := 0
	for {
		retryCount++
		_time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err = fsIns.GetSettleInfo(&bind.CallOpts{
			From: fs.addr,
		}, pIndex, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return _time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return _time, size, price, maxPay, hasPaid, canPay, lost, lostPaid, managePay, endPaid, linearPaid, nil
	}
}

// GetBalance Get income balance of rindex.
func (fs *ContractModule) GetBalance(fsAddr common.Address, rIndex uint64, tIndex uint32) (*big.Int, *big.Int, error) {
	var avail, tmp *big.Int

	fsIns, err := newFileSys(fsAddr)
	if err != nil {
		return avail, tmp, err
	}

	retryCount := 0
	for {
		retryCount++
		avail, tmp, err = fsIns.GetBalance(&bind.CallOpts{
			From: fs.addr,
		}, rIndex, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return avail, tmp, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return avail, tmp, nil
	}
}
