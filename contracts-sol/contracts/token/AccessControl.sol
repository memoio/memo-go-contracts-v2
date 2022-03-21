// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/erc20In/IAccessControl.sol";

// 本合约中的role指代的是下面三个constant值，表示访问控制的3种角色等级
contract AccessControl is IAccessControl {
    mapping(uint8 => mapping(address => bool)) accessRoles;
    bool private paused; // 为true表示禁止转账操作

    uint8 public constant DEFAULT_ADMIN_ROLE = 0;
    uint8 public constant MINTER_ROLE = 1;
    uint8 public constant PAUSER_ROLE = 2;

    constructor() {
        accessRoles[DEFAULT_ADMIN_ROLE][msg.sender] = true;
        accessRoles[MINTER_ROLE][msg.sender] = true;
        accessRoles[PAUSER_ROLE][msg.sender] = true;
    }

    function setUpRole(uint8 role, address account) external override returns (bool) {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "NAR"); // caller does not has accessRole
        accessRoles[role][account] = true;
        return true;
    }

    function hasRole(uint8 role, address account) public view override returns (bool) {
        return accessRoles[role][account];
    }

    // admin revoke other account's role
    function revokeRole(uint8 role, address account) external override returns (bool) {
        require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "NAR"); // caller does not has accessRole
        accessRoles[role][account] = false;
        return true;
    }

    // account renounce its role 
    function renounceRole(uint8 role) external override returns (bool) {
        accessRoles[role][msg.sender] = false;
        return true;
    }

    function pause() external override returns (bool) {
        require(hasRole(PAUSER_ROLE, msg.sender), "CNP"); // cann't pause
        paused = true;
        return true;
    }

    function unpause() external override returns (bool) {
        require(hasRole(PAUSER_ROLE, msg.sender), " CNUP"); // cann't unpause
        paused = false;
        return true;
    }

    function getPaused() public view override returns (bool) { //可见性如果是external，则继承该合约的ERC20合约将不能直接调用该函数
        return paused;
    }
}