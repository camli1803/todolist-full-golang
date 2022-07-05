package models

import (
	"gorm.io/gorm"
)

// entity
type User struct {
	gorm.Model
	UserName string `gorm:"column:username;not null"`
	Email    string `gorm:"column:email;not null;unique"`
	PassWord string `gorm:"column:password;not null"`
	Todos    []Todo
}
