package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/leetrent/bookings/internal/config"
	"github.com/leetrent/bookings/internal/driver"
	"github.com/leetrent/bookings/internal/handlers"
	"github.com/leetrent/bookings/internal/helpers"
	"github.com/leetrent/bookings/internal/models"
	"github.com/leetrent/bookings/internal/render"
)

const port = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	server := &http.Server{
		Addr:    port,
		Handler: routes(&appConfig),
	}

	fmt.Printf("Starting application on port %s", port)
	err = server.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	appConfig.InProduction = false
	appConfig.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	appConfig.Session = session

	///////////////////////////////////////////////////////
	// Connect to database
	///////////////////////////////////////////////////////
	log.Println("Attempting to connect to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=test_connect user=postgres password=Ca$eyPo0h")
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot connect to database: %v", err))
	}
	log.Println("Attempt to connect to database was successful.")

	///////////////////////////////////////////////////////
	// Register datatype models.Reservation
	// so it can be stored in HTTP Session
	///////////////////////////////////////////////////////
	gob.Register(models.Reservation{})

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache.")
		return nil, err
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	repo := handlers.NewRepo(&appConfig, db)
	handlers.NewHandlers(repo)
	render.NewTemplates(&appConfig)
	helpers.NewHelpers(&appConfig)

	return db, nil
}
