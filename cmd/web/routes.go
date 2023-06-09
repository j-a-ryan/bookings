package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/j-a-ryan/bookings/pkg/config"
	"github.com/j-a-ryan/bookings/pkg/handlers"
)

// get rid of go get github.com/bmizerany/pat .
// routes uses github.com/go-chi/chi
// go get github.com/justinas/nosurf
func routes(app *config.AppConfig) http.Handler { // Returns a MUX (multiplexer)
	// Get rid of pat's mux
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))       // Cast Home method as a http.Handler
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About)) // Cast About method as a http.Handler

	mux := chi.NewRouter()

	// Chi offers middleware.
	mux.Use(middleware.Recoverer) // Recover gracefully from panics
	mux.Use(WriteToConsole)       // Dummy middleware we wrote that just logs to console.
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/kennys", handlers.Repo.Kennys)
	mux.Get("/ace", handlers.Repo.Ace)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)

	// File server for static files
	// Mux will handle calls to the /static folder using the server.
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
