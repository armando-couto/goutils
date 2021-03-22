package utils

import (
	"math/big"
	"strconv"
)

func Subtract(valo1, valor2 float64) float64 {
	const prec = 200
	a := new(big.Float).SetPrec(prec).SetFloat64(valo1)
	b := new(big.Float).SetPrec(prec).SetFloat64(valor2)
	result := new(big.Float).Sub(a, b)
	retorno, _ := strconv.ParseFloat(result.String(), 64)
	return retorno
}
