package users_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/db_emulation"
)

var SetIdFunc = func(user *models.User, id int) {
	user.Id = id
}

type Repository struct {
	db db_emulation.DbEmulation[models.User]
}

func NewRepository() *Repository {
	return &Repository{
		db: NewDbEmulation(),
	}
}

func (r *Repository) GetByUsername(username string) *models.User {
	for _, user := range r.db {
		if user.Username == username {
			return &user
		}
	}

	return nil
}

func (r *Repository) Create(user *models.User) (*models.User, error) {
	user = r.db.Create(user, SetIdFunc)

	return user, nil
}
