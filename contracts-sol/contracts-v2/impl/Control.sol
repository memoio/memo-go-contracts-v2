// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IControl.sol";
import "../interfaces/IRole.sol";
import "../interfaces/IFileSys.sol";
import "../interfaces/IIssuance.sol";
import "../interfaces/IToken.sol";
import "../interfaces/IERC20.sol";
import "../interfaces/IAuth.sol";
import "../interfaces/IPledge.sol";
import "./Owner.sol";
import "../Recover.sol";

/// @dev This contract contains functions controlling Role, Pledge, FileSys
contract Control is IControl, Owner {
    using Recover for bytes32;

    uint16 public version = 2;

    constructor(address _a) Owner(msg.sender, _a) {

    }

    receive() external payable {}

    function activate(uint64 _index, bool _active, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "activate", _index, _active));
        ia.perm(h, signs);

        IRoleSetter(instances[6]).activate(_index, _active);
    }

    function ban(uint64 _index, bool _banned, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "ban", _index, _banned));
        ia.perm(h, signs);

        IRoleSetter(instances[6]).ban(_index, _banned);
    } 

    function addT(address _t, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "addT", _t));
        ia.perm(h, signs);

        uint32 ti = ITokenSetter(instances[7]).addT(_t);
        IPledgeSetter(instances[7]).addT(ti);
    }

    function createGroup(uint16 _level, address fsAddr, uint64[] memory indexes, bytes[] memory signs) external override {
        IAuth ia = IAuth(instances[2]);
        bytes32 h = keccak256(abi.encodePacked(address(this), "createGroup", _level, fsAddr, indexes));
        ia.perm(h, signs);

        IRoleSetter r = IRoleSetter(instances[6]);
        return r.createGroup(_level,fsAddr, indexes);
    }

    // register self to get index
    function registerAccount() external override {
        IRoleSetter r = IRoleSetter(instances[6]);
        return r.registerAccount(msg.sender);
    } 

    // register as user/keeper/provider
    function registerRole(uint64 _index, uint8 _rtype, bytes memory extra) external override {
        address a = IRoleGetter(instances[6]).getAddr(_index);
        require(a==msg.sender, "IC");// should called self
        return IRoleSetter(instances[6]).registerRole(_index, _rtype, extra);
    }
    
    function addToGroup(uint64 _index, uint64 _gIndex) external override {
        address a = IRoleGetter(instances[6]).getAddr(_index);
        require(a==msg.sender, "IC");// should called self
        return IRoleSetter(instances[6]).addToGroup(_index, _gIndex);
    }

    
}