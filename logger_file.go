package goutils

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
	CreateFileDayInfo o antigo nome era: CreateFileDay
*/
func CreateFileDayInfo(message string) {
	f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	logger.Println(message)
	fmt.Println(message)
}

func CreateFileDayError(message string) {
	f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(message)
	fmt.Println(message)
}

func CreateFileDayInfoNotDate(message string) {
	f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", 0)
	logger.Println(message)
}