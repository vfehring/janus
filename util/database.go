package util

import (
	"os"
	"vfehring/janus/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	DB.AutoMigrate(&models.User{})
}

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Fatal("Failed to connect to the database.")
	}
	Log.Infof("Database connection successful.")
}
