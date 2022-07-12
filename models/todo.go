package models

import (
	"todolist/database"

	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `gorm:"column:content;not null"`
	Done    bool   `gorm:"column:done"`
	UserID  uint64 `gorm:"column:userid;not null"`
}

type TodoCreateRequest struct {
	Content string `json:"content" validate:"required"`
	Done    bool   `json:"done"`
}

type TodoUpdateRequest struct {
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type TodoResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateTodos(todo Todo) error {
	result := database.DB.Create(&todo)
	return result.Error
}

func GetAllTodosByDoneUserID(userID uint64, done bool) (todos []Todo, err error) {
	result := database.DB.Where("userid = ? AND done = ?", userID, done).Find(&todos)
	return todos, result.Error
}

func GetATodoByUserIDTodoID(userID, todoID uint64) (todo Todo, err error) {
	result := database.DB.Where("userid = ? AND id = ?", userID, todoID).First(&todo)
	return todo, result.Error
}

func UpdateATodoByUserIDTodoID(userID uint64, todoID uint64, todo Todo) error {
	result := database.DB.Model(&todo).Where("userid = ? AND id = ?", userID, todoID).Updates(&todo)
	return result.Error
}

func DeleteATodoByUserIDTodoID(userID uint64, todoID uint64) error {
	var todo Todo
	result := database.DB.Where("userid = ? AND id = ?", userID, todoID).Delete(&todo)
	return result.Error
}
