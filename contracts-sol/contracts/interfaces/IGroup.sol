// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IGroupGetter.sol";
import "./IGroupSetter.sol";

interface IGroup is IGroupSetter,IGroupGetter {}