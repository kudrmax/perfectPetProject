package testdb

import (
	"fmt"

	"github.com/stretchr/testify/assert"

	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/storage"
)

func MustInit(a *assert.Assertions) *storage.Storage {
	st, err := Init()
	if err != nil {
		a.FailNowf(err.Error(), "failed to init storage")
	}

	return st
}

func Init() (*storage.Storage, error) {
	userStorageConfig := storage.Config{ // TODO перенести в config и от туда же брать данные для создания БД для интеграционных тестов в docker compose
		Host:     "localhost", // TODO брать из env переменных
		Port:     "5432",
		User:     "postgres",
		DbName:   "postgres",
		Password: "postgres",
	}

	db, err := storage.New(userStorageConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	return db, nil
}
