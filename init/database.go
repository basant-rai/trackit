package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectToDB initializes the database connection
func ConnectToDB() {
	dsn := os.Getenv("DB_URL") // Make sure to set this environment variable
	if dsn == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	log.Println("Database connection established")
}
