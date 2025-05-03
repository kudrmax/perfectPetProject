package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mvrilo/go-redoc"

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

	rootRouter.Mount("/docs", getRedocRouter())
	rootRouter.Mount("/", getApiRouter())

	log.Println("Server started at http://localhost:8080")
	log.Println("API: http://localhost:8080/api/1/tweets")
	log.Println("Feed: http://localhost:8080/api/1/tweets/feed")
	log.Println("OpenAPI docs at http://localhost:8080/docs/openapi")
	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		log.Fatalf("❌ server exited with error: %v", err)
	}
}

func getApiRouter() http.Handler {
	// config
	jwtTokenDuration := time.Minute * 15
	jwtSecret := "super-secret"

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
	_ = authService

	// handlers

	handler := myHttp.NewHandler(
		tweetService,
	)

	// routers

	router := chi.NewRouter()
	server := api.NewStrictHandler(handler, nil)

	//router.Use(nethttpmiddleware.OapiRequestValidator(swagger)) // валидация API
	router.Mount("/", api.Handler(server))

	return router
}

func getRedocRouter() http.Handler {
	doc := redoc.Redoc{
		Title:       "API Documentation",
		Description: "Интерактивная документация для API",
		SpecFile:    "./openapi/openapi.gen.yaml",
		SpecPath:    "/docs/openapi.yaml", // относительный путь внутри /docs
		DocsPath:    "/docs/openapi",      // будет доступен как /docs/openapi
	}

	router := chi.NewRouter()

	router.Get("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, doc.SpecFile)
	})
	router.Get("/openapi", doc.Handler())

	return router
}

func getSwaggerRouter() http.Handler {
	// использовать гайд отсюда:
	// https://youtu.be/87au30fl5e4?t=1376
	return nil
}
