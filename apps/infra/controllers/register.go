package controllers

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
	InfraLinuxRepo         *repos.InfraLinuxRepo
	InfraHostHeartbeatRepo *repos.InfraHostHeartbeatRepo
}

func DefaultRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to infra service\n"))
}

// Refer: https://google.aip.dev/158
func Register(deps registerDeps) {
	deps.Mux.Route("/apis/infra", func(r chi.Router) {
		r.HandleFunc("/", DefaultRouter)

		r.Route("/v1", func(r chi.Router) {
			r.HandleFunc("/", DefaultRouter)

			r.HandleFunc("/motd/{hostname}", deps.CreateVisitRecord)
			r.HandleFunc("/linux/{hostname}", deps.CreateVisitRecord)
			r.HandleFunc("/linux/count", deps.Count)

			r.HandleFunc("/host/heartbeat", deps.CreateInfraHostHeartbeat)
		})
	})
}
