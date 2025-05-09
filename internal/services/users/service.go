package users

import "github.com/kudrmax/perfectPetProject/internal/models"

type (
	userRepository interface {
		GetByUsername(username string) *models.User
		Create(user *models.User) (*models.User, error)
	}
)

type Service struct {
	userRepository userRepository
}

func NewService(
	userRepository userRepository,
) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) Create(user *models.User) (*models.User, error) {
	return s.userRepository.Create(user)
}

func (s *Service) GetByUsername(username string) (user *models.User) {
	return s.userRepository.GetByUsername(username)
}
