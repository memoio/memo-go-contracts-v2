package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	callconts "memoContract/callcontracts"
	"memoContract/test"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ethEndPoint  string
	qethEndPoint string
)

func main() {
	//--eth=http://119.147.213.219:8101 --qeth=http://119.147.213.219:8101      testnet网
	eth := flag.String("eth", "http://119.147.213.220:8193", "eth api Address;")   //dev网
	qeth := flag.String("qeth", "http://119.147.213.220:8194", "eth api Address;") //dev网，用于keeper、provider连接
	flag.Parse()
	ethEndPoint = *eth
	qethEndPoint = *qeth
	callconts.EndPoint = ethEndPoint

	// 查看余额，支付交易Gas费
	bal := callconts.QueryEthBalance(test.AdminAddr, ethEndPoint)
	fmt.Println("test admin-account has balance: ", bal, " in Ethereum")
	bal = callconts.QueryEthBalance(test.Acc1, ethEndPoint)
	fmt.Println("test common-account has balance: ", bal, " in Ethereum")

	var (
		name   string   = "memo"
		symbol string   = "M"
		tmp    *big.Int = big.NewInt(0)
	)

	txopts := &callconts.TxOpts{
		Nonce:    nil,
		GasPrice: nil,
		GasLimit: callconts.DefaultGasLimit,
	}
	var erc20Addr common.Address
	status := make(chan error)
	e := callconts.NewERC20(erc20Addr, common.HexToAddress(test.AdminAddr), test.AdminSk, txopts, ethEndPoint, status)

	fmt.Println("============1. begin test deploy ERC20 contract============")
	erc20Addr, _, err := e.DeployERC20(name, symbol)
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ERC20 contract Address:", erc20Addr.Hex()) // 0xa96303D074eF892F39BCF5E19CD25Eeff7A73BAA
	//erc20Addr := common.HexToAddress("0xa96303D074eF892F39BCF5E19CD25Eeff7A73BAA")

	fmt.Println("============2. begin test GetName============")
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.Acc1), test.Sk1, txopts, ethEndPoint, status)
	n, err := e.GetName()
	if err != nil {
		log.Fatal(err)
	}
	if n != name {
		log.Fatal("The name obtained ", n, " does not match the actual")
	}

	fmt.Println("============3. begin test GetSymbol============")
	s, err := e.GetSymbol()
	if err != nil {
		log.Fatal(err)
	}
	if s != symbol {
		log.Fatal("The symbol obtained ", s, " does not match the actual")
	}

	fmt.Println("============4. begin test GetDecimals============")
	d, err := e.GetDecimals()
	if err != nil {
		log.Fatal(err)
	}
	if d != 18 {
		log.Fatal("The decimal obtained ", d, " does not match the actual")
	}

	fmt.Println("============5. begin test BalanceOf============")
	bal, err = e.BalanceOf(common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	if bal.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("The balance in ERC20 obtained ", bal, " does not match the actual")
	}

	fmt.Println("============6. begin test Allowance============")
	allo, err := e.Allowance(common.HexToAddress(test.AdminAddr), common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	if allo.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("The allowance obtained ", allo, " does not match the actual")
	}

	fmt.Println("============7. begin test HasRole in AccessControl============")
	hasAdminRole, err := e.HasRole(callconts.AdminRole, common.HexToAddress(test.AdminAddr))
	if err != nil {
		log.Fatal(err)
	}
	if !hasAdminRole {
		log.Fatal("The admin account doesn't has AdminRole")
	}

	fmt.Println("============8. begin test GetPaused in AccessControl============")
	paused, err := e.GetPaused()
	if err != nil {
		log.Fatal(err)
	}
	if paused {
		log.Fatal("The paused param obtained doesn't match the actual")
	}

	fmt.Println("============9. begin test Pause in AccessControl============")
	err = e.Pause()
	if err == nil {
		log.Fatal("The account has not AdminRole")
	}
	fmt.Println(err) // tmp
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.AdminAddr), test.AdminSk, txopts, ethEndPoint, status)
	err = e.Pause()
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	paused, err = e.GetPaused()
	if err != nil {
		log.Fatal(err)
	}
	if !paused {
		log.Fatal("Pause fails")
	}

	fmt.Println("============10. begin test Unpause in AccessControl============")
	err = e.Unpause()
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	paused, err = e.GetPaused()
	if err != nil {
		log.Fatal(err)
	}
	if paused {
		log.Fatal("Unpause fails")
	}

	fmt.Println("============11. begin test SetUpRole in AccessControl============")
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.Acc1), test.Sk1, txopts, ethEndPoint, status)
	err = e.SetUpRole(callconts.AdminRole, common.HexToAddress(test.Acc1))
	if err == nil {
		log.Fatal("should fail")
	}
	fmt.Println(err)
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.AdminAddr), test.AdminSk, txopts, ethEndPoint, status)
	err = e.SetUpRole(callconts.AdminRole, common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	hasAdminRole, err = e.HasRole(callconts.AdminRole, common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	if !hasAdminRole {
		log.Fatal("SetUpRole fails")
	}

	fmt.Println("============12. begin test RevokeRole in AccessControl============")
	err = e.RevokeRole(callconts.AdminRole, common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	hasAdminRole, err = e.HasRole(callconts.AdminRole, common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	if hasAdminRole {
		log.Fatal("RevokeRole fails")
	}

	ts0, err := e.GetTotalSupply()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TotalSupply0 is: ", ts0)

	fmt.Println("============13. begin test MintToken in ERC20============")
	bal0, err := e.BalanceOf(common.HexToAddress(test.AdminAddr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Old balance of AdminAddr is: ", bal0)
	err = e.MintToken(common.HexToAddress(test.AdminAddr), big.NewInt(test.MoneyTo))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	bal, err = e.BalanceOf(common.HexToAddress(test.AdminAddr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance of AdminAddr is: ", bal)
	if tmp.Sub(bal, bal0).Cmp(big.NewInt(test.MoneyTo)) != 0 {
		log.Fatal("MintToken fails")
	}

	fmt.Println("============14. begin test GetTotalSupply============")
	ts, err := e.GetTotalSupply()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TotalSupply is:", ts)
	if tmp.Sub(ts, ts0).Cmp(big.NewInt(test.MoneyTo)) != 0 {
		log.Fatal("The totalSupply obtained ", ts, " does not match the actual")
	}

	fmt.Println("============15. begin test Transfer in ERC20============")
	bal0, err = e.BalanceOf(common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	err = e.Transfer(common.HexToAddress(test.Acc1), big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	bal, err = e.BalanceOf(common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance of Addr is: ", bal)
	if tmp.Sub(bal, bal0).Cmp(big.NewInt(100)) != 0 {
		log.Fatal("Transfer fails")
	}

	fmt.Println("============16. begin test Approve in ERC20============")
	err = e.Approve(common.HexToAddress(test.Acc1), big.NewInt(500))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	allo, err = e.Allowance(common.HexToAddress(test.AdminAddr), common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(test.AdminAddr, "=>", test.Acc1, "allowance is: ", allo)
	if allo.Cmp(big.NewInt(500)) != 0 {
		log.Fatal("Approve fails")
	}

	fmt.Println("============17. begin test TransferFrom in ERC20============")
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.Acc1), test.Sk1, txopts, ethEndPoint, status)
	bal0, err = e.BalanceOf(common.HexToAddress(test.Acc2))
	if err != nil {
		log.Fatal(err)
	}
	err = e.TransferFrom(common.HexToAddress(test.AdminAddr), common.HexToAddress(test.Acc2), big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	allo, err = e.Allowance(common.HexToAddress(test.AdminAddr), common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("allowance of Addr is: ", allo)
	if allo.Cmp(big.NewInt(400)) != 0 {
		log.Fatal("TransferFrom allowance fails")
	}
	bal, err = e.BalanceOf(common.HexToAddress(test.Acc2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance of Acc2 is: ", bal)
	if tmp.Sub(bal, bal0).Cmp(big.NewInt(100)) != 0 {
		log.Fatal("TransferFrom fails")
	}

	fmt.Println("============18. begin test IncreaseAllowance in ERC20============")
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.AdminAddr), test.AdminSk, txopts, ethEndPoint, status)
	err = e.IncreaseAllowance(common.HexToAddress(test.Acc1), big.NewInt(500))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	allo, err = e.Allowance(common.HexToAddress(test.AdminAddr), common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("allowance of Addr is: ", allo)
	if allo.Cmp(big.NewInt(900)) != 0 {
		log.Fatal("IncreaseAllowance fails")
	}

	fmt.Println("============19. begin test DecreaseAllowance in ERC20============")
	err = e.DecreaseAllowance(common.HexToAddress(test.Acc1), big.NewInt(900))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	allo, err = e.Allowance(common.HexToAddress(test.AdminAddr), common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("allowance of Addr is: ", allo)
	if allo.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("DecreaseAllowance fails")
	}

	fmt.Println("============20. begin test AirDrop in ERC20============")
	bal0, err = e.BalanceOf(common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	bal01, err := e.BalanceOf(common.HexToAddress(test.Acc2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance of Addr、Acc2 is: ", bal0, bal01)
	err = e.AirDrop([]common.Address{common.HexToAddress(test.Acc1), common.HexToAddress(test.Acc2)}, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	bal, err = e.BalanceOf(common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	bal11, err := e.BalanceOf(common.HexToAddress(test.Acc2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance of Addr、Acc2 is: ", bal, bal11)
	if tmp.Sub(bal, bal0).Cmp(big.NewInt(100)) != 0 || tmp.Sub(bal11, bal01).Cmp(big.NewInt(100)) != 0 {
		log.Fatal("AirDrop fails")
	}

	fmt.Println("============21. begin test Burn in ERC20============")
	err = e.SetUpRole(callconts.AdminRole, common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.Acc1), test.Sk1, txopts, ethEndPoint, status)
	err = e.Burn(bal)
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	bal, err = e.BalanceOf(common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance of Addr is: ", bal)
	if bal.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("Burn by Addr fails")
	}

	err = e.SetUpRole(callconts.AdminRole, common.HexToAddress(test.Acc2))
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.Acc2), test.Sk2, txopts, ethEndPoint, status)
	err = e.Burn(bal11)
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	bal, err = e.BalanceOf(common.HexToAddress(test.Acc2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance of Acc2 is: ", bal)
	if bal.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("Burn by Acc2 fails")
	}

	fmt.Println("============22. begin test RenounceRole in AccessControl============")
	err = e.RenounceRole(callconts.AdminRole)
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	hasAdminRole, err = e.HasRole(callconts.AdminRole, common.HexToAddress(test.Acc2))
	if err != nil {
		log.Fatal(err)
	}
	if hasAdminRole {
		log.Fatal("Acc2 RenounceRole fails")
	}

	e = callconts.NewERC20(erc20Addr, common.HexToAddress(test.Acc1), test.Sk1, txopts, ethEndPoint, status)
	err = e.RenounceRole(callconts.AdminRole)
	if err != nil {
		log.Fatal(err)
	}
	err = <-status
	if err != nil {
		log.Fatal(err)
	}
	hasAdminRole, err = e.HasRole(callconts.AdminRole, common.HexToAddress(test.Acc1))
	if err != nil {
		log.Fatal(err)
	}
	if hasAdminRole {
		log.Fatal("Addr RenounceRole fails")
	}

	fmt.Println("============test success!============")
}
