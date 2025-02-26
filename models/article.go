package models

type Article struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
