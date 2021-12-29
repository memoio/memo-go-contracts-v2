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
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewIssu new a instance of ContractModule. issuAddr: Issuance contract address
func NewIssu(issuAddr, addr common.Address, hexSk string, txopts *TxOpts) iface.IssuanceInfo {
	issu := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: issuAddr,
	}

	return issu
}

// DeployIssuance deploy an Issuance contract, called by admin
func (issu *ContractModule) DeployIssuance(rolefsAddr common.Address) (common.Address, *role.Issuance, error) {
	var issuAddr, issuAddress common.Address
	var issuInstance, issuIns *role.Issuance
	var err error

	log.Println("begin deploy Issuance contract...")
	client := getClient(EndPoint)
	defer client.Close()
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(issu.hexSk, nil, issu.txopts)
		if errMA != nil {
			return issuAddr, nil, errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		issuAddress, tx, issuInstance, err = role.DeployIssuance(auth, client, rolefsAddr)
		if issuAddress.String() != InvalidAddr {
			issuAddr = issuAddress
			issuIns = issuInstance
		}
		if err != nil {
			retryCount++
			log.Println("deploy Issuance Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(DefaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return issuAddr, issuIns, err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err == ErrTxFail {
			checkRetryCount++
			log.Println("deploy Issuance transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return issuAddr, issuIns, err
			}
			continue
		}
		if err != nil {
			return issuAddr, issuIns, err
		}
		break
	}
	log.Println("Issuance has been successfully deployed! The address is ", issuAddr.Hex())
	return issuAddr, issuIns, nil
}

func newIssuance(issuAddr common.Address, client *ethclient.Client) (*role.Issuance, error) {
	issuIns, err := role.NewIssuance(issuAddr, client)
	if err != nil {
		return nil, err
	}
	return issuIns, nil
}

// MintLevel get mintLevel in Issuance contract
func (issu *ContractModule) MintLevel() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(EndPoint)
	defer client.Close()
	issuIns, err := newIssuance(issu.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.MintLevel(&bind.CallOpts{
			From: issu.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// LastMint get lastMint in Issuance contract
func (issu *ContractModule) LastMint() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(EndPoint)
	defer client.Close()
	issuIns, err := newIssuance(issu.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.LastMint(&bind.CallOpts{
			From: issu.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// Price get price in Issuance contract
func (issu *ContractModule) Price() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(EndPoint)
	defer client.Close()
	issuIns, err := newIssuance(issu.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.Price(&bind.CallOpts{
			From: issu.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// Size get size in Issuance contract
func (issu *ContractModule) Size() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(EndPoint)
	defer client.Close()
	issuIns, err := newIssuance(issu.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.Size(&bind.CallOpts{
			From: issu.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// SpaceTime get spaceTime in Issuance contract
func (issu *ContractModule) SpaceTime() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(EndPoint)
	defer client.Close()
	issuIns, err := newIssuance(issu.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.SpaceTime(&bind.CallOpts{
			From: issu.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// TotalPay get totalPay in Issuance contract
func (issu *ContractModule) TotalPay() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(EndPoint)
	defer client.Close()
	issuIns, err := newIssuance(issu.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.TotalPay(&bind.CallOpts{
			From: issu.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}

// TotalPaid get totalPaid in Issuance contract
func (issu *ContractModule) TotalPaid() (*big.Int, error) {
	m := big.NewInt(0)

	client := getClient(EndPoint)
	defer client.Close()
	issuIns, err := newIssuance(issu.contractAddress, client)
	if err != nil {
		return m, err
	}

	retryCount := 0
	for {
		retryCount++
		m, err = issuIns.TotalPaid(&bind.CallOpts{
			From: issu.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return m, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return m, nil
	}
}
