// controllers/user_controller.go
package controllers

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/models"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Register route
	router.POST("/register", func(c *gin.Context) {
		// Create a new database connection for this request
		db, err := config.SetupDatabase()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		// Deallocate any existing prepared statements to prevent conflicts
		db.Exec("DEALLOCATE ALL")

		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the user already exists
		if err := db.Where("username = ?", user.Username).First(&user).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		// Save user to database
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	// Login route
	router.POST("/login", func(c *gin.Context) {
		// Create a new database connection for this request
		db, err := config.SetupDatabase()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := db.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate JWT token for the user
		token, err := utils.GenerateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// controllers/user_controller.go
	router.GET("/profile", utils.AuthMiddleware(), func(c *gin.Context) {
		// Create a new database connection for this request
		db, err := config.SetupDatabase()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
			return
		}

		// Get the username from the context (set by the JWT middleware)
		username := c.MustGet("username").(string)

		// Log the username extracted from the JWT token
		log.Println("Username from JWT:", username)

		// Fetch user from the database using the username
		var user models.User
		if err := db.Where("LOWER(username) = LOWER(?)", username).First(&user).Error; err != nil {
			log.Println("User not found in database for username:", username)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// If the user is found, return the username
		c.JSON(http.StatusOK, gin.H{
			"message":  "Welcome " + user.Username,
			"username": user.Username,
		})
	})

}
