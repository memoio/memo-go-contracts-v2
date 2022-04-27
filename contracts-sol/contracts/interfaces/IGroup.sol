// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct GroupOut {
    bool isActive;  // true when keeper count >= level 
    bool isBan;     // 组是否已被禁用
    uint16 level;   // security level
    uint256 size;   // storeSize
    uint256 price;  // storePrice
    address pool;
}

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IGroupSetter {
    // called by admin
    function ban(uint64 _gi, bool _isBan) external;

    function activate(uint64 _gi, uint16 _kc) external;
    function create(uint16 _level, uint8 _mr, uint256 _kr, uint256 _pr) external;
}

interface IGroupGetter {
    function add(uint8 _rType, uint64 _gi, uint256 _pm) external view;
    function checkG(uint64 i) external view;
    
    function getGCnt() external view returns (uint64);
    function getLevel(uint64 i) external view returns (uint16);
    function getGInfo(uint64 _i) external view returns (GroupOut memory);
    function getPInfo(uint64 _i) external view returns (uint256, uint256);

    function getKManage(uint64 _i) external view returns (address);
    function getPool(uint64 _i) external view returns (address);
}

interface IGroup is IGroupSetter,IGroupGetter {}