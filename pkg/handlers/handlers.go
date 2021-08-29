package handlers

import (
	"net/http"

	"github.com/leetrent/bookings/pkg/render"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(rw, "about.page.tmpl")
}
