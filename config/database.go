package config

import (
	"github.com/AlejandroJorge/url-shortener-go/util"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	util.PanicIfError(err)

	return db
}
