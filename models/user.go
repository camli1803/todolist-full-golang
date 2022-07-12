package models

import (
	"time"
	"todolist/database"

	"gorm.io/gorm"
)

// entity
type User struct {
	gorm.Model
	UserName string `gorm:"column:username;not null"`
	Email    string `gorm:"column:email;not null;unique"`
	Password string `gorm:"column:password;not null"`
	Todos    []Todo
}

// dtos
type UserInfoSignUp struct {
	UserName string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserInfoSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	ID        uint64    `json:"id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	UserName string `json:"username" validate:"required,max=100"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" validate:"required,min=8"`
}

// Create new user
func Create(user User) error {
	result := database.DB.Create(&user)
	return result.Error
}

// take user by email
func TakeUserByEmail(email string) (User, error) {
	var user User
	result := database.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

// take user by id
func TakeUserByID(id uint64) (User, error) {
	var user User
	result := database.DB.Take(&user, id)
	return user, result.Error
}

// update user by id
func UpdateUserByID(id uint64, user User) error {
	result := database.DB.Model(&user).Where("id = ?", id).Updates(&user)
	return result.Error
}

// delete user by id
func DeleteUserByID(id uint64) error {
	var user User
	result := database.DB.Delete(&user, id)
	return result.Error
}

// todo api
func GetAllTodosByUserID(userID uint64) (todos []Todo, err error) {
	var user User
	result := database.DB.Preload("Todos").Find(&user, "id = ?", userID)
	return user.Todos, result.Error
}
