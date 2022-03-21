//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/fileSysIn/IFileSys.sol";
import "../interfaces/pledgePoolIn/IPledgePool.sol";
import "../interfaces/erc20In/IERC20.sol";
import "../interfaces/roleIn/IIssuance.sol";
import "./RToken.sol";
import "../Recover.sol";
import "../Owner.sol";
import "../interfaces/roleIn/IRole.sol";

/**
 *@author MemoLabs
 *@title Used to register account roles for people participating in the memo system
 */
contract Role is IRole, Owner {
    using Recover for bytes32;

    struct RoleInfo {
        bool isActive; // user注册后即生效，k-p注册后还需要质押才能生效
        bool isBanned; // 是否被admin禁止
        uint8 roleType; // 0 account; 1 user; 2 provider; 3 keeper
        uint64 index; // 账户序列号，从1开始
        uint64 gIndex; // 所属group，从1开始
        bytes extra; // 用于存储额外的信息，一般是公钥信息
    }

    struct GroupInfo {
        bool isActive;
        bool isBanned; // 组是否已被禁用
        bool isReady;  // 是否已在线下成组; 由签名触发
        uint16 level;  // security level
        uint64[] keepers; // 里面有哪些Keeper
        uint64[] providers; // 有哪些provider
        uint64[] users;
        uint256 size;  // storeSize
        uint256 price; //storePrice
        address fsAddr;  // fs contract address
    }

    address public override pledgePool; // pledge pool 合约地址
    address public foundation; // foundation address 基金会账户地址，索引默认为0,不需要进行register

    address[] addrs; // all roles 序列号即为index,从1开始
    mapping(address => RoleInfo) info;
    uint256 public override pledgeK; // keeper质押的最小金额
    uint256 public override pledgeP; // provider质押的最小金额

    // manage group
    GroupInfo[] groups;

    address public rToken;
    address public issuance;
    address public rolefs;

    uint16 public version;

    event RUser(uint64 index, address addr);
    event RKeeper(uint64 index, address addr);
    event RProvider(uint64 index, address addr);
    event CreateGroup(uint64 gIndex);

    // called by admin, specify foundation、primaryToken、pledgeK、pledgeP, and create RToken-contract
    // 创建后需要根据该合约地址创建PledgePool合约及Issuance合约，然后调用setPI函数将PledgePool合约地址及Issuance合约地址赋值给该合约
    constructor(address f, address t, uint256 pk, uint256 ppro, uint16 _version) {
        foundation = f;
        pledgeK = pk;
        pledgeP = ppro;
        version = _version;

        // create RToken-contract by "create2"
        bytes32 salt = keccak256(abi.encodePacked(f, t));
        address rtoken;
        bytes memory b = type(RToken).creationCode;
        assembly {
            rtoken := create2(0, add(b, 0x20), mload(b), salt)
        }
        rToken = rtoken;
        RToken r = RToken(rtoken);
        r.addT(t);
    }

    receive() external payable {}

    /// @dev set pledgePool and issuance and rolefs
    function setPI(address _p, address i, address rfs) external onlyOwner {
        pledgePool = _p;
        issuance = i;
        rolefs = rfs;
    }

    // 账户本身调用register()获得序号，或者通过签名信息由其他账户代为调用
    // hash(caller,"role-register"), signed by addr
    // 不需要防止多次利用该签名
    function register(address addr, bytes memory sign) external override returns (uint64) {
        address caller = msg.sender;

        if(caller!=addr){ // 由其他账户代为调用
            bytes32 h = keccak256(abi.encodePacked(caller, "role-register"));
            verifySign(addr, h, sign);
        }

        if(info[addr].index != 0) {
            return info[addr].index;
        } else {
            addrs.push(addr);
            uint64 len = uint64(addrs.length);
            info[addr].index = len;
            return len;
        }
    }

    /**
     * @dev check '_index' param in functions about register
     */
    function checkRI(uint64 _index) private view returns (address) {
        address addr = addrs[_index-1];
        require(info[addr].roleType == 0, "NF"); // not first,this account has been user、keeper or provider
        return addr;
    }

    /// @dev check '_index' param in functions about group
    function checkIG(uint64 _index, uint8 _rType) private view returns (address) {
        address addr = addrs[_index-1];
        require(info[addr].roleType==_rType && !info[addr].isActive && !info[addr].isBanned, "AE"); // account error
        return addr;
    }

    /// @dev check if 'index' indicates user(roleType:1) or provider(roleType:2) or keeper(roleType:3)
    function checkIR(uint64 _index, uint8 _rType) external view override returns (address) {
        address a = addrs[_index-1];
        require(info[a].roleType == _rType, "RE"); // role error
        return a;
    }

    function verifySign(address addr, bytes32 h, bytes memory sign) private pure {
        require(h.recover(sign) == addr, "IC"); // illegal caller
    }

    /// @dev check '_gIndex' param in functions
    function checkG(uint64 _gIndex) private view {
        require(groups[_gIndex-1].isActive && !groups[_gIndex-1].isBanned, "GB"); // group is banned
    }

    /// @dev check whether 'tIndex' is valid
    function checkT(uint32 tIndex) public view override returns (address) {
        RToken rt = RToken(rToken);
        require(rt.isValid(tIndex), "TE"); // payToken error
        return rt.getTA(tIndex);
    }

    // 账户本身调用或由其他账户代为调用
    // hash(caller,_blsKey,"keeper")
    // 不需要防止多次利用该签名
    function registerKeeper(uint64 _index, bytes memory _blsKey, bytes memory sign) external override {
        address caller = msg.sender;
        address addr = checkRI(_index);

        if(caller != addr){ // 由其他账户代为调用，需要验证签名
            bytes32 h = keccak256(abi.encodePacked(caller, _blsKey, "keeper"));
            verifySign(addr, h, sign);
        }

        // 检查质押金额是否满足要求
        uint256 pledgedMoney = IPledgePool(payable(pledgePool)).getBalance(_index,0); // 获得_index账户在主代币上的质押金额
        require(pledgedMoney >= pledgeK, "NE"); // pledged money is not enough
        
        info[addr].roleType = 3;
        info[addr].extra = _blsKey;

        emit RKeeper(_index, addr);
    }

    // 账户本身调用或由其他账户代为调用
    // hash(caller,"provider")
    // 不需要防止多次利用该签名
    function registerProvider(uint64 _index, bytes memory sign) external override {
        address caller = msg.sender;
        address addr = checkRI(_index);

        if(caller != addr){ // 由其他账户代为调用，需要验证签名
            bytes32 h = keccak256(abi.encodePacked(caller, "provider"));
            verifySign(addr, h, sign);
        }

        // 检查质押金额是否满足要求
        uint256 pledgedMoney = IPledgePool(payable(pledgePool)).getBalance(_index, 0);
        require(pledgedMoney >= pledgeP, "NE"); // pledged money is not enough

        info[addr].roleType = 2;

        emit RProvider(_index, addr);
    }

    // 账户本身调用或由其他账户代为调用
    // hash(caller,gIndex,blsKey)
    // 不需要防止多次利用该签名
    function registerUser(uint64 _index, uint64 _gIndex, bytes memory blsKey, bytes memory sign) external override {
        // check params
        address addr = checkRI(_index);
        address caller = msg.sender;

        checkG(_gIndex);

        if(caller != addr){ // 由其他账户代为调用，需要验证签名
            bytes32 h = keccak256(abi.encodePacked(caller, _gIndex, blsKey));
            verifySign(addr, h, sign);
        }

        // fs
        IFileSys(payable(groups[_gIndex-1].fsAddr)).createFs(_index);

        groups[_gIndex-1].users.push(_index);

        info[addr].roleType = 1;
        info[addr].gIndex = _gIndex;
        info[addr].extra = blsKey;
        info[addr].isActive = true;

        emit RUser(_index, addr);
    }

    // called by owner
    // 前提是：PledgePool合约被部署并将地址告知Role合约
    function registerToken(address tAddr) external override onlyOwner {
        // check if exist
        RToken rt = RToken(rToken);
        uint32 tIndex = rt.addT(tAddr);

        // update PledgePool.sol
        IPledgePool(payable(pledgePool)).addToken(tAddr, tIndex);
    }

    // called by owner
    // indexes是keepers的索引
    // 调用该函数时，需要同步部署一个FileSys合约，并将FileSys合约地址赋值给该group（通过调用setGF）
    function createGroup(uint64[] memory indexes, uint16 _level) external override onlyOwner returns (uint64) {
        GroupInfo memory g;
        groups.push(g);
        uint64 _gIndex = uint64(groups.length);

        for(uint8 i = 0; i<indexes.length; i++) {
            address addr = checkIG(indexes[i], 3);
            info[addr].gIndex = _gIndex;
            info[addr].isActive = true;
        }

        groups[_gIndex-1].level = _level;
        groups[_gIndex-1].keepers = indexes;

        if(indexes.length >= uint(_level)) {
            groups[_gIndex-1].isActive = true;
        }
        emit CreateGroup(_gIndex);
        return _gIndex;
    }

    // 给group赋值fsAddr,由admin createGroup之后，部署一个FileSys合约，之后调用setGF函数将FileSys合约地址赋值给Role合约
    function setGF(uint64 _gIndex, address _fsAddr) external onlyOwner {
        groups[_gIndex-1].fsAddr = _fsAddr;
    }

    // called by owner
    function addKeeperToGroup(uint64 _index, uint64 _gIndex) external override onlyOwner {
        require(!groups[_gIndex-1].isBanned, "GB"); // group is banned
        
        address addr = checkIG(_index, 3);

        IFileSys(payable(groups[_gIndex-1].fsAddr)).addKeeper(_index);

        groups[_gIndex-1].keepers.push(_index);
        info[addr].gIndex = _gIndex;
        info[addr].isActive = true;
        if(groups[_gIndex-1].keepers.length >= groups[_gIndex-1].level) {
            groups[_gIndex-1].isActive = true;
        }
    }

    // 由账户自己调用或由其他账户代为调用
    // hash(caller,gIndex)
    // 不需要防止多次利用该签名
    function addProviderToGroup(uint64 _index, uint64 _gIndex, bytes memory sign) external override {
        // check params
        address addr = checkIG(_index, 2);

        checkG(_gIndex);

        address caller = msg.sender;

        if(caller != addr){ // 由其他账户代为调用，需要验证签名
            bytes32 h = keccak256(abi.encodePacked(caller, _gIndex));
            verifySign(addr, h, sign);
        }

        groups[_gIndex-1].providers.push(_index);
        info[addr].gIndex = _gIndex;
        info[addr].isActive = true;
    }

    function setPledgeMoney(uint256 kPledge, uint256 pPledge) external override onlyOwner {
        pledgeK = kPledge;
        pledgeP = pPledge;
    }


    // 可由账户本身调用或由其他账户代为调用
    // 调用前需要uAddr先approve合约地址（FileSys.sol）账户指定金额
    // hash(caller, uIndex, tIndex, money)
    // user往文件系统FileSys合约中充值
    function recharge(uint64 uIndex, uint32 tIndex, uint256 money, bytes memory sign) external payable override {
        // check params
        address addr = addrs[uIndex-1];
        require(info[addr].roleType==1, "NU"); // not user

        address tAddr = checkT(tIndex);
        
        address caller = msg.sender;

        if(caller != addr){ // 由其他账户代为调用，需要验证签名
            bytes32 h = keccak256(abi.encodePacked(caller,uIndex, tIndex, money));
            verifySign(addr, h, sign);
        }

        // TODO：如何防止该签名被重复利用。如果只approve了money金额，那么再次利用该签名将会失败

        IFileSys(payable(groups[info[addr].gIndex-1].fsAddr)).recharge(uIndex, tIndex, addr, tAddr, money); // msg.value有问题
    }

    // 可由账户本身调用或由其他账户代为调用
    // hash(caller, tIndex, amount)
    // user、keeper、foundation从FileSys合约中取回余额(caller为foundation时，即为foundation取余额)
    function withdrawFromFs(uint64 index, uint32 tIndex, uint256 amount, bytes memory sign) external override {
        // check params
        address tAddr = checkT(tIndex); // tIndex error
        require(amount > 0); // should not be zero
        require(index>0);

        address acc = addrs[index-1];
        address fsAddr = groups[info[acc].gIndex-1].fsAddr;

        uint8 rtype = info[acc].roleType;
        address caller = msg.sender;
        if(caller == foundation){
            index = 0;
            rtype = 0;
            acc = foundation;
        }else{
            require(!info[acc].isBanned, "II"); // illegal index
            require(rtype!=2); // 不能是provider
            if(caller != acc){ // 由其他账户代为调用，需要验证签名
                bytes32 h = keccak256(abi.encodePacked(caller, tIndex, amount));
                verifySign(acc, h, sign);
            }
        }

        // TODO：如何防止该签名被重复利用

        IFileSys(payable(fsAddr)).withdraw(index, tIndex, rtype, tAddr, acc, amount);
    }

    /// @dev 'index' indicates index of 'groups' rather than gIndex
    //function setGInfo(uint64 index, bool isAdd, uint256 _size, uint256 sPrice) external override {
    function setGInfo(SGParams memory ps) external {
        require(msg.sender == rolefs, "N");
        if(ps.isAdd){
            groups[ps.index].size += ps._size;
            groups[ps.index].price += ps.sPrice;
        }else{
            groups[ps.index].size -= ps._size; // 隐含 require(groups[index].size >= _size)
            groups[ps.index].price -= ps.sPrice; // 隐含 require(groups[index].price >= sPrice)
        }
    }

    // ===================get===================

    function getAddrsNum() external view returns (uint64) {
        return uint64(addrs.length);
    }

    // 根据数组索引值（不是账户index）获取相应的账户地址，超出数组范围将revert
    function getAddr(uint64 i) external view override returns (address) {
        return addrs[i];
    }

    function getRoleIndex(address acc) external view returns (uint64) {
        return info[acc].index;
    }

    function getRoleInfo(address acc) external view override returns (bool, bool, uint8, uint64, uint64, bytes memory) {
        return (info[acc].isActive, info[acc].isBanned, info[acc].roleType, info[acc].index, info[acc].gIndex, info[acc].extra);
    }

    function getGroupsNum() external view returns (uint64) {
        return uint64(groups.length);
    }

    function getGroupInfo(uint64 i) external view returns (bool, bool, bool, uint16, uint256,uint256, address) {
        return (groups[i].isActive, groups[i].isBanned, groups[i].isReady, groups[i].level, groups[i].size, groups[i].price, groups[i].fsAddr);
    }

    function getFsAddr(uint64 i) external view override returns (address) {
        return groups[i].fsAddr;
    }

    function getAddrGindex(uint64 i) external view override returns (address,uint64) {
        address acc = addrs[i];
        return (acc,info[acc].gIndex);
    }

    function getGKNum(uint64 i) external view override returns (uint64) {
        return uint64(groups[i].keepers.length);
    }

    function getGUPNum(uint64 i) external view returns (uint64, uint64) {
        return (uint64(groups[i].users.length),uint64(groups[i].providers.length));
    }

    // 根据groups的数组索引（不是group-index）、以及keepers的数组索引查找相应的keeper-index
    function getGroupK(uint64 ig, uint64 ik) external view returns (uint64) {
        return groups[ig].keepers[ik];
    }

    function getGroupP(uint64 ig, uint64 ip) external view returns (uint64) {
        return groups[ig].providers[ip];
    }

    function getGU(uint64 ig, uint64 iu) external view returns (uint64) {
        return groups[ig].users[iu];
    }
}