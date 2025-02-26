package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"theme" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" gorm:"default:'user'"` // user, admin
}
