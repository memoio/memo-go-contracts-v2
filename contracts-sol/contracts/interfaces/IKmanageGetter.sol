// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;


interface IKmanageGetter {
    function getRate() external view returns (uint8);
    function getKCnt() external view returns (uint16);
    function getK(uint64 _i) external view returns (uint64);
    function getSP(uint8 _ti) external view returns (uint64, uint256); 
    function balanceOf(uint64 _ki, uint8 _ti) external view returns(uint256, uint256);
}