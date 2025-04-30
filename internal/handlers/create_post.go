package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/kudrmax/perfectPetProject/internal/api"
)

// CreatePost Создать новый пост
// (POST /api/1/posts/create_post)
func (*Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := parseBody[api.PostCreate](r)
	if err != nil {
		writeBadRequest(w, err)
		return
	}

	if err = validateBody(body); err != nil {
		writeBadRequest(w, err)
		return
	}

	writeJson(w, http.StatusCreated, api.Post{
		Id:   666,
		Text: "Некий текст",
	})
}

func validateBody(body api.PostCreate) error {
	if strings.TrimSpace(body.Text) == "" {
		return errors.New("empty text")
	}

	return nil
}
