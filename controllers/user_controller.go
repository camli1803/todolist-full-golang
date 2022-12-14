package controllers

import (
	"fmt"

	"todolist/models"

	jwt "todolist/auth"
)

func SignUp(userInfo models.UserInfoSignUp) error {
	var user models.User
	user.UserName = userInfo.UserName
	user.Email = userInfo.Email
	user.Password = userInfo.Password
	err := models.Create(user)
	return err
}

func SignIn(userInfo models.UserInfoSignIn) (string, error) {
	user, err := models.TakeUserByEmail(userInfo.Email)
	if err != nil {
		err = fmt.Errorf("your email is incorrect")
		return "", err
	}
	if userInfo.Password != user.Password {
		err = fmt.Errorf("your password is incorrect")
		return "", err
	}

	// create token
	token, err := jwt.Create(user.ID, user.Email)
	return token, err
}

func ResetForgotPassword(userResetPassword models.UserResetPasswordRequest) error {
	user, err := models.TakeUserByEmail(userResetPassword.Email)
	if err != nil {
		return err
	}

	user.Password = userResetPassword.NewPassword
	return models.UpdateUserByID(uint64(user.ID), user)
}

func GetUserByID(id uint64) (models.UserResponse, error) {
	user, err := models.TakeUserByID(id)
	userRes := convertUserEntityToUserResponse(user)
	return userRes, err
}

func UpdateUserByID(id uint64, userUpdateRequest models.UserUpdateRequest) error {
	var user models.User
	user.UserName = userUpdateRequest.UserName
	err := models.UpdateUserByID(id, user)
	return err
}

func ChangePasswordByID(id uint64, changePasswordReq models.ChangePasswordRequest) error {
	var user models.User
	user, err := models.TakeUserByID(id)
	if err != nil {
		return err
	}
	if changePasswordReq.CurrentPassword != user.Password {
		return fmt.Errorf("your current password is incorrect")
	}

	if changePasswordReq.NewPassword == changePasswordReq.CurrentPassword {
		return fmt.Errorf("your password has not been changed")
	}

	user.Password = changePasswordReq.NewPassword

	return models.UpdateUserByID(id, user)
}

func DeleteUserByID(id uint64) error {
	err := models.DeleteUserByID(id)
	return err
}

func convertUserEntityToUserResponse(user models.User) (userResponse models.UserResponse) {
	userResponse.ID = uint64(user.ID)
	userResponse.UserName = user.UserName
	userResponse.Email = user.Email
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt
	return userResponse
}

// Todo Api
func GetAllTodosByUserID(userID uint64) (todosRes []models.TodoResponse, err error) {
	todos, err := models.GetAllTodosByUserID(userID)
	for _, todo := range todos {
		todosRes = append(todosRes, convertTodoEntityToTodoResponse(todo))
	}
	return todosRes, err
}