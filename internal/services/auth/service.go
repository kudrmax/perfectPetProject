package auth

import (
	"errors"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

var (
	UserAlreadyExistsErr    = errors.New("user already exists")
	UserNotFoundErr         = errors.New("user not found")
	FailedCreateUserErr     = errors.New("failed to create user")
	FailedHGenerateTokenErr = errors.New("failed to generate token")
	WrongPasswordErr        = errors.New("wrong password")
)

type (
	userService interface {
		Create(user *models.User) (*models.User, error)
		GetByUsername(username string) *models.User
	}

	passwordHasherService interface {
		GenerateHashPassword(password string) (string, error)
		CompareHashAndPassword(passwordHash, password string) bool
	}

	jwtProviderService interface {
		GenerateToken(userId int) (string, error)
		ParseToken(token string) (userId int, err error)
	}
)

type Service struct {
	userService           userService
	jwtProviderService    jwtProviderService
	passwordHasherService passwordHasherService
}

func NewService(
	userService userService,
	jwtProvider jwtProviderService,
	passwordChecker passwordHasherService,
) *Service {
	return &Service{
		userService:           userService,
		jwtProviderService:    jwtProvider,
		passwordHasherService: passwordChecker,
	}
}

func (s *Service) Register(name, username, password string) (accessToken string, err error) {
	if s.isUserExists(username) {
		return "", UserAlreadyExistsErr
	}

	passwordHash, err := s.passwordHasherService.GenerateHashPassword(password)
	user := &models.User{
		Name:         name,
		Username:     username,
		PasswordHash: passwordHash,
	}

	newUser, err := s.userService.Create(user)
	if err != nil {
		return "", FailedCreateUserErr
	}

	return s.generateAccessToken(newUser)
}

func (s *Service) Login(username, password string) (accessToken string, err error) {
	user := s.userService.GetByUsername(username)
	if user == nil {
		return "", UserNotFoundErr
	}

	if !s.passwordHasherService.CompareHashAndPassword(user.PasswordHash, password) {
		return "", WrongPasswordErr
	}

	return s.generateAccessToken(user)
}

func (s *Service) ValidateTokenAndGetUserId(token string) (userId int, err error) {
	return s.jwtProviderService.ParseToken(token)
}

func (s *Service) isUserExists(username string) bool {
	return s.userService.GetByUsername(username) != nil
}

func (s *Service) generateAccessToken(user *models.User) (string, error) {
	token, err := s.jwtProviderService.GenerateToken(user.Id)
	if err != nil {
		return "", FailedHGenerateTokenErr
	}

	return token, nil
}
