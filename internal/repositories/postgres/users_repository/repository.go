package users_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/storage"
)

type Repository struct {
	db *storage.Storage
}

func New(db *storage.Storage) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, nil
	}

	query := `
		SELECT id, name, username, passwordHash 
		FROM users 
		WHERE username = $1
	`

	var user models.User
	err := r.db.QueryRow(query, username).
		Scan(&user.Id, &user.Name, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, r.processGetErrors(err)
	}

	return &user, nil
}

func (r *Repository) Create(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (name, username, passwordHash) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, username, passwordHash
	`

	var newUser models.User
	err := r.db.QueryRow(query, user.Name, user.Username, user.PasswordHash).
		Scan(&newUser.Id, &newUser.Name, &newUser.Username, &newUser.PasswordHash)
	if err != nil {
		return nil, r.processCreateErrors(err)
	}

	return &newUser, nil
}

func (r *Repository) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (r *Repository) UpdateByUsername(username string, newUser *models.User) (*models.User, error) {
	return nil, nil
}

func (r *Repository) DeleteByUsername(username string) error {
	return nil
}
