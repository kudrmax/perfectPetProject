package handlers

import (
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/api"
)

// GetFeed Получить ленту постов
// (GET /api/1/posts/feed)
func (*Handler) GetFeed(w http.ResponseWriter, r *http.Request) {
	writeJson(w, http.StatusOK, []api.Post{
		{
			Id:   666,
			Text: "Некий текст",
		},
		{
			Id:   777,
			Text: "Некий текст 2",
		},
	})
}
