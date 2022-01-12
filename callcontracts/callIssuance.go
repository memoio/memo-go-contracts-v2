package callconts

import (
	"log"
	"math/big"
	"memoContract/contracts/role"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewIssu new a instance of ContractModule. issuAddr: Issuance contract address
func NewIssu(issuAddr, addr common.Address, hexSk string, txopts *TxOpts, endPoint string, status chan error) iface.IssuanceInfo {
	issu := &ContractModule{
		addr:            addr,
		hexSk:           hexSk,
		txopts:          txopts,
		contractAddress: issuAddr,
		endPoint:        endPoint,
		Status:          status, // 用于接收：后台goroutine检查交易是否执行成功， nil代表成功
	}

	return issu
}

// DeployIssuance deploy an Issuance contract, called by admin
func (issu *ContractModule) DeployIssuance(rolefsAddr common.Address) (common.Address, *role.Issuance, error) {
	var issuAddr common.Address
	var issuIns *role.Issuance

	log.Println("begin deploy Issuance contract...")
	client := getClient(issu.endPoint)
	defer client.Close()

	// txopts.gasPrice参数赋值为nil
	auth, errMA := makeAuth(issu.hexSk, nil, issu.txopts)
	if errMA != nil {
		return issuAddr, nil, errMA
	}
	// 构建交易，通过 sendTransaction 将交易发送至 pending pool
	issuAddr, tx, issuIns, err := role.DeployIssuance(auth, client, rolefsAddr)
	// ====面临的失败场景====
	// 交易参数通过abi打包失败;payable检测失败;构造types.Transaction结构体时遇到的失败问题（opt默认值字段通过预言机获取）；
	// 交易发送失败，直接返回错误
	if err != nil {
		log.Println("DeployIssuance Err:", err)
		return issuAddr, nil, err
	}
	log.Println("transaction hash:", tx.Hash().Hex())
	log.Println("send transaction successfully!")
	// 交易成功发送至 pending pool , 后台检查交易是否成功执行,执行失败则将错误传入 ContractModule 中的 status 通道
	// 交易若由于链上拥堵而短时间无法被打包，不再增加gasPrice重新发送
	go checkTx(tx, issu.Status, "DeployIssuance")

	log.Println("Issuance address is ", issuAddr.Hex())
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

	client := getClient(issu.endPoint)
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

	client := getClient(issu.endPoint)
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

	client := getClient(issu.endPoint)
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

	client := getClient(issu.endPoint)
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

	client := getClient(issu.endPoint)
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

	client := getClient(issu.endPoint)
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

	client := getClient(issu.endPoint)
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
