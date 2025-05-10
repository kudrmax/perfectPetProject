package tweets

import (
	"time"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	tweetRepository interface {
		GetAll() []*models.Tweet
		Create(tweet *models.Tweet) (*models.Tweet, error)
	}
)

type Service struct {
	tweetRepository tweetRepository
}

func NewService(
	tweetRepository tweetRepository,
) *Service {
	return &Service{
		tweetRepository: tweetRepository,
	}
}

func (s *Service) GetAll() []*models.Tweet {
	return s.tweetRepository.GetAll()
}

func (s *Service) Create(text string, userId int) (*models.Tweet, error) {
	tweet, err := s.tweetRepository.Create(&models.Tweet{
		Text:      text,
		UserId:    userId,
		CreatedAt: time.Now(),
	})
	_ = err // TODO добавить обработку ошибок

	return tweet, nil
}
