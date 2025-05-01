package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/oapi-codegen/nethttp-middleware"

	"github.com/kudrmax/perfectPetProject/internal/api"
	"github.com/kudrmax/perfectPetProject/internal/handlers"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/posts_repository"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/users_repository"
	"github.com/kudrmax/perfectPetProject/internal/services/posts"
)

func main() {
	userRepository := users_repository.NewRepository()
	postRepository := posts_repository.NewRepository()
	postService := posts.NewService(
		postRepository,
		userRepository,
	)

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("❌ failed to load swagger: %v", err)
	}
	swagger.Servers = nil

	router := chi.NewRouter()
	router.Use(nethttpmiddleware.OapiRequestValidator(swagger)) // валидация API

	handler := handlers.NewHandler(postService)
	router.Mount("/", api.Handler(handler))

	log.Println("Server started at http://localhost:8080")
	if err = http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("❌ server exited with error: %v", err)
	}
}
