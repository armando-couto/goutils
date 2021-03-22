package goutils

import (
	"github.com/joho/godotenv"
	"os"
	"runtime"
)

func Godotenv(key string) string {
	var err error
	if runtime.GOOS == "linux" {
		err = godotenv.Load(".env.production")
	} else {
		err = godotenv.Load(".env")
	}
	if err != nil {
		err = godotenv.Load("../.env.test")
		if err != nil {
		}
	}
	return os.Getenv(key)
}
