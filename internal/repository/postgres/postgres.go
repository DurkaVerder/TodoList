package postgres

import (
	"database/sql"
	"log"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{db: initDataBase()}
}

func initDataBase() *sql.DB {
	var err error
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
