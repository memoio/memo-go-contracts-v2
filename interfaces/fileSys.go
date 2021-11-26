package iface

import (
	"math/big"
	filesys "memoContract/contracts/filesystem"

	"github.com/ethereum/go-ethereum/common"
)

// FileSysInfo contains operations related to payment about storage order
type FileSysInfo interface {
	DeployFileSys(founder uint64, gIndex uint64, r common.Address, rfs common.Address, keepers []uint64) (common.Address, *filesys.FileSys, error)

	GetFsInfo(fsAddr common.Address, uIndex uint64) (isActive bool, tokenIndex uint32, err error)
	GetFsInfoAggOrder(fsAddr common.Address, uIndex uint64, provider uint64) (uint64, uint64, error)
	GetStoreInfo(fsAddr common.Address, uIndex uint64, pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, error)
	GetChannelInfo(fsAddr common.Address, uIndex uint64, pIndex uint64, tIndex uint32) (*big.Int, uint64, uint64, error)
	GetSettleInfo(fsAddr common.Address, pIndex uint64, tIndex uint32) (uint64, uint64, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
	GetBalance(fsAddr common.Address, rIndex uint64, tIndex uint32) (*big.Int, *big.Int, error)
}
