package test

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	MoneyTo   = 1e18
	WaitTime  = 3 * time.Second
	AdminAddr = "0x1c111472F298E4119150850c198C657DA1F8a368"
	AdminSk   = "0a95533a110ee10bdaa902fed92e56f3f7709a532e22b5974c03c0251648a5d4"
	Acc1      = "0x9011B1c901d330A63d029B6B325EdE69aeEe11d4"
	Sk1       = "7000cd6cee7cdb6bcc7eda212d1c5aea1d8d35321895f4d601fdd49be96fbc7b"
	Acc2      = "0x7B024b830B0a9a315Baf9C76c0E53ec2dD19cf85"
	Sk2       = "a6d933d11c64217989e8a9449c2ad498ff33df51825a425db71c220e6158ee87"
	Acc3      = "0xeA1652B7a432C3A5607C13aEcDa44BE0Ab24692C"
	Sk3       = "376187ebad236c9e45827d6fee5f00573c6d690f3600f3b3cb4159b9a89446c9"
	Acc4      = "0xC7DFcac0e19e67561e80aEc20474cC3D5eC88ad0"
	Sk4       = "c46a4316171ee0070b4f7c8e08229c63d9f5e683b2ca895d38f2e25e0a27ea4f"
	Acc5      = "0xe23627B7c85A55afA2dF7689A10e7a880723ab8D"
	Sk5       = "457a3b8a46f0990b8a75eaa1bf7a157813dd0908a106b216a1af8de6e3a2881f"
)

var (
	Foundation   = common.HexToAddress("0x30F6551c2F970b21C1A9426aeb289c4ED6d570Fd")
	PrimaryToken = common.HexToAddress("0xa96303D074eF892F39BCF5E19CD25Eeff7A73BAA")
	RTokenAddr = common.HexToAddress("0x7a424f9aF3A69e19fe2A839Cf564d620B6C984d7")
)