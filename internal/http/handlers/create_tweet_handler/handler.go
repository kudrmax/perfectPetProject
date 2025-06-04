package create_tweet_handler

import (
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/http/const"
	"github.com/kudrmax/perfectPetProject/internal/http/handlers/http_common"
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

func New(
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

	tweetCreate, err := http_common.GetRequestBody[http_model.Tweet](r)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userID := ctx.Value(_const.UserIdContextKey).(int)
	if userID == 0 {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}

	tweet, err := h.tweetService.Create(tweetCreate.Text, userID)
	if err != nil {
		http.Error(w, "failed to create tweet", http.StatusInternalServerError)
		// TODO log
		return
	}

	err = http_common.WriteResponse(w, http.StatusCreated, convertToDto(tweet))
	if err != nil {
		// TODO log error
	}
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
