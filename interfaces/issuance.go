package iface

import (
	"math/big"
	"memoContract/contracts/issuance"

	"github.com/ethereum/go-ethereum/common"
)

// IssuanceInfo contains deploy Isuance-contract function
type IssuanceInfo interface {
	DeployIssuance(rfs common.Address) (common.Address, *issuance.Issuance, error)
	MintLevel() (*big.Int, error)
	LastMint() (*big.Int, error)
	Price() (*big.Int, error)
	Size() (*big.Int, error)
	SpaceTime() (*big.Int, error)
	TotalPay() (*big.Int, error)
	TotalPaid() (*big.Int, error)
}
