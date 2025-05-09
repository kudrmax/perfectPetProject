package users_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/db_emulation"
	"github.com/kudrmax/perfectPetProject/internal/services/password_hasher"
)

func NewDbEmulation() db_emulation.DbEmulation[models.User] {
	db := db_emulation.NewDbEmulation[models.User]()
	addDummyData(&db)
	return db
}

func addDummyData(db *db_emulation.DbEmulation[models.User]) {
	passwordHasher := password_hasher.NewService()

	passwordHash, _ := passwordHasher.GenerateHashPassword("kudrmax")
	db.Create(&models.User{Id: 1, Name: "Max", Username: "kudrmax", PasswordHash: passwordHash}, SetIdFunc)

	passwordHash, _ = passwordHasher.GenerateHashPassword("elina")
	db.Create(&models.User{Id: 2, Name: "Elina", Username: "elina", PasswordHash: passwordHash}, SetIdFunc)

	passwordHash, _ = passwordHasher.GenerateHashPassword("string")
	db.Create(&models.User{Id: 3, Name: "string", Username: "string", PasswordHash: passwordHash}, SetIdFunc)
}
