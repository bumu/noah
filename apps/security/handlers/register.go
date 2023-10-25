package handlers

import (
	"net/http"

	// "noah/apps/apigw/data/repos"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux
}

func Register(deps registerDeps) {
	// Internal management apis.
	deps.Mux.Route("/internal", func(r chi.Router) {
		r.HandleFunc("/security/check", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok\n"))
		})
	})
}
