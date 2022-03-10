package goutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type MongoObject struct {
	Date    time.Time
	Message string
}

/*
	CreateFileDayInfo o antigo nome era: CreateFileDay
*/
func CreateFileDayInfo(message string) {
	message = strings.ReplaceAll(message, "\n", "")
	if ConvertStringToBool(Godotenv("logger")) {
		connection := ConnectionMongoDB()
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		defer connection.Client().Disconnect(ctx)
		_, err := connection.Collection("infos").InsertOne(ctx, MongoObject{Date: time.Now(), Message: message})
		if err != nil {
			log.Println(err)
		}
	} else {
		f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		logger := log.New(f, "INFO\t", log.Ldate|log.Ltime)
		logger.Println(message)
		fmt.Println(message)
	}
}

func CreateFileDayError(message string) {
	message = strings.ReplaceAll(message, "\n", "")
	if ConvertStringToBool(Godotenv("logger")) {
		connection := ConnectionMongoDB()
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		defer connection.Client().Disconnect(ctx)
		_, err := connection.Collection("errors").InsertOne(ctx, MongoObject{Date: time.Now(), Message: message})
		if err != nil {
			log.Println(err)
		}
	} else {
		f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		logger := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
		fmt.Println(message)
	}
}

type MongoObject2 struct {
	Date   time.Time
	Object interface{}
}

func InsertAudit(object interface{}) {
	if ConvertStringToBool(Godotenv("block_logger")) {
		return
	}

	if ConvertStringToBool(Godotenv("logger")) {
		connection := ConnectionMongoDB()
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		defer connection.Client().Disconnect(ctx)
		_, err := connection.Collection("audits").InsertOne(ctx, MongoObject2{Date: time.Now(), Object: object})
		if err != nil {
			log.Println(err)
		}
	} else {
		f, err := os.OpenFile(fmt.Sprint(time.Now().Format("20060102"), ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		logger := log.New(f, "AUDIT\t", log.Ldate|log.Ltime)
		logger.Println(object)
		fmt.Println(object)
	}
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
