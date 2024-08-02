package helper

import (
	"math/big"
)

func TokenAmountStrToFloat64(amount string) float64 {
	a, _ := big.NewFloat(0).SetString(amount)
	a = a.Quo(a, big.NewFloat(1e18))
	f, _ := a.Float64()

	return f
}
