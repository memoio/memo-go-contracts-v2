//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/erc20In/IERC20.sol";
import "../interfaces/pledgePoolIn/IPledgePool.sol";
import "../interfaces/roleIn/IRole.sol";
import "../role/RToken.sol";
import "../Recover.sol";

/**
 *@author MemoLabs
 *@title Any who has registered can introduce the pledge pool and get income
 */
contract PledgePool is IPledgePool {
    using Recover for bytes32;

    bytes4 private constant SELECTOR0 = bytes4(keccak256(bytes('balanceOf(address)')));

    struct RewardInfo {
        uint256 rewardAccu; // 本代币的accumulator
        uint256 lastReward; // 上一次变更时的本代币奖励总量。等于余额加奖励
    }

    // uint32 token;  // largest token-index，指示要添加进来的代币索引tokenIndex
    // address[] tokens; // 所有支持的代币的信息,0为主代币的代币地址
    address public rToken;
    address public role;
    uint256 public totalPledge; // 账户质押量
    mapping(uint32 => RewardInfo) tInfo;  // 每种代币的信息, tokenIndex=>RewardInfo
    mapping(uint64 => mapping(uint32 => RewardInfo)) allAmount; // 所有质押的人的信息，mapping(roleIndex => tokenIndex => RewardInfo)

    event Pledge(address indexed from, uint256 money);
    event Withdraw(address indexed from, uint256 money);

    /// @dev created by admin; 'r' indicates role-contract address
    constructor(address primeToken, address _rToken, address r) {
        rToken = _rToken;

        role = r;

        // address.call(bytes  memory) returns (bool, bytes memory)
        (bool success, bytes memory data) = primeToken.staticcall(abi.encodeWithSelector(SELECTOR0, address(this)));
        require(success && data.length >= 32);
        uint256 balance = abi.decode(data, (uint256)); // PledgePool合约账户在primeToken代币中的余额

        tInfo[0].lastReward = balance;
    }

    receive() external payable {}

    function verifySign(address addr, bytes32 h, bytes memory sign) private pure {
        require(h.recover(sign) == addr, "IC"); // illegal caller
    }

    function updateReward(uint256 amount, uint32 i, uint64 index) internal {
        uint256 bal = _getBalance(RToken(rToken).getTA(i), address(this)); // 获取质押池在该代币上的余额

        uint256 tv; // 增发量
        if (bal > tInfo[i].lastReward){ // 有增发，针对主代币即addOrder过程触发了mintToken
            tv = bal - tInfo[i].lastReward;
        }

        // 代币i有增发，需更新tInfo[i].rewardAccu
        if(tv>0 && totalPledge>0){
            tv = tv * 1e9 / totalPledge; // 存在精度问题，可能一直为0，所以将totalPledge除以1e9
            tInfo[i].rewardAccu += tv;
        }
        tInfo[i].lastReward = bal; // update to latest

        // 将账户余额加上应得的分润值
        // res表示：距离上次更新分润值期间，每质押主代币应得的代币i分润值的累积总和
        uint256 res = tInfo[i].rewardAccu - allAmount[index][i].rewardAccu;
        res = res * amount / 1e9; // amount应大于1e9
        allAmount[index][i].lastReward += res; // 添加分润值
        allAmount[index][i].rewardAccu = tInfo[i].rewardAccu; // 更新accu
    }

    // 非流动性质押，账户本身调用或由其他账户代为调用，质押主代币
    // 调用前需要index指示的账户approve本合约（也就是pledgePool合约）账户指定的金额（也就是money）
    // 用户调用register()后，如果想要申请keeper、provider角色，就需要先质押
    // 质押量过小时（具体值不能确定，但小于1e9），由于合约整除，可能前几次的分润值为0，随着代币的多次分发，后面才能获得分润值
    /**
     *@notice Pledge money.
     *@dev Called by the account itself or by another account on its behalf.
     *@param index The index of the account.
     *@param money The value that the account want to pledge.
     *@param sign The signature by index-account, hash(caller,money,"pledge").
     */
    function pledge(uint64 index, uint256 money, bytes memory sign) external override payable {
        address caller = msg.sender;
        IRole r = IRole(payable(role));
        address acc = r.getAddr(index-1);
        (,bool isBanned,,,,) = r.getRoleInfo(acc);
        require(!isBanned, "B"); // the index account is banned

        if(caller != acc){ // 由其他账户代为调用，需要验证签名
            bytes32 h = keccak256(abi.encodePacked(caller, money, "pledge"));
            verifySign(acc, h, sign);
        }

        // TODO：如何防止该签名被重复利用

        RewardInfo memory reward = allAmount[index][0];
        uint256 amount = reward.lastReward;

        // 更新token accumulator，结算奖励
        uint32 tLen = RToken(rToken).getTNum();
        for(uint32 i=0; i<tLen; i++){
            reward = allAmount[index][i];
            if(reward.rewardAccu==0 && reward.lastReward==0){
                allAmount[index][i].rewardAccu = tInfo[i].rewardAccu;
            }

            updateReward(amount, i, index);
        }

        IERC20(RToken(rToken).getTA(0)).transferFrom(acc, address(this), money); // msg.sender is address(this); tx.origin is acc

        // update
        tInfo[0].lastReward += money;
        allAmount[index][0].lastReward += money;
        totalPledge += money;

        emit Pledge(acc, money);
    }

    // 由账户本身调用或由其他账户代为调用
    // hash(caller, tIndex, money),  从质押池取回余额
    function withdraw(uint64 index, uint32 tIndex, uint256 money, bytes memory sign) external override {
        // check params
        address addr = IRole(payable(role)).getAddr(index-1);
        uint256 lock;

        {
            (,bool isBanned,uint8 rtype,,,) = IRole(payable(role)).getRoleInfo(addr);
            require(!isBanned, "B"); // the index account is banned

            if(msg.sender != addr){ // 由其他账户代为调用，需要验证签名
                bytes32 h = keccak256(abi.encodePacked(msg.sender, tIndex, money));
                verifySign(addr, h, sign);
            }

            if(rtype==2){
                lock = IRole(payable(role)).pledgeP();
            }
            if(rtype==3){
                lock = IRole(payable(role)).pledgeK();
            }
        }

        // TODO：如何防止该签名被重复利用

        address tAddr = RToken(rToken).getTA(tIndex);
        require(tAddr != address(0), "TE"); // tIndex error

        RewardInfo memory reward = allAmount[index][0];
        if(reward.lastReward==0 && reward.rewardAccu==0){
            return;
        }

        uint256 amount = reward.lastReward;

        if(totalPledge>0){
            // 更新token accumulator，结算奖励
            uint32 tLen = RToken(rToken).getTNum();
            for(uint32 i=0; i<tLen; i++) {
                if(tIndex!=0 && tIndex!=i){ //只更新主代币和tIndex代币
                    continue;
                }

                updateReward(amount, i, index);
            }
        }

        reward = allAmount[index][tIndex];
        if(reward.rewardAccu==0 && reward.lastReward==0){
            return;
        }
        uint256 rw = reward.lastReward;
        if(tIndex==0){
            if (rw > lock){
                rw -= lock;
            }else {
                rw = 0;
            }
        }
        if(money>0 && money<rw){
            rw = money;
        }

        if(rw>0){
            IERC20(tAddr).transfer(addr, rw); // TF:transfer fail；此合约账户转账给acc账户

            // update token
            if (tInfo[tIndex].lastReward > rw){
                tInfo[tIndex].lastReward -= rw;
            }else {
                tInfo[tIndex].lastReward = 0;
            }

            // update value
            if (allAmount[index][tIndex].lastReward > rw){
                allAmount[index][tIndex].lastReward -= rw;
            }else {
                allAmount[index][tIndex].lastReward = 0;
            }
        }

        if(tIndex==0){
            totalPledge += allAmount[index][tIndex].lastReward;
            totalPledge -= amount;
        }

        emit Withdraw(addr, rw);
    }

    // by owner, that is role-contract
    function addToken(address tAddr, uint32 tIndex) external override {
        require(msg.sender == role, "N");
        tInfo[tIndex].rewardAccu = 0;
        tInfo[tIndex].lastReward = _getBalance(tAddr, address(this));
    }

    // ========== get ===========

    // 获得账户在指定代币上的余额
    function _getBalance(address tAddr, address account) internal view returns (uint256) {
        (bool success, bytes memory data) = tAddr.staticcall(abi.encodeWithSelector(SELECTOR0, account));
        require(success && data.length >= 32, "CE"); // staticcall error
        return abi.decode(data, (uint256));
    }

    // 获得指定代币的质押总额，即该质押池账户在指定代币上的余额。除了账户质押的部分，还有增发的部分。
    // 目前质押函数只开放了主代币的质押
    function getPledge(uint32 tIndex) external view override returns (uint256) {
        address tAddr = RToken(rToken).getTA(tIndex);
        require(tAddr != address(0), "TE"); // tIndex error

        return _getBalance(tAddr, address(this));
    }

    // 获得指定账户index对于指定代币tindex的分润值
    // 仍有一些逻辑问题和设计问题需要完善
    function getBalance(uint64 index, uint32 tIndex) external view override returns (uint256) {
        address tAddr = RToken(rToken).getTA(tIndex);
        require(tAddr != address(0), "TE"); // tIndex error

        // 没有质押过，返回0
        RewardInfo memory reward0 = allAmount[index][0];
        if (reward0.rewardAccu==0 && reward0.lastReward==0) {
            return 0;
        }
        uint256 amount = reward0.lastReward;

        RewardInfo memory rewardt = tInfo[tIndex];
        uint256 val = rewardt.rewardAccu;

        uint256 bal = _getBalance(tAddr, address(this));
        bal -= rewardt.lastReward;
        if(bal > 0 && totalPledge > 0) {
            bal = bal * 1e9 / totalPledge;
            val = val + bal;
        }

        RewardInfo memory rewardi = allAmount[index][tIndex];
        if(rewardi.rewardAccu==0 && rewardi.lastReward==0) {
            val = val * amount / 1e9;
        }else{
            val = val - rewardi.rewardAccu;
            val = val * amount / 1e9;
            val = val + rewardi.lastReward;
        }

        return val;
    }
}