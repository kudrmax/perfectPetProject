package main

import (
	"my/perfectPetProject/internal/repositories/postgres/posts_repo"
	"my/perfectPetProject/internal/repositories/postgres/users_repo"
	"my/perfectPetProject/internal/services/posts"
)

func main() {
	userRepository := users_repo.NewRepository()
	postRepository := posts_repo.NewRepository()
	postService := posts.NewService(postRepository, userRepository)
	_ = postService
}
