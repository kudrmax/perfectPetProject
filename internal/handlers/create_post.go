package handlers

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/kudrmax/perfectPetProject/internal/api"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

var (
	EmptyTextErr        = errors.New("empty text")
	CannotCreatePostErr = errors.New("cannot create post")
)

// CreatePost Создать новый пост
// (POST /api/1/posts/create_post)
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := parseBody[api.PostCreate](r)
	if err != nil {
		writeBadRequest(w, err)
		return
	}

	if err = validateBody(body); err != nil {
		writeBadRequest(w, err)
		return
	}

	post, err := h.postRepository.Create(convertPostToModel(&body))
	if err != nil {
		writeInternalError(w, CannotCreatePostErr)
	}

	writeJson(w, http.StatusCreated, convertModelToDto(post))
}

func validateBody(body api.PostCreate) error {
	if strings.TrimSpace(body.Text) == "" {
		return EmptyTextErr
	}

	return nil
}

func convertPostToModel(postDto *api.PostCreate) *models.Post {
	return &models.Post{
		Text:     postDto.Text,
		Datetime: time.Now(),
	}
}

func convertModelToDto(post *models.Post) *api.Post {
	return &api.Post{
		Id:        int(post.Id),
		Text:      post.Text,
		CreatedAt: post.Datetime,
	}
}
