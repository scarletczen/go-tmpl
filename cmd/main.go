package main

import (

  "github.com/abhinavthapa1998/go-tmpl/handler"
  "github.com/labstack/echo/v4"
)

type DB struct{}

func main(){
  app := echo.New()
  var db DB
  userHandler := handler.UserHandler{db}
  app.GET("/user",userHandler.HandleUserShow)
  app.Start(":3000")
}
