package posts

import "my/perfectPetProject/internal/models"

type postRepository interface {
	GetAll() []*models.Post
	Create(post *models.Post) (*models.Post, error)
	Delete(id int64) error
}

type userRepository interface {
	GetById(id int64) *models.User
	Create(user *models.User) (*models.User, error)
	Delete(id int64) error
}
