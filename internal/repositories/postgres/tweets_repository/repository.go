package tweets_repository

import (
	"database/sql"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAll() ([]*models.Tweet, error) {
	query := `
		SELECT id, user_id, text, created_at, updated_at 
		FROM twits 
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, r.processGetAllErrors(err)
	}

	tweets := make([]*models.Tweet, 0)
	for rows.Next() {
		tweet := new(models.Tweet)
		err = rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Text, &tweet.CreatedAt, nil)
		if err != nil {
			return nil, r.processGetAllErrors(err)
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}

func (r *Repository) Create(twit *models.Tweet) (*models.Tweet, error) {
	if empty(twit) {
		return nil, ErrEmptyTwit
	}

	query := `
		INSERT INTO twits (user_id, text) 
		VALUES ($1, $2) 
		RETURNING id, created_at, updated_at
	`

	var newTwit models.Tweet
	err := r.db.
		QueryRow(query, newTwit.UserId, newTwit.Text).
		Scan(&newTwit.Id, &newTwit.CreatedAt, nil)
	if err != nil {
		return nil, r.processCreateErrors(err)
	}

	return &newTwit, nil
}

func empty(twit *models.Tweet) bool {
	return twit == nil || twit.Text == "" || twit.UserId == 0
}
