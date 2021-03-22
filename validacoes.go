package goutils

import "strings"

func MaskCard6DigitosIniciais(cardString string) string {
	s := strings.SplitN(cardString, " ", 4)
	return s[0] + s[1][0:2] + "XXXXXX" + s[3]
}

func MaskUltimosDigitos(card string) string {
	sp := strings.Split(card, "XXXXXX")
	return "**********" + sp[1]
}

func MaskCard(cardString string) string {
	return "************" + strings.SplitN(cardString, " ", 4)[3]
}

func ValidadeCartao(valor string) (string, string) {
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

func ValidarTelefone(telefone string) string {
	r := strings.NewReplacer(" ", "",
		"-", "",
		"(", "", ")", "", " ", "")

	// Replace all pairs.
	return r.Replace(telefone)
}