package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() error {
	var err error
	connectStr := "dbname=TodoList user=postgres password=durka sslmode=disable port=5432"
	if DB, err = sql.Open("postgres", connectStr); err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}
