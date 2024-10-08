package router

import (
	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo) {
	e.GET("/api/task/:id", )
	e.GET("/api/tasks", )
	e.POST("/api/addtask", )
	e.POST("/api/login", )
	e.POST("/api/register", )
	e.PUT("/api/task/:id", )
	e.PUT("/api/user/:id", )
	e.DELETE("/api/deleteuser/:id", )
	e.DELETE("/api/deletetask/:id", )
}
