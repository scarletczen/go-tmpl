package main

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/abhinavthapa1998/go-tmpl/handler"
)

type DB struct{}

func main() {
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "abhinav@gmail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
