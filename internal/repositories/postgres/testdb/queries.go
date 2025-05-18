package testdb

import (
	"database/sql"

	"github.com/stretchr/testify/require"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

func MustAddUser(r *require.Assertions, db *sql.DB, user *models.User) {
	query := `
		INSERT INTO users
			(name,
			username,
			passwordhash)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := db.
		QueryRow(query, user.Name, user.Username, user.PasswordHash).
		Scan(&user.Id)
	r.NoError(err)
}

func MustGetUserByUsername(r *require.Assertions, db *sql.DB, username string) *models.User {
	query := `
		SELECT id, name, username, passwordhash
		FROM users
		WHERE username = $1
	`

	var user models.User
	err := db.
		QueryRow(query, username).
		Scan(&user.Id, &user.Name, &user.Username, &user.PasswordHash)
	r.NoError(err)

	return &user
}
