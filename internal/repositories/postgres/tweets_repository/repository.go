package tweets_repository

import (
	"database/sql"

	"github.com/kudrmax/perfectPetProject/internal/models"
	"github.com/kudrmax/perfectPetProject/internal/utils"

	sq "github.com/Masterminds/squirrel"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

const (
	tableName = "twits"

	colID        = "id"
	colUserID    = "user_id"
	colText      = "text"
	colCreatedAt = "created_at"
)

func (r *Repository) GetAll() ([]*models.Tweet, error) {
	sb := sq.
		Select(colID, colUserID, colText, colCreatedAt).
		From(tableName)
	query, _ := sb.MustSql()

	tx, err := r.db.Begin() // TODO перейти на BeginTx
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback() // TODO убедиться, что такой rollback работает
		}
	}()

	rows, err := tx.Query(query)
	if err != nil {
		return nil, r.processGetAllErrors(err)
	}

	tweets := make([]*models.Tweet, 0)
	for rows.Next() {
		tweet := new(models.Tweet)
		err = rows.Scan(&tweet.Id, &tweet.UserId, &tweet.Text, &tweet.CreatedAt) // TODO как не принимать некоторые параметры?
		if err != nil {
			return nil, r.processGetAllErrors(err)
		}

		tweets = append(tweets, tweet)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return tweets, nil
}

func (r *Repository) Create(twit *models.Tweet) (*models.Tweet, error) {
	if empty(twit) {
		return nil, ErrEmptyTwit
	}

	sb := sq.
		Insert(tableName).
		Columns(colUserID, colText).
		Values(twit.UserId, twit.Text).
		Suffix(utils.ReturningSQL(colID, colCreatedAt)).
		PlaceholderFormat(sq.Dollar)
	query, args := sb.MustSql()

	tx, err := r.db.Begin() // TODO перейти на BeginTx
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback() // TODO убедиться, что такой rollback работает
		}
	}()

	err = tx.
		QueryRow(query, args...).
		Scan(&twit.Id, &twit.CreatedAt)
	if err != nil {
		return nil, r.processCreateErrors(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return twit, nil
}

func empty(twit *models.Tweet) bool {
	return twit == nil || twit.Text == "" || twit.UserId == 0
}
