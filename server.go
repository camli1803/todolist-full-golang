package main

import (
	"net/http"
	"todolist/database"
	"todolist/migrations"

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

	// api
	e.GET("/", homepage)

	e.Logger.Fatal(e.Start(":1710"))
}

func homepage(c echo.Context) error {
	return c.JSON(http.StatusOK, "Todolist")
}
