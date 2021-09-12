package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/leetrent/bookings/pkg/config"
	"github.com/leetrent/bookings/pkg/handlers"
)

func routes(ac *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
