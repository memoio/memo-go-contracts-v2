package iface

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// ERC20Info contains information about memo primary token
type ERC20Info interface {
	DeployERC20(string, string)
	GetName() string
	GetSymbol() string
	GetDecimals() uint8
	GetTotalSupply() *big.Int
	BalanceOf(common.Address) (*big.Int)
	Allowance(common.Address, common.Address) (*big.Int)

	Transfer(common.Address, *big.Int) (bool)
	Approve(common.Address, *big.Int) (bool)
	TransferFrom(common.Address, common.Address, *big.Int) (bool)
	IncreaseAllowance(common.Address, *big.Int) (bool)
	DecreaseAllowance(common.Address, *big.Int) (bool)
	MintToken(common.Address, *big.Int) (bool)
	Burn(*big.Int) (bool)
	AirDrop([]common.Address, *big.Int) (bool)
}

// AccessControlInfo contains information about token operation permissions
type AccessControlInfo interface {
	SetUpRole([32]byte, common.Address) (bool)
	HasRole([32]byte, common.Address) (bool)
	RevokeRole([32]byte, common.Address) (bool)
	RenounceRole([32]byte) (bool)
	Pause() (bool)
	Unpause() (bool)
	GetPaused() (bool)
}
