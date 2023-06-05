package util

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {

}

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Log.Fatal("Failed to connect to the database.")
	}
}
