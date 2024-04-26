package controllers

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/EmmaUmeh/Goka-CMS/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Signup(c *gin.Context, db *gorm.DB) {
	var user models.User


    log.Printf("user data: %+v", user)
	// Decode request body (assuming JSON format)
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding user data"})
		return
	}

	// Check if the user already exists in the database
	if models.UserExistsByEmail(db, user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with email already exists"})
		return
	}

	// If the user does not exist, proceed with signup logic
	if err := user.CreateUser(db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	users := []models.User{user}

	c.JSON(http.StatusOK, gin.H{"message": "Signup successful", "users": users})
}


// func LoginUser (context *gorm.Context, db *gorm.DB) {
    

// }