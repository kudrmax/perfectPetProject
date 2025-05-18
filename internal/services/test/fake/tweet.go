package fake

import "github.com/kudrmax/perfectPetProject/internal/models"

func Twit(options ...twitOption) *models.Tweet {
	twit := &models.Tweet{
		Text: RandString(),
	}

	for _, opt := range options {
		opt(twit)
	}

	return twit
}

type twitOption Option[models.Tweet]
