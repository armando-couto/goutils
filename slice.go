package goutils

// Função para verificar se uma string não está na lista
func NotInList(value string, list []string) bool {
	for _, item := range list {
		if item == value {
			return false
		}
	}
	return true
}

// Função para verificar se uma string está na lista
func IsInList(value string, list []string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Função para adicionar um elemento a uma slice apontada por um ponteiro
func AddElementStringWithReference(slicePointer *[]string, element string) {
	*slicePointer = append(*slicePointer, element)
}
