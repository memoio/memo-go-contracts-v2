// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct RoleOut {
    bool isBan;  // 是否被admin禁止
    bool isActive;
    uint8 rType; // 0 account; 1 user; 2 provider; 3 keeper
    uint64 index;   // 账户序列号，从1开始
    uint64 gIndex;  // 所属group，从1开始
    address owner;  // acoount for receiving income
    bytes extra;    // 用于存储额外的信息，0 empty; 1 verify key; 2 proxy address?; 3 bls public key
}

struct PledgeOut {
    uint256 kpledge;
    uint256 ppledge;
    address at;
}

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IRoleSetter {
    // called by admin
    function ban(uint64 _i, bool _ban) external;

    function reAcc(address _a) external; 
    function reRole(uint64 _i, uint8 _rtype, bytes memory extra) external;
    function setG(uint64 _i, uint64 _gi) external;
    function activate(uint64 _i) external returns (uint64);
}

interface IRoleGetter {
    // index is rType and in some group
    function checkIG(uint64 _i, uint8 _rType) external view returns (address, address, uint64, uint8);
    
    function getIndex(address _a) external view returns (uint64);
    function getAddr(uint64 _i) external view returns (address);
    function getRInfo(address _a) external view returns (RoleOut memory);
}

interface IRole is IRoleSetter,IRoleGetter {}