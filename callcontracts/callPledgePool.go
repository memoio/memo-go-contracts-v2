// only prepare to complete calling basic functions of contract now, after that, will judge the input parameters of functions
// And also need to add Getter functions.
// Also need to add 'Approve' function in some functions related to transfer operations.

package callconts

import (
	"log"
	"math/big"
	"memoContract/contracts/pledgepool"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewPledgePool new an instance of ContractModule
func NewPledgePool(addr common.Address, hexSk string, txopts *TxOpts) iface.PledgePoolInfo {
	p := &ContractModule{
		addr:   addr,
		hexSk:  hexSk,
		txopts: txopts,
	}

	return p
}

// newPledgePool new an instance of PledgePool contract, 'pledgepAddr' indicates PledgePool contract address.
func newPledgePool(pledgepAddr common.Address) (*pledgepool.PledgePool, error) {
	ppIns, err := pledgepool.NewPledgePool(pledgepAddr, getClient(EndPoint))
	if err != nil {
		return nil, err
	}
	return ppIns, nil
}

// DeployPledgePool deploy a PledgePool contract, called by admin.
// primeToken、rToken contract address、role contract address.
func (p *ContractModule) DeployPledgePool(primeToken common.Address, rToken common.Address, role common.Address) (common.Address, *pledgepool.PledgePool, error) {
	var pledgepAddr, pledgepAddress common.Address
	var pledgepInstance, pledgepIns *pledgepool.PledgePool
	var err error

	log.Println("begin deploy PledgePool contract...")
	client := getClient(EndPoint)
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0
	for {
		auth, errMA := makeAuth(p.hexSk, nil, p.txopts)
		if errMA != nil {
			return pledgepAddr, nil, errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		pledgepAddress, tx, pledgepInstance, err = pledgepool.DeployPledgePool(auth, client, primeToken, rToken, role)
		if pledgepAddress.String() != InvalidAddr {
			pledgepAddr = pledgepAddress
			pledgepIns = pledgepInstance
		}
		if err != nil {
			retryCount++
			log.Println("deploy PledgePool Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return pledgepAddr, pledgepIns, err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("deploy PledgePool transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return pledgepAddr, pledgepIns, err
			}
			continue
		}
		break
	}
	log.Println("PledgePool has been successfully deployed! The address is ", pledgepAddr.Hex())
	return pledgepAddr, pledgepIns, nil
}

// Pledge money.
// Called by the account itself or by another account on its behalf.
// 调用前需要index指示的账户approve本合约（也就是pledgePool合约）账户指定的金额（也就是value）
func (p *ContractModule) Pledge(pledgepAddr, erc20Addr, roleAddr common.Address, rindex uint64, value *big.Int, sign []byte) error {
	pledgepIns, err := newPledgePool(pledgepAddr)
	if err != nil {
		return err
	}

	addr, err := p.GetAddr(roleAddr, rindex)
	if err != nil {
		return err
	}

	// check the rindex is not being banned
	_, isBanned, _, _, _, _, err := p.GetRoleInfo(roleAddr, addr)
	if isBanned {
		return ErrIsBanned
	}

	// check whether the allowance[addr][pledgePoolAddr] is not less than value, if not, will set allowance automatically by code.
	allo, err := p.Allowance(erc20Addr, addr, pledgepAddr)
	if err != nil {
		return err
	}
	if allo.Cmp(value) < 0 {
		tmp := big.NewInt(0)
		tmp.Sub(value, allo)
		log.Println("The allowance of ", addr.Hex(), " to ", pledgepAddr.Hex(), " is not enough, also need to add allowance", tmp)
		// if called by the account itself， then call IncreaseAllowance directly.
		if sign == nil && p.addr.Hex() == addr.Hex() {
			err = p.IncreaseAllowance(erc20Addr, pledgepAddr, tmp)
			if err != nil {
				return err
			}
		} else { // otherwise, quit Pledge
			return ErrAlloNotE
		}
	}

	log.Println("begin Pledge in PledgePool contract with value", value, " and rindex", rindex, " ...")

	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0
	for {
		auth, errMA := makeAuth(p.hexSk, nil, p.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = pledgepIns.Pledge(auth, rindex, value, sign)
		if err != nil {
			retryCount++
			log.Println("Pledge Err:", err)
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
			log.Println("Pledge transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Pledge in PledgePool contract has been successful!")
	return nil
}

// Withdraw Called by the account itself or by another account on its behalf.
// withdraw its balance from PledgePool.
func (p *ContractModule) Withdraw(pledgepAddr, roleAddr, rTokenAddr common.Address, rindex uint64, tindex uint32, value *big.Int, sign []byte) error {
	pledgepIns, err := newPledgePool(pledgepAddr)
	if err != nil {
		return err
	}

	// check if rindex is banned
	addr, err := p.GetAddr(roleAddr, rindex)
	if err != nil {
		return err
	}
	_, isBanned, _, _, _, _, err := p.GetRoleInfo(roleAddr, addr)
	if isBanned {
		return ErrIsBanned
	}
	// check if tindex is valid
	isValid, err := p.IsValid(rTokenAddr, tindex)
	if !isValid {
		return ErrTIndex
	}

	log.Println("begin Withdraw in PledgePool contract with value", value, " and rindex", rindex, " and tindex", tindex, " ...")

	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0
	for {
		auth, errMA := makeAuth(p.hexSk, nil, p.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = pledgepIns.Withdraw(auth, rindex, tindex, value, sign)
		if err != nil {
			retryCount++
			log.Println("Withdraw Err:", err)
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
			log.Println("Withdraw transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Withdraw in PledgePool contract has been successful!")
	return nil
}

// GetPledge Get all pledge amount in specified token.
func (p *ContractModule) GetPledge(pledgepAddr common.Address, tindex uint32) (*big.Int, error) {
	var amount *big.Int

	pledgepIns, err := newPledgePool(pledgepAddr)
	if err != nil {
		return amount, err
	}

	retryCount := 0
	for {
		retryCount++
		amount, err = pledgepIns.GetPledge(&bind.CallOpts{
			From: p.addr,
		}, tindex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return amount, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return amount, nil
	}
}

// GetBalanceInPPool Get balance of the account related rindex in specified token.
func (p *ContractModule) GetBalanceInPPool(pledgepAddr common.Address, rindex uint64, tindex uint32) (*big.Int, error) {
	var amount *big.Int

	pledgepIns, err := newPledgePool(pledgepAddr)
	if err != nil {
		return amount, err
	}

	retryCount := 0
	for {
		retryCount++
		amount, err = pledgepIns.GetBalance(&bind.CallOpts{
			From: p.addr,
		}, rindex, tindex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return amount, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return amount, nil
	}
}
