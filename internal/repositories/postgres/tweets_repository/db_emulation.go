package tweets_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/testdb"
)

func NewDbEmulation() testdb.DbEmulation[models.Tweet] {
	db := testdb.NewDbEmulation[models.Tweet]()
	addDummyData(&db)
	return db
}

func addDummyData(db *testdb.DbEmulation[models.Tweet]) {
	db.Create(&models.Tweet{Id: 1, Text: "First tweet"}, SetIdFunc)
	db.Create(&models.Tweet{Id: 2, Text: "Second tweet"}, SetIdFunc)
}
