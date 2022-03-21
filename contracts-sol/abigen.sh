abigen --sol contracts/role/RoleFS.sol --out ../memo-go-contracts-v2/contracts/rolefs/RoleFS.go --pkg rolefs --type RoleFS --bin ~/Documents/ContractABI/RoleFS.bin
abigen --sol contracts/fileSystem/FileSys.sol --out ../memo-go-contracts-v2/contracts/fileSystem/FileSys.go --pkg filesys --type FileSys --bin ~/Documents/ContractABI/FileSys.bin
abigen --sol contracts/role/Issuance.sol --out ../memo-go-contracts-v2/contracts/issuance/Issuance.go --pkg issuance --type Issuance --bin ~/Documents/ContractABI/Issuance.bin
abigen --sol contracts/role/Role.sol --out ../memo-go-contracts-v2/contracts/role/Role.go --pkg role --type Role --bin ~/Documents/ContractABI/Role.bin
