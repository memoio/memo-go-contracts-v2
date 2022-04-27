//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IRole.sol";
import "../interfaces/IAuth.sol";
import "./Owner.sol";
import "./Pool.sol";
import "./Kmanage.sol";

/**
 *@author MemoLabs
 *@title Manage account, roles and groups in the memo system.
 */
contract Role is IRole, Owner {
    struct RoleInfo {
        bool isBan;     // 是否被admin禁止
        bool isActive;  // true when add to group; keeper need set by admin 
        uint8 rType;    // 0 account; 1 user; 2 provider; 3 keeper
        uint64 index;   // 账户序列号，从1开始
        uint64 gIndex;  // 所属group，从1开始
        address payee;  // acoount for receiving income
        bytes extra;    // 0 empty; 1 verify key; 2 ; 3 bls public key
    }

    uint16 public version = 2;

    address[] addrs; // all roles 序列号即为index,从0开始, addrs[0]为foundation地址
    mapping(address => RoleInfo) info;

    event ReAcc(address addr, uint64 index); // to get all registered account by filter logs
    event ReRole(uint8 indexed _rType, uint64 index); // to get all users/keepers/providers by filter logs

    constructor(address _ctl, address _a, address f) Owner(_ctl, _a) {
        addrs.push(f);
        info[f].payee = f;
        emit ReAcc(f, 0);
    }

    function ban(uint64 _index, bool _banned) external onlyOwner override {
        address a = addrs[_index];
        info[a].isBan = _banned;
    }

    function setG(uint64 _index, uint64 _gi) external onlyOwner override {
        address a = addrs[_index];
        info[a].gIndex = _gi;
        if (info[a].rType != 3) {
            info[a].isActive = true;
        }
    }

    function activate(uint64 _index) external onlyOwner override returns (uint64) {
        address a = addrs[_index];
        require(!info[a].isActive && !info[a].isBan && info[a].rType == 3 && info[a].gIndex > 0, "AE"); // is banned
        
        info[a].isActive = true;

        return info[a].gIndex;
    }

    // role is already in some group
    function checkIG(uint64 _index, uint8 _rType) external view override returns (address, address, uint64, uint8) {
        address a = addrs[_index];
        require(!info[a].isBan, "AE"); // account error
        //if (info[a].gIndex > 0) {
        //    require(info[a].isActive, "NA"); // not active
        //}
        if (_rType > 0) {
            require(info[a].isActive && info[a].rType==_rType, "TE"); // Type error
        }
        
        return (a,info[a].payee, info[a].gIndex, info[a].rType);
    }

    function reAcc(address a) external onlyOwner override {
        if(info[a].index == 0 && info[a].gIndex == 0) {
            uint64 len = uint64(addrs.length);
            info[a].index = len;
            info[a].payee = a;
            addrs.push(a);
            emit ReAcc(a, len);
        }
    }

    function reRole(uint64 _index, uint8 _rType, bytes memory _extra) external onlyOwner override {
        address a = addrs[_index];
        require(info[a].rType==_rType && info[a].gIndex == 0 && !info[a].isBan, "AE");

        info[a].rType = _rType;
        info[a].extra = _extra;
        emit ReRole(_rType, _index);
    }

    // ===================get===================
    function getACnt() external view returns (uint64) {
        return uint64(addrs.length);
    }

    function getAddr(uint64 _i) external view override returns (address) {
        return addrs[_i];
    }

    function getIndex(address _a) external view override returns (uint64) {
        return info[_a].index;
    }

    function getRInfo(address acc) external view override returns (RoleOut memory) {
        RoleOut memory r;
        r.isBan = info[acc].isBan;
        return r;
    }
}