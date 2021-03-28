package goutils

/*
	CheckIfIdIfNotZero o antigo nome era: ChecaSeIdSeNaoEstaZero
*/
func CheckIfIdIfNotZero(objetoId int) interface{} {
	if objetoId == 0 {
		return nil
	}
	return objetoId
}