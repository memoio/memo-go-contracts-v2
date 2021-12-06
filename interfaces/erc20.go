package iface

import (
	"math/big"

	"memoContract/contracts/erc20"

	"github.com/ethereum/go-ethereum/common"
)

// ERC20Info contains information about memo primary token
type ERC20Info interface {
	DeployERC20(string, string) (common.Address, *erc20.ERC20, error)
	GetName(common.Address) (string, error)
	GetSymbol(common.Address) (string, error)
	GetDecimals(common.Address) (uint8, error)
	GetTotalSupply(common.Address) (*big.Int, error)
	BalanceOf(common.Address, common.Address) (*big.Int, error)
	Allowance(common.Address, common.Address, common.Address) (*big.Int, error)

	Transfer(common.Address, common.Address, *big.Int) error
	Approve(common.Address, common.Address, *big.Int) error
	TransferFrom(common.Address, common.Address, common.Address, *big.Int) error
	IncreaseAllowance(common.Address, common.Address, *big.Int) error
	DecreaseAllowance(common.Address, common.Address, *big.Int) error
	MintToken(common.Address, common.Address, *big.Int) error
	Burn(common.Address, *big.Int) error
	AirDrop(common.Address, []common.Address, *big.Int) error

	// AccessControl
	SetUpRole(common.Address, uint8, common.Address) error
	RevokeRole(common.Address, uint8, common.Address) error
	RenounceRole(common.Address, uint8) error
	Pause(common.Address) error
	Unpause(common.Address) error

	GetPaused(common.Address) (bool, error)
	HasRole(common.Address, uint8, common.Address) (bool, error)
}
