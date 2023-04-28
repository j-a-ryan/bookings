package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/j-a-ryan/bookings/pkg/config"
	"github.com/j-a-ryan/bookings/pkg/handlers"
	"github.com/j-a-ryan/bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// Now refactored into packages, the program must be run by:
// go run ./cmd/web/ .
// from project root.
// Seesions from go get github.com/alexedwards/scs/v2
// sc uses cookies by default as the store.
func main() {

	app.InProduction = false // Set to true when deploying for production

	session = scs.New()
	app.Session = session
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true                  // Cookie persists after browser closed.
	session.Cookie.SameSite = http.SameSiteLaxMode // Not nec same site
	session.Cookie.Secure = app.InProduction       // https for production, http for dev

	tc, err := render.CreateTemplateCacheAutomatic()
	if err != nil {
		log.Fatal("Cannot create template cache.")
	}

	app.TemplateCache = tc
	app.UseCache = true

	// These two should be one method, one line here.
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// Instead of this server code....
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)/
	// _ = http.ListenAndServe(port, nil) // No need to pass in a handler. Our handler is above.

	// ...Use this server code, including pat's mux
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app), // pat mux routes
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
