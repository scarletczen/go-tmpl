package main

import (

  "github.com/abhinavthapa1998/go-tmpl/handler"
  "github.com/labstack/echo/v4"
)

type DB struct{}

func main(){
  app := echo.New()
  userHandler := handler.UserHandler{}
  app.GET("/user",userHandler.HandleUserShow)
  app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc{
  return func(c echo.Context) error{
    c.Set("user","abhinav@gmail.com")
    return next(c)
  }
}
