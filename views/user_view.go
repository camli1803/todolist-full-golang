package views

import (
	"fmt"
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
	return c.JSON(http.StatusCreated, "your account has been registered")
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

func ResetForgotPassword(c echo.Context) error {
	var userResetPassword models.UserResetPasswordRequest
	err := c.Bind(&userResetPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(userResetPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = controllers.ResetForgotPassword(userResetPassword)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = fmt.Errorf("your email is invalid")
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "your password has been reset")
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
	return c.JSON(http.StatusOK, "your account has been changed")
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
	return c.JSON(http.StatusOK, "your password has been changed")
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
	return c.JSON(http.StatusOK, "your account has been deleted")
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