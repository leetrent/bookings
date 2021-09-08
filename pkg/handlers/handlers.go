package handlers

import (
	"net/http"

	"github.com/leetrent/bookings/pkg/config"
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
	render.RenderTemplate(rw, "home.page.tmpl")
}

func (m *Repository) About(rw http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(rw, "about.page.tmpl")
}
