// cmd/main.go
package main

import (
	"log"
	"user-service/config"
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := config.SetupDatabase()

	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	// Initialize routes, no need to pass db connection anymore
	controllers.InitializeRoutes(router, db)

	// Start the server
	log.Fatal(router.Run(":8081"))
}
