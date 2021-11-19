// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addr\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AirDrop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"airDrop\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acc\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"renounceRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setUpRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002ad138038062002ad1833981810160405281019062000037919062000439565b60016000808060001b815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000807f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060016000807f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508160059080519060200190620001ca929190620001ec565b508060069080519060200190620001e3929190620001ec565b50505062000523565b828054620001fa90620004ed565b90600052602060002090601f0160209004810192826200021e57600085556200026a565b82601f106200023957805160ff19168380011785556200026a565b828001600101855582156200026a579182015b82811115620002695782518255916020019190600101906200024c565b5b5090506200027991906200027d565b5090565b5b80821115620002985760008160009055506001016200027e565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200030582620002ba565b810181811067ffffffffffffffff82111715620003275762000326620002cb565b5b80604052505050565b60006200033c6200029c565b90506200034a8282620002fa565b919050565b600067ffffffffffffffff8211156200036d576200036c620002cb565b5b6200037882620002ba565b9050602081019050919050565b60005b83811015620003a557808201518184015260208101905062000388565b83811115620003b5576000848401525b50505050565b6000620003d2620003cc846200034f565b62000330565b905082815260208101848484011115620003f157620003f0620002b5565b5b620003fe84828562000385565b509392505050565b600082601f8301126200041e576200041d620002b0565b5b815162000430848260208601620003bb565b91505092915050565b60008060408385031215620004535762000452620002a6565b5b600083015167ffffffffffffffff811115620004745762000473620002ab565b5b620004828582860162000406565b925050602083015167ffffffffffffffff811115620004a657620004a5620002ab565b5b620004b48582860162000406565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200050657607f821691505b602082108114156200051d576200051c620004be565b5b50919050565b61259e80620005336000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c80638bb9c5bf116100c3578063d53913931161007c578063d53913931461043f578063d547741f1461045d578063dd62ed3e1461048d578063e63ab1e9146104bd578063f0141d84146104db578063fd1fc4a0146104f957610158565b80638bb9c5bf1461034357806391d1485414610373578063a217fddf146103a3578063a457c2d7146103c1578063a9059cbb146103f1578063c4e41b221461042157610158565b80633f4ba83a116101155780633f4ba83a1461025957806342966c68146102775780636805b84b146102a757806370a08231146102c557806379c65068146102f55780638456cb591461032557610158565b8063095ea7b31461015d578063150704011461018d57806317d7de7c146101ab57806323b872dd146101c95780632d958471146101f95780633950935114610229575b600080fd5b61017760048036038101906101729190611838565b610529565b6040516101849190611893565b60405180910390f35b610195610540565b6040516101a29190611947565b60405180910390f35b6101b36105d2565b6040516101c09190611947565b60405180910390f35b6101e360048036038101906101de9190611969565b610664565b6040516101f09190611893565b60405180910390f35b610213600480360381019061020e91906119f2565b61074d565b6040516102209190611893565b60405180910390f35b610243600480360381019061023e9190611838565b61080d565b6040516102509190611893565b60405180910390f35b6102616108ab565b60405161026e9190611893565b60405180910390f35b610291600480360381019061028c9190611a32565b610938565b60405161029e9190611893565b60405180910390f35b6102af610adc565b6040516102bc9190611893565b60405180910390f35b6102df60048036038101906102da9190611a5f565b610af3565b6040516102ec9190611a9b565b60405180910390f35b61030f600480360381019061030a9190611838565b610b3c565b60405161031c9190611893565b60405180910390f35b61032d610cf6565b60405161033a9190611893565b60405180910390f35b61035d60048036038101906103589190611ab6565b610d82565b60405161036a9190611893565b60405180910390f35b61038d600480360381019061038891906119f2565b610df4565b60405161039a9190611893565b60405180910390f35b6103ab610e5b565b6040516103b89190611af2565b60405180910390f35b6103db60048036038101906103d69190611838565b610e62565b6040516103e89190611893565b60405180910390f35b61040b60048036038101906104069190611838565b610f3f565b6040516104189190611893565b60405180910390f35b610429610f56565b6040516104369190611a9b565b60405180910390f35b610447610f60565b6040516104549190611af2565b60405180910390f35b610477600480360381019061047291906119f2565b610f84565b6040516104849190611893565b60405180910390f35b6104a760048036038101906104a29190611b0d565b611044565b6040516104b49190611a9b565b60405180910390f35b6104c56110cb565b6040516104d29190611af2565b60405180910390f35b6104e36110ef565b6040516104f09190611b69565b60405180910390f35b610513600480360381019061050e9190611ccc565b6110f8565b6040516105209190611893565b60405180910390f35b600061053633848461130f565b6001905092915050565b60606006805461054f90611d57565b80601f016020809104026020016040519081016040528092919081815260200182805461057b90611d57565b80156105c85780601f1061059d576101008083540402835291602001916105c8565b820191906000526020600020905b8154815290600101906020018083116105ab57829003601f168201915b5050505050905090565b6060600580546105e190611d57565b80601f016020809104026020016040519081016040528092919081815260200182805461060d90611d57565b801561065a5780601f1061062f5761010080835404028352916020019161065a565b820191906000526020600020905b81548152906001019060200180831161063d57829003601f168201915b5050505050905090565b600080600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610729576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072090611dd5565b60405180910390fd5b6107348585856114da565b610741853385840361130f565b60019150509392505050565b600061075c6000801b33610df4565b61079b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079290611e41565b60405180910390fd5b600160008085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905092915050565b60006108a1338484600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461089c9190611e90565b61130f565b6001905092915050565b60006108d77f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a33610df4565b610916576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090d90611f32565b60405180910390fd5b6000600160006101000a81548160ff0219169083151502179055506001905090565b60006109476000801b33610df4565b610986576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097d90611f9e565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610a0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a049061200a565b60405180910390fd5b828103600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508260046000828254610a65919061202a565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610aca9190611a9b565b60405180910390a36001915050919050565b6000600160009054906101000a900460ff16905090565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000610b687f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633610df4565b610ba7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9e906120aa565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0e90612116565b60405180910390fd5b8160046000828254610c299190611e90565b9250508190555081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c7f9190611e90565b925050819055508273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ce49190611a9b565b60405180910390a36001905092915050565b6000610d227f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a33610df4565b610d61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5890612182565b60405180910390fd5b60018060006101000a81548160ff0219169083151502179055506001905090565b60008060008084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b600080600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b600080600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610f27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1e906121ee565b60405180910390fd5b610f34338585840361130f565b600191505092915050565b6000610f4c3384846114da565b6001905092915050565b6000600454905090565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6000610f936000801b33610df4565b610fd2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fc990611e41565b60405180910390fd5b600080600085815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905092915050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b60006012905090565b60006111076000801b33610df4565b611146576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113d9061225a565b60405180910390fd5b825182611153919061227a565b600460008282546111649190611e90565b9250508190555060005b835181101561130457600073ffffffffffffffffffffffffffffffffffffffff168482815181106111a2576111a16122d4565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff161415611201576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f89061234f565b60405180910390fd5b8260026000868481518110611219576112186122d4565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461126a9190611e90565b92505081905550838181518110611284576112836122d4565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516112e99190611a9b565b60405180910390a380806112fc9061236f565b91505061116e565b506001905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561137f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161137690612404565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156113ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e690612470565b60405180910390fd5b80600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516114cd9190611a9b565b60405180910390a3505050565b6114e2610adc565b15611522576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611519906124dc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611592576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161158990612470565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611602576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115f990612548565b60405180910390fd5b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611689576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168090611dd5565b60405180910390fd5b818103600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461171e9190611e90565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516117829190611a9b565b60405180910390a350505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006117cf826117a4565b9050919050565b6117df816117c4565b81146117ea57600080fd5b50565b6000813590506117fc816117d6565b92915050565b6000819050919050565b61181581611802565b811461182057600080fd5b50565b6000813590506118328161180c565b92915050565b6000806040838503121561184f5761184e61179a565b5b600061185d858286016117ed565b925050602061186e85828601611823565b9150509250929050565b60008115159050919050565b61188d81611878565b82525050565b60006020820190506118a86000830184611884565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156118e85780820151818401526020810190506118cd565b838111156118f7576000848401525b50505050565b6000601f19601f8301169050919050565b6000611919826118ae565b61192381856118b9565b93506119338185602086016118ca565b61193c816118fd565b840191505092915050565b60006020820190508181036000830152611961818461190e565b905092915050565b6000806000606084860312156119825761198161179a565b5b6000611990868287016117ed565b93505060206119a1868287016117ed565b92505060406119b286828701611823565b9150509250925092565b6000819050919050565b6119cf816119bc565b81146119da57600080fd5b50565b6000813590506119ec816119c6565b92915050565b60008060408385031215611a0957611a0861179a565b5b6000611a17858286016119dd565b9250506020611a28858286016117ed565b9150509250929050565b600060208284031215611a4857611a4761179a565b5b6000611a5684828501611823565b91505092915050565b600060208284031215611a7557611a7461179a565b5b6000611a83848285016117ed565b91505092915050565b611a9581611802565b82525050565b6000602082019050611ab06000830184611a8c565b92915050565b600060208284031215611acc57611acb61179a565b5b6000611ada848285016119dd565b91505092915050565b611aec816119bc565b82525050565b6000602082019050611b076000830184611ae3565b92915050565b60008060408385031215611b2457611b2361179a565b5b6000611b32858286016117ed565b9250506020611b43858286016117ed565b9150509250929050565b600060ff82169050919050565b611b6381611b4d565b82525050565b6000602082019050611b7e6000830184611b5a565b92915050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611bc1826118fd565b810181811067ffffffffffffffff82111715611be057611bdf611b89565b5b80604052505050565b6000611bf3611790565b9050611bff8282611bb8565b919050565b600067ffffffffffffffff821115611c1f57611c1e611b89565b5b602082029050602081019050919050565b600080fd5b6000611c48611c4384611c04565b611be9565b90508083825260208201905060208402830185811115611c6b57611c6a611c30565b5b835b81811015611c945780611c8088826117ed565b845260208401935050602081019050611c6d565b5050509392505050565b600082601f830112611cb357611cb2611b84565b5b8135611cc3848260208601611c35565b91505092915050565b60008060408385031215611ce357611ce261179a565b5b600083013567ffffffffffffffff811115611d0157611d0061179f565b5b611d0d85828601611c9e565b9250506020611d1e85828601611823565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611d6f57607f821691505b60208210811415611d8357611d82611d28565b5b50919050565b7f4145420000000000000000000000000000000000000000000000000000000000600082015250565b6000611dbf6003836118b9565b9150611dca82611d89565b602082019050919050565b60006020820190508181036000830152611dee81611db2565b9050919050565b7f4e41520000000000000000000000000000000000000000000000000000000000600082015250565b6000611e2b6003836118b9565b9150611e3682611df5565b602082019050919050565b60006020820190508181036000830152611e5a81611e1e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611e9b82611802565b9150611ea683611802565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611edb57611eda611e61565b5b828201905092915050565b7f20434e5550000000000000000000000000000000000000000000000000000000600082015250565b6000611f1c6005836118b9565b9150611f2782611ee6565b602082019050919050565b60006020820190508181036000830152611f4b81611f0f565b9050919050565b7f434e420000000000000000000000000000000000000000000000000000000000600082015250565b6000611f886003836118b9565b9150611f9382611f52565b602082019050919050565b60006020820190508181036000830152611fb781611f7b565b9050919050565b7f4241454200000000000000000000000000000000000000000000000000000000600082015250565b6000611ff46004836118b9565b9150611fff82611fbe565b602082019050919050565b6000602082019050818103600083015261202381611fe7565b9050919050565b600061203582611802565b915061204083611802565b92508282101561205357612052611e61565b5b828203905092915050565b7f434e4d0000000000000000000000000000000000000000000000000000000000600082015250565b60006120946003836118b9565b915061209f8261205e565b602082019050919050565b600060208201905081810360008301526120c381612087565b9050919050565b7f4d5a000000000000000000000000000000000000000000000000000000000000600082015250565b60006121006002836118b9565b915061210b826120ca565b602082019050919050565b6000602082019050818103600083015261212f816120f3565b9050919050565b7f434e500000000000000000000000000000000000000000000000000000000000600082015250565b600061216c6003836118b9565b915061217782612136565b602082019050919050565b6000602082019050818103600083015261219b8161215f565b9050919050565b7f41425a0000000000000000000000000000000000000000000000000000000000600082015250565b60006121d86003836118b9565b91506121e3826121a2565b602082019050919050565b60006020820190508181036000830152612207816121cb565b9050919050565b7f434e414400000000000000000000000000000000000000000000000000000000600082015250565b60006122446004836118b9565b915061224f8261220e565b602082019050919050565b6000602082019050818103600083015261227381612237565b9050919050565b600061228582611802565b915061229083611802565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156122c9576122c8611e61565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f41445a0000000000000000000000000000000000000000000000000000000000600082015250565b60006123396003836118b9565b915061234482612303565b602082019050919050565b600060208201905081810360008301526123688161232c565b9050919050565b600061237a82611802565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156123ad576123ac611e61565b5b600182019050919050565b7f494f000000000000000000000000000000000000000000000000000000000000600082015250565b60006123ee6002836118b9565b91506123f9826123b8565b602082019050919050565b6000602082019050818103600083015261241d816123e1565b9050919050565b7f4953000000000000000000000000000000000000000000000000000000000000600082015250565b600061245a6002836118b9565b915061246582612424565b602082019050919050565b600060208201905081810360008301526124898161244d565b9050919050565b7f5446000000000000000000000000000000000000000000000000000000000000600082015250565b60006124c66002836118b9565b91506124d182612490565b602082019050919050565b600060208201905081810360008301526124f5816124b9565b9050919050565b7f4952000000000000000000000000000000000000000000000000000000000000600082015250565b60006125326002836118b9565b915061253d826124fc565b602082019050919050565b6000602082019050818103600083015261256181612525565b905091905056fea2646970667358221220f1d3d9699f7853d7377baa393582e68238cae7a05aeb35be182b04c249446e2864736f6c634300080a0033",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MetaData.Bin instead.
var ERC20Bin = ERC20MetaData.Bin

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20Bin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20 *ERC20Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20 *ERC20Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20.Contract.DEFAULTADMINROLE(&_ERC20.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ERC20 *ERC20CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ERC20.Contract.DEFAULTADMINROLE(&_ERC20.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20 *ERC20Caller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20 *ERC20Session) MINTERROLE() ([32]byte, error) {
	return _ERC20.Contract.MINTERROLE(&_ERC20.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ERC20 *ERC20CallerSession) MINTERROLE() ([32]byte, error) {
	return _ERC20.Contract.MINTERROLE(&_ERC20.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20 *ERC20Caller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20 *ERC20Session) PAUSERROLE() ([32]byte, error) {
	return _ERC20.Contract.PAUSERROLE(&_ERC20.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ERC20 *ERC20CallerSession) PAUSERROLE() ([32]byte, error) {
	return _ERC20.Contract.PAUSERROLE(&_ERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address acc) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, acc common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", acc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address acc) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(acc common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, acc)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address acc) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(acc common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, acc)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_ERC20 *ERC20Caller) GetDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_ERC20 *ERC20Session) GetDecimals() (uint8, error) {
	return _ERC20.Contract.GetDecimals(&_ERC20.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xf0141d84.
//
// Solidity: function getDecimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) GetDecimals() (uint8, error) {
	return _ERC20.Contract.GetDecimals(&_ERC20.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ERC20 *ERC20Caller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ERC20 *ERC20Session) GetName() (string, error) {
	return _ERC20.Contract.GetName(&_ERC20.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_ERC20 *ERC20CallerSession) GetName() (string, error) {
	return _ERC20.Contract.GetName(&_ERC20.CallOpts)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_ERC20 *ERC20Caller) GetPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_ERC20 *ERC20Session) GetPaused() (bool, error) {
	return _ERC20.Contract.GetPaused(&_ERC20.CallOpts)
}

// GetPaused is a free data retrieval call binding the contract method 0x6805b84b.
//
// Solidity: function getPaused() view returns(bool)
func (_ERC20 *ERC20CallerSession) GetPaused() (bool, error) {
	return _ERC20.Contract.GetPaused(&_ERC20.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ERC20 *ERC20Caller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ERC20 *ERC20Session) GetSymbol() (string, error) {
	return _ERC20.Contract.GetSymbol(&_ERC20.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() view returns(string)
func (_ERC20 *ERC20CallerSession) GetSymbol() (string, error) {
	return _ERC20.Contract.GetSymbol(&_ERC20.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) GetTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "getTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) GetTotalSupply() (*big.Int, error) {
	return _ERC20.Contract.GetTotalSupply(&_ERC20.CallOpts)
}

// GetTotalSupply is a free data retrieval call binding the contract method 0xc4e41b22.
//
// Solidity: function getTotalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) GetTotalSupply() (*big.Int, error) {
	return _ERC20.Contract.GetTotalSupply(&_ERC20.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20 *ERC20Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20 *ERC20Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20.Contract.HasRole(&_ERC20.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ERC20 *ERC20CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ERC20.Contract.HasRole(&_ERC20.CallOpts, role, account)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfd1fc4a0.
//
// Solidity: function airDrop(address[] addrs, uint256 money) returns(bool)
func (_ERC20 *ERC20Transactor) AirDrop(opts *bind.TransactOpts, addrs []common.Address, money *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "airDrop", addrs, money)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfd1fc4a0.
//
// Solidity: function airDrop(address[] addrs, uint256 money) returns(bool)
func (_ERC20 *ERC20Session) AirDrop(addrs []common.Address, money *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.AirDrop(&_ERC20.TransactOpts, addrs, money)
}

// AirDrop is a paid mutator transaction binding the contract method 0xfd1fc4a0.
//
// Solidity: function airDrop(address[] addrs, uint256 money) returns(bool)
func (_ERC20 *ERC20TransactorSession) AirDrop(addrs []common.Address, money *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.AirDrop(&_ERC20.TransactOpts, addrs, money)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_ERC20 *ERC20Transactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "burn", burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_ERC20 *ERC20Session) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Burn(&_ERC20.TransactOpts, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Burn(&_ERC20.TransactOpts, burnAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address target, uint256 mintedAmount) returns(bool)
func (_ERC20 *ERC20Transactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "mintToken", target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address target, uint256 mintedAmount) returns(bool)
func (_ERC20 *ERC20Session) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.MintToken(&_ERC20.TransactOpts, target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(address target, uint256 mintedAmount) returns(bool)
func (_ERC20 *ERC20TransactorSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.MintToken(&_ERC20.TransactOpts, target, mintedAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ERC20 *ERC20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ERC20 *ERC20Session) Pause() (*types.Transaction, error) {
	return _ERC20.Contract.Pause(&_ERC20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ERC20 *ERC20TransactorSession) Pause() (*types.Transaction, error) {
	return _ERC20.Contract.Pause(&_ERC20.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns(bool)
func (_ERC20 *ERC20Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns(bool)
func (_ERC20 *ERC20Session) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _ERC20.Contract.RenounceRole(&_ERC20.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns(bool)
func (_ERC20 *ERC20TransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _ERC20.Contract.RenounceRole(&_ERC20.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns(bool)
func (_ERC20 *ERC20Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns(bool)
func (_ERC20 *ERC20Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.RevokeRole(&_ERC20.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns(bool)
func (_ERC20 *ERC20TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.RevokeRole(&_ERC20.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x2d958471.
//
// Solidity: function setUpRole(bytes32 role, address account) returns(bool)
func (_ERC20 *ERC20Transactor) SetUpRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "setUpRole", role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x2d958471.
//
// Solidity: function setUpRole(bytes32 role, address account) returns(bool)
func (_ERC20 *ERC20Session) SetUpRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetUpRole(&_ERC20.TransactOpts, role, account)
}

// SetUpRole is a paid mutator transaction binding the contract method 0x2d958471.
//
// Solidity: function setUpRole(bytes32 role, address account) returns(bool)
func (_ERC20 *ERC20TransactorSession) SetUpRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ERC20.Contract.SetUpRole(&_ERC20.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ERC20 *ERC20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ERC20 *ERC20Session) Unpause() (*types.Transaction, error) {
	return _ERC20.Contract.Unpause(&_ERC20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ERC20 *ERC20TransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC20.Contract.Unpause(&_ERC20.TransactOpts)
}

// ERC20AirDropIterator is returned from FilterAirDrop and is used to iterate over the raw logs and unpacked data for AirDrop events raised by the ERC20 contract.
type ERC20AirDropIterator struct {
	Event *ERC20AirDrop // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20AirDropIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20AirDrop)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20AirDrop)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20AirDropIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20AirDropIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20AirDrop represents a AirDrop event raised by the ERC20 contract.
type ERC20AirDrop struct {
	Addr  []common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAirDrop is a free log retrieval operation binding the contract event 0xeeb9052f86d5d21d099b20a1a39748440e720a832b3b99482afa7152fa6bd906.
//
// Solidity: event AirDrop(address[] addr, uint256 value)
func (_ERC20 *ERC20Filterer) FilterAirDrop(opts *bind.FilterOpts) (*ERC20AirDropIterator, error) {

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "AirDrop")
	if err != nil {
		return nil, err
	}
	return &ERC20AirDropIterator{contract: _ERC20.contract, event: "AirDrop", logs: logs, sub: sub}, nil
}

// WatchAirDrop is a free log subscription operation binding the contract event 0xeeb9052f86d5d21d099b20a1a39748440e720a832b3b99482afa7152fa6bd906.
//
// Solidity: event AirDrop(address[] addr, uint256 value)
func (_ERC20 *ERC20Filterer) WatchAirDrop(opts *bind.WatchOpts, sink chan<- *ERC20AirDrop) (event.Subscription, error) {

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "AirDrop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20AirDrop)
				if err := _ERC20.contract.UnpackLog(event, "AirDrop", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAirDrop is a log parse operation binding the contract event 0xeeb9052f86d5d21d099b20a1a39748440e720a832b3b99482afa7152fa6bd906.
//
// Solidity: event AirDrop(address[] addr, uint256 value)
func (_ERC20 *ERC20Filterer) ParseAirDrop(log types.Log) (*ERC20AirDrop, error) {
	event := new(ERC20AirDrop)
	if err := _ERC20.contract.UnpackLog(event, "AirDrop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
