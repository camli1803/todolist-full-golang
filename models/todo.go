package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `gorm:"column:content;not null"`
	Done    bool   `gorm:"column:done"`
	UserID  uint64 `gorm:"column:userid;not null"`
}
