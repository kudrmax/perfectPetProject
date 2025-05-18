package testdb

import (
	"database/sql"
	"testing"

	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/database"
)

// TODO перенести в config и от туда же брать данные для создания БД для интеграционных тестов в docker compose
// TODO завести отдельную БД для тестов
var testDbConfig = database.Config{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	DbName:   "postgres",
	Password: "postgres",
}

func MustInit(t testing.TB) *sql.DB {
	t.Helper()

	return database.MustInit(testDbConfig)
}
