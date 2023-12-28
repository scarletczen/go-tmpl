package main

import (
	"net/http"

	"github.com/abhinavthapa1998/go-tmpl/handler"
	"github.com/labstack/echo/v4"
)

func main(){
  app := echo.New()
  userHandler := handler.UserHandler{}
  app.GET("/user",userHandler.HandleUserShow)
  app.Start(":3000")
 }
