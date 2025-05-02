package users_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/db_emulation"
)

type Repository struct {
	db db_emulation.DbEmulation[models.User]
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetById(id int) *models.User {
	user := r.db.GetById(id)

	return user
}

func (r *Repository) Create(user *models.User) (*models.User, error) {
	user = r.db.Create(user)

	return user, nil
}
