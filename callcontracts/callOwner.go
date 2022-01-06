package callconts

import (
	"log"
	"math/big"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewOwn new a instance of ContractModule. roleAddr: Role contract address
func NewOwn(roleAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string) iface.OwnerInfo {
	own := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: roleAddr,
		endPoint:        endPoint,
	}

	return own
}

// AlterOwner called by admin, to alter Role-contract's owner
func (own *ContractModule) AlterOwner(newOwnerAddr common.Address) error {
	client := getClient(own.endPoint)
	defer client.Close()
	roleIns, err := newRole(own.contractAddress, client)
	if err != nil {
		return err
	}

	if newOwnerAddr.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	owner, err := own.GetOwner()
	if err != nil {
		return err
	}
	if owner.Hex() != own.addr.Hex() {
		log.Println("own.addr:", own.addr.Hex())
		return errNotOwner
	}

	log.Println("begin AlterOwner in Role contract...")
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

		tx, err = roleIns.AlterOwner(auth, newOwnerAddr)
		if err != nil {
			retryCount++
			log.Println("alterOwner Err:", err)
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
			log.Println("AlterOwner in Role transaction fails:", err)
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
	log.Println("AlterOwner in Role has been successful!")
	return nil
}

// GetOwner get the owner-address of Role contract
// 'own.contractAddress' indicates the Role contract address
func (own *ContractModule) GetOwner() (common.Address, error) {
	var ownAddr common.Address

	client := getClient(own.endPoint)
	defer client.Close()
	roleIns, err := newRole(own.contractAddress, client)
	if err != nil {
		return ownAddr, err
	}

	log.Println("begin GetOwner in Role contract...")
	retryCount := 0

	for {
		retryCount++
		ownAddr, err = roleIns.GetOwner(&bind.CallOpts{
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
