package controllers

import (
	"todolist/models"
)

func CreateTodos(userID uint64, todoCreateRequest models.TodoCreateRequest) error {
	var todo models.Todo
	todo.Content = todoCreateRequest.Content
	todo.Done = todoCreateRequest.Done
	todo.UserID = userID
	return models.CreateTodos(todo)
}

func GetAllTodosByDoneUserID(userID uint64, done bool) (todosRes []models.TodoResponse, err error) {
	todos, err := models.GetAllTodosByDoneUserID(userID, done)
	if err != nil {
		return nil, err
	}
	for _, todo := range todos {
		todosRes = append(todosRes, convertTodoEntityToTodoResponse(todo))
	}
	return todosRes, err
}

func GetATodoByUserIDTodoID(userID, todoID uint64) (todoRes models.TodoResponse, err error) {
	todo, err := models.GetATodoByUserIDTodoID(userID, todoID)
	todoRes = convertTodoEntityToTodoResponse(todo)
	return todoRes, err
}

func UpdateATodoByUserIDTodoID(userID uint64, todoID uint64, todoUpdateRequest models.TodoUpdateRequest) error {
	var todo models.Todo
	todo.Content = todoUpdateRequest.Content
	todo.Done = todoUpdateRequest.Done
	return models.UpdateATodoByUserIDTodoID(userID, todoID, todo)
}

func DeleteATodoByUserIDTodoID(userID uint64, todoID uint64) error {
	return models.DeleteATodoByUserIDTodoID(userID, todoID)
}

func convertTodoEntityToTodoResponse(todo models.Todo) (todoResponse models.TodoResponse) {
	todoResponse.ID = todo.ID
	todoResponse.Content = todo.Content
	todoResponse.Done = todo.Done
	todoResponse.CreatedAt = todo.CreatedAt
	todoResponse.UpdatedAt = todo.UpdatedAt
	return todoResponse
}
