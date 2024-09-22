package main

import (
	"TodoList/config"
	"TodoList/internal/database"
	"TodoList/internal/server"
	"log"
)

func main() {
	var err error
	// Init Config
	cfg, err := config.InitCfg()
	if err != nil {
		log.Fatal("Error init database:", err)
	}

	// Init DataBase
	if err = database.InitDb(cfg); err != nil {
		log.Fatal("Error init database:", err)
	}

	// Run App
	server.Run(cfg)
}
