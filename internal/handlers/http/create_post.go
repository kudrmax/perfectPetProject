package http

import (
	"context"
	"errors"

	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/converters/posts_converter"
)

var (
	EmptyTextErr         = errors.New("empty text")
	CannotCreateTweetErr = errors.New("cannot create tweet")
)

func (h *Handler) CreateTweet(ctx context.Context, request api.CreateTweetRequestObject) (api.CreateTweetResponseObject, error) {
	text := request.Body.Text
	userId := 666

	post, err := h.postService.Create(text, userId)
	if err != nil {
		return api.CreateTweet500JSONResponse{InternalErrorJSONResponse: api.InternalErrorJSONResponse{
			Error: err.Error(),
		}}, nil
	}

	return api.CreateTweet201JSONResponse(posts_converter.ToApiModel(*post)), nil
}
