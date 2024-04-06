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

	mux.Get("/", app.Home)
	mux.Get("/{page}", app.Page)

	return mux
}
