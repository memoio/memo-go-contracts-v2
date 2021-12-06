package callconts

import (
	"memoContract/contracts/role"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NewRT new a instance of ContractModule
func NewRT(addr common.Address, hexSk string, txopts *TxOpts) iface.RTokenInfo {
	rt := &ContractModule{
		addr:   addr,
		hexSk:  hexSk,
		txopts: txopts,
	}

	return rt
}

// IsValid check whether the tokenIndex is valid, rtokenAddr indicates RToken contract address, get it by RToken() in callRole.go
func (rt *ContractModule) IsValid(rTokenAddr common.Address, tIndex uint32) (bool, error) {
	var isvalid bool
	rToken, err := role.NewRToken(rTokenAddr, getClient(EndPoint))
	if err != nil {
		return false, err
	}

	retryCount := 0
	for {
		retryCount++
		isvalid, err = rToken.IsValid(&bind.CallOpts{
			From: rt.addr,
		}, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return false, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return isvalid, nil
	}
}

// GetTA get the address of tokenIndex
func (rt *ContractModule) GetTA(rtokenAddr common.Address, tIndex uint32) (common.Address, error) {
	var taddr common.Address
	rToken, err := role.NewRToken(rtokenAddr, getClient(EndPoint))
	if err != nil {
		return taddr, err
	}

	retryCount := 0
	for {
		retryCount++
		taddr, err = rToken.GetTA(&bind.CallOpts{
			From: rt.addr,
		}, tIndex)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return taddr, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return taddr, nil
	}
}

// GetTI get the tokenIndex by token address and if it is valid
func (rt *ContractModule) GetTI(rtokenAddr common.Address, taddr common.Address) (uint32, bool, error) {
	var tindex uint32
	var isvalid bool
	rToken, err := role.NewRToken(rtokenAddr, getClient(EndPoint))
	if err != nil {
		return tindex, isvalid, err
	}

	retryCount := 0
	for {
		retryCount++
		tindex, isvalid, err = rToken.GetTI(&bind.CallOpts{
			From: rt.addr,
		}, taddr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return tindex, isvalid, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return tindex, isvalid, nil
	}
}

// GetTNum get the number of tokens that memo supports
func (rt *ContractModule) GetTNum(rtokenAddr common.Address) (uint32, error) {
	var tnum uint32
	rToken, err := role.NewRToken(rtokenAddr, getClient(EndPoint))
	if err != nil {
		return tnum, err
	}

	retryCount := 0
	for {
		retryCount++
		tnum, err = rToken.GetTNum(&bind.CallOpts{
			From: rt.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return tnum, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return tnum, nil
	}
}
