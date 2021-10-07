package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/leetrent/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var appConfig config.AppConfig
	mux := routes(&appConfig)
	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing, test passed
	default:
		t.Error(fmt.Sprintf("type is not of chi.Mux but is of %T instead.", v))
	}
}
