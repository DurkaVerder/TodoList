package main

import (
	"TodoList/internal/db"
	"log"
)

func main() {
	var err error
	// Init Config

	// Init DataBase
	if err = db.InitDb(); err != nil {
		log.Fatal("Error init database:", err)
	}

	// Init Handlers

	// Run App
}
