package controllers

import (
	"community/back/global"
	"community/back/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateForms(c *gin.Context) {
	var form models.Form

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := range form.OptionList {
		form.OptionList[i].ParentID = form.ID
	}

	// 打印表单信息
	fmt.Printf("创建表单 - ID: %v, 标题: %v\n", form.ID, form.Title)
	fmt.Println("表单选项:")
	for i, option := range form.OptionList {
		fmt.Printf("  选项 %d: %+v\n", i+1, option)
	}

	if err := global.Db.AutoMigrate(&models.Form{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.AutoMigrate(&models.Option{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Create(&form).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, form)
}

func GetForms(c *gin.Context) {
	var forms []models.Form

	// 使用Preload加载关联的选项列表
	if err := global.Db.Preload("OptionList").Find(&forms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 打印所有表单信息
	fmt.Println("获取所有表单:")
	for i, form := range forms {
		fmt.Printf("表单 %d - ID: %v, 标题: %v\n", i+1, form.ID, form.Title)
		fmt.Printf("选项数量: %d\n", len(form.OptionList))
	}

	c.JSON(http.StatusOK, forms)
}

func GetFormsById(c *gin.Context) {
	var form models.Form
	id := c.Param("id")

	// 使用Preload加载关联的选项列表
	if err := global.Db.Preload("OptionList").First(&form, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 打印单个表单信息
	fmt.Printf("获取表单 ID: %v\n", id)
	fmt.Printf("表单详情: %+v\n", form)
	fmt.Printf("选项数量: %d\n", len(form.OptionList))
	for i, option := range form.OptionList {
		fmt.Printf("选项 %d: %+v\n", i+1, option)
	}

	c.JSON(http.StatusOK, form)
}
