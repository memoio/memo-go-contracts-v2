//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IRole.sol";
import "../interfaces/IAuth.sol";
import "../interfaces/IFileSys.sol";
import "./Owner.sol";
import "../Recover.sol";

/**
 *@author MemoLabs
 *@title Used to register account roles for people participating in the memo system
 */
contract Role is IRoleSetter, IRoleGetter, Owner {
    using Recover for bytes32;

    struct RoleInfo {
        bool isActive;  // user注册后即生效，k-p注册后还需要质押才能生效
        bool isBanned;  // 是否被admin禁止
        uint8 roleType; // 0 account; 1 user; 2 provider; 3 keeper
        uint64 index;   // 账户序列号，从1开始
        uint64 gIndex;  // 所属group，从1开始
        address payee;  // acoount for receiving income
        bytes extra;    // 0 empty; 1 verify key; 2 ; 3 bls public key
    }

    struct GroupInfo {
        bool isActive; // true when keeper count >= level 
        bool isBanned; // 组是否已被禁用
        bool isReady;  // 是否已在线下成组; 由签名触发
        uint16 level;  // security level
        uint64[] keepers;   // 里面有哪些keeper
        uint64[] providers; // 有哪些provider
        uint64[] users;     // 有哪些provider
        uint256 size;   // storeSize
        uint256 price;  // storePrice
        address fsAddr; // fs contract address
    }

    uint16 public version = 2;

    address public foundation; //基金会账户地址，索引默认为0,不需要进行register

    address[] addrs; // all roles 序列号即为index,从1开始
    mapping(address => RoleInfo) info;

    // manage group
    GroupInfo[] groups;

    event ReAcc(address addr, uint64 index);
    event ReRole(uint64 index, uint8 _rType);
    event CreateGroup(uint64 gIndex);

    constructor(address f, address _rfs, address _a) Owner(_rfs, _a) {
        foundation = f;
    }

    // used for keeper
    function activate(uint64 _index, bool _active) external onlyOwner override {
        address a = addrs[_index-1];
        require(!info[a].isBanned, "IB"); // is banned
        if (info[a].roleType == 3 && _active && !info[a].isActive && info[a].gIndex > 0) {
            uint64 gIndex = info[a].gIndex; 
            groups[gIndex-1].keepers.push(_index);
            if (groups[gIndex-1].keepers.length >= groups[gIndex-1].level) {
                groups[gIndex-1].isActive = true;
                IFileSysSetter(groups[gIndex-1].fsAddr).addKeeper(_index);
            }
        }
        info[a].isActive = _active;
    }

    function ban(uint64 _index, bool _banned) external onlyOwner override {
        address a = addrs[_index-1];
        info[a].isBanned = _banned;
    }

    /// @dev check if 'index' is rType and not in some group
    function checkIR(uint64 _index, uint8 _rType) external view override returns (address) {
        address a = addrs[_index-1];
        require(info[a].roleType==_rType && !info[a].isActive && !info[a].isBanned, "AE"); // account error
        return a;
    }

    // role is already in some group
    function checkIG(uint64 _index, uint8 _rType) external view override returns (address, address, uint64) {
        address a = addrs[_index-1];
        require(info[a].roleType==_rType && info[a].isActive && !info[a].isBanned, "AE"); // account error
        return (a,info[a].payee, info[a].gIndex);
    }

    function registerAccount(address a) external onlyOwner override {
        if(info[a].index == 0) {
            addrs.push(a);
            uint64 len = uint64(addrs.length);
            info[a].index = len;
            info[a].payee = a;
            emit ReAcc(a, len);
        }
    }

    function registerRole(uint64 _index, uint8 _rType, bytes memory _extra) external onlyOwner override {
        address a = this.checkIR(_index, 0);
        info[a].roleType = _rType;
        info[a].extra = _extra;
        emit ReRole(_index, _rType);
    }

    function createGroup(uint16 _level, address fsAddr, uint64[] memory indexes) external onlyOwner override {
        GroupInfo memory g;
        groups.push(g);
        uint64 _gIndex = uint64(groups.length);

        IFileSysSetter ifs = IFileSysSetter(fsAddr);
        for(uint8 i = 0; i<indexes.length; i++) {
            address addr = this.checkIR(indexes[i], 3);
            info[addr].gIndex = _gIndex;
            info[addr].isActive = true;
            ifs.addKeeper(indexes[i]);
        }

        groups[_gIndex-1].level = _level;
        groups[_gIndex-1].fsAddr = fsAddr;
        groups[_gIndex-1].keepers = indexes;

        if(indexes.length >= uint(_level)) {
            groups[_gIndex-1].isActive = true;
        }
        emit CreateGroup(_gIndex);
    }

    function addToGroup(uint64 _index, uint64 _gIndex) external onlyOwner override {
        require(!groups[_gIndex-1].isBanned, "GB"); // group is banned
    
        address a = addrs[_index-1];
        require(!info[a].isActive && !info[a].isBanned, "AE"); // account error

        info[a].gIndex = _gIndex;
        if (info[a].roleType == 1) {
            info[a].isActive = true;
            groups[_gIndex-1].users.push(_index);
        }

        if (info[a].roleType == 2) {
            info[a].isActive = true;
            groups[_gIndex-1].providers.push(_index);
        }
    }

    function setGF(uint64 _gIndex, address _fsAddr) external onlyOwner {
        require(groups[_gIndex-1].fsAddr == address(0), "NE"); // not empty
        groups[_gIndex-1].fsAddr = _fsAddr;
    }

    // ===================get===================
    function getAddrCnt() external view returns (uint64) {
        return uint64(addrs.length);
    }

    // 根据数组索引值（不是账户index）获取相应的账户地址，超出数组范围将revert
    function getRoleIndex(address acc) external view returns (uint64) {
        return info[acc].index;
    }

    function getAddr(uint64 i) external view override returns (address) {
        return addrs[i-1];
    }

    function getRoleInfo(address acc) external view override returns (RoleOut memory) {
        RoleOut memory r;
        r.isActive = info[acc].isActive;
        return r;
    }

    function getGroupInfo(uint64 i) external view override returns (GroupOut memory) {
        GroupOut memory g;
        return g;
    }

    function getFsAddr(uint64 i) external view override returns (address) {
        return groups[i-1].fsAddr;
    }

    function getGroupK(uint64 ig, uint64 ik) external view override returns (uint64) {
        return 0;
    }

    function getGroupP(uint64 ig, uint64 ip) external view override returns (uint64) {
        return 0;
    }
}