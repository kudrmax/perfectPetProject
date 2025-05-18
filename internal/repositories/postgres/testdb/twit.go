package testdb

import (
	"database/sql"

	"github.com/stretchr/testify/require"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

func MustDeleteAllTweets(r *require.Assertions, db *sql.DB) {
	query := `
		DELETE
		FROM twits
		WHERE true
	`

	_, err := db.Query(query)
	r.NoError(err)
}

func MustAddTweet(r *require.Assertions, db *sql.DB, tweet *models.Tweet) {
	query := `
		INSERT INTO twits (user_id, text)
		VALUES ($1, $2)
		RETURNING id
	`

	err := db.
		QueryRow(query, tweet.UserId, tweet.Text).
		Scan(&tweet.Id)
	r.NoError(err)
}
