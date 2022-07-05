package main

import (
	"net/http"
	"todolist/database"
	"todolist/migrations"
	"todolist/views"

	todolist_validator "todolist/validator"

	"todolist/middleware"

	goplayground_validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	DB, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	err = migrations.MigrateDB(DB)
	if err != nil {
		return
	}
	e := echo.New()
	e.Validator = &todolist_validator.CustomValidator{Validator: goplayground_validator.New()}

	// api
	e.GET("/", homepage)

	userApi := e.Group("/users")
	{
		//public api
		userApi.POST("/signUp", views.SignUp)
		userApi.POST("/signIn", views.SignIn)

		// private api
		userApi.GET("/:id", views.GetUserByID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		userApi.PATCH("/:id", views.UpdateUserByID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		userApi.PATCH("/:id/changePassword", views.ChangePasswordByID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		userApi.DELETE("/:id", views.DeleteUserByID, middleware.CheckAuthentication, middleware.CheckAuthorization)

		// todo api
		todoApi := userApi.Group("/:id/todos")
		todoApi.GET("", views.GetAllTodosByUserID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		todoApi.POST("", views.CreateTodosByUserID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		todoApi.GET("/completed", views.GetAllTodosCompletedByUserID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		todoApi.GET("/unfinished", views.GetAllTodosUnFinishedByUserID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		todoApi.GET("/:todo_id", views.GetATodoByUserIDTodoID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		todoApi.PATCH("/:todo_id", views.UpdateATodoByUserIDTodoID, middleware.CheckAuthentication, middleware.CheckAuthorization)
		todoApi.DELETE("/:todo_id", views.DeleteATodoByUserIDTodoID, middleware.CheckAuthentication, middleware.CheckAuthorization)
	}

	e.Logger.Fatal(e.Start(":1710"))
}

func homepage(c echo.Context) error {
	return c.JSON(http.StatusOK, "Todolist")
}
