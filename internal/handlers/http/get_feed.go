package http

import (
	"context"

	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/converters/posts_converter"
)

func (h *Handler) GetFeed(ctx context.Context, request api.GetFeedRequestObject) (api.GetFeedResponseObject, error) {
	posts := h.postService.GetAll()
	return api.GetFeed200JSONResponse(posts_converter.ToApiModelSlice(posts)), nil
}
