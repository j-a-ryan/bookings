package handlers

import (
	"net/http"

	"github.com/j-a-ryan/bookings/pkg/config"
	"github.com/j-a-ryan/bookings/pkg/models"
	"github.com/j-a-ryan/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home provides the home page. Has a Repository receiver.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr // Get the user's IP address
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About produces the about page. Has a Repository receiver.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// Dummy test
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}