package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"

	callconts "memoContract/callcontracts"
	test "memoContract/test"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var errHexskFormat = errors.New("the hexsk'format is wrong")

func main() {
	fmt.Println("welcome to test contract!")

	encodePackedHex := "00000000000000030000000061bc37780000000061bc3779000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000561"
	fmt.Println("len:", len(encodePackedHex))

	encodePacked := hexToByte(encodePackedHex)
	solh := hexToByte("85c6ceefb89fb4e3e80dd6db29861de012e185df27d365278b5187e990088736")
	fmt.Println("encodePacked:", encodePacked)
	fmt.Println("sol-hash:", solh)
	fmt.Println("sol-hash:", bytesToHex(solh))

	h, err := GetHash(3, 1639724920, 1639724921, 10, 0, 0, big.NewInt(5), "a")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("h:", h)

	sign, err := callconts.SignForRepair(test.AdminSk, 3, 1639724920, 1639724921, 10, 0, 0, big.NewInt(5), "a")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signature:", sign)
	fmt.Println("signature:", bytesToHex(sign))
}

func hexToByte(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func bytesToHex(b []byte) string {
	s := hex.EncodeToString(b)
	return s
}

func GetHash(pIndex, start, end, size, nonce uint64, tIndex uint32, sprice *big.Int, label string) ([]byte, error) {
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

	return hash, nil
}

func bytesCombine(b ...[]byte) []byte {
	return bytes.Join(b, []byte(""))
}
