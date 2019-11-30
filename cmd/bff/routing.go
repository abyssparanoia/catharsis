package main

import (
	"net/http"

	"github.com/abyssparanoia/catharsis/bff/handler"
	"github.com/go-chi/chi"
)

// Routing ... define routing
func routing(r chi.Router, d dependency) {

	// Ping
	r.Get("/ping", handler.Ping)

	http.Handle("/", r)
}
