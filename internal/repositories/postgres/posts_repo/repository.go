package posts_repo

import (
	"my/perfectPetProject/internal/models"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetAll() []*models.Post {
	return []*models.Post{}
}

func (r *Repository) Create(post *models.Post) (*models.Post, error) {
	return nil, nil
}

func (r *Repository) Delete(id int64) error {
	return nil
}
