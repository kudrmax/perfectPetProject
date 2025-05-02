package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mvrilo/go-redoc"

	myHttp "github.com/kudrmax/perfectPetProject/internal/handlers/http"
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/tweets_repository"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/users_repository"
	"github.com/kudrmax/perfectPetProject/internal/services/tweets"
)

func main() {
	rootRouter := chi.NewRouter()

	rootRouter.Mount("/docs", getRedocRouter())
	rootRouter.Mount("/", getApiRouter())

	log.Println("Server started at http://localhost:8080")
	log.Println("API: http://localhost:8080/api/1/tweets")
	log.Println("OpenAPI docs at http://localhost:8080/docs/openapi")
	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		log.Fatalf("❌ server exited with error: %v", err)
	}
}

func getApiRouter() http.Handler {
	tweetService := tweets.NewService(
		tweets_repository.NewRepository(),
		users_repository.NewRepository(),
	)
	handler := myHttp.NewHandler(tweetService)

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
