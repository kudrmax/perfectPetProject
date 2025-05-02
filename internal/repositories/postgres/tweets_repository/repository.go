package tweets_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/db_emulation"
)

type Repository struct {
	db db_emulation.DbEmulation[models.Tweet]
}

func NewRepository() *Repository {
	return &Repository{
		db: NewDbEmulation(),
	}
}

func (r *Repository) GetAll() []*models.Tweet {
	return r.db.GetAll()
}

func (r *Repository) Create(tweet *models.Tweet) (*models.Tweet, error) {
	tweet = r.db.Create(tweet)

	return tweet, nil
}
