// cmd/main.go
package main

import (
	"log"
	"product-service/config"
	"product-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize the database
	db := config.SetupDatabase()

	// Initialize product routes
	controllers.InitializeProductRoutes(router, db)

	// Start the server
	log.Fatal(router.Run(":8082"))
}
