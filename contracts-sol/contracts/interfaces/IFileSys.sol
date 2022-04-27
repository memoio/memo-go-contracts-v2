// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IFileSysSetter.sol";
import "./IFileSysGetter.sol";

interface IFileSys is IFileSysSetter, IFileSysGetter {}