// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IAuth.sol";
import "../Recover.sol";

// hash(contract address, nonce, method name, params)

contract Auth is IAuth {
    using Recover for bytes32;

    uint16 public version = 2;
    
    address[] public controls; // five address
    uint256 public nonce;
    
    mapping(address => mapping(address => bool)) private access;

    // how to contruct?
    constructor(address[] memory _addrs) {
        controls = _addrs;
    }

    function set(address _contract, address account, bool isSet, bytes[] memory signs) external override {
        uint size;
        assembly {
            size := extcodesize(account)
        }
        require(size != 0, "NE"); // need ext addr

        bytes32 h = keccak256(abi.encodePacked(address(this), nonce, "set", _contract, account, isSet));
        uint8 valid = 0;
        for(uint256 i=0;i<signs.length;i++){
            if(h.recover(signs[i])==controls[i]){
                valid++;
            }
        }
        require(valid>=3,"SNE"); // sign not enough

        access[_contract][account] = isSet;
        nonce+=1;
    }

    function perm(bytes32 ha, bytes[] memory signs) external override {
        bytes32 h = keccak256(abi.encodePacked(address(this), nonce, "perm", ha));
        uint8 valid = 0;
        for(uint256 i=0;i<signs.length;i++){
            if(h.recover(signs[i])==controls[i]){
                valid++;
            }
        }
        require(valid>=3,"SNE"); // sign not enough
        nonce+=1;
    }

    function can(address _contract, address account) external view override returns (bool) {
        return access[_contract][account];
    }
}