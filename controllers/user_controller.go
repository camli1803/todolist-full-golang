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
	user.PassWord = userInfo.PassWord
	err := models.Create(user)
	return err
}

func SignIn(userInfo models.UserInfoSignIn) (string, error) {
	user, err := models.TakeUserByEmail(userInfo.Email)
	if err != nil {
		err = fmt.Errorf("email is incorrect")
		return "", err
	}
	if userInfo.PassWord != user.PassWord {
		err = fmt.Errorf("password is incorrect")
		return "", err
	}

	// create token
	token, err := jwt.Create(user.ID, user.Email)
	return token, err
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
	user.PassWord = changePasswordReq.PassWord
	err := models.UpdateUserByID(id, user)
	return err
}

func DeleteUserByID(id uint64) error {
	err := models.DeleteUserByID(id)
	return err
}

func convertUserEntityToUserResponse(user models.User) (userResponse models.UserResponse) {
	userResponse.ID = user.ID
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
