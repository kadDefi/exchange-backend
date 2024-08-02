package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	EmptyAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	EmptyHash    = common.HexToHash("0x0")
	TxRetryCount = 5
	Timeout      = 600
	ExtraGas     = 20
)
