package goutils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

/**
USE mysql;
CREATE USER 'root'@'localhost' IDENTIFIED BY '';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost';
UPDATE user SET plugin='auth_socket' WHERE User='root';
FLUSH PRIVILEGES;
exit;
*/
func ConexaoBDMySQL() *sql.DB {
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
