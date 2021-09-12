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
		log.Fatal("Cannot create template cache.")
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)
	render.NewTemplates(&appConfig)

	server := &http.Server{
		Addr:    port,
		Handler: routes(&appConfig),
	}

	fmt.Printf("Starting application on port %s", port)
	err = server.ListenAndServe()
	log.Fatal(err)
}
