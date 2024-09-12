// controllers/order_controller.go
package controller

import (
	"encoding/json"
	"net/http"
	"order-service/models"
	"order-service/utils"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeOrderRoutes(router *gin.Engine, db *gorm.DB) {

	// Protected routes
	protected := router.Group("/")
	protected.Use(utils.AuthMiddleware()) // JWT Middleware to protect routes

	// Public route: View a specific order by ID
	router.GET("/orders/:id", func(c *gin.Context) {
		var order models.Order
		orderID := c.Param("id")
		if err := db.Where("id = ?", orderID).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"order": order})
	})

	// Public route: View all orders (Admin can view all, users can see their own orders)
	router.GET("/orders", func(c *gin.Context) {
		var orders []models.Order
		userID := c.Query("user_id")

		if userID != "" {
			userIDInt, _ := strconv.Atoi(userID)
			db.Where("user_id = ?", userIDInt).Find(&orders)
		} else {
			db.Find(&orders) // Admin case
		}

		c.JSON(http.StatusOK, gin.H{"orders": orders})
	})

	// Create an order (Authenticated route)
	protected.POST("/orders", func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Fetch product details from the product-service
		productID := order.ProductID
		productServiceURL := os.Getenv("PRODUCT_SERVICE_URL")
		productURL := productServiceURL +"/products/" + strconv.Itoa(int(productID))

		resp, err := http.Get(productURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product details"})
			return
		}
		defer resp.Body.Close()

		// Parse product details response
		var productResponse struct {
			Product struct {
				Price float64 `json:"price"`
			} `json:"product"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&productResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse product details"})
			return
		}

		// Calculate total amount based on product price and quantity
		order.TotalAmount = productResponse.Product.Price * float64(order.Quantity)

		// Save the order to the database
		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order": order})
	})

	// Update order status (Admin-only route)
	protected.PUT("/orders/:id/status", func(c *gin.Context) {
		role := c.GetString("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have the necessary permissions to update order status"})
			return
		}

		var order models.Order
		orderID := c.Param("id")

		// Fetch the order by ID
		if err := db.Where("id = ?", orderID).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		// Get the new status from the request body
		var req struct {
			Status string `json:"status" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update the order status
		order.Status = req.Status

		if err := db.Save(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully", "order": order})
	})
}
