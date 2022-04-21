pragma solidity ^0.8.0;

struct Check {
    address tokenAddr;      // token address, point out which token to pay
    address fromAddr;       // buyer of this check, should be check's signer
    uint64  nonce;          // nonce of the check, check's nonce should not smaller than it.
    uint256 value;          // value of the check, payvalue shoud not exceed value
}

struct Paycheck {
	Check check;
    uint256 pay;       // money to pay, should not exceed value
}

struct BatchCheck  {
    address tokenAddr;   
    uint64 min;        
	uint64 max;
	uint256 value;
}

interface IChannel {
    // sigop hash(contract, op, token, to, from, nonce, value)
    // sig hash(contract, op, token, to, from, nonce, value, payVaue)
    function withdraw(Paycheck memory _ck, bytes memory _sigop, bytes memory _sig) external;
    // sig hash(contract, op, token, to, min, max, value)
    function withdrawBatch(BatchCheck memory _bc, bytes memory _sig) external;

    function getNonce(address _t, address _a) external view returns(uint64);
}