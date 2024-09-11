// utils/jwt_middleware.go
package utils

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// utils/jwt_middleware.go
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		log.Println("Authorization header:", tokenString) // <-- Add this to log the token

		// Check if token has the "Bearer" prefix
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid token format") // <-- Log invalid format
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString = parts[1]

		// Validate the token
		claims, err := ValidateToken(tokenString)
		if err != nil {
			log.Println("Invalid or expired token:", err) // <-- Log token validation issues
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		log.Println("Token is valid, claims:", claims) // <-- Log valid token claims

		// Store the username in the context
		c.Set("username", claims.Username)
		c.Next()
	}
}
