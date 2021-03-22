package utils

func ChecaSeIdSeNaoEstaZero(objetoId int) interface{} {
	if objetoId == 0 {
		return nil
	}
	return objetoId
}