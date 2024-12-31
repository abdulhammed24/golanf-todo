package database

import (
	"fmt"
	"log"
	"os"
	"todo-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database
func Init() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Build DSN (Data Source Name)
    dsn := os.Getenv("DB_URL")

    // Connect to the database
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // AutoMigrate the models
    if err := DB.AutoMigrate(&models.Todo{}); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    fmt.Println("Database connected and migrated successfully")
}
