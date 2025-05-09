package password_hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	FailedHashPasswordErr = errors.New("failed to hash password")
)

type Service struct {
	cost int
}

func NewService() *Service {
	return &Service{
		cost: bcrypt.DefaultCost,
	}
}

func (s *Service) GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", FailedHashPasswordErr
	}

	return string(hashedPassword), nil
}

func (s *Service) CompareHashAndPassword(passwordHash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
