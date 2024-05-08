package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EmmaUmeh/Goka-CMS/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)
var SECRET_KEY string

func init() {
	err := godotenv.Load()
	if err!= nil {
		log.Fatalf("Error loading.env file: %v", err)
	}

	// Get database connection details from environment variables
	SECRET_KEY = os.Getenv("SECRET_KEY")
}

func Signup(c *gin.Context, db *gorm.DB) {
	var user models.User

	if err := c.BindJSON(&user); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding user data"})
		return
	}

	log.Printf("Received user data: %+v", user)
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	if UserExistsByEmail(db, user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with email already exists"})
		return
	}

	if err := CreateUser(db, &user); err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)
	userArray := []models.User{user}

	token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["userId"] = user.ID // Assuming you want to use the user's ID in the token
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    tokenString, err := token.SignedString([]byte(SECRET_KEY))
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Return the token in the response
    c.JSON(http.StatusOK, gin.H{
        "message": "Signup successful",
        "user":    userArray,
        "token":   tokenString,
    })
}

func Login(c *gin.Context, db *gorm.DB) {
    var user models.User

    if err := c.BindJSON(&user); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding credentials"})
        return
    }

    var foundUser models.User
    if err := db.Where("email =?", user.Email).First(&foundUser).Error; err!= nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Verify the password
    if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err!= nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate a new token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["userId"] = foundUser.ID // Use the user's ID, not email
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    tokenString, err := token.SignedString([]byte(SECRET_KEY))
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "User Login Successfully": foundUser,
        "token": tokenString, // Include the token in the response
    })
}


// Placeholder function for password verification
func verifyPassword(email, password string) bool {
	// Implement your password verification logic here
	return true
}

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func UserExistsByEmail(db *gorm.DB, email string) bool {
	var user models.User
	if err := db.Where("email =?", email).First(&user).Error; err!= nil {
		return false
	}
	return true
}
