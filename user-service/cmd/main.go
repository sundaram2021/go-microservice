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

	
	controllers.InitializeRoutes(router, db)

	log.Fatal(router.Run(":8081"))
}
