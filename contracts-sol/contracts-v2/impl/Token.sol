// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IToken.sol";
import "../interfaces/IAuth.sol";
import "./Owner.sol";

// base

/// @dev This contract is about the token addresses that memo supported.
contract Token is ITokenSetter, ITokenGetter, Owner {
    struct TokenInfo {
        uint32 index;  // 代币序列号
        bool isBanned; // 该币种是否被禁止
    }

    uint16 public version = 2;

    // 代币信息,0为主代币的代币地址
    address[] tokens;
    mapping(address => TokenInfo) tInfo;

    event AddT(address t, uint32 tIndex);

    constructor(address _rfs, address _a, address _t) Owner(_rfs, _a) {
        instances[1] = _rfs;
        instances[2] = _a;
        tokens.push(_t);
        tInfo[_t].index = 0;
    }

    function addT(address t) external override onlyOwner returns(uint32) {
        // check if exist
        require(!_hasToken(t), "TE"); // token exist 

        // update storage
        uint32 tIndex = uint32(tokens.length);
        tokens.push(t);
        tInfo[t].index = tIndex;

        emit AddT(t, tIndex);
        return tIndex;
    }

    // need banT 
    function banT(address tAddr) external override onlyOwner {
        tInfo[tAddr].isBanned = true;
    }

    function _hasToken(address tAddr) internal view returns (bool) {
        if (tokens.length == 0) {
            return false;
        } 

        if (tInfo[tAddr].index > 0) {
            return true;
        }

        if (tokens[0] == tAddr) {
            return true;
        }

        return false;
    }

    // =========get==========
    function getTA(uint32 tIndex) external view override returns (address, bool) {
        if(tIndex < tokens.length){
            return (tokens[tIndex],tInfo[tokens[tIndex]].isBanned) ;
        }
        return (address(0), false);
    }

    // 返回tIndex以及该代币地址是否有效
    function getTI(address t) external view override returns (uint32, bool) {
        uint32 tIndex = tInfo[t].index;
        if(t==tokens[tIndex]){ // 注册过
            return (tIndex, tInfo[t].isBanned);
        }
        return (tIndex, false);
    }

    function getTNum() external view returns (uint32) {
        return uint32(tokens.length);
    }
}