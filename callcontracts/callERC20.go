package callconts

import (
	"log"
	"math/big"
	"memoContract/contracts/erc20"
	iface "memoContract/interfaces"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

// NewERC20 new a instance of ContractModule
func NewERC20(addr common.Address, hexSk string, txopts *TxOpts) iface.ERC20Info {
	e := &ContractModule{
		addr:   addr,
		hexSk:  hexSk,
		txopts: txopts,
	}

	return e
}

// DeployERC20 deploy an ERC20 contract, called by admin, specify name and symbol.
func (e *ContractModule) DeployERC20(name, symbol string) (common.Address, *erc20.ERC20, error) {
	var erc20Addr, erc20Address common.Address
	var erc20Instance, erc20Ins *erc20.ERC20
	var err error

	log.Println("begin deploy ERC20 contract...")
	client := getClient(EndPoint)
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return erc20Addr, nil, errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		erc20Address, tx, erc20Instance, err = erc20.DeployERC20(auth, client, name, symbol)
		if erc20Address.String() != InvalidAddr {
			erc20Addr = erc20Address
			erc20Ins = erc20Instance
		}
		if err != nil {
			retryCount++
			log.Println("deploy ERC20 Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return erc20Addr, erc20Ins, err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("deploy ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return erc20Addr, erc20Ins, err
			}
			continue
		}
		break
	}
	log.Println("ERC20 has been successfully deployed! The address is ", erc20Addr.Hex())
	return erc20Addr, erc20Ins, nil
}

// newERC20 new an instance of ERC20 contract, 'erc20Addr' indicates ERC20 contract address
func newERC20(erc20Addr common.Address) (*erc20.ERC20, error) {
	erc20Ins, err := erc20.NewERC20(erc20Addr, getClient(EndPoint))
	if err != nil {
		return nil, err
	}
	return erc20Ins, nil
}

