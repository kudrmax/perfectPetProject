package users_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/db_emulation"
)

func NewDbEmulation() db_emulation.DbEmulation[models.User] {
	db := db_emulation.NewDbEmulation[models.User]()
	addDummyData(&db)
	return db
}

func addDummyData(db *db_emulation.DbEmulation[models.User]) {
	db.Create(&models.User{Id: 1, Name: "Max"})
	db.Create(&models.User{Id: 2, Name: "Elina"})
}
