package posts_converter

import (
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/models"
)

func ToApiModel(post models.Post) api.Post {
	return api.Post{
		Id:        post.Id,
		Text:      post.Text,
		CreatedAt: post.Datetime,
	}
}

func ToApiModelSlice(posts []*models.Post) []api.Post {
	apiPosts := make([]api.Post, 0, len(posts))

	for _, post := range posts {
		apiPosts = append(apiPosts, ToApiModel(*post))
	}

	return apiPosts
}
