// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IGroupSetter {
    // called by admin
    function ban(uint64 _gi, bool _isBan) external;

    function activate(uint64 _gi, uint16 _kc) external;
    function create(uint16 _level, uint8 _mr, uint256 _kr, uint256 _pr) external;
}