package controller

import "github.com/labstack/echo"

type UserController interface {
	HandleLogin(ctx echo.Context) error
	HandleRegister(ctx echo.Context) error
	HandleGetUser(ctx echo.Context) error
	HandleUpdateUser(ctx echo.Context) error
	HandleDeleteUser(ctx echo.Context) error
}

func (c *ControllerManager) HandleLogin(ctx echo.Context) error {
	return nil
}

func (c *ControllerManager) HandleRegister(ctx echo.Context) error {
	return nil
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
