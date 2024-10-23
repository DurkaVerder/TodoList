package controller

import "github.com/labstack/echo"

type TaskController interface {
	HandleGetAllTasks(ctx echo.Context) error
	HandleGetTask(ctx echo.Context) error
	HandleCreateTask(ctx echo.Context) error
	HandleUpdateTask(ctx echo.Context) error
	HandleDeleteTask(ctx echo.Context) error
}

func (c *ControllerManager) HandleGetAllTasks(ctx echo.Context) error {
	return nil
}

func (c *ControllerManager) HandleGetTask(ctx echo.Context) error {
	return nil
}

func (c *ControllerManager) HandleCreateTask(ctx echo.Context) error {
	return nil
}

func (c *ControllerManager) HandleUpdateTask(ctx echo.Context) error {
	return nil
}

func (c *ControllerManager) HandleDeleteTask(ctx echo.Context) error {
	return nil
}
