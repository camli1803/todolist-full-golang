package views

import (
	"net/http"
	"todolist/controllers"
	"todolist/models"

	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SignUp(c echo.Context) error {
	var userInfo models.UserInfoSignUp
	err := c.Bind(&userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = c.Validate(userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = controllers.SignUp(userInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "success")
}

func SignIn(c echo.Context) error {
	var userInfo models.UserInfoSignIn
	err := c.Bind(&userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = c.Validate(userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	token, err := controllers.SignIn(userInfo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, token)
}

// get user by id
func GetUserByID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := controllers.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userRes)
}

// update user by id
func UpdateUserByID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var userUpdateRequest models.UserUpdateRequest
	err = c.Bind(&userUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = c.Validate(userUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = controllers.UpdateUserByID(userID, userUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}

// Change Password
func ChangePasswordByID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var changePasswordReq models.ChangePasswordRequest
	err = c.Bind(&changePasswordReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(changePasswordReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.ChangePasswordByID(userID, changePasswordReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}

// Delete user
func DeleteUserByID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.DeleteUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "success")
}

// todo api
func GetAllTodosByUserID(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todosRes, err := controllers.GetAllTodosByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todosRes)
}
