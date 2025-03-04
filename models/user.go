package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" gorm:"default:'user'"`                 // user, admin
	FormList []uint `json:"form_list" gorm:"type:json;serializer:json"` // 存储已经投过票的表单ID
}
