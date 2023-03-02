package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sagarhande/roomifyr/pkg/config"
	"github.com/sagarhande/roomifyr/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/rooms/{roomType}", handlers.Repo.Rooms)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
