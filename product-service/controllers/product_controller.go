package controllers

import (
	"net/http"
	"product-service/models"
	"product-service/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeProductRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/products", func(c *gin.Context) {
		var products []models.Product
		if err := db.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"products": products})
	})

	router.GET("/products/:id", func(c *gin.Context) {
		var product models.Product
		if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"product": product})
	})

	protected := router.Group("/")
	protected.Use(utils.AuthMiddleware())
	protected.POST("/products", func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
	})

	protected.PUT("/products/:id", func(c *gin.Context) {
		var product models.Product
		if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": product})
	})

	protected.DELETE("/products/:id", func(c *gin.Context) {
		var product models.Product
		if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if err := db.Delete(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	})
}
