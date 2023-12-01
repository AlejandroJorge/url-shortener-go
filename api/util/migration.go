package util

import (
	"github.com/AlejandroJorge/url-shortener-go/model"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(&model.URL{})
}
