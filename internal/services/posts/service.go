package posts

import (
	"time"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	postRepository interface {
		GetAll() []*models.Post
		Create(post *models.Post) (*models.Post, error)
		Delete(id int) error
	}

	userRepository interface {
		GetById(id int) *models.User
		Create(user *models.User) (*models.User, error)
		Delete(id int) error
	}
)

type Service struct {
	postRepository postRepository
	userRepository userRepository
}

func NewService(
	postRepository postRepository,
	userRepository userRepository,
) *Service {
	return &Service{
		postRepository: postRepository,
		userRepository: userRepository,
	}
}

func (s *Service) GetAllPosts() []*models.Post {
	return s.postRepository.GetAll()
}

func (s *Service) CreatePost(text string, userId int) (*models.Post, error) {
	post, _ := s.postRepository.Create(&models.Post{
		Text:     text,
		UserId:   userId,
		Datetime: time.Now(),
	})

	return post, nil
}

func (s *Service) DeletePost(id int) error {
	_ = s.postRepository.Delete(id)
	return nil
}
