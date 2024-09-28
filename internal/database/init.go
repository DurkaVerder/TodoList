package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	open := "user=postgres password=durka dbname=TodoList sslmode=disable"
	db, err := sql.Open("postgres", open)
	if err != nil {
		log.Fatal("Error open db: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connect db: ", err)
	}
	return db
}
