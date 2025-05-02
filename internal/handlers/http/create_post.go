package http

import (
	"context"
	"errors"

	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/converters/posts_converter"
)

var (
	EmptyTextErr        = errors.New("empty text")
	CannotCreatePostErr = errors.New("cannot create post")
)

func (h *Handler) CreatePost(ctx context.Context, request api.CreatePostRequestObject) (api.CreatePostResponseObject, error) {
	text := request.Body.Text
	userId := 666

	post, err := h.postService.CreatePost(text, userId)
	if err != nil {
		return api.CreatePost500JSONResponse{InternalErrorJSONResponse: api.InternalErrorJSONResponse{
			Error: err.Error(),
		}}, nil
	}

	return api.CreatePost201JSONResponse(posts_converter.ToApiModel(*post)), nil
}
