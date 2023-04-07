package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/sagarhande/roomifyr/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:

	default:
		t.Error("type is not *chi.Mux, type is", v)
	}
}
