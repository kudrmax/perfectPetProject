package posts_converter

import (
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

func ToApiModel(post models.Tweet) api.Tweet {
	return api.Tweet{
		Id:        post.Id,
		Text:      post.Text,
		CreatedAt: post.CreatedAt,
	}
}

func ToApiModelSlice(posts []*models.Tweet) []api.Tweet {
	apiPosts := make([]api.Tweet, 0, len(posts))

	for _, post := range posts {
		apiPosts = append(apiPosts, ToApiModel(*post))
	}

	return apiPosts
}
