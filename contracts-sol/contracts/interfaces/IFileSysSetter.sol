// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;


// struct for addOrder params
struct OrderIn {
    uint64 uIndex;
    uint64 pIndex;
    uint64 start;
    uint64 end;
    uint64 size;
    uint64 nonce;
    uint8 tIndex;
    uint256 sPrice;
    bytes usign;
    bytes psign;
}

// struct for proWithdraw params
struct PWIn {
    uint64 pIndex;
    uint8 tIndex;
    uint256 pay;
    uint256 lost;
}


interface IFileSysSetter {
    function addOrder(OrderIn memory ps, uint256 _mr) external;
    function subOrder(OrderIn memory ps, uint256 _mr) external returns (uint256) ;

    function recharge(uint64 _i, uint8 _ti, uint256 money, bool isLock) external;
    function withdraw(uint64 _i, uint8 _ti, uint256 money) external returns (uint256);
    function proWithdraw(PWIn memory ps, uint256 _mr) external returns(uint256, uint256);
}