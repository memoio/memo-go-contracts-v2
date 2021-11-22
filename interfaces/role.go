package iface

import (
	"math/big"
	"memoContract/contracts/rolefs"

	"github.com/ethereum/go-ethereum/common"
)

// OwnerInfo contains information about the owner of the Role contract
type OwnerInfo interface {
	// called by owner
	AlterOwner(common.Address)

	// called by anyone
	GetOwner() common.Address
}

// RoleInfo contains operations related to memo roles
type RoleInfo interface {
	// deploy Role contract. specify foundation、primaryToken、pledgeK、pledgeP
	DeployRole(common.Address, common.Address, *big.Int, *big.Int)

	// callled by owner, set pledgePool-address and issuance-address
	SetPI(common.Address, common.Address)

	// called by anyone to get index
	Register(common.Address, []byte) uint64

	// called by anyone, check if tIndex is valid
	CheckT(uint32) common.Address

	// called by anyone to register Keeper
	RegisterKeeper(uint64, []byte, []byte)

	// called by anyone to register Provider
	RegisterProvider(uint64, []byte)

	// called by anyone to register User
	RegisterUser(uint64, uint64, uint32, []byte, []byte)

	// called by owner to add token
	RegisterToken(common.Address)

	// called by owner
	CreateGroup([]uint64, uint16) uint64

	// called by owner to set fsAddress for group after CreateGroup and deployFileSys
	SetGF(uint64, common.Address)

	// called by owner
	AddKeeperToGroup(uint64, uint64)

	// called by provider or called by others
	AddProviderToGroup(uint64, uint64, []byte)

	// called by owner to set the amount that the keeper and provider needs to pledge
	SetPledgeMoney(*big.Int, *big.Int)

	// called by user or called by others
	Recharge(uint64, uint32, *big.Int, []byte)

	// called by memo-role or called by others
	WithdrawFromFs(uint64, uint32, *big.Int, []byte)

	// get the number of registered addresses
	GetAddrsNum() uint64

	// get address by array index value(not role index, the array index value is the role index minus 1)
	GetAddr(uint64) common.Address

	// get the account role index by address
	GetRoleIndex(common.Address) uint64

	// get account information by address
	GetRoleInfo(common.Address) (bool, bool, uint8, uint64, uint64, []byte)

	// get the number of groups
	GetGroupsNum() uint64

	// get group information by gIndex
	GetGroupInfo(uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address)

	// get FileSys-contract address by group array index(not gIndex, group array index is gIndex minus 1)
	GetFsAddr(uint64) common.Address

	// get address and gIndex by array index value
	GetAddrGindex(uint64) (common.Address, uint64)

	// get the number of keepers in the group
	GetGKNum(uint64) uint64

	// get the number of providers in the group
	GetGPNum(uint64) uint64

	// get keeper role index by group array index and keeper array index
	GetGroupK(uint64, uint64) uint64

	// get provider role index by group array index and provider array index
	GetGroupP(uint64, uint64) uint64
}

// RoleFSInfo contains operations related to memo roles and filesystem-payment
type RoleFSInfo interface {
	DeployRoleFS() (common.Address, *rolefs.RoleFS, error)
	// called by owner, which is the deployer
	SetAddr(common.Address, common.Address, common.Address, common.Address, common.Address) error
	AddOrder(common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, []byte, [][]byte) error
	SubOrder(common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, []byte, [][]byte) error
	AddRepair(common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, [][]byte) error
	SubRepair(common.Address, uint64, uint64, uint64, uint64, uint64, uint64, uint32, *big.Int, []byte, [][]byte) error
	ProWithdraw(common.Address, uint64, uint32, *big.Int, *big.Int, [][]byte) error
}

// RTokenInfo contains operations related to tokens that memo supported
type RTokenInfo interface {
	IsValid(common.Address,uint32) (bool, error)
	GetTA(common.Address,uint32) (common.Address, error)
	GetTI(common.Address,common.Address) (uint32, bool, error)
	GetTNum(common.Address) (uint32, error)
}

// IssuanceInfo contains deploy Isuance-contract function
type IssuanceInfo interface {
	DeployIssuance(rfs common.Address)
}
