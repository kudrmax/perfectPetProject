package handlers

import (
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/api"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

// GetFeed Получить ленту постов
// (GET /api/1/posts/feed)
func (h *Handler) GetFeed(w http.ResponseWriter, r *http.Request) {
	posts := h.postRepository.GetAll()
	writeJson(w, http.StatusOK, convertModelsToDto(posts))
}

func convertModelsToDto(posts []*models.Post) []api.Post {
	if len(posts) == 0 {
		return []api.Post{}
	}

	out := make([]api.Post, 0, len(posts))

	for _, post := range posts {
		out = append(out, *convertModelToDto(post))
	}

	return out
}
