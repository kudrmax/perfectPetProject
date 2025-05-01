package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/api"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	postRepository interface {
		GetAll() []*models.Post
		Create(post *models.Post) (*models.Post, error)
	}
)

type Handler struct {
	postRepository postRepository
}

func NewHandler(
	postRepository postRepository,
) *Handler {
	return &Handler{
		postRepository: postRepository,
	}
}

func parseBody[T any](r *http.Request) (T, error) {
	var body T
	err := json.NewDecoder(r.Body).Decode(&body)
	return body, err
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeBadRequest(w http.ResponseWriter, err error) {
	writeJson(w, http.StatusBadRequest, api.BadRequest{
		Error: err.Error(),
	})
}

func writeInternalError(w http.ResponseWriter, err error) {
	writeJson(w, http.StatusInternalServerError, api.InternalError{
		Error: err.Error(),
	})
}
