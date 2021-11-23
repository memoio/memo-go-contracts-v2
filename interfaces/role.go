package iface

import (
	"math/big"
	"memoContract/contracts/role"
	"memoContract/contracts/rolefs"

	"github.com/ethereum/go-ethereum/common"
)

// OwnerInfo contains information about the owner of the Role contract
type OwnerInfo interface {
	// called by owner
	AlterOwner(common.Address, common.Address) error

	// called by anyone
	GetOwner(common.Address) (common.Address, error)
}

// RoleInfo contains operations related to memo roles
type RoleInfo interface {
	// deploy Role contract. specify foundation、primaryToken、pledgeK、pledgeP
	DeployRole(common.Address, common.Address, *big.Int, *big.Int) (common.Address, *role.Role, error)

	// callled by owner, set pledgePool-address、issuance-address and rolefs-address
	SetPI(common.Address, common.Address, common.Address, common.Address) error

	// called by anyone to get index
	Register(common.Address, common.Address, []byte) (uint64, error)

	// called by anyone to register Keeper
	RegisterKeeper(common.Address, uint64, []byte, []byte) error

	// called by anyone to register Provider
	RegisterProvider(common.Address, uint64, []byte) error

	// called by anyone to register User
	RegisterUser(common.Address, uint64, uint64, uint32, []byte, []byte) error

	// called by owner to add token
	RegisterToken(common.Address, common.Address) error

	// called by owner
	CreateGroup(common.Address, []uint64, uint16) (uint64, error)

	// called by owner to set fsAddress for group after CreateGroup and deployFileSys
	SetGF(common.Address, uint64, common.Address) error

	// called by owner
	AddKeeperToGroup(common.Address, uint64, uint64) error

	// called by provider or called by others
	AddProviderToGroup(common.Address, uint64, uint64, []byte) error

	// called by owner to set the amount that the keeper and provider needs to pledge
	SetPledgeMoney(common.Address, *big.Int, *big.Int) error

	// called by user or called by others
	Recharge(common.Address, uint64, uint32, *big.Int, []byte) error

	// called by memo-role or called by others
	WithdrawFromFs(common.Address, uint64, uint32, *big.Int, []byte) error

	// get the number of registered addresses
	GetAddrsNum(common.Address) (uint64, error)

	// get address by array index value(not role index, the array index value is the role index minus 1)
	GetAddr(common.Address, uint64) (common.Address, error)

	// get the account role index by address
	GetRoleIndex(common.Address, common.Address) (uint64, error)

	// get account information by address
	GetRoleInfo(common.Address, common.Address) (bool, bool, uint8, uint64, uint64, []byte, error)

	// get the number of groups
	GetGroupsNum(common.Address, ) (uint64, error)

	// get group information by gIndex
	GetGroupInfo(common.Address, uint64) (bool, bool, bool, uint16, *big.Int, *big.Int, common.Address, error)

	// get FileSys-contract address by group array index(not gIndex, group array index is gIndex minus 1)
	GetFsAddr(common.Address, uint64) (common.Address, error)

	// get address and gIndex by array index value
	GetAddrGindex(common.Address, uint64) (common.Address, uint64, error)

	// get the number of keepers in the group
	GetGKNum(common.Address, uint64) (uint64, error)

	// get the number of providers in the group
	GetGPNum(common.Address, uint64) (uint64, error)

	// get keeper role index by group array index and keeper array index
	GetGroupK(common.Address, uint64, uint64) (uint64, error)

	// get provider role index by group array index and provider array index
	GetGroupP(common.Address, uint64, uint64) (uint64, error)
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
	IsValid(common.Address, uint32) (bool, error)
	GetTA(common.Address, uint32) (common.Address, error)
	GetTI(common.Address, common.Address) (uint32, bool, error)
	GetTNum(common.Address) (uint32, error)
}

// IssuanceInfo contains deploy Isuance-contract function
type IssuanceInfo interface {
	DeployIssuance(rfs common.Address) (common.Address, *role.Issuance, error)
}
