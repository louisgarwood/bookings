package main

import (
	"github.com/louisgarwood/bookings/pkg/config"
	"github.com/louisgarwood/bookings/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func DefineRoutes(app *config.AppConfig) http.Handler {

	multiplexer := chi.NewRouter()

	multiplexer.Use(middleware.Recoverer)
	multiplexer.Use(WriteToConsole)
	multiplexer.Use(NoSurf)
	multiplexer.Use(SessionLoad)

	multiplexer.Get("/", handlers.Repo.Home)
	multiplexer.Get("/about", handlers.Repo.About)
	multiplexer.Get("/lou", handlers.Repo.Lou)

	return multiplexer
}