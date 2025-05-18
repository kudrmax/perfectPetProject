package users_repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUsernameAlreadyExists = fmt.Errorf("username already exists")
	ErrEmptyUser             = fmt.Errorf("some field of user is empty")
)

func (r *Repository) processGetErrors(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	return err
}

func (r *Repository) processCreateErrors(err error) error {
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") { // TODO заменить на код ошибки через pgx
		return ErrUsernameAlreadyExists
	}

	return err
}
