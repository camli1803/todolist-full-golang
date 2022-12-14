package main

import (
	"net/http"
	"todolist/database"
	"todolist/middleware"
	"todolist/migrations"
	todolist_validator "todolist/validator"
	"todolist/views"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load("tmp/.env")
	if err != nil {
		panic("load file .env failed")
	}

	DB, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	err = migrations.MigrateDB(DB)
	if err != nil {
		return
	}
	e := echo.New()

	var customValidator todolist_validator.CustomValidator
	validate := customValidator.New()
	e.Validator = &todolist_validator.CustomValidator{Validator: validate}

	// api
	e.GET("/", homepage)

	userApi := e.Group("/auth")
	{
		//public api
		userApi.POST("/signup", views.SignUp)
		userApi.POST("/signin", views.SignIn)
		userApi.PATCH("/reset_forgot_password", views.ResetForgotPassword)

		// private api
		userApiPrivate := userApi.Group("/users/:id")
		userApiPrivate.Use(middleware.CheckAuthentication, middleware.CheckAuthorization)
		userApiPrivate.GET("", views.GetUserByID)
		userApiPrivate.PATCH("", views.UpdateUserByID)
		userApiPrivate.PATCH("/change_password", views.ChangePasswordByID)
		userApiPrivate.DELETE("", views.DeleteUserByID)

		// todo api
		todoApi := userApiPrivate.Group("/todos")
		todoApi.GET("", views.GetAllTodosByUserID)
		todoApi.POST("", views.CreateTodosByUserID)
		todoApi.GET("/show_by_done", views.GetAllTodosByDoneUserID)
		todoApi.GET("/:todo_id", views.GetATodoByUserIDTodoID)
		todoApi.PATCH("/:todo_id", views.UpdateATodoByUserIDTodoID)
		todoApi.DELETE("/:todo_id", views.DeleteATodoByUserIDTodoID)
	}

	e.Logger.Fatal(e.Start(":1710"))
}

func homepage(c echo.Context) error {
	return c.JSON(http.StatusOK, "Todolist")
}
