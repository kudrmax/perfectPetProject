package users_repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUsernameAlreadyExists = fmt.Errorf("username already exists")
)

func (r *Repository) processGetErrors(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	return nil
}

func (r *Repository) processCreateErrors(err error) error {
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") { // TODO заменить на код ошибки через pgx
		return ErrUsernameAlreadyExists
	}

	return nil
}
