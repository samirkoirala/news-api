package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Auto-migrate the News model
	err = DB.AutoMigrate(&News{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
}

type News struct {
    ID      uint   `json:"id" gorm:"primaryKey"`
    Title   string `json:"title"`
    Content string `json:"content"`
}
