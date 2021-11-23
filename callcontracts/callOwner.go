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

// NewOwn new a instance of ContractModule
func NewOwn(addr common.Address, hexSk string, txopts *TxOpts) iface.OwnerInfo {
	own := &ContractModule{
		addr:   addr,
		hexSk:  hexSk,
		txopts: txopts,
	}

	return own
}

// newOwner new a instance of Owner contract
func newOwner(ownerAddr common.Address) (*role.Owner, error) {
	ownerIns, err := role.NewOwner(ownerAddr, getClient(EndPoint))
	if err != nil {
		return nil, err
	}
	return ownerIns, nil
}

// AlterOwner called by admin, to alter Role-contract's owner
// 'ownerAddr' indicates the Owner contract address
func (own *ContractModule) AlterOwner(ownerAddr common.Address, newOwnerAddr common.Address) error {
	ownerIns, err := newOwner(ownerAddr)
	if err != nil {
		return err
	}

	log.Println("begin AlterOwner in Owner contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(own.hexSk, nil, own.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = ownerIns.AlterOwner(auth, newOwnerAddr)
		if err != nil {
			retryCount++
			log.Println("alterOwner Err:", err)
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
			log.Println("AlterOwner in Owner transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("AlterOwner in Owner has been successful!")
	return nil
}

// GetOwner get the owner-address of Role contract
// 'ownerAddr' indicates the Owner contract address
func (own *ContractModule) GetOwner(ownerAddr common.Address) (common.Address, error) {
	var ownAddr common.Address

	ownerIns, err := newOwner(ownerAddr)
	if err != nil {
		return ownAddr, err
	}

	log.Println("begin GetOwner in Owner contract...")
	retryCount := 0

	for {
		retryCount++
		ownAddr, err = ownerIns.GetOwner(&bind.CallOpts{
			From: own.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return ownAddr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return ownAddr, nil
	}
}
