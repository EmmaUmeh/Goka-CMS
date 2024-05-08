// api/middleware/middleware.go
package middleware

import (
    "net/http"
    "github.com/golang-jwt/jwt/v5"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract the JWT token from the Authorization header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
            c.Abort()
            return
        }

        // Validate the JWT token and extract the user ID
        token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
            // Replace with your secret key
            return []byte("your_secret_key"), nil
        })

        if err!= nil ||!token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Assert the Claims to jwt.MapClaims
        claims, ok := token.Claims.(jwt.MapClaims)
        if!ok ||!token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Now you can safely access the claims
        userID, ok := claims["userId"].(string)
        if!ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Attach the user ID to the request context
        c.Set("userID", userID)

        c.Next()
    }
}
