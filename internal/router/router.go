package router

import (
	"TodoList/internal/controller"

	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo, controller controller.Controller) {

	e.GET("/api/task/:id", controller.HandleGetTask)
	e.GET("/api/tasks", controller.HandleGetAllTasks)
	e.GET("/api/profile", controller.HandleProfileUser)
	e.POST("/api/addtask", controller.HandleCreateTask)
	e.POST("/api/login", controller.HandleLogin)
	e.POST("/api/register", controller.HandleRegister)
	e.PUT("/api/task/:id", controller.HandleUpdateTask)
	e.PUT("/api/user/:id", controller.HandleUpdateUser)
	e.DELETE("/api/deleteuser/:id", controller.HandleDeleteUser)
	e.DELETE("/api/deletetask/:id", controller.HandleDeleteTask)
}
