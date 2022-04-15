// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IProxy.sol";
import "../interfaces/IControl.sol";
import "./Owner.sol";

/// @dev This contract contains functions controlling Role, Pledge, FileSys
contract Proxy is IProxy, Owner {
    uint16 public version = 2;

    constructor(address _o,address _a) Owner(_o, _a) {
    }

    receive() external payable {}

    function activate(uint64 _i, bool _active, bytes[] memory signs) external override {
        IControl(instances[10]).activate(_i, _active, signs);
    }

    function ban(uint64 _i, bool _ban, bytes[] memory signs) external override {
        IControl(instances[10]).ban(_i, _ban, signs);
    }

    function addT(address _t, bytes[] memory signs) external override {
        IControl(instances[10]).addT(_t, signs);
    }

    function createGroup(uint16 _level, address _fs, uint256 _kr, uint256 _pr) external override {
        IControl(instances[10]).createGroup(_level, _fs, _kr, _pr);
    }

    function registerAccount() external override {
        IControl(instances[10]).registerAccount(msg.sender);
    }
    
    function registerRole(uint64 _i, uint8 _rtype, bytes memory _extra) external override {
        IControl(instances[10]).registerRole(msg.sender, _i, _rtype, _extra);
    }

    function addToGroup(uint64 _i, uint64 _gi) external override {
        IControl(instances[10]).addToGroup(msg.sender, _i, _gi);
    }
    
    function pledge(uint64 _i, uint256 _money) external override {
        IControl(instances[10]).pledge(msg.sender, _i, _money);
    }

    function unpledge(uint64 _i, uint8 _ti, uint256 _money) external override {
        IControl(instances[10]).unpledge(msg.sender, _i, _ti, _money);
    }

    function addOrder( OrderIn memory _oi) external override {
        IControl(instances[10]).addOrder(msg.sender, _oi);
    }

    function subOrder(OrderIn memory _oi) external override {
        IControl(instances[10]).subOrder(msg.sender, _oi);
    }

    function recharge(uint64 _i, uint8 _ti, uint256 _money) external override {
        IControl(instances[10]).recharge(msg.sender, _i, _ti, _money);
    }

    function withdraw(uint64 _i, uint8 _ti, uint256 _money) external override {
        IControl(instances[10]).withdraw(msg.sender, _i, _ti, _money);
    }

    function proWithdraw(PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external override {
        IControl(instances[10]).proWithdraw(msg.sender, _ps, _kis, ksigns);
    }

    function get(uint8 _type) external view override returns(address) {
        return IControl(instances[10]).get(_type);
    }
}