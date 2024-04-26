//utils/database.go

package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func ConnectDB() (*gorm.DB, error) {
	// Replace these values with your PostgreSQL connection details
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}


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
		return nil, err
	}

	// Return the database connection
	return db, nil
}