// Transfer the account represented by the e.hexsk transfers the specified amount to recipient.
func (e *ContractModule) Transfer(erc20Addr, recipient common.Address, value *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	if recipient.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	// need to determine whether the account balance is enough to transfer.
	bal, err := e.BalanceOf(erc20Addr, e.addr)
	if err != nil {
		return err
	}
	log.Println("Your balance is ", bal)
	if bal.Cmp(value) < 0 {
		log.Println("Balance is not enough, please reconfirm the transfer amount.")
		return ErrBalNotE
	}

	log.Println("begin Transfer to", recipient.Hex(), " with value", value, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.Transfer(auth, recipient, value)
		if err != nil {
			retryCount++
			log.Println("Transfer Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("Transfer in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Transfer in ERC20 has been successful!")
	return nil
}

// Approve The account represented by the e.hexsk authorizes the balance of the specified amount to addr.
func (e *ContractModule) Approve(erc20Addr, addr common.Address, value *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	if addr.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	// need to determine whether the account balance is enough to approve.
	bal, err := e.BalanceOf(erc20Addr, e.addr)
	if err != nil {
		return err
	}
	log.Println("Your balance is ", bal)
	if bal.Cmp(value) < 0 {
		log.Println("Balance is not enough, please reconfirm the approve amount.")
		return ErrBalNotE
	}

	log.Println("begin Approve", addr.Hex(), " with value", value, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.Approve(auth, addr, value)
		if err != nil {
			retryCount++
			log.Println("Approve Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("Approve in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Approve in ERC20 has been successful!")
	return nil
}

// TransferFrom The account represented by the e.hexsk transfer value of the specified amount from sender to recipient.
func (e *ContractModule) TransferFrom(erc20Addr, sender, recipient common.Address, value *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	// need to determine whether the account allowance is enough to TransferFrom.
	allo, err := e.Allowance(erc20Addr, sender, e.addr)
	if err != nil {
		return err
	}
	log.Println("Your allowance of", sender.Hex(), " is ", allo)
	if allo.Cmp(value) < 0 {
		log.Println("Allowance is not enough, please reconfirm the TransferFrom amount.")
		return ErrBalNotE
	}

	// need to determine whether the sender balance is enough to transfer.
	bal, err := e.BalanceOf(erc20Addr, sender)
	if err != nil {
		return err
	}
	log.Println("Sender ", sender.Hex(), " balance is", bal)
	if bal.Cmp(value) < 0 {
		log.Println("Sender balance is not enough.")
		return ErrBalNotE
	}

	log.Println("begin TransferFrom from", sender.Hex(), "to", recipient.Hex(), " with value", value, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.TransferFrom(auth, sender, recipient, value)
		if err != nil {
			retryCount++
			log.Println("TransferFrom Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("TransferFrom in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("TransferFrom in ERC20 has been successful!")
	return nil
}

// IncreaseAllowance The account represented by the e.hexsk increase the allowance for recipient.
func (e *ContractModule) IncreaseAllowance(erc20Addr, recipient common.Address, value *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	if recipient.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	log.Println("begin IncreaseAllowance to", recipient.Hex(), " with value", value, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.IncreaseAllowance(auth, recipient, value)
		if err != nil {
			retryCount++
			log.Println("IncreaseAllowance Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("IncreaseAllowance in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("IncreaseAllowance in ERC20 has been successful!")
	return nil
}

// DecreaseAllowance The account represented by the e.hexsk decrease the allowance for recipient.
func (e *ContractModule) DecreaseAllowance(erc20Addr, recipient common.Address, value *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	if recipient.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	log.Println("begin DecreaseAllowance to", recipient.Hex(), " with value", value, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.DecreaseAllowance(auth, recipient, value)
		if err != nil {
			retryCount++
			log.Println("DecreaseAllowance Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("DecreaseAllowance in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("DecreaseAllowance in ERC20 has been successful!")
	return nil
}

// MintToken The account represented by the e.hexsk mint token to target. Called by who has MINTER_ROLE.
func (e *ContractModule) MintToken(erc20Addr, acAddr, target common.Address, mintValue *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	hasMinterRole, err := e.HasRole(acAddr, uint8(1), e.addr)
	if err != nil {
		return err
	}
	if !hasMinterRole {
		return ErrNoMintRight
	}

	if target.Hex() == InvalidAddr {
		return ErrInValAddr
	}

	log.Println("begin MintToken to", target.Hex(), " with value", mintValue, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.MintToken(auth, target, mintValue)
		if err != nil {
			retryCount++
			log.Println("MintToken Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("MintToken in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("MintToken in ERC20 has been successful!")
	return nil
}

// Burn The account represented by the e.hexsk burn it's balance. Called by who has DEFAULT_ADMIN_ROLE.
func (e *ContractModule) Burn(acAddr, erc20Addr common.Address, burnValue *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	hasAdminRole, err := e.HasRole(acAddr, uint8(0), e.addr)
	if err != nil {
		return err
	}
	if !hasAdminRole {
		return ErrNoAdminRight
	}

	bal, err := e.BalanceOf(erc20Addr, e.addr)
	if err != nil {
		return err
	}
	if bal.Cmp(burnValue) < 0 {
		return ErrBalNotE
	}

	log.Println("begin Burn with value", burnValue, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.Burn(auth, burnValue)
		if err != nil {
			retryCount++
			log.Println("Burn Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("Burn in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("Burn in ERC20 has been successful!")
	return nil
}

// AirDrop The account represented by the e.hexsk airdrop to targets. Called by who has DEFAULT_ADMIN_ROLE.
func (e *ContractModule) AirDrop(acAddr, erc20Addr common.Address, targets []common.Address, value *big.Int) error {
	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return err
	}

	hasAdminRole, err := e.HasRole(acAddr, uint8(0), e.addr)
	if err != nil {
		return err
	}
	if !hasAdminRole {
		return ErrNoAdminRight
	}

	var tmp []string
	var tmpAddr string
	for _, t := range targets {
		tmpAddr = t.Hex()
		if tmpAddr == InvalidAddr {
			return ErrInValAddr
		}
		tmp = append(tmp, tmpAddr)
	}

	log.Println("begin AirDrop to", tmp, " with value", value, " in ERC20 contract...")
	tx := &types.Transaction{}
	retryCount := 0
	checkRetryCount := 0

	for {
		auth, errMA := makeAuth(e.hexSk, nil, e.txopts)
		if errMA != nil {
			return errMA
		}

		// generally caused by too low gasprice
		rebuild(err, tx, auth)

		tx, err = erc20Ins.AirDrop(auth, targets, value)
		if err != nil {
			retryCount++
			log.Println("AirDrop Err:", err)
			if err.Error() == core.ErrNonceTooLow.Error() && auth.GasPrice.Cmp(big.NewInt(defaultGasPrice)) > 0 {
				log.Println("previously pending transaction has successfully executed")
				break
			}
			if retryCount > sendTransactionRetryCount {
				return err
			}
			time.Sleep(retryTxSleepTime)
			continue
		}

		err = checkTx(tx)
		if err != nil {
			checkRetryCount++
			log.Println("AirDrop in ERC20 transaction fails:", err)
			if checkRetryCount > checkTxRetryCount {
				return err
			}
			continue
		}
		break
	}
	log.Println("AirDrop in ERC20 has been successful!")
	return nil
}

// GetName get the name of erc20 token.
func (e *ContractModule) GetName(erc20Addr common.Address) (string, error) {
	var name string

	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return name, err
	}

	retryCount := 0
	for {
		retryCount++
		name, err = erc20Ins.GetName(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return name, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return name, nil
	}
}

// GetSymbol get the symbol of erc20 token.
func (e *ContractModule) GetSymbol(erc20Addr common.Address) (string, error) {
	var name string

	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return name, err
	}

	retryCount := 0
	for {
		retryCount++
		name, err = erc20Ins.GetSymbol(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return name, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return name, nil
	}
}

// GetDecimals get the decimals of erc20 token. For example, it's 18 in eth.
func (e *ContractModule) GetDecimals(erc20Addr common.Address) (uint8, error) {
	var decimals uint8

	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return decimals, err
	}

	retryCount := 0
	for {
		retryCount++
		decimals, err = erc20Ins.GetDecimals(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return decimals, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return decimals, nil
	}
}

// GetTotalSupply get the total supply of erc20 token.
func (e *ContractModule) GetTotalSupply(erc20Addr common.Address) (*big.Int, error) {
	var totalSupply *big.Int

	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return totalSupply, err
	}

	retryCount := 0
	for {
		retryCount++
		totalSupply, err = erc20Ins.GetTotalSupply(&bind.CallOpts{
			From: e.addr,
		})
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return totalSupply, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return totalSupply, nil
	}
}

// BalanceOf get the balance of addr in erc20 token.
func (e *ContractModule) BalanceOf(erc20Addr, addr common.Address) (*big.Int, error) {
	var balance *big.Int

	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return balance, err
	}

	retryCount := 0
	for {
		retryCount++
		balance, err = erc20Ins.BalanceOf(&bind.CallOpts{
			From: e.addr,
		}, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return balance, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return balance, nil
	}
}

// Allowance get the allowance of sender to addr.
func (e *ContractModule) Allowance(erc20Addr, sender, addr common.Address) (*big.Int, error) {
	var allowance *big.Int

	erc20Ins, err := newERC20(erc20Addr)
	if err != nil {
		return allowance, err
	}

	retryCount := 0
	for {
		retryCount++
		allowance, err = erc20Ins.Allowance(&bind.CallOpts{
			From: e.addr,
		}, sender, addr)
		if err != nil {
			if retryCount > sendTransactionRetryCount {
				return allowance, err
			}
			time.Sleep(retryGetInfoSleepTime)
			continue
		}

		return allowance, nil
	}
}
