package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql(dsn string) {
	if db != nil {
		return
	}
	db, err = sql.Open("mysql", dsn)
}
