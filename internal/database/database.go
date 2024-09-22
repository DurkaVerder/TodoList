package database

import (
	"TodoList/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb(cfg *config.Config) error {
	var err error
	openBdStr := fmt.Sprintf("dbname=%v password=%v port=%v sslmode=%v user=%v", cfg.DB.DbName, cfg.DB.Password, cfg.DB.Port, cfg.DB.SSLMode, cfg.DB.User)
	if DB, err = sql.Open("postgres", openBdStr); err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}
