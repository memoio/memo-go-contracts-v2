pragma solidity ^0.8.0;

import "../interfaces/IProxy.sol";
import "../interfaces/IToken.sol";
import "../interfaces/IToken.sol";

/// @dev This contract contains functions controlling Role, Pledge, FileSys
contract Getter is IGetter {
    uint16 public version = 2;
    address public owner;
    constructor(address _proxy ){
        owner = _proxy;
    }

    function getTA(uint8 _ti) external view override returns (address, bool) {
        return ITokenGetter(IProxy(owner).get(7)).getTA(_ti);
    }

    function getTI(address _t) external view override returns (uint8, bool) {
        return ITokenGetter(IProxy(owner).get(7)).getTI(_t);
    }
}