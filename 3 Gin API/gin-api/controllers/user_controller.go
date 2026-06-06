package controllers

import (
	"net/http"

	"gin-api/models"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := services.GetUsers()

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdUser := services.CreateUser(user)

	c.JSON(http.StatusCreated, createdUser)
}
