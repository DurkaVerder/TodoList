package controller

import (
	"TodoList/internal/model"
	"net/http"

	"github.com/labstack/echo"
)

type UserController interface {
	HandleLogin(ctx echo.Context) error
	HandleRegister(ctx echo.Context) error
	HandleGetUser(ctx echo.Context) error
	HandleUpdateUser(ctx echo.Context) error
	HandleDeleteUser(ctx echo.Context) error
}

func (c *ControllerManager) HandleLogin(ctx echo.Context) error {
	data := model.EnterDataUser{}

	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid json format")
	}

	ok, err := c.Service.Login(data)
	if err != nil || !ok {
		return ctx.JSON(http.StatusInternalServerError, "Error login user")
	}

	return ctx.JSON(http.StatusOK, "login user")
}

func (c *ControllerManager) HandleRegister(ctx echo.Context) error {
	data := model.EnterDataUser{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid json format")
	}
	if err := c.Service.Register(data); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error add user")
	}

	return ctx.JSON(http.StatusOK, "Added user")
}

func (c *ControllerManager) HandleGetUser(ctx echo.Context) error {
	return nil
}

func (c *ControllerManager) HandleUpdateUser(ctx echo.Context) error {
	return nil
}

func (c *ControllerManager) HandleDeleteUser(ctx echo.Context) error {
	return nil
}
