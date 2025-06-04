package get_feed_handler

import (
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/http/handlers/http_common"
	"github.com/kudrmax/perfectPetProject/internal/http/http_model"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	tweetService interface {
		GetAll() ([]*models.Tweet, error)
	}
)

type Handler struct {
	tweetService tweetService
}

func New(
	tweetService tweetService,
) *Handler {
	return &Handler{
		tweetService: tweetService,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tweets, err := h.tweetService.GetAll()
	if err != nil {
		http.Error(w, "failed to get all tweets", http.StatusBadRequest)
		return
	}

	if err = http_common.WriteResponse(w, http.StatusOK, convertToDto(tweets)); err != nil {
		// TODO log
	}
}

func convertToDto(tweets []*models.Tweet) []http_model.Tweet {
	out := make([]http_model.Tweet, 0, len(tweets))

	for i := range tweets {
		out = append(out, http_model.Tweet{
			Id:        tweets[i].Id,
			Text:      tweets[i].Text,
			CreatedAt: tweets[i].CreatedAt,
		})
	}

	return out
}
