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
        uint32 tIndex;
        uint256 sPrice;
        bytes usign;
        bytes psign;
}

// struct for proWithdraw params
struct PWIn {
    uint64 pIndex;
    uint32 tIndex;
    uint256 pay;
    uint256 lost;
}

struct StoreOut {
    uint64 start; // last start 
    uint64 end;   // 什么时刻的状态，last end time
    uint64 size;   // 在该存储节点上的存储总量，byte
    uint256 price; // 按周期计费; per cycle
}

struct FsOut {
    bool isActive;
    uint64 nonce;    // 防止order重复提交
    uint64 subNonce; // 用于订单到期
}

// Settlement indicates billing information
struct SettleOut {
    uint64 time; // store状态改变或支付的时间, 与 epoch 对齐
    uint64 size; // 在该存储节点上的存储总量
    uint256 price; // 累积的sprice(即sizePrice)

    uint256 maxPay;  // 对此provider所有user聚合总额度； expected 加和
    uint256 hasPaid; // 已经支付
    uint256 canPay;  // 最近一次store/pay时刻，可以支付的金额
    uint256 lost;    // lost due to unable response to challenge
    uint256 lostPaid;// pay to repair

    uint256 managePay; // pay for group keepers >= endPaid+linearPaid
    uint256 endPaid;   // release when order expire
    uint256 linearPaid;// release when pay for provider
}

interface IFileSysSetter {
    function addUser(uint64 _uIndex) external;
    function addKeeper(uint64 _kIndex) external;
    function addOrder(OrderIn memory ps) external;
    function subOrder(uint64 _kIndex, OrderIn memory ps) external;

    function recharge(uint64 uIndex, uint32 tIndex, uint256 money) external;
    function withdraw(uint64 uIndex, uint32 tIndex, uint256 money) external returns (uint256);
    function proWithdraw(PWIn memory ps) external returns(uint256);
}

interface IFileSysGetter {
    function balanceOf(uint64 index, uint32 tIndex) external view returns(uint256, uint256);
    function getFsInfo(uint64 user, uint64 provider) external view returns (FsOut memory);
    function getStoreInfo(uint64 user, uint64 provider, uint32 token) external view returns (StoreOut memory);
    function getSettleInfo(uint64 index, uint32 tIndex) external view returns (SettleOut memory);
}