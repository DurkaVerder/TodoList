package main

import (
	"TodoList/internal/repository/postgres"

	"github.com/labstack/echo"
)

func main() {
	repoPostgres := postgres.PostgresRepo{}
	repoPostgres.Init()

	e := echo.New()

	e.Start(":2222")
}
