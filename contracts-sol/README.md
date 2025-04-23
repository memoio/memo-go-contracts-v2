# memo-contracts-v2

`@author:Memolabs`

This repository contains the smart contract files in the memo project.

## Environment

**Language** solidity

**Editor** vscode

**Syntax and functional testing** solc, remix

**Integrated debugging tool** Hardhat

## Files

Contract files are in the contracts directory. The role directory contains role-related contracts; the fileSystem directory contains file system and settlement-related contracts; the pledgePool directory contains pledge pool-related contracts; the token directory contains token-related contracts; the interfaces directory contains contract interfaces; Owner.sol mainly records the owner information of the Role contract, including changing the owner; Recover.sol is a library file that recovers the signer based on the signature information.

Currently, the admin deploys the Role contract (the Owner contract will be automatically deployed at the same time, and the RToken contract will also be created by the Role contract), FileSys contract, PledgePool contract, ERC20 contract (the AccessControl contract will be automatically deployed at the same time), Issue contract, and RoleFS contract.

admin deploys contracts in the following order: ERC20, Role, RoleFS, PledgePool, Issue, (setPI), FileSys, (setAddr)

To optimize gas consumption, consider:

- Assign msg.sender, msg.value, tx.origin to local variables to avoid multiple accesses to msg and tx information

- Limit the use of arrays
- Limit the modification and access to storage

### Role

```sol
constructor(address f, address t, uint256 pk, uint256 ppro){}
```

You need to specify the foundation address, main token address, and pledgeK and pledgeP parameters. The main token address is the ERC20 contract address mentioned above.
After creation, you need to create a PledgePool contract based on the contract address, and you need to create an Issuance contract and a RoleFS contract, and then call the **setPI function** to assign the PledgePool contract address, the Issuance contract address, and the RoleFS contract address to the contract.

Create an RToken contract through the salt value when deploying the Role contract.
The salt value is set to:

```sol
bytes32 salt = keccak256(abi.encodePacked(foundation, primaryToken));
```

### FileSys

```sol
/// @dev created by admin; 'r' indicates role-contract address, 'rfs' indicates RoleFS-contract address
constructor(uint64 founder, uint64 _gIndex, address r, address rfs, uint64[] memory _keepers){}
```

You need to specify the foundation address, group index parameter, Role contract address, RoleFS contract address, and keepers index.
After admin calls the createGroup function in the Role contract, a FileSys contract is deployed synchronously (\_gIndex is the return value of the createGroup function, and \_keepers is the same as the first parameter when calling the createGroup function), and then the **setGF function** in the Role contract is called to assign the FileSys contract address to the Role contract.
Note: Every time CreateGroup is called, a FileSys contract needs to be deployed, and a RoleFS contract also needs to be deployed.

File system contract, which contains storage orders and payment information.

### PledgePool

```sol
/// @dev created by admin; 'r' indicates role-contract address
constructor(address primeToken, address _rToken, address r){}
```

It is necessary to specify the prime token address, RToken contract address (which can be obtained from the Role contract), and Role contract address.

Staking pool contract, any account with a serial number (call the register function in the Role contract to obtain the serial number) can stake and participate in profit sharing. The amount staked by the account when registering a role is also staked in the staking pool.

### ERC20

```sol
constructor(string memory _name, string memory _symbol){}
```

The token name and symbol need to be specified.

As the main token paid by the memo system, it is issued when addOrder is triggered. The addOrder function is in the RoleFS.sol file.

deploy Gas Used: 2212493 gas (according to the test results in Remix, the same below)

### Issuance

```sol
constructor(address rfs){}
```

The RoleFS contract address needs to be specified.

### RoleFS

```sol
constructor(){}
```

No parameters need to be specified. However, after deploying the RoleFS contract, the admin needs to call the **setAddr function** in the contract and specify the Issuance, Role, FileSys, and RToken contract addresses.

## Role registration process

In the Role contract, the account first calls the register function to obtain the serial number. The account that registers the User can directly call the registerUser function. To register the Keeper and Provider accounts, you need to first call the pledge function in the PledgePool contract to pledge, and then call the registerKeeper and registerProvider functions in the Role contract.

