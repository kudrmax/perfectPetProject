package create_tweet_handler

import (
	"encoding/json"
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/http/http_model"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type (
	tweetService interface {
		Create(text string, userId int) (*models.Tweet, error)
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

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var tweetCreate http_model.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweetCreate); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	text := tweetCreate.Text
	userId := 666

	tweet, err := h.tweetService.Create(text, userId)
	if err != nil {
		http.Error(w, "Failed to create tweet", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(convertToDto(tweet))
}

func convertToDto(tweet *models.Tweet) http_model.Tweet {
	if tweet == nil {
		return http_model.Tweet{}
	}

	return http_model.Tweet{
		Id:        tweet.Id,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt,
	}
}
