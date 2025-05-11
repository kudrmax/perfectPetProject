package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	*sql.DB
}

func New(cfg Config) (*Storage, error) {
	db, err := sql.Open("postgres", cfg.ToDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return &Storage{DB: db}, nil
}
