package goutils

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

/*
	Subtract
*/
func Subtract(valo1, valor2 float64) float64 {
	const prec = 200
	a := new(big.Float).SetPrec(prec).SetFloat64(valo1)
	b := new(big.Float).SetPrec(prec).SetFloat64(valor2)
	result := new(big.Float).Sub(a, b)
	retorno, _ := strconv.ParseFloat(result.String(), 64)
	return retorno
}

/*
	ConvertFloatToFloatScale2
*/
func ConvertFloatToFloatScale2(valor float64) float64 {
	value := strconv.FormatFloat(valor, 'f', 2, 64)
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

/*
	ConvertFloat64ToString
*/
func ConvertFloat64ToString(value float64) string {
	s := fmt.Sprintf("%.2f", value)
	return s
}

/*
	ConvertkeepZeroToFloat64
*/
func ConvertkeepZeroToFloat64(value KeepZero) float64 {
	s, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return s
}

func ConvertkeepZeroToFloat64To4Decimal(value KeepZero) float64 {
	s, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return s
}

/*
	ConvertFloat64ToString4Decimal
*/
func ConvertFloat64ToString4Decimal(value float64) string {
	s := fmt.Sprintf("%.4f", value)
	return s
}

/*
	ConvertStringToFloat64
*/
func ConvertStringToFloat64(value string) float64 {
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

/*
	ConvertStringToFloatScale2Comma o antigo nome era: ConvertStringToFloatScale2Virgula
*/
func ConvertStringToFloatScale2Comma(value string) float64 {
	value = strings.Replace(value, "%", "", 1)
	value = strings.Replace(value, ".", "", 1)
	value = strings.Replace(value, "R", "", 1)
	value = strings.Replace(value, "$", "", 1)
	value = strings.Replace(value, " ", "", 1)
	value = strings.Replace(value, ",", ".", 1)
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

/*
	ConvertStringToFloatScale2Format o antigo nome era: ConvertStringToFloatScale2
*/
func ConvertStringToFloatScale2FormatNumber(value string) float64 {
	value = strings.Replace(value, ".", "", 1)
	value = strings.Replace(value, ",", "", 1)
	if len(value) == 1 {
		value = fmt.Sprint("0", value)
	}
	value = fmt.Sprint(value[:len(value)-2], ".", value[len(value)-2:])
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

/*
	ConvertStringToFloatScale2
*/
func ConvertStringToFloatScale2(value string) float64 {
	if value == "" {
		return 0
	}
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

/*
	ConvertFloatToBrMoneyString
*/
func ConvertFloatToBrMoneyString(value float64) string{
	str := 	fmt.Sprintf("%.2f", value)
	split := strings.Split(str, ".")

	number := split[0]
	decimal := "00"
	if len(split) > 0 {
		decimal = split[1]
	}

	var values []string

	if len(number) < 3 {
		values = append(values, number)
	} else {
		for i := len(number); i > 0; i -= 3 {
			j := i-3
			if j < 0 {
				j = 0
			}
			values = append([]string{number[j:i]}, values...)
		}
	}
	return strings.Join(values[:], ".") + `,` + decimal
}
