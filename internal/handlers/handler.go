package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/api"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	postService interface {
		GetAllPosts() []*models.Post
		CreatePost(text string, userId int64) (*models.Post, error)
	}
)

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
