// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/fileSysIn/IFileSys.sol";
import "../interfaces/erc20In/IERC20.sol";

/// @dev This contract is related to data storage and data reading payment in the file system.
contract FileSys is IFileSys {

    // StoreInfo is at some time
    struct StoreInfo {
        uint64 time; // 什么时刻的状态，last end time
        uint64 size; // 在该存储节点上的存储总量，byte
        uint256 price; // 按周期计费; per cycle
    }

    // ChannelInfo for user pay read to provider
    // TODO: when to change the 'amount' and 'nonce'
    struct ChannelInfo {
        uint256 amount; // available amount
        uint64 nonce; // 防止channel重复提交，pro提交后+1
        uint64 expire; // 用于channel到期，user取回
    }

    // AggOrder is AggregatedOrder is user->provider order and channel
    struct AggOrder {
        uint64 nonce; // 防止order重复提交
        uint64 subNonce; // 用于订单到期
        mapping(uint32 => StoreInfo) sInfo; // 不同代币的支付信息，tokenIndex => StoreInfo
        mapping(uint32 => ChannelInfo) channel; // 不同代币的Channel支付信息，tokenaddr->ChannelInfo
    }

    // FsInfo each user have at most one group
    struct FsInfo {
        bool isActive;
        uint64[] providers; // 指示provider索引
        mapping(uint64 => AggOrder) ao; // 该User对每个Provider的订单信息
    }

    // Settlement indicates billing information
    struct Settlement {
        uint64 time; // store状态改变或支付的时间, 与 epoch 对齐
        uint64 size; // 在该存储节点上的存储总量
        uint256 price; // 累积的sprice(即sizePrice)

        uint256 maxPay; // 对此provider所有user聚合总额度； expected 加和
        uint256 hasPaid; // 已经支付
        uint256 canPay; // 最近一次store/pay时刻，可以支付的金额
        uint256 lost; // lost due to unable response to challenge
        uint256 lostPaid; // pay to repair

        uint256 managePay; // pay for group keepers >= endPaid+linearPaid
        uint256 endPaid; // release when order expire
        uint256 linearPaid; // release when pay for provider
    }

    // fs合约状态变量
    uint8 manageRate = 4; // group分得资金的百分比；4% for group, 其中3% linear and 1% at end;
    uint8 taxRate = 1; // 基金会分得资金的百分比；1% for foundation;
    uint64 gIndex; // 指代所属的group
    uint64 foundation; // 基金会账户的索引（固定为0）

    mapping(uint64 => mapping(uint32 => uint256)) balances; // 账户可用的余额
    mapping(uint64 => mapping(uint32 => uint256)) penalty;  // 由于没有回应挑战而受到的惩罚

    //uint64[] users; 
    mapping(uint64 => FsInfo) fs; // user => FsInfo
    FsInfo repairFs; // 

    uint64[] keepers;
    uint64 period; // keeper根据比例获取收益的时间间隔
    uint64 lastTime; // 上次分利润时间
    mapping(uint32 => uint256) tAcc; // 记录分润值，每次分润后归0，tokenIndex=>num
    uint64 totalCount; // 记录所有keeper触发order相关函数的总次数
    mapping(uint64 => uint64) count; // 记录keeper触发Order相关函数的次数，用于分润

    // uint64[] providers;  // 没用上
    mapping(uint64 => mapping(uint32 => Settlement)) proInfo;

    uint32[] tokens; // user使用某token时候加进来

    address public role; // role合约地址
    address public rolefs; // roleFS合约地址

    /// @dev created by admin; 'r' indicates role-contract address, 'rfs' indicates RoleFS-contract address
    constructor(uint64 founder, uint64 _gIndex, address r, address rfs, uint64[] memory _keepers) {
        role = r;
        rolefs = rfs;

        foundation = founder;
        gIndex = _gIndex;

        repairFs.isActive = true;

        keepers = _keepers;
        period = 1;
        lastTime = uint64(block.timestamp);

        for(uint64 i=0; i<keepers.length; i++){
            count[keepers[i]] = 1;
        }
        totalCount = uint64(keepers.length);
    }

    receive() external payable {}

    //函数修饰符，判断是不是owner调用
    modifier onlyRole(){
        require(msg.sender == role, "N");
        _;
    }

    modifier onlyRoleFS(){
        require(msg.sender == rolefs, "N");
        _;
    }

    function _settlementAdd(uint64 _pIndex, uint32 _tokenIndex, uint64 start, uint64 size, uint256 sprice, uint256 pay, uint256 manage) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        uint256 hp;
        if(se.time < start){
            if(se.time==0){ // 首次addOrder
                hp = 0;
            }else {
                hp = uint256(start-se.time);
            }
            proInfo[_pIndex][_tokenIndex].time = start;
        }else if(se.time > start){
            hp = uint256(se.time - start);
        }
        hp = hp * se.price;
        proInfo[_pIndex][_tokenIndex].canPay += hp;

        // update price and size
        proInfo[_pIndex][_tokenIndex].price += sprice;
        proInfo[_pIndex][_tokenIndex].size += size;

        // update maxPay for pIndex
        proInfo[_pIndex][_tokenIndex].maxPay += pay;

        // pay to keeper, 4% of pay
        proInfo[_pIndex][_tokenIndex].managePay += manage;
    }

    function _settlementSub(uint64 _pIndex, uint32 _tokenIndex, uint64 end, uint64 size, uint256 sprice) internal {
        // update canPay
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        uint256 hp;
        if(se.time < end){
            hp = end - se.time;
            proInfo[_pIndex][_tokenIndex].time = end;
        }else {
            hp = se.time - end;
        }
        hp = hp * se.price;
        proInfo[_pIndex][_tokenIndex].canPay += hp;

        // update size and price
        // 内含判断：
        // require(proInfo[_pIndex][_tokenIndex].price > sprice)
        // require(proInfo[_pIndex][_tokenIndex].size > size)
        proInfo[_pIndex][_tokenIndex].price -= sprice;
        proInfo[_pIndex][_tokenIndex].size -= size;
    }

    // _settlementCal called by func withdraw
    function _settlementCal(uint64 _pIndex, uint32 _tokenIndex, uint256 pay, uint256 lost) internal returns (uint256) {
        Settlement memory se = proInfo[_pIndex][_tokenIndex];
        // 'has paid', or 'lost' is not right
        if(se.hasPaid>pay || lost < se.lost){
            return 0;
        }
        proInfo[_pIndex][_tokenIndex].lost = lost;

        uint64 ntime = uint64(block.timestamp);
        if(se.time < ntime){
            uint256 hp = ntime - se.time;
            hp = hp * se.price;
            proInfo[_pIndex][_tokenIndex].canPay += hp;
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

    // called by Role-contract，在user注册时被role合约调用
    function createFs(uint64 _uIndex) external override onlyRole {
        // 不需要判断_uIndex，因为下面会自动判断,solidity的算术运算是安全运算
        require(!fs[_uIndex].isActive, "E"); // the fs already exists

        //users.push(_uIndex);
        fs[_uIndex].isActive = true;
    }

    // called by Role-contract
    function addKeeper(uint64 _kIndex) external override onlyRole {
        keepers.push(_kIndex);
        count[_kIndex] = 1;
        totalCount++;
    }


    // called by RoleFS-contract
    function addOrder(AOParams memory ps) external override onlyRoleFS {
        require(fs[ps.uIndex].isActive, "NE"); // the fs does not exist

        if(!_hasToken(ps.tIndex)){
            tokens.push(ps.tIndex);
        }

        // 验证金额是否足够
        // 内含判断：
        // require(end >= start)
        uint256 pay = uint256(ps._end-ps._start);
        pay = pay * ps.sPrice;
        uint256 manage = pay / 100 * uint256(manageRate);
        uint256 tax = pay / 100 * uint256(taxRate);
        uint256 payAndTax = pay + manage + tax;
        require(balances[ps.uIndex][ps.tIndex] >= payAndTax, "NE"); // balance is not enough

        if(fs[ps.uIndex].ao[ps.pIndex].channel[ps.tIndex].expire == 0){
            fs[ps.uIndex].providers.push(ps.pIndex);
            fs[ps.uIndex].ao[ps.pIndex].channel[ps.tIndex].expire = ps._end; // cann't be added
        }
        require(fs[ps.uIndex].ao[ps.pIndex].nonce == ps.nonce, "NE"); // nonce error
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].time <= ps._end, "EE"); // end error, end shouldn't less than last order's end

        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].price += ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size += ps._size;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].time = ps._end;

        _settlementAdd(ps.pIndex, ps.tIndex, ps._start, ps._size, ps.sPrice, pay, manage);

        fs[ps.uIndex].ao[ps.pIndex].nonce++;

        // add to foundation
        balances[foundation][ps.tIndex] += tax;
        balances[ps.uIndex][ps.tIndex] -= payAndTax;
    }

    // called by RoleFS-contract
    function subOrder(SOParams memory ps) external override onlyRoleFS {
        // check params
        require(fs[ps.uIndex].ao[ps.pIndex].subNonce == ps.nonce, "EN"); // nonce error
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size >= ps._size, "ES"); // size error

        // update size and price
        // 内含判断：
        // require(fs[uIndex].ao[_pIndex].sInfo[_tokenIndex].price >= sprice)
        // require(fs[uIndex].ao[_pIndex].sInfo[_tokenIndex].size >= size)
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].price -= ps.sPrice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size -= ps._size;

        // update settlement
        _settlementSub(ps.pIndex, ps.tIndex, ps._end, ps._size, ps.sPrice);

        // pay to keeper, 其中的1%在结束时才支付,存储在endPaid中
        uint256 endPaid = ps.sPrice * uint256(ps._end-ps._start);
        endPaid = endPaid / 100;
        proInfo[ps.pIndex][ps.tIndex].endPaid += endPaid;
        tAcc[ps.tIndex] += endPaid;

        // pay to keeper, 3% for linear, do in proWithdraw

        fs[ps.uIndex].ao[ps.pIndex].subNonce++;
        if(ps.kIndex!=ps.uIndex){
            count[ps.kIndex]++;
        }
        totalCount++;
    }

    // 由Role合约调用
    // 调用前需要uAddr先approve本合约地址（FileSys.sol）账户指定金额
    // user往文件系统FileSys合约中充值
    function recharge(uint64 uIndex, uint32 tIndex, address uAddr, address tAddr, uint256 money) external override onlyRole {
        require(fs[uIndex].isActive, "NA"); // the fs with user is not active

        // add tIndex
        if(!_hasToken(tIndex)){
            tokens.push(tIndex);
        }

        IERC20(tAddr).transferFrom(uAddr, address(this), money);

        balances[uIndex][tIndex] += money;
    }

    // provider withdraw money, called by owner
    function proWithdraw(PWParams memory ps) external override onlyRoleFS {
        // pay to provider
        uint256 thisPay = _settlementCal(ps.pIndex, ps.tIndex, ps.pay, ps.lost);
        if(thisPay==0){
            return;
        }

        // linear pay to keepers
        uint256 lpay = thisPay * 3 / 100; // 3% thisPay for linearPaid
        proInfo[ps.pIndex][ps.tIndex].linearPaid += lpay;
        tAcc[ps.tIndex] += lpay;

        IERC20(ps.tAddr).transfer(ps.pAddr, thisPay); // transfer fail, the previous operation will be rolled back 
    }

    // user、keeper、foundation取回余额, called by owner
    function withdraw(uint64 index, uint32 tokenIndex, uint8 roleType, address tAddr, address addr, uint256 amount) external override onlyRole {
        if(roleType == 3){ // keeper
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
                            uint64 kc = count[keepers[j]]; // 触发的次数
                            if(kc==0){
                                count[keepers[j]] = 1;
                            }else{
                                uint256 pro = per * kc;
                                balances[keepers[j]][tokens[i]] += pro;
                            }
                        }
                        tAcc[tokens[i]] = 0;
                    }
                }

                lastTime = ntime;
            }
        }
        uint256 bal = balances[index][tokenIndex];
        if(amount>bal){
            amount=bal;
        }

        balances[index][tokenIndex] -= amount;
        IERC20(tAddr).transfer(addr, amount);
    }

    // called by owner
    function addRepair(uint64 kIndex, uint64 pIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) external override onlyRoleFS {
        // verify money is enough
        Settlement memory se = proInfo[pIndex][tokenIndex];
        uint256 bal = se.lost - se.lostPaid;
        uint256 pay = sprice * uint256(end-start);
        uint256 per = pay / 100;
        uint256 manage = per * manageRate;
        require(bal >= pay, "NE"); // balance not enough

        // add newProvider
        if(!_hasP(newPro)){
            repairFs.providers.push(newPro);
        }
        require(repairFs.ao[newPro].nonce == nonce, "IN"); // illegal nonce, nonce should be same
        // start > current - 2*epoch
        repairFs.ao[newPro].sInfo[tokenIndex].price += sprice;
        repairFs.ao[newPro].sInfo[tokenIndex].size += size;

        _settlementAdd(newPro, tokenIndex, start, size, sprice, pay, manage);

        repairFs.ao[newPro].nonce++;
        proInfo[pIndex][tokenIndex].lostPaid += pay;
        count[kIndex]++;
        totalCount++;
    }

    // called by owner
    function subRepair(uint64 kIndex, uint64 newPro, uint64 start, uint64 end, uint64 size, uint64 nonce, uint32 tokenIndex, uint256 sprice) external override onlyRoleFS {
        require(repairFs.ao[newPro].subNonce == nonce, "EN"); // error nonce

        require(repairFs.ao[newPro].sInfo[tokenIndex].size >= size, "ES"); // error size
        // change
        repairFs.ao[newPro].sInfo[tokenIndex].price -= sprice; // 等价于 require(repairFs.ao[newPro].sInfo[tokenIndex].price > sprice)
        repairFs.ao[newPro].sInfo[tokenIndex].size -= size; // 等价于 require(repairFs.ao[newPro].sInfo[tokenIndex].size > size)

        _settlementSub(newPro, tokenIndex, end,size,sprice);

        // pay to keeper, 1% for endpay
        uint256 endPaid = sprice * uint256(end-start);
        endPaid = endPaid / 100;
        proInfo[newPro][tokenIndex].endPaid += endPaid;
        tAcc[tokenIndex] += endPaid;

        // pay to keeper, 4% for linear;due to pro no trigger pay

        repairFs.ao[newPro].subNonce++;
        count[kIndex]++;
        totalCount++;
    }

    // ================get=============

    // get some parameters in FsInfo
    function getFsInfo(uint64 user) external view override returns (bool){
        if(user==0){
            return repairFs.isActive;
        }
        return fs[user].isActive;
    }

    // get providers sum in fs info
    function getFsPSum(uint64 user) external view override returns (uint64) {
        if(user==0){
            return uint64(repairFs.providers.length);
        }
        return uint64(fs[user].providers.length);
    }

    // get provider in fs info by array index
    function getFsPro(uint64 user, uint64 i) external view override returns (uint64) {
        if(user==0){
            return repairFs.providers[i];
        }
        return fs[user].providers[i];
    }

    // get provider's aggregate order in FsInfo
    function getFsInfoAggOrder(uint64 user, uint64 provider) external view override returns (uint64, uint64) {
        if(user==0){
            return (repairFs.ao[provider].nonce, repairFs.ao[provider].subNonce);
        }
        return (fs[user].ao[provider].nonce,fs[user].ao[provider].subNonce);
    }

    // get storeInfo in fs
    function getStoreInfo(uint64 user, uint64 provider, uint32 token) external view override returns (uint64, uint64, uint256){
        if(user==0){
            return (repairFs.ao[provider].sInfo[token].time, repairFs.ao[provider].sInfo[token].size, repairFs.ao[provider].sInfo[token].price);
        }
        return (fs[user].ao[provider].sInfo[token].time, fs[user].ao[provider].sInfo[token].size, fs[user].ao[provider].sInfo[token].price);
    }
    
    // get channelInfo in fs
    function getChannelInfo(uint64 user, uint64 provider, uint32 token) external view override returns (uint256, uint64, uint64){
        if(user==0){
            return (repairFs.ao[provider].channel[token].amount, repairFs.ao[provider].channel[token].nonce, repairFs.ao[provider].channel[token].expire);
        }
        return (fs[user].ao[provider].channel[token].amount, fs[user].ao[provider].channel[token].nonce, fs[user].ao[provider].channel[token].expire);
    }

    // judge if tokens has the _tokenIndex
    function _hasToken(uint32 _tokenIndex) internal view returns (bool) {
        for(uint32 i = 0; i<tokens.length; i++){
            if(tokens[i]==_tokenIndex){
                return true;
            }
        }
        return false;
    }

    // judge if FsInfo.providers include the pIndex
    function _hasP(uint64 pIndex) internal view returns (bool) {
        for(uint256 i = 0; i<repairFs.providers.length; i++){
            if(repairFs.providers[i]==pIndex){
                return true;
            }
        }
        return false;
    }

    // 获得支付计费相关的信息
    function getSettleInfo(uint64 pIndex, uint32 tIndex) external view override returns (uint64, uint64, uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256){
        Settlement memory se = proInfo[pIndex][tIndex];
        return (se.time,se.size,se.price,se.maxPay,se.hasPaid,se.canPay,se.lost,se.lostPaid,se.managePay,se.endPaid,se.linearPaid);
    }

    // 获得账户收益余额
    function getBalance(uint64 index, uint32 tIndex) external view override returns(uint256, uint256){
        uint256 avail = balances[index][tIndex];
        Settlement memory se = proInfo[index][tIndex];
        uint256 canPay = se.canPay;
        uint256 tmp = uint256(block.timestamp - se.time);
        tmp = tmp * se.size;
        tmp = tmp * se.price;

        canPay += tmp;
        tmp = se.maxPay - se.lost;
        if(canPay > tmp){
            canPay = tmp;
        }
        tmp = canPay - se.hasPaid;

        if(totalCount == 0){
            return (avail, tmp);
        }

        if(count[index]!=0){ // 如果index是keeper,就需要根据比例计算分润值
            uint256 sum = tAcc[tIndex] * count[index];
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