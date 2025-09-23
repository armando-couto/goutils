package goutils

import "github.com/shopspring/decimal"

func ConvertDecimalToKeepZero(value decimal.Decimal) KeepZero {
	aux, _ := value.Float64()
	return KeepZero(aux)
}

func ConvertKeepZeroToDecimal(value KeepZero) decimal.Decimal {
	return decimal.NewFromFloat(float64(value))
}
