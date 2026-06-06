package main

import (
	"gin-api/config"
	"gin-api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	godotenv.Load()

	routes.RegisterRoutes(router)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
