// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface IControl {
    // called by admin 
    function activate(uint64 _index, bool _active, bytes[] memory signs) external;
    function ban(uint64 _index, bool _banned, bytes[] memory signs) external;
    function createGroup(uint16 _level,address fsAddr, uint64[] memory indexes, bytes[] memory signs) external;

    function addT(address _a, bytes[] memory signs) external;

    // register self to get index
    function registerAccount() external; 
    function registerRole(uint64 index, uint8 rtype, bytes memory extra) external;
    // add a user/keeper/provider to group
    function addToGroup(uint64 _index, uint64 _gIndex) external;
    

    function pledge(uint64 index, uint256 money) external;
    function withdraw(uint64 index, uint32 tIndex, uint256 money, uint256 lock) external returns (uint256);
}