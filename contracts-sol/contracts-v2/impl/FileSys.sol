// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IFileSys.sol";
import "../interfaces/IAuth.sol";
import "../interfaces/IERC20.sol";
import "./Owner.sol";

/// @dev This contract is related to data storage and data reading payment in the file system.
contract FileSys is IFileSys, Owner {

    struct StoreInfo {
        uint64 start; // last start 
        uint64 end;   // 什么时刻的状态，last end time
        uint64 size;   // 在该存储节点上的存储总量，byte
        uint256 price; // 按周期计费; per cycle
    }

    struct AggOrder {
        uint64 nonce;    // 防止order重复提交
        uint64 subNonce; // 用于订单到期
        mapping(uint8 => StoreInfo) sInfo; // 不同代币的支付信息，tokenIndex => StoreInfo
    }

    struct FsInfo {
        uint64[] providers;             // provider索引
        mapping(uint64 => AggOrder) ao; // 该User对每个Provider的订单信息
    }

    // Settlement indicates billing information
    struct Settlement {
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

    uint16 public version = 2;

    uint8 public constant taxRate = 1;

    mapping(uint64 => mapping(uint8 => uint256)) balances; // 账户可用的余额

    mapping(uint64 => FsInfo) fs; // user => FsInfo; user 0 is repair fs

    mapping(uint64 => mapping(uint8 => Settlement)) proInfo; // pro => token => income

    constructor(address _rfs, address _a) Owner(_rfs, _a) {
    }

    function _settlementAdd(uint64 _pIndex, uint8 _tokenIndex, uint64 start, uint64 size, uint256 sprice, uint256 pay, uint256 manage) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(se.time < start){
            if(se.time!=0){ // 非首次addOrder
                proInfo[_pIndex][_tokenIndex].canPay += (start-se.time) * se.price;
            }
            proInfo[_pIndex][_tokenIndex].time = start;
        }else if(se.time > start){
            proInfo[_pIndex][_tokenIndex].canPay += uint256(se.time - start)*sprice;
        }

        proInfo[_pIndex][_tokenIndex].price += sprice;
        proInfo[_pIndex][_tokenIndex].size += size;
        proInfo[_pIndex][_tokenIndex].maxPay += pay; // update maxPay; hardlimit
        proInfo[_pIndex][_tokenIndex].managePay += manage; // pay to keeper, 4% of pay
    }

    // roughly
    function _settlementSub(uint64 _pIndex, uint8 _tokenIndex, uint64 end, uint64 size, uint256 sprice) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(se.time < end){
            if (se.time != 0) {
                proInfo[_pIndex][_tokenIndex].canPay += (end - se.time) * se.price;
            }
            proInfo[_pIndex][_tokenIndex].time = end;
        } else if(se.time > end) {
            // should sub it
            uint256 hp = (se.time - end) * sprice;
            if (proInfo[_pIndex][_tokenIndex].canPay > hp) {
                proInfo[_pIndex][_tokenIndex].canPay -= hp;
            } else {
                proInfo[_pIndex][_tokenIndex].canPay = 0;
            }
        }

        // update size and price
        if (proInfo[_pIndex][_tokenIndex].price > sprice) {
            proInfo[_pIndex][_tokenIndex].price -= sprice;  
        } else {
            proInfo[_pIndex][_tokenIndex].price = 0; 
        }

        if (proInfo[_pIndex][_tokenIndex].size > size) {
            proInfo[_pIndex][_tokenIndex].size -= size;
        } else {
            proInfo[_pIndex][_tokenIndex].size = 0;
        }
    }

    // _settlementCal called by func withdraw
    function _settlementCal(uint64 _pIndex, uint8 _tokenIndex, uint256 pay, uint256 lost) internal returns (uint256) {
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        if(proInfo[_pIndex][_tokenIndex].maxPay<pay){
            return 0;
        }
        
        // 'has paid', or 'lost' is not right
        if(se.hasPaid > pay || se.lost > lost){
            return 0;
        }
        proInfo[_pIndex][_tokenIndex].lost = lost;

        uint64 ntime = uint64(block.timestamp);
        if(se.time < ntime){
            proInfo[_pIndex][_tokenIndex].canPay += (ntime - se.time) * se.price;
            proInfo[_pIndex][_tokenIndex].time = ntime;
        }

        // can pay is not right
        if(proInfo[_pIndex][_tokenIndex].canPay<pay){
            return 0;
        }

        uint256 res = pay - se.hasPaid;
        proInfo[_pIndex][_tokenIndex].hasPaid = pay;
        return res;
    }

    function addOrder(OrderIn memory ps, uint256 mr) external override onlyOwner {
        uint256 pay = (ps.end-ps.start) * ps.sPrice;
        uint256 manage = pay / 100 * uint256(mr);
        uint256 tax = pay / 100 * uint256(taxRate);
        uint256 payAndTax = pay + manage + tax;
        require(balances[ps.uIndex][ps.tIndex] >= payAndTax, "BNE"); // balance not enough

        // 验证nonce
        require(fs[ps.uIndex].ao[ps.pIndex].nonce == ps.nonce, "NE"); // nonce error
        // start不减, end不减
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start <= ps.start, "SE"); // start error, start shouldn't less than last order's start
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end <= ps.end, "EE"); // end error, end shouldn't less than last order's end

        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].price += ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size += ps.size;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start = ps.start;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end = ps.end;

        _settlementAdd(ps.pIndex, ps.tIndex, ps.start, ps.size, ps.sPrice, pay, manage);

        fs[ps.uIndex].ao[ps.pIndex].nonce++;

        // add to foundation
        balances[0][ps.tIndex] += tax;
        balances[ps.uIndex][ps.tIndex] -= payAndTax;
    }

    function subOrder(OrderIn memory ps) external override onlyOwner returns(uint256) {

        require(fs[ps.uIndex].ao[ps.pIndex].subNonce == ps.nonce, "EN"); // nonce error

        // update size and price
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].price -= ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size -= ps.size;

        // update settlement
        _settlementSub(ps.pIndex, ps.tIndex, ps.end, ps.size, ps.sPrice);

        // pay to keeper, 其中的1%在结束时才支付,存储在endPaid中
        uint256 endPaid = ps.sPrice * uint256(ps.end-ps.start) / 100;
        proInfo[ps.pIndex][ps.tIndex].endPaid += endPaid;

        fs[ps.uIndex].ao[ps.pIndex].subNonce++;
        
        // pay to keeper, 3% for linear, do in proWithdraw
        return endPaid;
    }

    function recharge(uint64 uIndex, uint8 tIndex, uint256 money) external override onlyOwner {
        balances[uIndex][tIndex] += money;
    }

    // provider withdraw money, called by owner
    function proWithdraw(PWIn memory ps) external override onlyOwner returns(uint256, uint256) {
        // pay to provider
        uint256 thisPay = _settlementCal(ps.pIndex, ps.tIndex, ps.pay, ps.lost);
        if(thisPay==0){
            return (0,0);
        }

        // linear pay to keepers
        uint256 lpay = thisPay * 3 / 100; // 3% thisPay for linearPaid
        proInfo[ps.pIndex][ps.tIndex].linearPaid += lpay;
        
        return (thisPay, lpay);
    }

    function withdraw(uint64 index, uint8 tokenIndex, uint256 amount) external override onlyOwner returns(uint256){
        uint256 bal = balances[index][tokenIndex];
        if(amount>bal){
            amount=bal;
        }

        balances[index][tokenIndex] -= amount;
        return amount;
    }

    // ================get=============
    function getFsInfo(uint64 user, uint64 pro) external view override returns (FsOut memory){
        FsOut memory f;
        return f;
    }

    // get storeInfo in fs
    function getStoreInfo(uint64 user, uint64 provider, uint8 token) external view override returns (StoreOut memory){
        StoreOut memory s;
        s.start = fs[user].ao[provider].sInfo[token].start;
        return s;
    }

    // 获得支付计费相关的信息
    function getSettleInfo(uint64 pIndex, uint8 tIndex) external view override returns (SettleOut memory){
        SettleOut memory s;
        s.time = proInfo[pIndex][tIndex].time;
        return s;
    }

    // 获得账户收益余额
    function balanceOf(uint64 index, uint8 tIndex) external view override returns(uint256, uint256){
        uint256 avail = balances[index][tIndex];
        Settlement memory se = proInfo[index][tIndex];
        uint256 canPay = se.canPay;
        if (block.timestamp > se.time) {
            canPay += uint256(block.timestamp - se.time) * se.price;
        }

        uint256 tmp = se.maxPay - se.lost;
        if(canPay > tmp){
            canPay = tmp;
        }
        tmp = canPay - se.hasPaid;

        return (avail, tmp);
    }
}