package tweets_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/db_emulation"
)

func NewDbEmulation() db_emulation.DbEmulation[models.Tweet] {
	db := db_emulation.NewDbEmulation[models.Tweet]()
	addDummyData(&db)
	return db
}

func addDummyData(db *db_emulation.DbEmulation[models.Tweet]) {
	db.Create(&models.Tweet{Id: 1, Text: "First tweet"}, SetIdFunc)
	db.Create(&models.Tweet{Id: 2, Text: "Second tweet"}, SetIdFunc)
}
