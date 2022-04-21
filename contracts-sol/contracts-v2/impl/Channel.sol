// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IChannel.sol";
import "../interfaces/IERC20.sol";
import "../Recover.sol";
import "../Owner.sol";

contract Channel is IChannel, Owner  {
    mapping(address => mapping(address => uint64)) public nonces; // token => to => nonce
    
    // constructor
    constructor() payable {
    }

    function withdraw(Paycheck memory _ck, bytes memory _sigop, bytes memory _sig) external override {
        require(_ck.pay <= _ck.check.value, "pe");
        require(_ck.check.nonce >= nonces[_ck.check.tokenAddr][msg.sender], "ne");
        
        // verify check's signer
        bytes memory checkPack = 
        abi.encodePacked(
            address(this),
            getOwner(),
            _ck.check.tokenAddr,
            msg.sender,
            _ck.check.fromAddr,
            _ck.check.nonce,
            _ck.check.value
    		);
        bytes32 checkHash = keccak256(checkPack);      
        require(getOwner()== Recover.recover(checkHash, _sigop), "is");
    	
        // verify paycheck's signer
        bytes memory paycheckPack = 
        abi.encodePacked(
            checkPack,
            _ck.pay
        );
        bytes32 paycheckHash = keccak256(paycheckPack);
        require(_ck.check.fromAddr == Recover.recover(paycheckHash,_sig), "is");
        
        // pay
        IERC20(_ck.check.tokenAddr).transfer(msg.sender, _ck.pay);
        
        // update nonce after paid
        nonces[_ck.check.tokenAddr][msg.sender]++;
    }
    
    function withdrawBatch(BatchCheck memory _bc, bytes memory _sigop) external override {
        require(_bc.min == nonces[_bc.tokenAddr][msg.sender], "ne");
        require(_bc.min < _bc.max, "nm");

        // verify check's signer
        bytes memory bcPack = 
        abi.encodePacked(
            address(this),
            getOwner(),
            _bc.tokenAddr,
            msg.sender,
            _bc.min,
            _bc.max,
            _bc.value
    		);
        bytes32 bcHash = keccak256(bcPack);
        address bcSigner = Recover.recover(bcHash, _sigop);
        
        require(getOwner() == bcSigner, "is");
    	
        // pay
        IERC20(_bc.tokenAddr).transfer(msg.sender, _bc.value);
        
        // update nonce after paid
        nonces[_bc.tokenAddr][msg.sender] = _bc.max+1;
    }

    // get nonce of a specified node
    function getNonce(address _t, address _a) external override view returns(uint64) {
        return nonces[_t][_a];
    }
}