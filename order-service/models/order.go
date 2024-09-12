// models/order.go
package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      uint    `json:"user_id"`
	ProductID   uint    `json:"product_id"`
	Quantity    int     `json:"quantity"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"` // e.g., "Pending", "Shipped", "Delivered"
}
