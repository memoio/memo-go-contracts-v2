// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface ITokenSetter {
    function addT(address t) external returns(uint32);
    function banT(address t) external; 
}

interface ITokenGetter {
    function getTA(uint32 tIndex) external view returns (address, bool);
    function getTI(address t) external view returns (uint32, bool);
}