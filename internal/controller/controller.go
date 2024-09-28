package controller

import (
	"TodoList/internal/service"
	"net/http"

	"github.com/labstack/echo"
)

// web layer

type TaskController struct {
	TaskService *service.TaskService
}

func (taskController *TaskController) GetAllTasks(ctx echo.Context) error {
	tasks, err := taskController.TaskService.GetAllTask(ctx)
	if err != nil {
		//Render error
		return err
	}
	return ctx.JSON(http.StatusAccepted, tasks)
}

func (taskController *TaskController) GetTask(ctx echo.Context) error {
	task, err := taskController.TaskService.GetTask(ctx)
	if err != nil {
		// Render error
		return err
	}
	return ctx.JSON(http.StatusAccepted, task)
}

func (taskController *TaskController) AddTask(ctx echo.Context) error {
	if err := taskController.TaskService.AddTask(ctx); err != nil {
		// Render error
		return err
	}
	return ctx.JSON(http.StatusOK, "Added")
}

func (taskController *TaskController) DeleteTask(ctx echo.Context) error {
	if err := taskController.TaskService.DeleteTask(ctx); err != nil {
		// Render error
		return err
	}

	return ctx.JSON(http.StatusOK, "Deleted")
}

func (taskController *TaskController) ChangeTask(ctx echo.Context) error {
	if err := taskController.TaskService.ChangeTask(ctx); err != nil {
		// Render error
		return err
	}
	return ctx.JSON(http.StatusOK, "Changed")
}
