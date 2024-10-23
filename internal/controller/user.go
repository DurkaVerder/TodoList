package controller

import (
	"TodoList/internal/jwt"
	"TodoList/internal/model"
	"net/http"

	"github.com/labstack/echo"
)

type UserController interface {
	HandleLogin(ctx echo.Context) error
	HandleRegister(ctx echo.Context) error
	HandleProfileUser(ctx echo.Context) error
	HandleUpdateUser(ctx echo.Context) error
	HandleDeleteUser(ctx echo.Context) error
}

func (c *ControllerManager) HandleLogin(ctx echo.Context) error {
	data := model.EnterDataUser{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid json format")
	}

	token, err := c.Service.Login(data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error login user")
	}

	if err := c.Cookie.SaveJWTInCookie(token, ctx); err != nil {
		return ctx.JSON(http.StatusOK, "error create JWT")
	}

	return ctx.JSON(http.StatusOK, "login user")
}

func (c *ControllerManager) HandleRegister(ctx echo.Context) error {
	data := model.EnterDataUser{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid json format")
	}

	token, err := c.Service.Register(data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error add user")
	}

	if err := c.Cookie.SaveJWTInCookie(token, ctx); err != nil {
		return ctx.JSON(http.StatusOK, "error create JWT")
	}
	return ctx.JSON(http.StatusOK, "Added user")
}

func (c *ControllerManager) HandleProfileUser(ctx echo.Context) error {
	token, err := c.Cookie.GetJWTFromCookie(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error get JWT")
	}

	claims, err := jwt.ValidateJWT(token)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error validate JWT")
	}
	userIdFloat, ok := claims["userId"].(float64)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "Invalid userID type")
	}
	userId := int(userIdFloat)
	user, err := c.Service.GetUser(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error get user")
	}
	return ctx.JSON(http.StatusOK, user)
}

func (c *ControllerManager) HandleUpdateUser(ctx echo.Context) error {

	return nil
}

func (c *ControllerManager) HandleDeleteUser(ctx echo.Context) error {
	return nil
}
