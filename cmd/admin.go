package cmd

import (
	"errors"
	"fmt"
	"math/big"

	callconts "memoContract/callcontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// AdminCmd is about contracts functions called by admin
var AdminCmd = &cli.Command{
	Name:  "admin",
	Usage: "Admin deploy and call contracts",
	Subcommands: []*cli.Command{
		deployERC20Cmd,
		mintCmd,
	},
}

var deployERC20Cmd = &cli.Command{
	Name:  "deploy-erc20",
	Usage: "admin deploy ERC20 contract",
	Action: func(cctx *cli.Context) error {
		name := cctx.Args().Get(0)
		symbol := cctx.Args().Get(1)
		fmt.Println("name:", name, " symbol:", symbol)

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		e := callconts.NewERC20(callconts.AdminAddr, callconts.AdminAddr, callconts.AdminSk, txopts)
		erc20Addr, _, err := e.DeployERC20(name, symbol)
		if err != nil {
			return err
		}
		fmt.Println("ERC20 contract address:", erc20Addr.Hex())

		return nil
	},
}

var mintCmd = &cli.Command{
	Name:  "mint",
	Usage: "admin mint ERC20 token",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "erc20Addr",
			Usage: "indicate the ERC20 contract address",
			Value: "0xa96303D074eF892F39BCF5E19CD25Eeff7A73BAA",
		},
	},
	Action: func(cctx *cli.Context) error {
		target := cctx.Args().Get(0)
		amount := cctx.Args().Get(1)
		fmt.Println("amount", amount, "target:", target)

		// TODO:check target

		mintValue := big.NewInt(0)
		mintValue, ok := mintValue.SetString(amount, 10)
		if !ok {
			return errors.New("SetString error")
		}
		if mintValue.Cmp(big.NewInt(0)) <= 0 {
			fmt.Println("mintValue should be more than 0")
		}

		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		erc20Addr := common.HexToAddress(cctx.String("erc20Addr"))
		e := callconts.NewERC20(erc20Addr, callconts.AdminAddr, callconts.AdminSk, txopts)
		err := e.MintToken(common.HexToAddress(target), mintValue)
		if err != nil {
			return err
		}

		return nil
	},
}
