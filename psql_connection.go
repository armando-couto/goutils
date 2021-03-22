package goutils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"time"
)

func ConexaoBD() *sql.DB {
	var host, dbname, user, password string
	var port int

	host = Godotenv("host")
	dbname = Godotenv("dbname")
	port, _  = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}