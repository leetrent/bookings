package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leetrent/bookings/pkg/config"
	"github.com/leetrent/bookings/pkg/handlers"
	"github.com/leetrent/bookings/pkg/render"
)

const port = ":8080"

func main() {
	var appConfig config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template chache.")
	}

	appConfig.TemplateCache = templateCache
	render.NewTemplates(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)
}
