package goutils

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
	CreateFileDay
*/
func CreateFileDay(message string) {
	f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "producao: ", log.Ltime)
	logger.Println(message)
}
