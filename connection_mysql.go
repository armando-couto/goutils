package goutils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/*
ConnectionBDMySQL o antigo nome era: ConexaoBDMySQL
*/
func ConnectionBDMySQL() *sql.DB {
	db, err := sql.Open("mysql", Godotenv("database_mysql"))
	if err != nil {
		CreateFileDay(Message{File: "ConnectionBDMySQL()", Error: err.Error()}, &MessageGotifyGlobal)
		if db.Ping() != nil {
			CreateFileDay(Message{File: "ConnectionBDMySQL()", Error: "Erro ping DB"}, nil)
		}
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(time.Second * 10)
	return db
}
