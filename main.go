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
 // Import your middleware package
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	SecretKeyLength = 32 // Length of the secret key
)

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err!= nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func main() {
	// Generate a random secret key
	secretKey, err := generateRandomString(SecretKeyLength)
	if err!= nil {
		log.Fatalf("Error generating random string: %v", err)
	}

	fmt.Println("Generated Secret Key:", secretKey)

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Load.env file
	err = godotenv.Load()
	if err!= nil {
		log.Fatalf("Error loading.env file: %v", err)
	}

	// Connect to the database
	db, err := utils.ConnectDB()
	if err!= nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Auto migrate database models
	db.AutoMigrate(&models.User{}) // This will create the users and tokens tables if they don't exist
	db.AutoMigrate(&models.Task{})

	// Create a new Gin router
	router := gin.Default()

	// Apply your middleware to the router
	// router.Use(api.AuthMiddleware()) // Assuming AuthMiddleware is exported from your middleware package

	// Setup routes
	routers.AuthRoutes(router, db)
	routers.TaskRoutes(router, db)

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
