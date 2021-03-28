package goutils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

/*
	ConnectionBDMySQL o antigo nome era: ConexaoBDMySQL
*/
func ConnectionBDMySQL() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("database_mysql"))
	if err != nil {
		CreateFileDay(err.Error())
		if db.Ping() != nil {
			CreateFileDay("Erro ping DB")
		}
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(time.Second * 10)
	return db
}
