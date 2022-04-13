// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// for manage other contracts
interface IAuth {
    // change owner, control etc...
    function perm(bytes32 h, bytes[] memory signs) external;
    function set(address con, address account, bool isSet, bytes[] memory signs) external;
    function can(address con,address account) external view returns (bool);
}