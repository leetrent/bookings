package handlers

import (
	"net/http"

	"github.com/leetrent/bookings/pkg/config"
	"github.com/leetrent/bookings/pkg/models"
	"github.com/leetrent/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	AppConfig *config.AppConfig
}

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: ac,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	m.AppConfig.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(rw http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["about"] = "Using string map in template data. Preparing default data."

	remoteIP := m.AppConfig.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}
