package db

import (
	"book/pkg/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
