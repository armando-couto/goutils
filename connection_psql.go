package goutils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

/*
Esse método retorna uma conexão que possa dar INSERT, UPDATE e DELETE.
Podendo passar na variável host vários hosts, isso é muito bom para quando tem replicações, tipo assim:

	host = "localhost;127.0.0.1"
*/
func ConnectionBDPostgreSQL(applicationName, sslMode string, readOnly bool) *sql.DB {
	var dbname, user, password string
	var port int
	var db *sql.DB

	if sslMode == "" {
		sslMode = "disable"
	}

	hosts := strings.Split(Godotenv("host"), ";")
	dbname = Godotenv("dbname")
	port, _ = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	for _, host := range hosts {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s application_name=%s sslmode=%s", host, port, user, password, dbname, applicationName, sslMode)

		db, err := sql.Open("postgres", psqlInfo)
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(1)
		db.SetConnMaxLifetime(1 * time.Minute)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			time.Sleep(2 * time.Second)
			return ConnectionBDPostgreSQL(applicationName, sslMode, readOnly)
		}

		var notReadOnly bool
		// Quando a função pg_is_in_recovery() é chamada e retorna FALSE quer dizer que é uma instância de ESCRITA,
		// se retornar TRUE quer dizer que é somente de LEITURA.
		err = db.QueryRow(`SELECT pg_is_in_recovery();`).Scan(&notReadOnly)
		if err == nil && notReadOnly == readOnly {
			return db
		}
	}

	return db
}

/*
Esse método retorna uma conexão que possa dar INSERT, UPDATE e DELETE.
Podendo passar na variável host vários hosts, isso é muito bom para quando tem replicações, tipo assim:

	host = "localhost;127.0.0.1"
*/
func ConnectionBDPostgreSQLLog(applicationName, sslMode string, readOnly bool) *sql.DB {
	var dbname, user, password string
	var port int
	var db *sql.DB

	if sslMode == "" {
		sslMode = "disable"
	}

	hosts := strings.Split(Godotenv("host_log"), ";")
	dbname = Godotenv("dbname_log")
	port, _ = strconv.Atoi(Godotenv("port_banco_log"))
	user = Godotenv("user_log")
	password = Godotenv("password_log")

	for _, host := range hosts {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s application_name=%s sslmode=%s", host, port, user, password, dbname, applicationName, sslMode)

		db, err := sql.Open("postgres", psqlInfo)
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(1)
		db.SetConnMaxLifetime(1 * time.Minute)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			time.Sleep(2 * time.Second)
			return ConnectionBDPostgreSQLLog(applicationName, sslMode, readOnly)
		}

		var notReadOnly bool
		// Quando a função pg_is_in_recovery() é chamada e retorna FALSE quer dizer que é uma instância de ESCRITA,
		// se retornar TRUE quer dizer que é somente de LEITURA.
		err = db.QueryRow(`SELECT pg_is_in_recovery();`).Scan(&notReadOnly)
		if err == nil && notReadOnly == readOnly {
			return db
		}
	}

	return db
}

/*
ConnectionBDPostgreSQLORM o antigo nome era: ConexaoBDORM
*/
func ConnectionBDPostgreSQLORM() (DB *gorm.DB) {
	var host, dbname, user, password string
	var port int

	host = Godotenv("host")
	dbname = Godotenv("dbname")
	port, _ = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
