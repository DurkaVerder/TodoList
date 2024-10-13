package controller

import (
	"TodoList/internal/service"

	"github.com/labstack/echo"
)

type Controller interface {
	HandleLogin(ctx echo.Context) error
	HandleRegister(ctx echo.Context) error
	HandleGetUser(ctx echo.Context) error
	HandleUpdateUser(ctx echo.Context) error
	HandleDeleteUser(ctx echo.Context) error
	HandleGetAllTasks(ctx echo.Context) error
	HandleGetTask(ctx echo.Context) error
	HandleCreateTask(ctx echo.Context) error
	HandleUpdateTask(ctx echo.Context) error
	HandleDeleteTask(ctx echo.Context) error
}

type ControllerManager struct {
	Service service.Service
}

func NewController(s service.Service) *ControllerManager {
	return &ControllerManager{Service: s}
}
