package users_repository

import "github.com/kudrmax/perfectPetProject/internal/models"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetById(id int) *models.User {
	return &models.User{}
}

func (r *Repository) Create(user *models.User) (*models.User, error) {
	return nil, nil
}

func (r *Repository) Delete(id int) error {
	return nil
}
