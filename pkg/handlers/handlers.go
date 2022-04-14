package handlers

import (
	"github.com/nwohater/Go-Bookings/pkg/config"
	"github.com/nwohater/Go-Bookings/pkg/models"
	"github.com/nwohater/Go-Bookings/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository Is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// SuckIt is a world changing route
func (m *Repository) SuckIt(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "suckit.page.tmpl", &models.TemplateData{})
}
