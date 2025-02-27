package models

import "gorm.io/gorm"

type Form struct {
	gorm.Model
	Title      string   `json:"title" binding:"required"`                   // 标题
	Type       int      `json:"type" gorm:"default:1" binding:"required"`   // 类型 1: 单选 2: 多选
	Status     int      `json:"status" gorm:"default:1" binding:"required"` // 状态 1: 启用 2: 禁用
	OptionList []Option `json:"optionList" gorm:"foreignKey:ParentID"`      // 选项列表
}
