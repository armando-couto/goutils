package goutils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func TokenGenerator1() string {
	b := make([]byte, 36)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func RandSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func TokenGeneratorOrderReferenceId() string {
	rand.Seed(time.Now().UnixNano())
	return RandSeq(16)
}

func Contains(a []string, x string) int {
	count := 0
	for _, n := range a {
		if x == n {
			count += 1
		}
	}
	return count
}

func RetiraPontoHifen(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(s, ".", "", -1), "-", "", -1), "/", "", -1)
}