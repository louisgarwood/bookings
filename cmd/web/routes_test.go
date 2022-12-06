package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/louisgarwood/bookings/internal/config"
)

func TestDefineRoutes(t *testing.T) {
	var app config.AppConfig

	mux := DefineRoutes(&app)

	switch handlerType := mux.(type) {

	case *chi.Mux:
		//do nothing
	default:
		t.Errorf("the mux is not a chi.Mux but a %T", handlerType)
	}
}
