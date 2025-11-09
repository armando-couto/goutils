package tests

import (
	"fmt"
	"github.com/armando-couto/goutils"
)

func main() {
	m := goutils.MessageGotify{
		ServerURL: "http://54.146.11.58:8485",
		Token:     "AlDU1S0-NbjrRbX",
		Title:     "Erro no sistema",
		Message:   "Ocorreu um erro inesperado no processamento.",
		Priority:  10,
	}

	err := m.SendNotification()
	if err != nil {
		fmt.Printf("Erro ao enviar notificação: %v\n", err)
	} else {
		fmt.Println("Notificação enviada com sucesso!")
	}
}
