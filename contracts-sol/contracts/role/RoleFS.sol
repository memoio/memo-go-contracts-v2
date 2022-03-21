// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/roleIn/IRole.sol";
import "../interfaces/fileSysIn/IFileSys.sol";
import "../interfaces/roleIn/IIssuance.sol";
import "../interfaces/roleIn/IRToken.sol";
import "../interfaces/erc20In/IERC20.sol";
import "../Recover.sol";

/// @dev This contract contains functions involving Role and FileSys contract
contract RoleFS {
    using Recover for bytes32;

    address role;
    address issuance;
    address owner;
    address rToken;
    
    /// @dev created by admin
    constructor() {
        owner = msg.sender;
    }

    /// @dev set issuance-contract、role-contract and RToken-contract address, called by owner
    function setAddr(address iss,address r, address rt) external {
        require(msg.sender == owner, "N");
        issuance = iss;
        role = r;
        rToken = rt;
    }

    function checkParam(address caller, uint64 a, uint64 b, uint32 tIndex, uint8 label) internal view returns (uint64, uint64) {
        IRole r = IRole(role);
        r.checkT(tIndex); // tIndex error

        (bool isActive, bool isBanned, uint8 roleType,uint64 cIndex,uint64 cGIndex,) = r.getRoleInfo(caller);
        if(label==1){ // addOrder, caller is user
            require(cIndex == a, "CE"); // caller error
        }
        if(label==2) { // subOrder, caller is user or keeper
            if(cIndex!=a){
                require(isActive && !isBanned && roleType==3, "CE"); // caller error
            }
        }

        (,uint64 _gIndex) = r.getAddrGindex(a-1);
        (,uint64 pGIndex) = r.getAddrGindex(b-1);
        require(_gIndex == pGIndex && _gIndex == cGIndex, "GD"); // group different
        return (cIndex, _gIndex);
    }

    function verifySign(address addr, bytes32 h, bytes memory sign) private pure {
        require(h.recover(sign) == addr, "IC"); // illegal caller
    }

    // called by user
    // 拥有user、provider的签名信息，以及其他keeper的聚合签名信息，然后调用此函数
    // hash(uIndex, pIndex, nonce, _start, end, _size, tIndex, sPrice)
    // 调用该函数前需要admin调用AccessControl合约的setUpRole函数，为该合约赋予MINTER_ROLE权限
    function addOrder(AOParams memory ps) external {
        
        // acc index should not be 0, acc index start from 1
        assert(ps.uIndex != 0);
        assert(ps.pIndex != 0);

        //// prepair params for verify sig

        // get address from array index(start from 0) 
        uint64 uLoc = ps.uIndex-1;
        uint64 pLoc = ps.pIndex-1;
        address uAddr = IRole(role).getAddr(uLoc);
        address pAddr = IRole(role).getAddr(pLoc);

        // get gIndex from user location in array
        uint64 gIndex;
        (, gIndex) = IRole(role).getAddrGindex(uLoc);

        //// verify signatures
            
        // calc hash
        bytes32 h = keccak256(abi.encodePacked(
            ps.uIndex,
            ps.pIndex,
            ps.nonce,
            ps._start,
            ps._end,
            ps._size, 
            ps.tIndex, 
            ps.sPrice)
            );
        
        // for storing signer
        address signer;

        // verify user sig
        signer  = h.recover(ps.usign);

        assert(signer == uAddr);

        // verify provider sig
        signer = h.recover(ps.psign);

        assert(signer == pAddr);

        // 检查参数大小
        require(ps._size > 0, "sz"); // size zero
        require(ps._end > ps._start, "es"); // end should more than start
        // 与天保持一致
        require((ps._end/86400)*86400 == ps._end, "te"); // end time error

        // 检查uIndex、pIndex角色是否正确
        IRole(role).checkIR(ps.uIndex, 1);
        IRole(role).checkIR(ps.pIndex, 2);

        // 检查uIndex、pIndex、caller的isbanned参数以及他们是否为同一group; tIndex的有效性
        (, uint64 uGIndex) = checkParam(msg.sender, ps.uIndex, ps.pIndex, ps.tIndex, 1);

        IFileSys(payable(IRole(role).getFsAddr(gIndex-1))).addOrder(ps);


        // for calling issu()
        IssuParams memory issuPs;
        issuPs._start = ps._start;
        issuPs._end = ps._end;
        issuPs._size = ps._size;
        issuPs.sPrice = ps.sPrice;
        
        // 主代币增发
        if (ps.tIndex == 0) { // 0 为主代币
            IIssuance iss = IIssuance(issuance);
            uint256 reward = iss.issu(issuPs);
            
            // prepair params for calling setGInfo()
            SGParams memory sgPs;
            sgPs.index = uGIndex-1;
            sgPs.isAdd = true;
            sgPs._size = ps._size;
            sgPs.sPrice = ps.sPrice;

            // update info in group
            IRole(role).setGInfo(sgPs);
   
            // Role合约往PledgePool合约中转入reward数量的tokens[0]代币，即主代币
            // 此处应该为mint，而不是transfer，但是这样的话，需要为RoleFS合约账户赋予mint权限。
            if(reward==0){
                return;
            }
            IERC20 e = IERC20(IRToken(rToken).getTA(0));
            bytes[5] memory tmp;
            e.mintToken(IRole(role).pledgePool(), reward, tmp);
        }
    }

    // called by keeper or user
    // 拥有user、provider的签名信息，以及其他keeper的聚合签名信息，然后调用此函数
    // hash(uIndex, pIndex, nonce, _start, end, _size, tIndex, sPrice)
    function subOrder(SOParams memory ps) external {
        // verify sign
        // acc index should not be 0, acc index start from 1
        assert(ps.uIndex != 0);
        assert(ps.pIndex != 0);

        // get address from array index(start from 0) 
        uint64 uLoc = ps.uIndex-1;
        uint64 pLoc = ps.pIndex-1;
        address uAddr = IRole(role).getAddr(uLoc);
        address pAddr = IRole(role).getAddr(pLoc);

        // get gIndex from user location in array
        uint64 gIndex;
        (, gIndex) = IRole(role).getAddrGindex(uLoc);

        // verify signatures
            
        // calc hash
        bytes32 h = keccak256(abi.encodePacked(
            ps.uIndex,
            ps.pIndex,
            ps.nonce,
            ps._start,
            ps._end,
            ps._size,
            ps.tIndex,
            ps.sPrice)
            );
        
        // for storing signer
        address signer;

        // verify user sig
        signer  = h.recover(ps.usign);

        assert(signer == uAddr);

        // verify provider sig
        signer = h.recover(ps.psign);

        assert(signer == pAddr);
        
        /*
        // verify each keeper sig
        for(uint64 i = 0; i < kNum; i++) {
            // caution: both param of getGroupK is the array location, not gIndex

            // get kIndex from gIndex and location in group
            uint64 kIndex = IRole(role).getGroupK(gIndex-1,i);

            // get addr from loc(loc = index - 1)
            address kAddr = IRole(role).getAddr(kIndex-1);

            // recover signer from sig
            bytes[] memory signs = ps.ksigns; 
            signer = h.recover(signs[i]); 

            assert(signer == kAddr);
        }
        */

        // do some checks
        IRole(role).checkIR(ps.uIndex, 1);
        IRole(role).checkIR(ps.pIndex, 2);

        (uint64 cIndex, uint64 _gIndex) = checkParam(msg.sender, ps.uIndex, ps.pIndex, ps.tIndex, 2);

        require(ps._size > 0, "ES"); // size error
        require(ps._end > ps._start && ps._end <= block.timestamp, "ET"); // time error

        ps.kIndex = cIndex; // add kIndex for fs.subOrder
        IFileSys(payable(IRole(role).getFsAddr(_gIndex-1))).subOrder(ps);


        // prepair params for calling setGInfo()
        SGParams memory sgPs;
        sgPs.index = _gIndex-1;
        sgPs.isAdd = false;
        sgPs._size = ps._size;
        sgPs.sPrice = ps.sPrice;
    
        if(ps.tIndex==0){
            IRole(role).setGInfo(sgPs);
        }
        
    }

    // called by keeper
    // hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "a")
    // signed by newProvider、at least two thirds of the keeper
    function addRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sprice, bytes memory pSign, bytes[] memory kSigns) external {
        // check params
        IRole(role).checkIR(pIndex, 2);
        IRole(role).checkIR(nPIndex, 2);

        // check params
        require(_size>0 && end > _start, "IP"); // illegal params

        (uint64 cIndex, uint64 _gIndex) = checkParam(msg.sender, pIndex, nPIndex, tIndex, 2);

        {
            // verify pSign
            bytes32 h = keccak256(abi.encodePacked(pIndex, _start, end, _size, nonce, tIndex, sprice, "a"));
            verifySign(IRole(role).getAddr(nPIndex-1), h, pSign);

            // verify kSigns
            // 需要在合约中再判断签名信息是否有相同的吗？不需要，需要在线下确保签名不重复。
            // 因为在合约里判断代价大，无法新建动态大小的数组，使用mapping也必须是storage存储。
            uint256 kNum = IRole(role).getGKNum(_gIndex-1) * 2 / 3;
            require(kSigns.length >= kNum, "SNE"); // kSigns not enough
            uint64 account;
            address s;
            for(uint256 i=0;i<kSigns.length;i++){
                s = h.recover(kSigns[i]);
                (,,,,uint64 index,) = IRole(role).getRoleInfo(s);
                if(index == _gIndex){
                    account++;
                }
            }
            require(account >= kNum, "SNE"); // kSigns not enough
        }

        IFileSys(payable(IRole(role).getFsAddr(_gIndex-1))).addRepair(cIndex, pIndex, nPIndex, _start, end, _size, nonce, tIndex, sprice);

        // modify size

        // 增发？
        if(tIndex == 0) {
            // sub due to lost
            IIssuance iss = IIssuance(issuance);
            iss.setTP(0, sprice * uint256(end - _start));
        }
    }

    // keeper 调用
    // hash(pIndex, _start, end, _size, nonce, tIndex, sPrice, "s")
    // signed by newProvider、at least two thirds of the keeper
    function subRepair(uint64 pIndex, uint64 nPIndex, uint64 _start, uint64 end, uint64 _size, uint64 nonce, uint32 tIndex, uint256 sPrice, bytes memory pSign, bytes[] memory kSigns) external {
        // check params
        IRole(role).checkIR(pIndex, 2);
        IRole(role).checkIR(nPIndex, 2);

        // check params
        require(_size >0 && end > _start && end <= block.timestamp, "IP"); // illegal param

        (uint64 cIndex, uint64 _gIndex) = checkParam(msg.sender, pIndex, nPIndex, tIndex, 2);

        {
            // verify pSign
            bytes32 h = keccak256(abi.encodePacked(pIndex, _start, end, _size, nonce, tIndex, sPrice, "s"));
            verifySign(IRole(role).getAddr(nPIndex-1), h, pSign);

            // verify kSigns
            // 需要在合约中再判断签名信息是否有相同的吗？不需要，需要在线下确保签名不重复。
            // 因为在合约里判断代价大，无法新建动态大小的数组，使用mapping也必须是storage存储。
            uint256 kNum = IRole(role).getGKNum(_gIndex-1) * 2 / 3;
            require(kSigns.length >= kNum, "SNE"); // kSigns not enough

            uint64 account;
            address t;
            for(uint256 i=0;i<kSigns.length;i++){
                t = h.recover(kSigns[i]);
                (,,,,uint64 index,) = IRole(role).getRoleInfo(t);
                if(index == _gIndex){
                    account++;
                }
            }
            require(account >= kNum, "SNE"); // kSigns not enough
        }

        IFileSys(payable(IRole(role).getFsAddr(_gIndex-1))).subRepair(cIndex, nPIndex, _start, end, _size, nonce, tIndex, sPrice);

        // modify size
    }

    // hash（ps.pIndex,ps.tIndex,ps.pay,ps.lost）
    // provider从FileSys合约中取回余额
    function proWithdraw(PWParams memory ps) external {

        //// prepair params for verify sig
        
        // acc index should not be 0, acc index start from 1
        assert(ps.pIndex != 0);

        //// prepair params for verify sig

        // get address from array index(start from 0) 
        uint64 pLoc = ps.pIndex-1;

        // get provider address and gIndex from provider location in array
        address proAddr;
        uint64 gIndex;
        (proAddr, gIndex) = IRole(role).getAddrGindex(pLoc);

        //// verify each keeper sig

        // calc hash
        bytes32 h = keccak256(abi.encodePacked(ps.pIndex,ps.tIndex,ps.pay,ps.lost));

        // for storing signer
        address signer;

        uint256 indexNum = ps.kIndexes.length;
        uint256 sigNum = ps.ksigns.length;
        assert(indexNum == sigNum);

        for(uint64 i = 0; i < indexNum; i++) {
            // caution: both param of getGroupK is the array location, not gIndex

            // get addr from loc(loc = kIndex - 1)
            address kAddr = IRole(role).getAddr(ps.kIndexes[i]-1);

            // recover signer from sig
            bytes[] memory signs = ps.ksigns; 
            signer = h.recover(signs[i]); 

            assert(signer == kAddr);           
        }

        // after verify sig 

        // check params
        {
            // check pIndex
            (bool isActive, bool isBanned, uint8 roleType,,,) = IRole(role).getRoleInfo(proAddr);
            require(isActive && !isBanned && roleType==2, "IP"); // illegal pIndex

            // sig num should not less than 2*(N+1)/3, N: kNum of group
            uint64 minKNum = 2 * (IRole(role).getGKNum(gIndex-1) + 1) / 3;
            require(sigNum >= minKNum, "SE"); // kSigns error

            // sig num should not less than group level
            uint16 level; // level is uint16 in contract
            (,,,level,,,) = IRole(role).getGroupInfo(gIndex-1);
            require(sigNum >= level, "SE");
        }


        // call fs.proWithdraw
        address fsAddr = IRole(role).getFsAddr(gIndex-1);
        // ps.pAddr = proAddr;
        // ps.tAddr = IRole(role).checkT(ps.tIndex);
        IFileSys(payable(fsAddr)).proWithdraw(ps);

        // call issuance.setTP
        uint256 kNum;   // keeper num of group
        if(ps.tIndex==0){
            (,,,,,,kNum,,,,) = IFileSys(payable(fsAddr)).getSettleInfo(ps.pIndex,ps.tIndex);
            // add lost to paid
            IIssuance(issuance).setTP(ps.lost, kNum);
        }
        
    }
}