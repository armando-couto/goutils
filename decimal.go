package goutils

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func ConvertDecimalToKeepZero(value decimal.Decimal) KeepZero {
	aux, _ := value.Float64()
	return KeepZero(aux)
}

func ConvertKeepZeroToDecimal(value KeepZero) decimal.Decimal {
	return decimal.NewFromFloat(float64(value))
}

func CompareKeepZeroWithInt(a KeepZero, b int) bool {
	decimalA := ConvertKeepZeroToDecimal(a)

	valueB := KeepZero(ConvertStringToFloatScale2FormatNumber(fmt.Sprint(b)))
	decimalB := ConvertKeepZeroToDecimal(valueB)

	return decimalA.Equal(decimalB)
}
