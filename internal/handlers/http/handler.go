package http

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type postService interface {
	GetAllPosts() []*models.Post
	CreatePost(text string, userId int) (*models.Post, error)
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
