package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func MustInit(cfg Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.ToDSN())
	if err != nil {
		panic(fmt.Errorf("failed to open db: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("failed to ping db: %w", err))
	}

	return db
}
