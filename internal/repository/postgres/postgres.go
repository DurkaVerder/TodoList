package postgres

import (
	"database/sql"
	"log"
)

type PostgresRepo struct {
	db *sql.DB
}

func (repo *PostgresRepo) NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{db: repo.init()}
}

func (repo *PostgresRepo) init() *sql.DB {
	open := "user=postgres password=durka dbname=TodoList sslmode=disable"
	db, err := sql.Open("postgres", open)
	if err != nil {
		log.Fatal("Error open db: ", err)
	}

	if err = repo.db.Ping(); err != nil {
		log.Fatal("Error connect db: ", err)
	}
	return db
}
