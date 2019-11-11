package main

import (
	"fmt"
	"net/http"

	"github.com/cheapeone/goland/api"
	"github.com/cheapeone/goland/database"
	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	db := database.Connect()

	api := api.NewApi(db)
	apiRouter := app.Group("/api/v1")
	api.Register(apiRouter)

	fmt.Println("Running server...")
	app.Logger.Fatal(app.Start(":3000"))
}
