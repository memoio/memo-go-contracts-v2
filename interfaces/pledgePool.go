package iface

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// PledgePoolInfo contains operations related to pledge on chain
type PledgePoolInfo interface {
	// primeToken、rToken、role
	DeployPledgePool(common.Address, common.Address, common.Address)
	Pledge(uint64, *big.Int, []byte)
	Withdraw(uint64, uint32, *big.Int, []byte)
	GetPledge(uint32) *big.Int
	GetBalance(uint64, uint32) *big.Int
}
