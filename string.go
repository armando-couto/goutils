package goutils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
	RemoveZerosInLeft
*/
func RemoveZerosInLeft(value string) string {
	i, _ := strconv.Atoi(value)
	return ConvertIntToString(i)
}

/*
	RemoveSpaceString o antigo nome era: TirarEspacoString
*/
func RemoveSpaceString(value string) string {
	t := strings.ReplaceAll(value, " ", "")
	return t
}

/*
	RemoveCharacters o antigo nome era: RemoveCaracteres
*/
func RemoveCharacters(value string) string {
	value = strings.Replace(value, ".", "", 1)
	value = strings.Replace(value, ".", "", 1)
	value = strings.Replace(value, "-", "", 1)
	value = strings.Replace(value, "/", "", 1)
	return value
}

/*
	StandardizesMasksByCard o antigo nome era: PadronizaMascarasDeCartao
*/
func StandardizesMasksByCard(numeroCartao string) string {
	numero := strings.ReplaceAll(numeroCartao, "000", "")
	numero = strings.ReplaceAll(numero, "*", "")
	return numero
}

/*
	StringTrim
*/
func StringTrim(value string) string {
	return strings.Trim(value, " ")
}

/*
	RemoveSpecialCharacters
*/
func RemoveSpecialCharacters(value string) string {
	// Make a Regex to say we only want letters and numbers
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(value, "")
}

/*
	RandSeq
*/
func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/*
	TokenGeneratorNLength o antigo nome era: RandToken
*/
func TokenGeneratorNLength(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/*
	TokenGeneratorOrderReferenceId
*/
func TokenGeneratorOrderReferenceId() string {
	rand.Seed(time.Now().UnixNano())
	return RandSeq(16)
}

/*
	RemoveHeadHyphen o antigo nome era: RetiraPontoHifen
*/
func RemoveHeadHyphen(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(s, ".", "", -1), "-", "", -1), "/", "", -1)
}

/*
	ValidateIfNotEmptyNumber o antigo nome era: ValidaDecimal
*/
func ValidateIfNotEmptyNumber(valor string) string {
	if valor == "" {
		return "0"
	}
	return valor
}

/*
	ValidateIfNotEmptyDate o antigo nome era: ValidaData
*/
func ValidateIfNotEmptyDate(data string) string {
	if data == "" {
		return "0001-01-01"
	}
	return data
}

/*
	RemoveCNPJMask o antigo nome era: RemoveMascaraCNPJ
*/
func RemoveCNPJMask(cnpj string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(cnpj, ".", ""), "-", ""), "/", "")
}

/*
	ParseBinToHex
*/
func ParseBinToHex(s string) string {
	return strconv.FormatInt(int64(ConvertStringToInt(s)), 16)
}
