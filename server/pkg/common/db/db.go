package db

import (
	"log"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	// open a db connection
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

	// create the tables if they doesn't exist
    db.AutoMigrate(&models.Book{})

    return db
}