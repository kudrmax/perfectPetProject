package http

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type postService interface {
	GetAll() []*models.Tweet
	Create(text string, userId int) (*models.Tweet, error)
}

type Handler struct {
	postService postService
}

func NewHandler(
	postService postService,
) *Handler {
	return &Handler{
		postService: postService,
	}
}
