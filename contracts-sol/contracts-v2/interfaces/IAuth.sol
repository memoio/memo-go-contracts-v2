// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// for manage other contracts
interface IAuth {
    // change owner, control etc...
    function perm(bytes32 _h, bytes[] memory signs) external;
    function set(address _c, address _a, bool _set, bytes[] memory signs) external;
    function can(address _c,address _a) external view returns (bool);
}