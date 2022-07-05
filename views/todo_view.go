package views

import (
	"net/http"
	"strconv"
	"todolist/controllers"
	"todolist/models"

	"github.com/labstack/echo/v4"
)

func CreateTodosByUserID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var todoCreateRequest models.TodoCreateRequest
	err = c.Bind(&todoCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(todoCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.CreateTodos(userID, todoCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}

func GetAllTodosCompletedByUserID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todosRes, err := controllers.GetAllTodosCompletedByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todosRes)
}

func GetAllTodosUnFinishedByUserID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todosRes, err := controllers.GetAllTodosUnFinishedByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todosRes)
}

func GetATodoByUserIDTodoID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todoID, err := strconv.ParseUint(c.Param("todo_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoRes, err := controllers.GetATodoByUserIDTodoID(userID, todoID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todoRes)

}

func UpdateATodoByUserIDTodoID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todoID, err := strconv.ParseUint(c.Param("todo_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var todoUpdateRequest models.TodoUpdateRequest
	err = c.Bind(&todoUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.UpdateATodoByUserIDTodoID(userID, todoID, todoUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}

func DeleteATodoByUserIDTodoID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todoID, err := strconv.ParseUint(c.Param("todo_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.DeleteATodoByUserIDTodoID(userID, todoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}
