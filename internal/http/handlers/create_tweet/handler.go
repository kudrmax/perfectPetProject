package create_tweet

import (
	"encoding/json"
	"net/http"
	"time"

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

type TweetCreate struct {
	Text string `json:"text"`
}

type Tweet struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        int       `json:"id"`
	Text      string    `json:"text"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var tweetCreate TweetCreate
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

func convertToDto(tweet *models.Tweet) Tweet {
	if tweet == nil {
		return Tweet{}
	}

	return Tweet{
		Id:        tweet.Id,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt,
	}
}
