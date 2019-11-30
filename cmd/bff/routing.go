package main

import (
	"net/http"

	"github.com/abyssparanoia/catharsis/pkg/httpaccesslog"

	"github.com/abyssparanoia/catharsis/bff/handler"
	"github.com/go-chi/chi"
)

// Routing ... define routing
func routing(r chi.Router, d dependency) {

	r.Use(httpaccesslog.Middleware())

	// need to authenticate for production
	r.Route("/v1", func(r chi.Router) {
		r.Route("/me", func(r chi.Router) {
			r.Post("/sign_in", d.AuthenticationHandler.SignIn)
		})
	})

	// Ping
	r.Get("/ping", handler.Ping)

	http.Handle("/", r)
}
