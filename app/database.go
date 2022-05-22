package app

import (
	"database/sql"

	"github.com/rzldimam28/simple-notes/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:__@tcp(localhost:3306)/simple_notes?parseTime=true")
	helper.PanicIfError(err)
	return db
}
