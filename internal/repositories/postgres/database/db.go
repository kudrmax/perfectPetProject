package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func MustInit(cfg Config) *sql.DB {
	dsn := cfg.ToDSN()

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(fmt.Errorf("failed to open db: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("failed to ping db: %w", err))
	}

	return db
}
