package server

import (
	"TodoList/config"
	"TodoList/internal/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run(cfg *config.Config) {
	e := echo.New()

	e.Use(middleware.Logger())
	handlers.InitHandlers(e)

	e.Start(cfg.SV.Port)
}
