package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	myHttp "github.com/kudrmax/perfectPetProject/internal/handlers/http"
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
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

	// handlers

	handler := myHttp.NewHandler(
		tweetService,
		authService,
	)

	// routers

	//swagger, _ := api.GetSwagger()

	router := chi.NewRouter()
	server := api.NewStrictHandler(handler, nil)

	//router.Use(nethttpmiddleware.OapiRequestValidatorWithOptions(swagger, &nethttpmiddleware.Options{
	//	Options: openapi3filter.Options{
	//		AuthenticationFunc: handler.AuthMiddleware,
	//	},
	//}))
	router.Use(myHttp.AuthMiddleware2(authService))
	router.Mount("/", api.Handler(server))

	return router
}
