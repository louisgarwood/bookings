package main

import (
	"net/http"

	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func DefineRoutes(app *config.AppConfig) http.Handler {

	multiplexer := chi.NewRouter()

	multiplexer.Use(middleware.Recoverer)
	multiplexer.Use(NoSurf)
	multiplexer.Use(SessionLoad)

	multiplexer.Get("/", handlers.Repo.Home)
	multiplexer.Get("/about", handlers.Repo.About)
	multiplexer.Get("/contact", handlers.Repo.Contact)
	multiplexer.Get("/captains", handlers.Repo.Captains)
	multiplexer.Get("/oceanview", handlers.Repo.Oceanview)

	multiplexer.Get("/search-availability", handlers.Repo.SearchAvailability)
	multiplexer.Post("/search-availability", handlers.Repo.PostSearchAvailability)
	multiplexer.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	multiplexer.Get("/make-reservation", handlers.Repo.Reservation)
	multiplexer.Post("/make-reservation", handlers.Repo.PostReservation)

	multiplexer.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	multiplexer.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return multiplexer
}
