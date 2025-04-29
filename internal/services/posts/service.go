package posts

import (
	"time"

	"my/perfectPetProject/internal/models"
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

func (s *Service) CreatePost(text string, userId int64) (*models.Post, error) {
	post, _ := s.postRepository.Create(&models.Post{
		Text:     text,
		UserId:   userId,
		Datetime: time.Now(),
	})

	return post, nil
}

func (s *Service) DeletePost(id int64) error {
	_ = s.postRepository.Delete(id)
	return nil
}
