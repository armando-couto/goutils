package goutils

import "strings"

/*
	MaskCard6Initials o antigo nome era: MaskCard6DigitosIniciais
*/
func MaskCard6Initials(cardString string) string {
	s := strings.ReplaceAll(cardString, " ", "")
	return s[0:6] + "XXXXXX" + s[len(s)-4:]
}
/*
	MaskLastDigits o antigo nome era: MaskUltimosDigitos
*/
func MaskLastDigits(card string) string {
	sp := strings.Split(card, "XXXXXX")
	return "**********" + sp[1]
}

/*
	MaskCard
*/
func MaskCard(cardString string) string {
	return "************" + strings.SplitN(cardString, " ", 4)[3]
}

/*
	ExpiryDate o antigo nome era: ValidadeCartao
*/
func ExpiryDate(valor string) (string, string) {
	var validade []string
	if strings.Contains(valor, "/") {
		validade = strings.Split(valor, "/")
	}
	if strings.Contains(valor, "-") {
		validade = strings.Split(valor, "-")
	}

	mesValidade := strings.Split(validade[0], "0")
	mes := ""
	if mesValidade[0] == "" {
		mes = mesValidade[1]
	} else {
		mes = validade[0]
	}
	return mes, validade[1]
}

/*
	ValidateTelephone o antigo nome era: ValidarTelefone
*/
func ValidateTelephone(telefone string) string {
	r := strings.NewReplacer(" ", "",
		"-", "",
		"(", "", ")", "", " ", "")

	// Replace all pairs.
	return r.Replace(telefone)
}
