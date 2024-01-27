package controllers

import (
	"net/http"

	"noah/apps/security/data/repos"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux

	SecurityUserAgentRepo *repos.SecurityUserAgentRepo
}

func Register(deps registerDeps) {
	// Internal management apis.
	deps.Mux.Route("/internal/security", func(r chi.Router) {
		r.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok\n"))
		})
	})
	// /api/v1/sgw/waf/check
	deps.Mux.Route("/apis/v1/sgw/sec", func(r chi.Router) {
		r.Get("/check", deps.Health)
		r.Post("/check", deps.CheckSec)
	})

}
