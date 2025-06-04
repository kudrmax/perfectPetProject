package users_repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/utils"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

const (
	tableName = "users"

	colID           = "id"
	colName         = "name"
	colUsername     = "username"
	colPasswordHash = "passwordHash"
)

func (r *Repository) GetByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, nil
	}

	sb := sq.
		Select(colID, colName, colUsername, colPasswordHash).
		From(tableName).
		Where(sq.Eq{colUsername: username}).
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

	sb := sq.
		Insert(tableName).
		Columns(colName, colUsername, colPasswordHash).
		Values(user.Name, user.Username, user.PasswordHash).
		Suffix(utils.ReturningSQL(colID, colName, colUsername, colPasswordHash)).
		PlaceholderFormat(sq.Dollar)
	query, args := sb.MustSql()

	var newUser models.User
	err := r.db.
		QueryRow(query, args...).
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
