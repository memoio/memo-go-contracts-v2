//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IGroup.sol";
import "../interfaces/IAuth.sol";
import "./Owner.sol";
import "./Pool.sol";
import "./Kmanage.sol";

/**
 *@author MemoLabs
 *@title Manage account, roles and groups in the memo system.
 */
contract Group is IGroup, Owner {
    struct GroupInfo {
        bool isActive;      // true when keeper count >= level 
        bool isBan;      // 组是否已被禁用
        uint16 level;       // security level
        uint256 kpr;        // keeper pledge require 
        uint256 ppr;        // provider pledge require
        address pool;       // poolAddr
        address kManage;    // keeper manage contract address
    }

    uint16 public version = 2;

    GroupInfo[] groups; // manage group

    event CreateGroup(uint64 gIndex);

    constructor(address _ctl, address _a) Owner(_ctl, _a) {
        GroupInfo memory g;
        g.isBan = true;
        groups.push(g);
    }

    function ban(uint64 _gi, bool _isBan) external onlyOwner override {
        groups[_gi].isBan = _isBan;
    }

    function create(uint16 _level, uint8 mr, uint256 _kr, uint256 _pr) external onlyOwner override {
        uint64 _gIndex = uint64(groups.length);

        // create pool address; force each group has unique pool  
        Pool p = new Pool(instances[1],instances[2]);
        Kmanage k = new Kmanage(instances[1], instances[2], mr);
        
        GroupInfo memory g;
        g.level = _level;
        g.kpr = _kr;
        g.ppr = _pr; 
        g.pool = address(p);
        g.kManage = address(k);
        g.isBan = true;  // banned at first
        groups.push(g);

        emit CreateGroup(_gIndex);
    }

    function activate(uint64 _gi, uint16 _kc) external onlyOwner override {
        require(!groups[_gi].isBan, "GB");
        if (_kc >= groups[_gi].level) {
            groups[_gi].isActive = true;
        }
    }

    // check

    function add(uint8 _rType, uint64 _gi, uint256 _pm) external override view {
        require(!groups[_gi].isBan, "GB"); // group is banned
    

        if (_rType == 1) {
            require(groups[_gi].isActive, "GNE");
            return;
        }

        if (_rType == 2) {
            require(groups[_gi].isActive, "GNE");
            require(_pm >= groups[_gi].ppr, "PPI"); // pledge insuf
            return;
        }

        if (_rType == 3) {
            require(_pm >= groups[_gi].kpr, "KPI"); // pledge insuf
        }
    }

    function checkG(uint64 i) external view override {
        require(groups[i].isActive && !groups[i].isBan);
    }

    // ===================get===================

    function getGCnt() external view override returns (uint64) {
        return uint64(groups.length);
    }

    function getLevel(uint64 i) external view override returns (uint16) {
        return groups[i].level;
    }

     function getPInfo(uint64 i) external view override returns (uint256, uint256) {
        return (groups[i].kpr, groups[i].ppr);
    }

    function getKManage(uint64 _i) external view override returns (address) {
        return groups[_i].kManage;
    }

    function getPool(uint64 i) external view override returns (address) {
        return groups[i].pool;
    }

    function getGInfo(uint64 i) external view override returns (GroupOut memory) {
        GroupOut memory g;
        g.isActive = groups[i].isActive;
        return g;
    }
}