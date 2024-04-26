package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EmmaUmeh/Goka-CMS/models"
	"github.com/EmmaUmeh/Goka-CMS/routers"
	"github.com/EmmaUmeh/Goka-CMS/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func main() {
	// Define the length of the random string (you can adjust this according to your needs)
	length := 32

	// Generate a random string
	secretKey, err := generateRandomString(length)
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}

	fmt.Println("Generated Secret Key:", secretKey)

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// Connect to the database
	db, err := utils.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&models.User{}) // This will create the users and tokens tables if they don't exist

	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routers.SetupRoutes(router, db)

	// Enable CORS
	handler := cors.Default().Handler(router)

	// Get the port from environment variables or use default 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Start the server
	log.Printf("Goka Server listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
