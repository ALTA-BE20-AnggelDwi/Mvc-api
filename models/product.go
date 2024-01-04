package models

import (
	"gorm.io/gorm"
)

// struct product gorm model
type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	UserID      uint   `json:"user_id" form:"user_id" gorm:"index"`
	Description string `json:"description" form:"description"`
	User        User   `gorm:"foreignKey:UserID"`
}
