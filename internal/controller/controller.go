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

func (c *TaskController) GetAllTasks(ctx echo.Context) error {
	tasks, err := c.TaskService.GetAllTask(ctx)
	if err != nil {
		//Render error
		return err
	}
	return ctx.JSON(http.StatusAccepted, tasks)
}

func (c *TaskController) GetTask(ctx echo.Context) error {
	task, err := c.TaskService.GetTask(ctx)
	if err != nil {
		// Render error
		return err
	}
	return ctx.JSON(http.StatusAccepted, task)
}

func (c *TaskController) AddTask(ctx echo.Context) error {
	if err := c.TaskService.AddTask(ctx); err != nil {
		// Render error
		return err
	}
	return ctx.JSON(http.StatusOK, "Added")
}

func (c *TaskController) DeleteTask(ctx echo.Context) error {
	if err := c.TaskService.DeleteTask(ctx); err != nil {
		// Render error
		return err
	}

	return ctx.JSON(http.StatusOK, "Deleted")
}

func (c *TaskController) ChangeTask(ctx echo.Context) error {
	if err := c.TaskService.ChangeTask(ctx); err != nil {
		// Render error
		return err
	}
	return ctx.JSON(http.StatusOK, "Changed")
}
