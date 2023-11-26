package config

import (
	"database/sql"

	"github.com/AlejandroJorge/url-shortener-go/util"
	_ "github.com/mattn/go-sqlite3"
)

func NewDatabaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "data.db")
	util.PanicIfError(err)

	return db
}
