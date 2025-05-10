package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kudrmax/perfectPetProject/internal/http/handlers/auth_handler"
	"github.com/kudrmax/perfectPetProject/internal/http/handlers/create_tweet_handler"
	"github.com/kudrmax/perfectPetProject/internal/http/handlers/get_feed_handler"
	"github.com/kudrmax/perfectPetProject/internal/http/middlewares/auth_middleware"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/tweets_repository"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/users_repository"
	"github.com/kudrmax/perfectPetProject/internal/services/auth"
	"github.com/kudrmax/perfectPetProject/internal/services/jwt_token_generator"
	"github.com/kudrmax/perfectPetProject/internal/services/password_hasher"
	"github.com/kudrmax/perfectPetProject/internal/services/tweets"
	"github.com/kudrmax/perfectPetProject/internal/services/users"
)

func main() {
	mux := getApiRouter()

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
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
	jwtProviderService := jwt_token_generator.NewService(jwtSecret, jwtTokenDuration)
	passwordCheckerService := password_hasher.NewService()

	authService := auth.NewService(
		userService,
		jwtProviderService,
		passwordCheckerService,
	)

	// middlewares

	auth_mv := auth_middleware.New(authService).AuthMiddleware
	// TODO сделать удобное подключение auth middlewares
	// TODO добавить recover middleware
	// TODO добавить логирование
	// TODO добавить разные env

	// handlers

	handlerMap := map[string]http.HandlerFunc{
		"POST /api/1/auth/register": auth_handler.New(authService).Register,
		"POST /api/1/auth/login":    auth_handler.New(authService).Login,
		"POST /api/1/tweets/create": auth_mv(create_tweet_handler.New(tweetService).Handle),
		"GET /api/1/tweets/feed":    get_feed_handler.New(tweetService).Handle,
	}

	// routers

	// TODO посмотреть зачем нужны другие роутеры (chi, gin, gorilla/mux)
	mux := http.NewServeMux()
	for path, handler := range handlerMap {
		mux.HandleFunc(path, handler)
	}

	return mux
}
