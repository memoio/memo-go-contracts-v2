// only prepare to complete calling basic functions of contract now, after that, will judge the input parameters of functions

package callconts

import (
	"log"
	"math/big"
	"memoContract/contracts/role"
	iface "memoContract/interfaces"
	"time"

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

// DeployR deploy a Role contract, called by admin, specify foundation、primaryToken、pledgeK、pledgeP
func (r *ContractModule) DeployR(foundation, primaryToken common.Address, pledgeKeeper, pledgeProvider *big.Int) (common.Address, *role.Role, error) {
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

// Analyze output results based on transaction information
func analyzeOutput(tx *types.Transaction) (uint64, error) {
	
}

// Register called by anyone to get index
func (r *ContractModule) Register(roleAddr, addr common.Address, sign []byte) (uint64, error) {
	var index uint64
	roleIns, err := newRole(roleAddr)
	if err != nil {
		return index, err
	}

	log.Println("begin Register in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return index, errMA
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
				return index, err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("Register in Role transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return index, err
			}
			continue
		}
		break
	}
	log.Println("Register in Role has been successful! The index you get is ", index)
	return index, nil
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
func (r *ContractModule) CreateGroup(roleAddr common.Address, kindexes []uint64, level uint16) (uint64, error) {
	var gindex uint64

	roleIns, err := newRole(roleAddr)
	if err != nil {
		return gindex, err
	}

	log.Println("begin CreateGroup in Role contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(r.hexSk, nil, r.txopts)
		if errMA != nil {
			return gindex, errMA
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
				return gindex, err
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