## Prerequisites

Before the account calls the addOrder function in the RoleFS contract, the admin needs to call the setRole function of the AccessControl contract to grant the MINTER_ROLE permission to the RoleFS contract.

## TODO list

### FileSys.sol

1. For ChannelInfo data structure:
Each time addOrder is executed, ChannelInfo.expire will be updated according to the end value, but the other two fields: amount and nonce are not used yet;
2. The period state variable can be considered to be modifiable, which is currently a constant;
3. In the addOrder operation, fs[_uIndex].ao[_pIndex].sInfo[_tIndex].time should also be updated accordingly. Currently, this value has not been updated; (processed)
4. 1% of the amount is paid to the keeper at the end, and in the subOrder operation, it is recorded in proInfo[\_pIndex][_tokenindex].endPaid; the other 3% of the linear payment is processed during proWithdraw. However, there are problems with the processing logic, which can be made more streamlined. (Handled)
5. In proWithdraw, balances[pindex][tokenindex] should be added to withdraw balance, but balances[pindex][tokenindex] is always 0, so this operation is redundant and needs to be removed. (Handled)
6. When addOrder is first called, se.time is 0 when calling \_settlementAdd function to update Provider's storage settlement related information, so proInfo[\_pIndex][_tokenindex].canPay = se.price \* (start - 0) is calculated at this time, which is wrong. It is necessary to determine whether se.time is 0. If it is 0, canPay will not be updated. (Handled)
7. fs[uIndex].tokenIndex is useless, because any token supported by Memo can be specified each time addOrder is called, so it should be removed. (Handled)
8. The contract state variable penalty (recording penalty information) is not handled.

### PledgePool.sol

1. In the pledge function, the hash value of the signature information is tentatively set to keccak256(abi.encodePacked(caller, money, "pledge")). If the user is pledged by another account (caller), this signature information can theoretically be used infinitely by the caller, allowing the user to pledge repeatedly. How to prevent the signature from being reused?
A nonce value can be maintained off-chain to record the nonce of each signature, and the nonce value is accumulated in sequence.
2. The withdraw function also faces the above problems;

### Role.sol

1. In the recharge function, how to avoid the reuse of signatures;

2. In the withdrawFromFs function, how to avoid the reuse of signatures;

3. The function of modifying the `isBanned` parameter in the RoleInfo and GroupInfo structures has not yet been added;

4. The function of Keeper and Provider retrieving the role application fund and canceling their own roles has not yet been implemented;

### RoleFS.sol

1. In the addOrder function, the signature verification of uSign, pSign, and kSigns has not yet been implemented;

2. In the subOrder function, the signature verification of uSign, pSign, and kSigns has not yet been implemented;

3. In the proWithdraw function, the signature verification of kSigns has not yet been implemented;

## Handover reception

In addition to the issuance of tokens with the addition of new orders, they also need to be issued over time. It is necessary to add an issuance factor and a minimum issuance ratio. If the cumulative reward reward+totalreward exceeds the target reward, and the issuance ratio does not reach the minimum issuance ratio, the excess needs to be divided by 2.
The issuance of Memo will be halved after the target issuance value is triggered, so early ecological participants can get more dividends.
The first target release of incentive issuance allocation is half of the Genesis issuance, and it will be halved after reaching the target until the set minimum issuance rate is reached and remains stable.
That is to say: the mintLevel should add the attribute of the issuance quantity to limit the upper limit of tokens that can be issued in each stage. After exceeding, the issuance rate will be halved (but not less than the minimum value) until the total order volume reaches the requirements of the next stage and jumps to the next stage. (Solved)

In addition, the proportion of the account pledge amount to obtain profit is determined by the proportion of the account pledge amount, and it should also be related to the duration. The longer the pledge time, the more profit you will get. This can be achieved through the timing of profit sharing. For example, if the profit sharing is triggered every 3 days, the earlier the account pledged, the more times it can get profit sharing.
Currently, profit sharing is triggered when an account is pledged or withdrawn. The longer the pledge is made, the more profit sharing value will be.
