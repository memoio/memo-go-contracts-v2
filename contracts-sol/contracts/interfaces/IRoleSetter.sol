// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IRoleSetter {
    // called by admin
    function ban(uint64 _i, bool _ban) external;

    function reAcc(address _a) external; 
    function reRole(uint64 _i, uint8 _rtype, bytes memory extra) external;
    function setG(uint64 _i, uint64 _gi) external;
    function activate(uint64 _i) external returns (uint64);
}