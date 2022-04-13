//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IAuth.sol";
import "../interfaces/IERC20.sol";
import "../interfaces/IPool.sol";
import "../Recover.sol";
import "./Owner.sol";

contract Pool is IPoolSetter, Owner {
    using Recover for bytes32;

    uint16 public version = 2;

    event Inflow(address indexed from, uint256 money);
    event Outflow(address indexed to, uint256 money);

    constructor(address _rfs, address _a) Owner(_rfs, _a) {
        instances[1] = _rfs;
        instances[2] = _a;
    }

    receive() external payable {}

    // onlyOwner not need
    function inflow(address tAddr, address from, uint256 money) external override  payable {
        IERC20(tAddr).transferFrom(from, address(this), money);
        emit Inflow(from, money);
    }

    function outflow(address tAddr, address to, uint256 money) external onlyOwner override payable {
        IERC20(tAddr).transfer(to, money);
        emit Outflow(to, money);
    }
}