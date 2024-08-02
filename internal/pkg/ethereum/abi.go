package ethereum

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func ABIEncode(values ...interface{}) []byte {
	bs := make([]byte, 0)

	for _, v := range values {
		var b []byte

		switch v := v.(type) {
		case common.Hash:
			b = v.Bytes()
		case common.Address:
			b = v.Bytes()
		case uint64:
			b = big.NewInt(int64(v)).Bytes()
		case int64:
			b = big.NewInt(int64(v)).Bytes()
		case uint32:
			b = big.NewInt(int64(v)).Bytes()
		case int32:
			b = big.NewInt(int64(v)).Bytes()
		case uint16:
			b = big.NewInt(int64(v)).Bytes()
		case int16:
			b = big.NewInt(int64(v)).Bytes()
		case uint8:
			b = big.NewInt(int64(v)).Bytes()
		case int8:
			b = big.NewInt(int64(v)).Bytes()
		case uint:
			b = big.NewInt(int64(v)).Bytes()
		case int:
			b = big.NewInt(int64(v)).Bytes()
		case bool:
			if v {
				b = []byte{1}
			} else {
				b = []byte{0}
			}
		case string:
			if n, ok := big.NewInt(0).SetString(v, 0); ok {
				b = n.Bytes()
			} else {
				b = []byte(v)
			}
		case []byte:
			b = v
		default:
			continue
		}

		b = common.LeftPadBytes(b, 32)
		bs = append(bs, b...)
	}

	return bs
}

func Hex2Bytes(s string) []byte {
	return common.Hex2Bytes(strings.TrimPrefix(s, "0x"))
}
