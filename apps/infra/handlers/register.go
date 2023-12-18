package handlers

import (
	"net/http"
	"noah/apps/infra/data/repos"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux

	// GatewayDomainRepo *repos.GatewayDomainRepo
	InfraLinuxRepo *repos.InfraLinuxRepo
}

// Refer: https://google.aip.dev/158
func Register(deps registerDeps) {
	// method order: get, list, create, update, delete.
	deps.Mux.Route("/v1/infra/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to infra service\n"))
		})
	})

	deps.Mux.Route("/apis/v1", func(r chi.Router) {
		r.HandleFunc("/infra", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to infra service\n"))
		})

		r.HandleFunc("/infra/motd/{hostname}", deps.CreateVisitRecord)
	})
}
