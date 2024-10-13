package service

import (
	"github.com/labstack/echo"
)

type Service interface {
	TaskManager
	UserManager
}

type TaskManager interface {
	AllTask
	GetTask
	AddTask
	UpdateTask
	DeleteTask
}

type AllTask interface {
	GetAllTasks(ctx echo.Context) error
}

type GetTask interface {
	GetTask(ctx echo.Context) error
}

type AddTask interface {
	AddTask(ctx echo.Context) error
}

type UpdateTask interface {
	UpdateTask(ctx echo.Context) error
}

type DeleteTask interface {
	DeleteTask(ctx echo.Context) error
}

type UserManager interface {
}
