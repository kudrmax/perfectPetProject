package tweets_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type Repository struct {
	db DbEmulation
}

func NewRepository() *Repository {
	return &Repository{
		db: NewDbEmulation(),
	}
}

func (r *Repository) GetAll() []*models.Tweet {
	return r.db.GetAll()
}

func (r *Repository) Create(post *models.Tweet) (*models.Tweet, error) {
	post = r.db.Create(post)

	return post, nil
}
