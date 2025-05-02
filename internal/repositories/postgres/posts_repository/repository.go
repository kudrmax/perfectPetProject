package posts_repository

import (
	"github.com/kudrmax/perfectPetProject/internal/models"
)

type Repository struct {
	db DbEmulation
}

func NewRepository() *Repository {
	return &Repository{
		db: NewDbEmulation(),
	}
}

func (r *Repository) GetAll() []*models.Post {
	return r.db.GetAllPosts()
}

func (r *Repository) Create(post *models.Post) (*models.Post, error) {
	post = r.db.CreatePost(post)

	return post, nil
}

func (r *Repository) Delete(id int) error {
	return nil
}
