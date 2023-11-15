package handlers

import (
	"net/http"

	"github.com/aytacworld/go-bookings/pkg/config"
	"github.com/aytacworld/go-bookings/pkg/models"
	"github.com/aytacworld/go-bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repo type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repo for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home this is the homepage
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About this is the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// Send some data
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
