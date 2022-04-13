// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// for manage erc20 
interface IAccess {
    function set(address account,bool isSet, bytes[] memory signs) external;
    function can(address account) external returns (bool);
}