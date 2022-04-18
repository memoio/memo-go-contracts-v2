// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IKmanage.sol";
import "../interfaces/IAuth.sol";
import "./Owner.sol";

/// @dev This contract is related to data storage and data reading payment in the file system.
contract Kmanage is IKmanage, Owner {
    uint16 public version = 2;

    uint8 public manageRate = 4; // group分得资金的百分比；4% for group, 其中3% linear and 1% at end;

    uint64 period;    // keeper根据比例获取收益的时间间隔
    uint64 lastTime;  // 上次分利润时间

    uint64[] keepers; // for profits
    mapping(uint8 => uint256) tAcc; // 记录分润值，每次分润后归0，tokenIndex=>num
    uint64 totalCount; // 记录所有keeper触发order相关函数的总次数
    mapping(uint64 => uint64) count; // 记录keeper触发Order相关函数的次数，用于分润

    mapping(uint64 => mapping(uint8 => uint256)) balances; // 账户可用的余额
    uint8[] tokens;

    /// @dev created by admin; 'r' indicates role-contract address, 'rfs' indicates RoleFS-contract address
    constructor(address _o, address _a, uint8 mr) Owner(_o, _a) {
        manageRate = mr;
        period = 7*86400; // one week
        lastTime = uint64(block.timestamp);
    }

    // judge if tokens has the _tokenIndex
    function _hasToken(uint8 _ti) internal view returns (bool) {
        for(uint8 i = 0; i<tokens.length; i++){
            if(tokens[i]==_ti){
                return true;
            }
        }
        return false;
    }

    // after add keeper to group
    function addKeeper(uint64 _ki) external override onlyOwner {
        keepers.push(_ki);
        count[_ki] = 1;
        totalCount++;
    }

    // after sub order
    function add(uint64 _ki) external override onlyOwner {
        require(count[_ki] > 0, "IC");

        count[_ki]++;
        totalCount++;
    }

    // after pro withdraw
    function addProfit(uint8 _ti, uint256 _money) external override onlyOwner {
        tAcc[_ti] += _money;
    }

    function withdraw(uint64 _ki, uint8 _ti, uint256 amount) external override onlyOwner returns(uint256){
        if (count[_ki] > 0) {
            uint64 ntime = uint64(block.timestamp);
            if(ntime-lastTime > period){
                if(totalCount != 0){
                    for(uint i = 0; i< tokens.length; i++){
                        uint256 value = tAcc[tokens[i]]; // 分润值
                        if(value==0){
                            continue;
                        }
                        uint256 per = value / uint256(totalCount);
                        for(uint j =0; j<keepers.length; j++){
                            uint256 pro = per * count[keepers[j]];
                            balances[keepers[j]][tokens[i]] += pro;
                            tAcc[tokens[i]] -= pro;
                        }
                    }
                }
                lastTime = ntime;
            }
        }
    
        uint256 bal = balances[_ki][_ti];
        if(amount>bal){
            amount=bal;
        }

        balances[_ki][_ti] -= amount;
        return amount;
    }   

    // ================get=============
    function getRate() external view override returns (uint8) {
        return manageRate;
    }

    function balanceOf(uint64 _i, uint8 _ti) external view override returns(uint256, uint256){
        uint256 avail = balances[_i][_ti];
        uint256 tmp = 0;

        if(totalCount == 0){
            return (avail, tmp);
        }

        if(count[_i]!=0){
            uint256 sum = tAcc[_ti] * count[_i];
            uint256 pro = sum / totalCount;
            if((block.timestamp-lastTime)>=period){
                avail += pro;
            }else{
                tmp+=pro;
            }
        }

        return (avail, tmp);
    }
}