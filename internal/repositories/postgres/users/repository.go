package users

import "my/perfectPetProject/internal/models"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetById(id int64) *models.User {
	return &models.User{}
}

func (r *Repository) Create(user *models.User) (*models.User, error) {
	return nil, nil
}

func (r *Repository) Delete(id int64) error {
	return nil
}
