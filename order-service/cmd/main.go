// cmd/main.go
package main

import (
	"log"
	"order-service/config"
	"order-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize the database
	db := config.SetupDatabase()

	// Initialize order routes
	controller.InitializeOrderRoutes(router, db)

	// Start the server
	log.Fatal(router.Run(":8083"))
}
