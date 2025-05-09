package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/kudrmax/perfectPetProject/internal/handlers/newhttp/create_tweet"
	"github.com/kudrmax/perfectPetProject/internal/handlers/newhttp/get_feed"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/tweets_repository"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/users_repository"
	"github.com/kudrmax/perfectPetProject/internal/services/auth"
	"github.com/kudrmax/perfectPetProject/internal/services/jwt_provider"
	"github.com/kudrmax/perfectPetProject/internal/services/password_hasher"
	"github.com/kudrmax/perfectPetProject/internal/services/tweets"
	"github.com/kudrmax/perfectPetProject/internal/services/users"
)

func main() {
	rootRouter := chi.NewRouter()

	rootRouter.Mount("/", getApiRouter())

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		log.Fatalf("❌ server exited with error: %v", err)
	}
}

func getApiRouter() http.Handler {
	// config
	// TODO использовать какую-то библиотеку для настройки конфига
	const (
		jwtTokenDuration = time.Minute * 15
		jwtSecret        = "super-secret"
	)

	// repositories

	tweetRepository := tweets_repository.NewRepository()
	userRepository := users_repository.NewRepository()

	// services

	tweetService := tweets.NewService(tweetRepository)
	userService := users.NewService(userRepository)
	jwtProviderService := jwt_provider.NewService(jwtSecret, jwtTokenDuration)
	passwordCheckerService := password_hasher.NewService()

	authService := auth.NewService(
		userService,
		jwtProviderService,
		passwordCheckerService,
	)

	_ = tweetService
	_ = authService

	// handlers

	handlerMap := map[string]http.HandlerFunc{
		"POST /api/1/tweets/create_post": create_tweet.NewHandler(tweetService).Handle,
		"GET /api/1/tweets/feed":         get_feed.NewHandler(tweetService).Handle,
	}

	// routers

	mux := http.NewServeMux()
	for path, handler := range handlerMap {
		mux.HandleFunc(path, handler)
	}

	return mux
}
