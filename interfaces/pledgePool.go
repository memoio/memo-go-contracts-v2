package iface

import (
	"math/big"
	"memoContract/contracts/pledgepool"

	"github.com/ethereum/go-ethereum/common"
)

// PledgePoolInfo contains operations related to pledge on chain
type PledgePoolInfo interface {
	// primeToken、rToken、role
	DeployPledgePool(common.Address, common.Address, common.Address) (common.Address, *pledgepool.PledgePool, error)
	Pledge(common.Address, uint64, *big.Int, []byte) error
	Withdraw(common.Address, uint64, uint32, *big.Int, []byte) error
	GetPledge(common.Address, uint32) (*big.Int, error)
	GetBalanceInPPool(common.Address, uint64, uint32) (*big.Int, error)
}
