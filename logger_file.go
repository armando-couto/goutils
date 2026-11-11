package goutils

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"sync"
	"time"
)

type Message struct {
	File    string
	Script  string
	Info    string
	Error   string
	Objects []interface{}
}

var (
	logger *log.Logger
	once   sync.Once
)

func getLogger() *log.Logger {
	once.Do(func() {
		name := fmt.Sprintf("%s.log", time.Now().Format("20060102"))

		lj := &lumberjack.Logger{
			Filename:   name,
			MaxSize:    5, // MB
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   true,
		}

		logger = log.New(lj, "", log.Ldate|log.Ltime)
	})

	return logger
}

func CreateFileDay(message Message, m *MessageGotify) {
	logger := getLogger()
	
	// Imprime o Info se não estiver vazio
	if message.Info != "" {
		logger.Printf("INFO\tFile: %s Script: %s | %s | Objects: %v", message.File, message.Script, message.Info, message.Objects)
		fmt.Printf("INFO\tFile: %s Script: %s | %s | Objects: %v\n", message.File, message.Script, message.Info, message.Objects)
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

	// Se não teve Info nem Error, imprime uma linha neutra
	if message.Info == "" && message.Error == "" {
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
