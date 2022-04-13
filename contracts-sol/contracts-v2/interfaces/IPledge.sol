// SPDX-License-Identifier:UNLICENED
pragma solidity ^0.8.0;

interface IPledgeSetter {
    function pledge(uint64 index, uint256 money) external;
    function withdraw(uint64 index, uint32 tIndex, uint256 money, uint256 lock) external returns (uint256);

    function addT(uint32 tIndex) external;
}

interface IPledgeGetter {
    function balanceOf(uint64 index, uint32 tIndex) external view returns (uint256);
}