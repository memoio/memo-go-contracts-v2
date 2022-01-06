package callconts

import (
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignForRegister Used to call Register on behalf of other accounts
// hash(caller, register)
func SignForRegister(caller common.Address, regAccSK string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSK)
	if err != nil {
		log.Fatal(err)
	}
	//(caller, register)的哈希值
	//label := common.LeftPadBytes([]byte(register), 32)
	label := []byte(register)
	hash := crypto.Keccak256(caller.Bytes(), label)

	//私钥对上述哈希值签名
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,_blsKey,"keeper")
func SignForRegisterKeeper(caller common.Address, _blsKey []byte, regAccSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,_blsKey,"keeper")
	label := []byte(labelKeeper)
	hash := crypto.Keccak256(caller.Bytes(), _blsKey, label)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,"provider")
func SignForRegisterProvider(caller common.Address, regAccSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,"provider")
	label := []byte(labelProvider)
	hash := crypto.Keccak256(caller.Bytes(), label)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,gIndex,payToken,blsKey)
func SignForRegisterUser(caller common.Address, gIndex uint64, payToken uint32, _blsKey []byte, regAccSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(regAccSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,gIndex,payToken,blsKey)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	tmp4 := make([]byte, 4)
	// append caller
	b = append(b, caller.Bytes()...)
	// append gIndex
	binary.BigEndian.PutUint64(tmp8, gIndex)
	b = append(b, tmp8...)
	// append payToken
	binary.BigEndian.PutUint32(tmp4, payToken)
	b = append(b, tmp4...)
	// append blsKey
	b = append(b, _blsKey...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller,gIndex)
func SignForAddProviderToGroup(caller common.Address, gIndex uint64, accSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(accSk)
	if err != nil {
		log.Fatal(err)
	}

	//hash(caller,gIndex,payToken,blsKey)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	// append caller
	b = append(b, caller.Bytes()...)
	// append gIndex
	binary.BigEndian.PutUint64(tmp8, gIndex)
	b = append(b, tmp8...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller, uIndex, tIndex, money)
func SignForRecharge(caller common.Address, uIndex uint64, tIndex uint32, money *big.Int, accSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(caller, uIndex, tIndex, money)
	b := make([]byte, 0)
	tmp8 := make([]byte, 8)
	tmp4 := make([]byte, 4)

	// append caller
	b = append(b, caller.Bytes()...)
	// append uIndex
	binary.BigEndian.PutUint64(tmp8, uIndex)
	b = append(b, tmp8...)
	// append tIndex
	binary.BigEndian.PutUint32(tmp4, tIndex)
	b = append(b, tmp4...)
	// append money
	m := common.LeftPadBytes(money.Bytes(), 32)
	b = append(b, m...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// hash(caller, tIndex, amount)
func SignForWithdrawFromFs(caller common.Address, tIndex uint32, amount *big.Int, accSk string) ([]byte, error) {
	skEcdsa, err := HexSkToEcdsa(accSk)
	if err != nil {
		log.Fatal(err)
	}

	// hash(caller, tIndex, amount)
	b := make([]byte, 0)
	tmp4 := make([]byte, 4)

	// append caller
	b = append(b, caller.Bytes()...)
	// append tIndex
	binary.BigEndian.PutUint32(tmp4, tIndex)
	b = append(b, tmp4...)
	// append amount
	m := common.LeftPadBytes(amount.Bytes(), 32)
	b = append(b, m...)

	fmt.Printf("b: %x\n", b)

	hash := crypto.Keccak256(b)

	//sign
	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// SignForRepair used to call AddRepair or SubRepair, when subRepair, label is "s"; when addRepair, label is "a"
func SignForRepair(sk string, pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, label string) ([]byte, error) {
	ecdsaSk, err := HexSkToEcdsa(sk)
	if err != nil {
		log.Fatal(err)
	}

	// (npIndex, _start, end, _size, nonce, tIndex, sprice, label)的哈希值
	by := make([]byte, 77)
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, pIndex)
	for i, b := range tmp {
		by[i] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, start)
	for i, b := range tmp {
		by[i+8] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, end)
	for i, b := range tmp {
		by[i+16] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, size)
	for i, b := range tmp {
		by[i+24] = byte(b)
	}
	binary.BigEndian.PutUint64(tmp, nonce)
	for i, b := range tmp {
		by[i+32] = byte(b)
	}
	t := make([]byte, 4)
	binary.BigEndian.PutUint32(t, tIndex)
	for i, b := range tmp {
		by[i+40] = byte(b)
	}
	spriceNew := common.LeftPadBytes(sprice.Bytes(), 32)
	for i, b := range spriceNew {
		by[i+44] = byte(b)
	}
	labelNew := []byte(label)
	by[76] = byte(labelNew[0])

	hash := crypto.Keccak256(by)

	fmt.Println("hash:", hash)

	// 私钥签名
	sig, err := crypto.Sign(hash, ecdsaSk)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// HexSkToByte transfer hex string to byte
func HexSkToByte(hexsk string) ([]byte, error) {
	var src []byte
	skLengthNoPrefix := EthSkLength - 2
	skByteEthLength := skLengthNoPrefix / 2

	switch len(hexsk) {
	case EthSkLength:
		if hexsk[:2] == "0x" {
			src = []byte(hexsk[2:])
		} else {
			return nil, errHexskFormat
		}
	case skLengthNoPrefix:
		src = []byte(hexsk)
	default:
		return nil, errHexskFormat
	}

	skByteEth := make([]byte, skByteEthLength)

	_, err := hex.Decode(skByteEth, src)
	if err != nil {
		return nil, err
	}

	return skByteEth, nil
}

// HexSkToEcdsa transfer hex string to ecdsa type
func HexSkToEcdsa(hexsk string) (*ecdsa.PrivateKey, error) {
	skByteEth, err := HexSkToByte(hexsk)
	if err != nil {
		return nil, err
	}
	skECDSA, err := crypto.ToECDSA(skByteEth)
	if err != nil {
		return nil, err
	}
	return skECDSA, nil
}
