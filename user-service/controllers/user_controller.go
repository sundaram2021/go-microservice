// controllers/user_controller.go
package controllers

import (
	"net/http"
	"user-service/models"
	"user-service/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	// Register route with raw SQL queries
	router.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the username already exists
		var count int64
		db.Raw("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&count)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = hashedPassword

		// Insert the user using raw SQL query
		result := db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, user.Password, user.Role)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
	})

	// Login route with raw SQL queries
	router.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		db.Raw("SELECT * FROM users WHERE username = ?", req.Username).Scan(&user)
		if user.Username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Compare the hashed password
		if err := utils.CheckPasswordHash(req.Password, user.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate a JWT token
		token, err := utils.GenerateToken(user.Username, user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	router.GET("/profile", utils.AuthMiddleware(), func(c *gin.Context) {
		username := c.MustGet("username").(string) // Retrieve the username from the token
		var user models.User

		// Fetch the user details from the database using the username
		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})
}
