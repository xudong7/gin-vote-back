package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" gorm:"default:'user'"` // user, admin
	FormList []int  `json:"formList" gorm:"type:json"`  // 需要投票的form的id列表
}
