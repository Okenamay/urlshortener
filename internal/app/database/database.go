package database

import (
	"log"

	"github.com/Okenamay/urlshortener/internal/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("shortener.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	DB = db

	DB.AutoMigrate(&models.URL{})
}
