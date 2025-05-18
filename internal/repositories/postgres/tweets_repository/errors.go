package tweets_repository

import "errors"

var (
	ErrEmptyTwit = errors.New("empty twit")
)

func (r *Repository) processGetAllErrors(err error) error {
	return err
}

func (r *Repository) processCreateErrors(err error) error {
	return err
}
