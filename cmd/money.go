package cmd

import (
	"fmt"
	"math/big"

	callconts "memoContract/callcontracts"
	"memoContract/test"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

var transferMoney = big.NewInt(1e9)
var transferEth = big.NewInt(2e18)

// MoneyCmd is about transfer eth or ERC20-token, and get balance
var MoneyCmd = &cli.Command{
	Name:  "money",
	Usage: "transfer eth or ERC20-token",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20Addr",
			Aliases: []string{"e"},
			Value:   "0xa96303D074eF892F39BCF5E19CD25Eeff7A73BAA", //默认值为common.go中的erc20合约地址
			Usage:   "the ERC20 contract address",
		},
	},
	Subcommands: []*cli.Command{
		balanceCmd,
		transferCmd,
		transferEthCmd,
		balanceEthCmd,
	},
}

var balanceCmd = &cli.Command{
	Name:  "balance",
	Usage: "get balance in ERC20 token. Args0:account",
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}

		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		e := callconts.NewERC20(erc20Addr, callconts.AdminAddr, callconts.AdminSk, txopts)
		bal, err := e.BalanceOf(common.HexToAddress(acc))
		if err != nil {
			return err
		}
		fmt.Println("erc20 balance:", bal)
		return nil
	},
}

var transferCmd = &cli.Command{
	Name:  "transfer",
	Usage: "transfer erc20-token to target account. Args0:target",
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}

		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		fmt.Println("erc20Addr:", erc20Addr.Hex())

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		e := callconts.NewERC20(erc20Addr, callconts.AdminAddr, callconts.AdminSk, txopts)
		err := e.Transfer(common.HexToAddress(acc), transferMoney)
		if err != nil {
			return err
		}
		fmt.Println("admin transfer", transferMoney, " to", acc, " successfully!")
		return nil
	},
}

var transferEthCmd = &cli.Command{
	Name:  "transferEth",
	Usage: "transfer eth to target account. Args0:target",
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}

		err := test.TransferTo(transferEth, common.HexToAddress(acc), callconts.EndPoint, callconts.QEndPoint)
		if err != nil {
			return err
		}

		return nil
	},
}

var balanceEthCmd = &cli.Command{
	Name:  "balanceEth",
	Usage: "get eth balance. Args0:account",
	Action: func(cctx *cli.Context) error {
		acc := cctx.Args().Get(0)
		fmt.Println("account:", acc)
		if len(acc) != 42 || acc == callconts.InvalidAddr {
			fmt.Println("account should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}

		bal := callconts.QueryEthBalance(acc, callconts.EndPoint)
		fmt.Println("eth balance:", bal)

		return nil
	},
}
