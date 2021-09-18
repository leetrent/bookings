package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/leetrent/bookings/pkg/config"
	"github.com/leetrent/bookings/pkg/handlers"
	"github.com/leetrent/bookings/pkg/render"
)

const port = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.Session = session

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
