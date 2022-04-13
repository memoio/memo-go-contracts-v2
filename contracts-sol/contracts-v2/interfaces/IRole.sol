// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct RoleOut {
    bool isActive;  // user注册后即生效，k-p注册后还需要质押才能生效
    bool isBanned;  // 是否被admin禁止
    uint8 roleType; // 0 account; 1 user; 2 provider; 3 keeper
    uint64 index;   // 账户序列号，从1开始
    uint64 gIndex;  // 所属group，从1开始
    address owner;  // acoount for receiving income
    bytes extra;    // 用于存储额外的信息，0 empty; 1 verify key; 2 proxy address?; 3 bls public key
}

struct GroupOut {
    bool isActive; // true when keeper count >= level 
    bool isBanned; // 组是否已被禁用
    bool isReady;  // 是否已在线下成组; 由签名触发
    uint16 level;  // security level
    uint64 kCnt;   // keeper count
    uint64 pCnt;   // keeper count
    uint64 uCnt;   // keeper count
    uint256 size;   // storeSize
    uint256 price;  // storePrice
    address fsAddr; // fs contract address
}

struct PledgeOut {
    uint256 kpledge;
    uint256 ppledge;
    address at;
}

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IRoleSetter {
    // called by admin
    function activate(uint64 _index, bool _active) external;
    function ban(uint64 _index, bool _banned) external;

    // called by admin, create a group, owner新建一个 FileSys 合约
    function createGroup(uint16 _level, address fsAddr, uint64[] memory indexes) external;

    // register self to get index
    function registerAccount(address a) external; 
    // combine as registerRole?
    function registerRole(uint64 index, uint8 rtype, bytes memory extra) external;
    // 向组中新加一个user/keeper/provider
    function addToGroup(uint64 _index, uint64 _gIndex) external;
}

interface IRoleGetter {
    // index is rType and not in some group
    function checkIR(uint64 _index, uint8 _rType) external view returns (address);
    // index is rType and in some group
    function checkIG(uint64 _index, uint8 _rType) external view returns (address, address, uint64);
    
    
    function getAddr(uint64 i) external view returns (address);
    function getRoleInfo(address acc) external view returns (RoleOut memory);
    function getGroupInfo(uint64 i) external view returns (GroupOut memory);

    function getFsAddr(uint64 i) external view returns (address);
    function getGroupK(uint64 ig, uint64 ik) external view returns (uint64);
    function getGroupP(uint64 ig, uint64 ip) external view returns (uint64);
}