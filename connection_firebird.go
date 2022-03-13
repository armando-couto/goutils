package goutils

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func ConnectionBDFirebird() *sql.DB {
	db, err := sql.Open("firebirdsql", Godotenv("database"))
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(1 * time.Minute)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		time.Sleep(5 * time.Second)
		return ConnectionBDFirebird()
	}

	return db
}
