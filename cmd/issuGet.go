package cmd

import (
	"fmt"
	"math/big"
	callconts "memoContract/callcontracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

// GetERC20Cmd erc20 address and caller address set by flags
// input of method set by param
var ISGet = &cli.Command{
	Name:  "isget",
	Usage: "call get methods of issuance contract",
	Flags: []cli.Flag{
		//
		&cli.StringFlag{
			Name:    "issu",
			Aliases: []string{"is"},
			Value:   callconts.IssuanceAddr.Hex(), // default fs addr
			Usage:   "issu address",
		},
		// caller
		&cli.StringFlag{
			Name:    "caller",
			Aliases: []string{"c"},
			Value:   callconts.AdminAddr.Hex(), // default caller = admin
			Usage:   "tx caller",
		},
		// end point
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   callconts.EndPoint, //默认值为common.go中的endPoint
			Usage:   "the geth endPoint",
		},
	},
	Subcommands: []*cli.Command{
		mlCmd,
		lmCmd,
		prCmd,
		szCmd,
		stCmd,
	},
}

//
var mlCmd = &cli.Command{
	Name:  "ml",
	Usage: "get mintlevel. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		ml, err := e.MintLevel()
		if err != nil {
			return err
		}
		fmt.Printf("\nmint level: %v [0x%x]\n", ml, ml)

		return nil
	},
}

//
var lmCmd = &cli.Command{
	Name:  "lm",
	Usage: "get last mint. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		lm, err := e.LastMint()
		if err != nil {
			return err
		}
		fmt.Printf("\nlast mint: %v [0x%x]\n", lm, lm)

		return nil
	},
}

var prCmd = &cli.Command{
	Name:  "pr",
	Usage: "get price. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		pr, err := e.Price()
		if err != nil {
			return err
		}
		fmt.Printf("\nprice: %v [0x%x]\n", pr, pr)

		return nil
	},
}

var szCmd = &cli.Command{
	Name:  "sz",
	Usage: "get size. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		sz, err := e.Size()
		if err != nil {
			return err
		}
		fmt.Printf("\nsize: %v [0x%x]\n", sz, sz)

		return nil
	},
}

var stCmd = &cli.Command{
	Name:  "st",
	Usage: "get space time. ",
	Action: func(cctx *cli.Context) error {
		// parse flags
		issu := common.HexToAddress(cctx.String("issu"))
		fmt.Println("issu:", issu)
		caller := common.HexToAddress(cctx.String("caller"))
		fmt.Println("caller:", caller)
		endPoint := cctx.String("endPoint")
		fmt.Println("endPoint:", endPoint)

		// send tx
		txopts := &callconts.TxOpts{
			Nonce:    nil,
			GasPrice: big.NewInt(callconts.DefaultGasPrice),
			GasLimit: callconts.DefaultGasLimit,
		}

		// caller
		e := callconts.NewIssu(issu, caller, "", txopts, endPoint, make(chan error))
		st, err := e.SpaceTime()
		if err != nil {
			return err
		}
		fmt.Printf("\nspace time: %v [0x%x]\n", st, st)

		return nil
	},
}
