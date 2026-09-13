package goutils

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Message struct {
	File    string
	Script  string
	Info    string
	Warning string
	Error   string
	Objects []interface{}
}

func CreateFileDay(message Message, m *MessageGotify) {
	f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.Ldate|log.Ltime)

	// Imprime o Info se não estiver vazio
	if message.Info != "" {
		logger.Printf("INFO\tFile: %s Script: %s | %s | Objects: %v", message.File, message.Script, message.Info, message.Objects)
		fmt.Printf("INFO\tFile: %s Script: %s | %s | Objects: %v\n", message.File, message.Script, message.Info, message.Objects)
	}

	// Imprime o Warning se não estiver vazio: é um alerta, não um erro.
	// Vai para o arquivo do dia, para o stdout e para o Gotify com o rótulo WARNING.
	if message.Warning != "" {
		messageString := fmt.Sprintf("WARNING\tFile: %s Script: %s | %s | Objects: %v\n", message.File, message.Script, message.Warning, message.Objects)
		logger.Printf("WARNING\tFile: %s Script: %s | %s | Objects: %v", message.File, message.Script, message.Warning, message.Objects)
		fmt.Print(messageString)
		if m != nil {
			m.SendNotification(message.File, messageString)
		}
	}

	// Imprime o Error se não estiver vazio
	if message.Error != "" {
		messageString := fmt.Sprintf("ERROR\tFile: %s Script: %s | %s | Objects: %v\n", message.File, message.Script, message.Error, message.Objects)
		logger.Printf("ERROR\tFile: %s Script: %s | %s | Objects: %v", message.File, message.Script, message.Error, message.Objects)
		fmt.Println(message)
		if m != nil {
			m.SendNotification(message.File, messageString)
		}
	}

	// Se não teve Info, Warning nem Error, imprime uma linha neutra
	if message.Info == "" && message.Warning == "" && message.Error == "" {
		logger.Printf("OTHER\tFile: %s Script: %s | Objects: %v", message.File, message.Script, message.Objects)
		fmt.Printf("OTHER\tFile: %s Script: %s | Objects: %v\n", message.File, message.Script, message.Objects)
	}

	// Nó de permissão do banco, igual original
	if message.Error == "pq: cannot execute INSERT in a read-only transaction" ||
		message.Error == "pq: cannot execute UPDATE in a read-only transaction" ||
		message.Error == "pq: cannot execute DELETE in a read-only transaction" {
		os.Exit(0)
	}
}
