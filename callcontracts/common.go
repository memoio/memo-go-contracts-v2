package callconts

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

var (
	// ErrTxFail indicates that the transaction is not packaged or an error occurred during the packaging process
	ErrTxFail = errors.New("transaction fails")
	// ErrTxExecu indicates that an error occurred during packaging
	ErrTxExecu = errors.New("Transaction mined but execution failed")
)

// TxOpts contains some general parameters about sending ethereum transaction
type TxOpts struct {
	Nonce    *big.Int
	GasPrice *big.Int
	GasLimit uint64
}

//ContractModule  The basic information of node used for contract
type ContractModule struct {
	addr   common.Address //local address
	hexSk  string         //local privateKey
	txopts *TxOpts
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

//CheckTx check whether transaction is successful through receipt
func checkTx(tx *types.Transaction) error {
	log.Println("Check Tx hash:", tx.Hash().Hex(), "nonce:", tx.Nonce(), "gasPrice:", tx.GasPrice())

	var receipt *types.Receipt
	for i := 0; i < 20; i++ {
		receipt = getTransactionReceipt(tx.Hash())
		if receipt != nil {
			break
		}
		t := checkTxSleepTime * (i + 1)
		time.Sleep(time.Duration(t) * time.Second)
	}

	if receipt == nil { //245s获取不到交易信息，判定交易失败
		return ErrTxFail
	}

	if receipt.Status == 0 { //等于0表示交易失败，等于1表示成功
		log.Println("Transaction mined but execution failed")
		txReceipt, err := receipt.MarshalJSON()
		if err != nil {
			return err
		}
		log.Println("TxReceipt:", string(txReceipt))
		return ErrTxExecu
	}

	log.Println("GasUsed:", receipt.GasUsed, "CumulativeGasUsed:", receipt.CumulativeGasUsed)

	return nil
}

//GetTransactionReceipt 通过交易hash获得交易详情
func getTransactionReceipt(hash common.Hash) *types.Receipt {
	client, err := ethclient.Dial(EndPoint)
	if err != nil {
		log.Fatal("rpc.Dial err", err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	return receipt
}

func rebuild(err error, tx *types.Transaction, auth *bind.TransactOpts) {
	if err == ErrTxFail && tx != nil {
		auth.Nonce = big.NewInt(int64(tx.Nonce()))
		auth.GasPrice = new(big.Int).Add(tx.GasPrice(), big.NewInt(defaultGasPrice))
		log.Println("rebuild transaction... nonce is ", auth.Nonce, " gasPrice is ", auth.GasPrice)
	}
}