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
func ConnectionBDPostgreSQL(applicationName string) *sql.DB {
	var dbname, user, password string
	var port int
	var db *sql.DB

	hosts := strings.Split(Godotenv("host"), ";")
	dbname = Godotenv("dbname")
	port, _ = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	for _, host := range hosts {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable application_name=%s", host, port, user, password, dbname, applicationName)

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
			return ConnectionBDPostgreSQL(applicationName)
		}

		var tableName string
		// Buscamos uma tabela do banco de dados
		err = db.QueryRow(`SELECT table_name FROM information_schema.role_table_grants WHERE table_schema='public' order by table_name LIMIT 1;`).Scan(&tableName)
		if err == nil {
			// Executamos um update dentro de umsa transação e logo em seguida é dado rollback, caso não dê erro quer dizer que é uma instância de gravação.
			_, err = db.Exec(`BEGIN; UPDATE public.` + tableName + ` SET id = id WHERE TRUE AND id IN (SELECT id FROM public.` + tableName + ` ORDER BY id LIMIT 1); ROLLBACK;`)
			if err == nil {
				return db
			}
		}
	}

	return db
}

func ConnectionBDPostgreSQLRead(applicationName string) *sql.DB {
	var host, dbname, user, password string
	var port int

	host = Godotenv("host_read")
	dbname = Godotenv("dbname_read")
	port, _ = strconv.Atoi(Godotenv("port_banco_read"))
	user = Godotenv("user_read")
	password = Godotenv("password_read")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable application_name=%s", host, port, user, password, dbname, applicationName)

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
		return ConnectionBDPostgreSQL(applicationName)
	}

	return db
}

func ConnectionBDPostgreSQLWithSSL() *sql.DB {
	var host, dbname, user, password string
	var port int

	host = Godotenv("host")
	dbname = Godotenv("dbname")
	port, _ = strconv.Atoi(Godotenv("port_banco"))
	user = Godotenv("user")
	password = Godotenv("password")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full", host, port, user, password, dbname)

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
		return ConnectionBDPostgreSQLWithSSL()
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
