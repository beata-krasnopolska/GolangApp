package main

import (
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter() //create multiplex

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
