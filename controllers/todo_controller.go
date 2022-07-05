package controllers

import (
	"todolist/models"
)

func CreateTodos(userID uint64, todoCreateRequest models.TodoCreateRequest) error {
	var todo models.Todo
	todo.Content = todoCreateRequest.Content
	todo.Done = todoCreateRequest.Done
	todo.UserID = userID

	err := models.CreateTodos(todo)
	return err
}

func GetAllTodosCompletedByUserID(userID uint64) (todosRes []models.TodoResponse, err error) {
	todos, err := models.GetAllTodosCompletedByUserID(userID)
	for _, todo := range todos {
		todosRes = append(todosRes, convertTodoEntityToTodoResponse(todo))
	}
	return todosRes, err
}

func GetAllTodosUnFinishedByUserID(userID uint64) (todosRes []models.TodoResponse, err error) {
	todos, err := models.GetAllTodosUnFinishedByUserID(userID)
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
	err := models.UpdateATodoByUserIDTodoID(userID, todoID, todo)
	return err
}

func DeleteATodoByUserIDTodoID(userID uint64, todoID uint64) error {
	err := models.DeleteATodoByUserIDTodoID(userID, todoID)
	return err
}

func convertTodoEntityToTodoResponse(todo models.Todo) (todoResponse models.TodoResponse) {
	todoResponse.ID = todo.ID
	todoResponse.Content = todo.Content
	todoResponse.Done = todo.Done
	todoResponse.CreatedAt = todo.CreatedAt
	todoResponse.UpdatedAt = todo.UpdatedAt
	return todoResponse
}
