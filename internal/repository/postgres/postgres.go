package postgres

import (
	"database/sql"
	"log"
)

type PostgresRepo struct {
	db *sql.DB
}

func (repo PostgresRepo) Init() {
	var err error
	open := "user=postgres password=durka dbname=TodoList sslmode=disable"
	repo.db, err = sql.Open("postgres", open)
	if err != nil {
		log.Fatal("Error open db: ", err)
	}

	if err = repo.db.Ping(); err != nil {
		log.Fatal("Error connect db: ", err)
	}
}
