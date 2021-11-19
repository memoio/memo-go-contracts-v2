package callconts

import (
	"log"
	rolefs "memoContract/contracts/rolefs"
	iface "memoContract/interfaces"

	"github.com/ethereum/go-ethereum/common"
)

//RoleFS  The basic information of node used for RoleFS contract
type RoleFS struct {
	addr   common.Address //local address
	hexSk  string         //local privateKey
	txopts *TxOpts
}

//NewRFS new a instance of contractAdminOwned
func NewRFS(addr common.Address, hexSk string) iface.RoleFSInfo {
	rfs := &RoleFS{
		addr:  addr,
		hexSk: hexSk,
	}

	return rfs
}

//DeployRoleFS deploy an RoleFS contract
func (rfs *RoleFS) DeployRoleFS() (common.Address, error) {
	var roleFSAddr common.Address
	client := getClient(EndPoint)
	auth, err := makeAuth(rfs.hexSk, nil, rfs.txopts)
	if err != nil {
		return roleFSAddr, err
	}

	roleFSAddr, _, _, err = rolefs.DeployRoleFS(auth, client)
	if err != nil {
		log.Println("deployRoleFSErr:", err)
		return roleFSAddr, err
	}
	log.Println("RoleFS Contract Addr:", roleFSAddr.String())
	return roleFSAddr, nil
}