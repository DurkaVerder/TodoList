package controller

import (
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
	HandleUpdateUserPassword(ctx echo.Context) error
	HandleUpdateUserName(ctx echo.Context) error
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

	userId, err := c.Cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	user, err := c.Service.GetUser(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error get user")
	}
	return ctx.JSON(http.StatusOK, user)
}

func (c *ControllerManager) HandleUpdateUserPassword(ctx echo.Context) error {
	userId, err := c.Cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	password := ""
	if err := ctx.Bind(&password); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data type")
	}

	if err := c.Service.UpdateUserPassword(userId, password); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "")
}

func (c *ControllerManager) HandleUpdateUserName(ctx echo.Context) error {
	userId, err := c.Cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	name := ""
	if err := ctx.Bind(&name); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data type")
	}

	if err := c.Service.UpdateUserPassword(userId, name); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "")
}

func (c *ControllerManager) HandleDeleteUser(ctx echo.Context) error {
	userId, err := c.Cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := c.Service.DeleteUser(userId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error delete account")
	}
	// Delete cookie
	return ctx.JSON(http.StatusAccepted, "Deleted")
}
