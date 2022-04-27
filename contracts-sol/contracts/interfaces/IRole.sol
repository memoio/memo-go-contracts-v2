// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IRoleSetter.sol";
import "./IRoleGetter.sol";

interface IRole is IRoleSetter,IRoleGetter {}