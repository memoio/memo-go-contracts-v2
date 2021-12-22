package callconts

import (
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// SetUpRole Called by who has DEFAULT_ADMIN_ROLE. Set role to addr.
// role: DEFAULT_ADMIN_ROLE(0)、MINTER_ROLE(1)、PAUSER_ROLE(2)
func (ac *ContractModule) SetUpRole(role uint8, addr common.Address) error {
	acIns, err := newERC20(ac.contractAddress)
	if err != nil {
		return err
	}

	hasAdmin, err := ac.HasRole(uint8(0), ac.addr)
	if err != nil {
		return err
	}
	if !hasAdmin {
		return ErrNoAdminRight
	}

	if role > 2 {
		return errAccessControlRole
	}

	if addr.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	log.Println("begin SetUpRole to", addr.Hex(), " with role", role, " in AccessControl contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(ac.hexSk, nil, ac.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = acIns.SetUpRole(auth, role, addr)
		if err != nil {
			retryCount++
			log.Println("SetUpRole Err:", err)
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
		if err != nil {
			checkRetryCount++
			log.Println("SetUpRole in AccessControl transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("SetUpRole in AccessControl has been successful!")
	return nil
}

// RevokeRole Called by who has DEFAULT_ADMIN_ROLE. Revoke other account's role.
func (ac *ContractModule) RevokeRole(role uint8, addr common.Address) error {
	acIns, err := newERC20(ac.contractAddress)
	if err != nil {
		return err
	}

	hasAdmin, err := ac.HasRole(uint8(0), ac.addr)
	if err != nil {
		return err
	}
	if !hasAdmin {
		return ErrNoAdminRight
	}

	if role > 2 {
		return errAccessControlRole
	}

	if addr.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	log.Println("begin RevokeRole to", addr.Hex(), " with role", role, " in AccessControl contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(ac.hexSk, nil, ac.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = acIns.RevokeRole(auth, role, addr)
		if err != nil {
			retryCount++
			log.Println("RevokeRole Err:", err)
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
		if err != nil {
			checkRetryCount++
			log.Println("RevokeRole in AccessControl transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("RevokeRole in AccessControl has been successful!")
	return nil
}

// RenounceRole Account renounce its role .
func (ac *ContractModule) RenounceRole(role uint8) error {
	acIns, err := newERC20(ac.contractAddress)
	if err != nil {
		return err
	}

	log.Println("begin RenounceRole", role, " in AccessControl contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(ac.hexSk, nil, ac.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = acIns.RenounceRole(auth, role)
		if err != nil {
			retryCount++
			log.Println("RenounceRole Err:", err)
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
		if err != nil {
			checkRetryCount++
			log.Println("RenounceRole in AccessControl transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("RenounceRole in AccessControl has been successful!")
	return nil
}

// Pause Set to true to prohibit transfer operation in erc20. Called by who has PAUSER_ROLE.
func (ac *ContractModule) Pause() error {
	acIns, err := newERC20(ac.contractAddress)
	if err != nil {
		return err
	}

	hasPause, err := ac.HasRole(uint8(2), ac.addr)
	if err != nil {
		return err
	}
	if !hasPause {
		return ErrNoPauseRight
	}

	log.Println("begin Pause in AccessControl contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(ac.hexSk, nil, ac.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = acIns.Pause(auth)
		if err != nil {
			retryCount++
			log.Println("Pause Err:", err)
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
		if err != nil {
			checkRetryCount++
			log.Println("Pause in AccessControl transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Pause in AccessControl has been successful!")
	return nil
}

// Unpause Set to false to allow transfer operation in erc20. Called by who has PAUSER_ROLE.
func (ac *ContractModule) Unpause() error {
	acIns, err := newERC20(ac.contractAddress)
	if err != nil {
		return err
	}

	hasPause, err := ac.HasRole(uint8(2), ac.addr)
	if err != nil {
		return err
	}
	if !hasPause {
		return ErrNoPauseRight
	}

	log.Println("begin Unpause in AccessControl contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(ac.hexSk, nil, ac.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = acIns.Unpause(auth)
		if err != nil {
			retryCount++
			log.Println("Unpause Err:", err)
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
		if err != nil {
			checkRetryCount++
			log.Println("Unpause in AccessControl transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Unpause in AccessControl has been successful!")
	return nil
}

// HasRole Check whether the account has the right.
func (ac *ContractModule) HasRole(role uint8, addr common.Address) (bool, error) {
	var has bool

	acIns, err := newERC20(ac.contractAddress)
	if err != nil {
		return has, err
	}

	retryCount := 0
	for {
		retryCount++
		has, err = acIns.HasRole(&bind.CallOpts{
			From: ac.addr,
		}, role, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return has, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return has, nil
	}
}

// GetPaused Check whether the account has the right.
func (ac *ContractModule) GetPaused() (bool, error) {
	var isPaused bool

	acIns, err := newERC20(ac.contractAddress)
	if err != nil {
		return isPaused, err
	}

	retryCount := 0
	for {
		retryCount++
		isPaused, err = acIns.GetPaused(&bind.CallOpts{
			From: ac.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return isPaused, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isPaused, nil
	}
}
