package controller

import (
	"TodoList/internal/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type TaskController interface {
	HandleGetAllTasks(ctx echo.Context) error
	HandleCreateTask(ctx echo.Context) error
	HandleUpdateTask(ctx echo.Context) error
	HandleDeleteTask(ctx echo.Context) error
}

func (c *ControllerManager) HandleGetAllTasks(ctx echo.Context) error {
	userId, err := c.Cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(http.StatusNonAuthoritativeInfo, "User nod found")
	}
	tasks, err := c.Service.GetTasks(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error get tasks")
	}
	return ctx.JSON(http.StatusAccepted, tasks)
}

func (c *ControllerManager) HandleCreateTask(ctx echo.Context) error {
	userId, err := c.Cookie.GetUserIdByCookie(ctx)
	if err != nil {
		return ctx.JSON(http.StatusNonAuthoritativeInfo, "Error get user id")
	}
	data := model.EnterDataTask{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := c.Service.CreateTask(data, userId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error create task")
	}

	return ctx.JSON(http.StatusCreated, "")
}

func (c *ControllerManager) HandleUpdateTask(ctx echo.Context) error {
	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid task id")
	}
	data := model.EnterDataTask{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := c.Service.UpdateTask(data, taskId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error update server")
	}
	return ctx.JSON(http.StatusAccepted, "")
}

func (c *ControllerManager) HandleDeleteTask(ctx echo.Context) error {
	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid task id")
	}
	if err := c.Service.DestroyTask(taskId); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error delete task")
	}

	return ctx.JSON(http.StatusAccepted, "")
}
