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
		CreateFileDayError(err.Error())
		if db.Ping() != nil {
			CreateFileDayError("Erro ping DB")
		}
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(time.Second * 10)
	return db
}
