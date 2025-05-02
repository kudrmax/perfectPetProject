package posts_repository

import "github.com/kudrmax/perfectPetProject/internal/models"

type DbEmulation map[int]models.Post

func NewDbEmulation() DbEmulation {
	return make(DbEmulation)
}

func (db *DbEmulation) CreatePost(post *models.Post) *models.Post {
	newId := db.getNewId()
	post.Id = newId

	(*db)[newId] = *post

	return post
}

func (db *DbEmulation) GetAllPosts() []*models.Post {
	out := make([]*models.Post, 0, len(*db))

	for _, post := range *db {
		out = append(out, &post)
	}

	return out
}

func (db *DbEmulation) getNewId() int {
	return db.getMaxId() + 1
}

func (db *DbEmulation) getMaxId() int {
	maxId := 0

	for id := range *db {
		maxId = max(maxId, id)
	}

	return maxId
}
