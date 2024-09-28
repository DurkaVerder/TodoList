package service

import (
	"TodoList/internal/model"
	"TodoList/internal/repository"

	"github.com/labstack/echo"
)

type TaskService struct {
	TaskRepository *repository.TaskRepository
}

// logic layer

func (taskService *TaskService) GetAllTask(ctx echo.Context) ([]model.Task, error) {
	return nil, nil
}

func (taskService *TaskService) GetTask(ctx echo.Context) (model.Task, error) {
	return model.Task{}, nil
}

func (taskService *TaskService) AddTask(ctx echo.Context) error {
	return nil
}

func (taskService *TaskService) DeleteTask(ctx echo.Context) error {
	return nil
}

func (taskService *TaskService) ChangeTask(ctx echo.Context) error {
	return nil
}
