package get_feed

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	tweetService interface {
		GetAll() []*models.Tweet
	}
)

type Handler struct {
	tweetService tweetService
}

func NewHandler(
	tweetService tweetService,
) *Handler {
	return &Handler{
		tweetService: tweetService,
	}
}

type Tweet struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        int       `json:"id"`
	Text      string    `json:"text"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tweets := h.tweetService.GetAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(convertToDto(tweets))
}

func convertToDto(tweets []*models.Tweet) []Tweet {
	out := make([]Tweet, 0, len(tweets))

	for i := range tweets {
		out = append(out, Tweet{
			Id:        tweets[i].Id,
			Text:      tweets[i].Text,
			CreatedAt: tweets[i].CreatedAt,
		})
	}

	return out
}
