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

	userRepository interface {
		GetById(id int) *models.User
		Create(user *models.User) (*models.User, error)
	}
)

type Service struct {
	tweetRepository tweetRepository
	userRepository  userRepository
}

func NewService(
	tweetRepository tweetRepository,
	userRepository userRepository,
) *Service {
	return &Service{
		tweetRepository: tweetRepository,
		userRepository:  userRepository,
	}
}

func (s *Service) GetAll() []*models.Tweet {
	return s.tweetRepository.GetAll()
}

func (s *Service) Create(text string, userId int) (*models.Tweet, error) {
	tweet, _ := s.tweetRepository.Create(&models.Tweet{
		Text:      text,
		UserId:    userId,
		CreatedAt: time.Now(),
	})

	return tweet, nil
}
