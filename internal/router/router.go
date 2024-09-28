package router

import (
	"TodoList/internal/controller"

	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo, taskController *controller.TaskController) {
	e.GET("/task/:id", taskController.GetTask)
	e.GET("/tasks", taskController.GetAllTasks)
	e.POST("/add", taskController.AddTask)
	e.PUT("/change/:id", taskController.ChangeTask)
	e.DELETE("/delete/:id", taskController.DeleteTask)
}
