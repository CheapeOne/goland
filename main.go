package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/cheapeone/goland/api/resolvers"
	"github.com/cheapeone/goland/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	app := echo.New()

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	playground := handler.Playground("GraphQL playground", "/query")
	app.GET("/playground", echo.WrapHandler(playground))

	db := database.Connect()
	graphql := resolvers.GraphqlHandler(db)
	app.POST("/query", echo.WrapHandler(graphql))

	fmt.Println("Running server...")
	app.Logger.Fatal(app.Start(":3000"))
}
