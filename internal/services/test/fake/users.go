package fake

import "github.com/kudrmax/perfectPetProject/internal/models"

func User(options ...userOption) *models.User {
	user := &models.User{
		Name:         RandString(),
		Username:     RandString(),
		PasswordHash: RandString(),
	}

	for _, opt := range options {
		opt(user)
	}

	return user
}

type userOption Option[models.User]

func WithUsername(username string) userOption {
	return func(order *models.User) {
		order.Username = username
	}
}
