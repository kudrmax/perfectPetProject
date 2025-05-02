package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mvrilo/go-redoc"

	myHttp "github.com/kudrmax/perfectPetProject/internal/handlers/http"
	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/posts_repository"
	"github.com/kudrmax/perfectPetProject/internal/repositories/postgres/users_repository"
	"github.com/kudrmax/perfectPetProject/internal/services/posts"
)

func main() {
	rootRouter := chi.NewRouter()

	rootRouter.Mount("/docs", getSwaggerRouter())
	rootRouter.Mount("/", getApiRouter())

	log.Println("Server started at http://localhost:8080")
	log.Println("OpenAPI docs at http://localhost:8080/docs/openapi")
	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		log.Fatalf("❌ server exited with error: %v", err)
	}
}

func getApiRouter() http.Handler {

	userRepository := users_repository.NewRepository()
	postRepository := posts_repository.NewRepository()
	postService := posts.NewService(
		postRepository,
		userRepository,
	)
	handler := myHttp.NewHandler(postService)

	router := chi.NewRouter()
	server := api.NewStrictHandler(handler, nil)

	//router.Use(nethttpmiddleware.OapiRequestValidator(swagger)) // валидация API
	//router.Mount("/", api.Handler(handler))
	router.Mount("/", api.Handler(server))

	return router
}

func getSwaggerRouter() http.Handler {
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
