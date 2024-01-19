package handlers

import (
	"net/http"
	"noah/apps/sgw/data/repos"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux

	GatewayDomainRepo *repos.GatewayDomainRepo
}

// Refer: https://google.aip.dev/158
func Register(deps registerDeps) {
	// method order: get, list, create, update, delete.
	deps.Mux.Route("/apis/v1/sgw/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to apigw service\n"))
		})
		r.Get("/domain", deps.ListDomain)
		r.Get("/dns", deps.ListDomain)
	})
}
