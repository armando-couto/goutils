package goutils

import (
	"regexp"
	"strconv"
	"strings"
)

func RemoveZerosInLeft(value string) string {
	i, _ := strconv.Atoi(value)
	return ConvertIntToString(i)
}

func TirarEspacoString(value string) string {
	t := strings.ReplaceAll(value, " ", "")
	return t
}

func RemoveCaracteres(value string) string {
	value = strings.Replace(value, ".", "", 1)
	value = strings.Replace(value, ".", "", 1)
	value = strings.Replace(value, "-", "", 1)
	value = strings.Replace(value, "/", "", 1)
	return value
}

func PadronizaMascarasDeCartao(numeroCartao string) string {
	numero := strings.ReplaceAll(numeroCartao, "000", "")
	numero = strings.ReplaceAll(numero, "*", "")
	return numero
}

func StringTrim(value string) string {
	return strings.Trim(value, " ")
}

func RemoveSpecialCharacters(value string) string {
	// Make a Regex to say we only want letters and numbers
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(value, "")
}