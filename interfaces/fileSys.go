package iface

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// FileSysInfo contains operations related to payment about storage order
type FileSysInfo interface {
	DeployFileSys(founder uint64, gIndex uint64, r common.Address, rfs common.Address, keepers []uint64)
	GetFsInfo(uIndex uint64) (isActive bool, tokenIndex uint32)
	GetFsInfoAggOrder(uIndex uint64, provider uint64) (uint64, uint64)
	GetStoreInfo(uIndex uint64, pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int)
	GetChannelInfo(uIndex uint64, pIndex uint64, tIndex uint32) (*big.Int, uint64, uint64)
	GetSettleInfo(pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int)
	GetBalance(rIndex uint64, tIndex uint32) (*big.Int, *big.Int)
}
