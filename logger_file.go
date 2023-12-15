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
	Error   string
	Objects []interface{}
}

func CreateFileDay(message Message) {
	f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	var prefix string
	if message.Info != "" {
		prefix = "INFO\t"
	} else if message.Error != "" {
		prefix = "ERROR\t"
	} else {
		prefix = "OTHER\t"
	}

	logger := log.New(f, prefix, log.Ldate|log.Ltime)
	logger.Println(message)
	fmt.Println(message)

	// Sempre que dá essa mensagem é por que está sem permissão no banco para executar essas operacoes, ai a aplicação será encerrada
	if message.Error == "pq: cannot execute INSERT in a read-only transaction" ||
		message.Error == "pq: cannot execute UPDATE in a read-only transaction" ||
		message.Error == "pq: cannot execute DELETE in a read-only transaction" {
		os.Exit(0)
	}
}
