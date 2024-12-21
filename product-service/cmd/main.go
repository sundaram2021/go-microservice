package main

import (
	"log"
	"product-service/config"
	"product-service/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := config.SetupDatabase()

	controllers.InitializeProductRoutes(router, db)

	log.Fatal(router.Run(":8082"))
}
