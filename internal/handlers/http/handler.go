package http

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	postService interface {
		GetAll() []*models.Tweet
		Create(text string, userId int) (*models.Tweet, error)
	}

	authService interface {
		Register(name, username, password string) (accessToken string, err error)
		Login(username, password string) (accessToken string, err error)
	}
)

type Handler struct {
	postService postService
	authService authService
}

func NewHandler(
	postService postService,
	authService authService,
) *Handler {
	return &Handler{
		postService: postService,
		authService: authService,
	}
}
