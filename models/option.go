package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	ParentID uint   `json:"parentID" binding:"required"` // 父ID
	Content  string `json:"content" binding:"required"`  // 内容
	Votes    int    `json:"votes"`                       // 投票数
}
