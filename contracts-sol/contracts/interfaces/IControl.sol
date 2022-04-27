// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IFileSysSetter.sol";

interface IControl {
    // called by admin 
    function activate(uint64 _i, bytes[] memory signs) external;          // 6, 11;11
    function ban(uint64 _i, bool _ban, bytes[] memory signs) external;    //6
    function addT(address _t, bool _ban, bytes[] memory signs) external;  //7,8
    function banG(uint64 _gi, bool _ban, bytes[] memory signs) external;  //11 

    function createGroup(uint16 _level, uint8 _mr, uint256 _kr, uint256 _pr) external; //11
    function reAcc(address _a) external;  // 6
    function reRole(address _a, uint8 _rtype, bytes memory _extra) external; // 6;6
    function addToGroup(address _a, uint64 _gi) external; // 6; 6,8,11
    
    function pledge(address _a, uint64 _i, uint256 _money) external; // 8; 7,5
    function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) external; // 5,6,7,8,11

    function addOrder(address _a, OrderIn memory _oi) external;
    function subOrder(address _a, OrderIn memory _oi) external;

    function recharge(address _a, uint64 _i, uint8 _ti, uint256 _money, bool isLock) external;
    function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) external;
    function proWithdraw(address _a, PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external;

    function get(uint8 _type) external view returns(address); 
}