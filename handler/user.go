package handler

import (
	"github.com/abhinavthapa1998/go-templ/model"
	"github.com/abhinavthapa1998/go-templ/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "abhinav@gmail.com",
	}
	return render(c, user.Show(u))
}
