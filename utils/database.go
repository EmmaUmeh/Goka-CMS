// utils/database.go

package utils

import (
    "fmt"
    "log"

    "github.com/jinzhu/gorm"
    "github.com/joho/godotenv"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "os"
)

func ConnectDB() (*gorm.DB, error) {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("error loading .env file: %w", err)
    }

    // Get database connection details from environment variables
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbName := os.Getenv("DB_NAME")
    dbPassword := os.Getenv("DB_PASSWORD")

    // Construct the connection string
    dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

    // Open a connection to the database
    db, err := gorm.Open("postgres", dbURI)
    if err != nil {
        log.Printf("Failed to connect to database: %v", err)
        return nil, err
    }

    // Log a success message when the database connection is established
    log.Println("Successfully connected to database")
	

    // Return the database connection
    return db, nil
}
