package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	fs := http.FileServer(http.Dir("./static/"))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Handle("/static/*", http.StripPrefix("/static", fs))

	mux.Get("/test-patterns", app.TestPatterns)

	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)
	mux.Get("/api/dog-breeds", app.GetAllDogBreedsHandler)

	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)
	mux.Get("/api/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)

	mux.Get("/api/dog-from-builder", app.CreateDogWithBuilder)
	mux.Get("/api/cat-from-builder", app.CreateCatWithBuilder)

	mux.Get("/api/animal-from-abstract-factory/{species}/{breed}", app.AnimalFromAbstractFactory)

	mux.Get("/api/cat-breeds", app.GetAllCatBreeds)

	mux.Get("/dog-of-month", app.DogOfMonth)

	mux.Get("/", app.Home)
	mux.Get("/{page}", app.Page)

	return mux
}
