// cmd/main.go
package main

import (
	"log"
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes, no need to pass db connection anymore
	controllers.InitializeRoutes(router)

	// Start the server
	log.Fatal(router.Run(":8081"))
}
