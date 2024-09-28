package main

import (
	"TodoList/internal/controller"
	"TodoList/internal/database"
	"TodoList/internal/repository"
	"TodoList/internal/router"
	"TodoList/internal/service"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	controller := &controller.TaskController{&service.TaskService{&repository.TaskRepository{database.Init()}}}

	router.InitRouter(e, controller)

	e.Start(":2222")
}
