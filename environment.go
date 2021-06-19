package goutils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"runtime"
)

/*
	Godotenv
*/
func Godotenv(key string) string {
	var err error
	if runtime.GOOS == "linux" {
		err = godotenv.Load(".env.production")
	} else {
		err = godotenv.Load(".env")
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	return os.Getenv(key)
}
