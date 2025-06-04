package users_repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

const (
	tableName          = "users"
	columnID           = "id"
	columnName         = "name"
	columnUsername     = "username"
	columnPasswordHash = "passwordHash"
)

func (r *Repository) GetByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, nil
	}

	sb := sq.
		Select(columnID, columnName, columnUsername, columnPasswordHash).
		From(tableName).
		Where(sq.Eq{columnUsername: username}).
		PlaceholderFormat(sq.Dollar)
	query, args := sb.MustSql()

	var user models.User
	err := r.db.
		QueryRow(query, args...).
		Scan(&user.Id, &user.Name, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, r.processGetErrors(err)
	}

	return &user, nil
}

func (r *Repository) Create(user *models.User) (*models.User, error) {
	if emptyUser(user) {
		return nil, ErrEmptyUser
	}

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

func emptyUser(user *models.User) bool {
	return user == nil || user.Name == "" || user.Username == ""
}
