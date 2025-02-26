package controllers

import (
	"community/back/global"
	"community/back/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateArticles(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.AutoMigrate(&models.Article{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, article)
}

func GetArticles(c *gin.Context) {
	var articles []models.Article
	if err := global.Db.Find(&articles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, articles)
}

func GetArticlesById(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, article)
}
