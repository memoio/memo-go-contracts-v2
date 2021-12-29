# memo-go-contracts-v2

## contracts
合约文件通过`abigen`转换为go文件，保存在该目录中。

## interface
包含了所有合约的访问接口。

## callcontracts
定义了一个`ContractModule`的结构体，在该目录下，实现了ContractModule对上述所有合约接口的调用。

## cmd
该目录下提供了查询余额、转账（erc20 token以及eth token）的命令行操作。以及admin部署合约、调用合约的命令行操作。

## 注意事项
在callcontracts/common.go中定义了ERC20Addr、RoleAddr、RoleFSAddr、RTokenAddr、FileSysAddr、PledgePoolAddr、IssuanceAddr以及admin的账户地址和私钥信息，以供全局使用。

这些合约地址由admin部署合约从而得到，admin部署合约步骤：
1. DeployERC20 => erc20Addr
2. DeployRole => roleAddr、rTokenAddr
3. DeployRoleFS => rolefsAddr
4. DeployPledgePool => pledgePoolAddr
5. DeployIssuance => issuAddr
6. role.SetPI
7. role.CreateGroup => fsAddr
8. rolefs.SetAddr

admin选取一部分keeper(也可以为空，因为后续可以调用addKeeperToGroup增添keeper到group中)创建group(createGroup)，每个group对应一个filesys。所以在admin调用createGroup时，会同时部署一个FileSys合约。

## index
系统中存在三种index：角色索引（rIndex）、group索引（gIndex）、代币索引（tIndex）
rIndex从1开始，不能为0；
gIndex也从1开始，不能为0；
tIndex从0开始，tIndex=0表示主代币；

## nonce
涉及到的nonce值都需要从0开始依次累加，从而与合约中的nonce值匹配；