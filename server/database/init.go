package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	connectionString := os.Getenv("NEON_DATABASE_URL")
	if connectionString == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto migrate schemas
	err = db.AutoMigrate(&User{}, &SessionLog{})
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	DB = db
	log.Println("Database connected & migrated successfully ✅")
}
