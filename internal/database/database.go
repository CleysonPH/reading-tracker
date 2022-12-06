package database

import "database/sql"

var (
	db  *sql.DB
	err error
)

func GetDB() (*sql.DB, error) {
	return db, err
}
