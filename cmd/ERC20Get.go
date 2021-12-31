package cmd

import (
	"fmt"
	"math/big"
	callconts "memoContract/callcontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// erc20 address and caller address set by flags
// input of method set by param
var GetERC20Cmd = &cli.Command{
	Name:  "eget",
	Usage: "call get methods of erc20 contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "erc20",
			Aliases: []string{"e"},
			Value:   callconts.ERC20Addr.Hash().Hex(), // default caller = admin
			Usage:   "erc20 address",
		},
		&cli.StringFlag{
			Name:    "caller",
			Aliases: []string{"c"},
			Value:   callconts.AdminAddr.Hex(), // default caller = admin
			Usage:   "tx caller",
		},
	},
	Subcommands: []*cli.Command{
		nameCmd,
		symbolCmd,
		decCmd,
		tsCmd,
		boCmd,
		alCmd,
	},
}

// get erc20 name
var nameCmd = &cli.Command{
	Name:  "name",
	Usage: "get erc20 token name. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts)
		n, err := e.GetName()
		if err != nil {
			return err
		}
		fmt.Println("erc20 name:", n)

		return nil
	},
}

// get erc20 symbol
var symbolCmd = &cli.Command{
	Name:  "sym",
	Usage: "get erc20 symbol.",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts)
		n, err := e.GetSymbol()
		if err != nil {
			return err
		}
		fmt.Println("erc20 symbol:", n)

		return nil
	},
}

// get erc20 decimal
var decCmd = &cli.Command{
	Name:  "dec",
	Usage: "get erc20 decimal.",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts)
		dec, err := e.GetDecimals()
		if err != nil {
			return err
		}
		fmt.Println("erc20 decimal:", dec)

		return nil
	},
}

// get erc20 total supply
var tsCmd = &cli.Command{
	Name:  "ts",
	Usage: "get erc20 total supply.",
	Action: func(cctx *cli.Context) error {
		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts)
		ts, err := e.GetTotalSupply()
		if err != nil {
			return err
		}
		fmt.Println("erc20 total supply:", ts)

		return nil
	},
}

// get balance of acc
var boCmd = &cli.Command{
	Name:  "bo",
	Usage: "get balance of an acc. arg0: acc address",
	Action: func(cctx *cli.Context) error {
		// parse args
		acc := cctx.Args().Get(0)
		// use primary token as default
		if acc == "" {
			acc = callconts.AdminAddr.Hex()
		} else {
			// check addr
			if len(acc) != 42 || acc == callconts.InvalidAddr {
				fmt.Println("erc20Addr should be with prefix 0x, and shouldn't be 0x0")
				return nil
			}
		}
		fmt.Println("acc address:", acc)

		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20 address:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller address:", caller)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts)
		b, err := e.BalanceOf(common.HexToAddress(acc))
		if err != nil {
			return err
		}
		fmt.Println("balance of acc: ", b)

		return nil
	},
}

// get allowance of an acc pair
var alCmd = &cli.Command{
	Name:  "al",
	Usage: "get allowance of an acc pair.  arg0: sender address, arg1: owner address",
	Action: func(cctx *cli.Context) error {
		// parse args
		sender := cctx.Args().Get(0)
		// check addr
		if len(sender) != 42 || sender == callconts.InvalidAddr {
			fmt.Println("erc20Addr should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println("sender address:", sender)

		owner := cctx.Args().Get(1)
		// check addr
		if len(owner) != 42 || owner == callconts.InvalidAddr {
			fmt.Println("erc20Addr should be with prefix 0x, and shouldn't be 0x0")
			return nil
		}
		fmt.Println("owner address:", owner)

		// parse flags
		erc20 := common.HexToAddress(cctx.String("erc20"))
		fmt.Println("erc20 address:", erc20)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller address:", caller)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}
		// erc20 caller
		e := callconts.NewERC20(erc20, caller, "", txopts)
		al, err := e.Allowance(common.HexToAddress(owner), common.HexToAddress(sender))
		if err != nil {
			return err
		}
		fmt.Printf("allowance of %s to %s is : %v\n", sender, owner, al)

		return nil
	},
}
