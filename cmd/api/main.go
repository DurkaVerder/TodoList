package main

import (
	"TodoList/internal/controller"
	"TodoList/internal/repository/postgres"
	"TodoList/internal/router"
	"TodoList/internal/service"

	"github.com/labstack/echo"
)

func main() {
	// Init connect database
	postgresRepository := postgres.NewPostgresRepo()

	// Init service
	service := service.NewService(postgresRepository)

	// Init controller
	controller := controller.NewController(service)

	// Init server
	e := echo.New()

	// Init routers
	router.InitRouter(e, controller)

	// Start app
	e.Start(":2222")
}
