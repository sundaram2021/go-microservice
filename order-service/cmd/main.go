package main

import (
	"log"
	"order-service/config"
	"order-service/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := config.SetupDatabase()

	controller.InitializeOrderRoutes(router, db)

	log.Fatal(router.Run(":8083"))
}
