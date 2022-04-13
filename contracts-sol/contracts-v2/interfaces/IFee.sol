// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// record fee in all fs
interface IFeeSetter {
    function recharge(uint64 index, uint32 tokenIndex, uint256 money) external;
    function withdraw(uint64 index, uint32 tokenIndex, uint256 money) external;
}

interface IFeeGetter {
    function balanceOf(uint64 index, uint32 tIndex) external view returns (uint256);
}