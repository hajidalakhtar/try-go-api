package app

import (
	"GolangRestfulApi/helper"
	"database/sql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "DATABASE_CONNECTION")
	helper.PanicIfError(err)

	return db
}
