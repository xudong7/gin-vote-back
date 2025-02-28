package controllers

import (
	"community/back/global"
	"community/back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	if err := global.Db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := global.Db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
