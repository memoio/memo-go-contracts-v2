package callconts

import (
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// EndPoint is rpc endpoint of geth node
var EndPoint string

const (
	//InvalidAddr implements invalid contracts-address
	InvalidAddr               = "0x0000000000000000000000000000000000000000"
	spaceTimePayGasLimit      = uint64(8000000)
	spaceTimePayGasPrice      = 2 * defaultGasPrice
	defaultGasPrice           = 200
	defaultGasLimit           = uint64(8000000)
	sendTransactionRetryCount = 5
	checkTxRetryCount         = 8
	checkTxSleepTime          = 5
	retryTxSleepTime          = time.Minute
	retryGetInfoSleepTime     = time.Minute
	waitTime                  = 3 * time.Second
)

// TxOpts contains some general parameters about sending ethereum transaction
type TxOpts struct {
	Nonce    *big.Int
	GasPrice *big.Int
	GasLimit uint64
}

// getClient get rpc-client based the endPoint
func getClient(endPoint string) *ethclient.Client {
	client, err := rpc.Dial(EndPoint)
	if err != nil {
		log.Println(err)
	}
	return ethclient.NewClient(client)
}

// makeAuth make the transactOpts to call contract
func makeAuth(hexSk string, moneyToContract *big.Int, txopts *TxOpts) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		log.Println("HexToECDSA err: ", err)
		return auth, err
	}

	auth = bind.NewKeyedTransactor(sk)
	auth.GasPrice = txopts.GasPrice
	auth.Value = moneyToContract //放进合约里的钱
	auth.Nonce = txopts.Nonce
	auth.GasLimit = txopts.GasLimit
	return auth, nil
}
