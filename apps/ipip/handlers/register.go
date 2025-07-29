package handlers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux
	//Logger  *slog.Logger
}

func Register(deps registerDeps) {
	deps.Mux.Route("/ip", func(r chi.Router) {
		// User api management
		r.Get("/", deps.IpHandler)

		r.Get("/{search_ip}", deps.IpHandler)
	})
}
