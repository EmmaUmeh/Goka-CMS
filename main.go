// main.go

package main

import (
    "log"
    "net/http"
    "os"

    "github.com/EmmaUmeh/Goka-CMS/utils"
    "github.com/EmmaUmeh/Goka-CMS/routers"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "github.com/rs/cors"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("error loading .env file: %v", err)
    }

    // Connect to the database
    db, err := utils.ConnectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

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
    log.Printf("Server listening on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}
