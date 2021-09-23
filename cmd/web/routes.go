package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wickywaa/gocourse1/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	return mux
}
