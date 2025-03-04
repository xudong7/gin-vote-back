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

func DeleteFormsById(c *gin.Context) {
	var form models.Form
	id := c.Param("id")

	// 软删除

	if err := global.Db.Where("parent_id = ?", id).Delete(&models.Option{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.First(&form, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Delete(&form).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, form)
}

func UpdateFormsById(c *gin.Context) {
	// 1. 得到表单信息
	var form models.Form
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := form.ID

	// 2. 得到用户信息
	userId := c.Param("id")
	var user models.User
	if err := global.Db.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 3. 判断formList是否包含该表单ID
	if !contains(user.FormList, id) {
		user.FormList = append(user.FormList, id)
		// Use Save instead of Updates for JSON fields
		if err := global.Db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "投票失败"})
			return
		}
	}

	// Update form information
	if err := global.Db.Model(&form).Where("id = ?", id).Updates(&form).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Update options
	for i := range form.OptionList {
		if err := global.Db.Model(&form.OptionList[i]).Where("id = ?", form.OptionList[i].ID).Updates(&form.OptionList[i]).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, form)
}

// 辅助函数：检查元素是否存在
func contains(slice []uint, item uint) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
