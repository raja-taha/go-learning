package routes

import (
	"gin-api/controllers"
	"gin-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/users", controllers.GetUsers)

	router.POST(
		"/users",
		middleware.AuthMiddleware(),
		controllers.CreateUser,
	)
}
