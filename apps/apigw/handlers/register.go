package handlers

import (
	"log"
	"net/http"

	"apigw/apps/apigw/data/repos"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux
	//Logger  *slog.Logger
	KeyRepo *repos.KeyRepo
	// Checker *sensitivemod.Checker
}

func Register(deps registerDeps) {
	deps.Mux.Route("/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to apigw service\n"))
		})
		r.Handle("/metrics", promhttp.Handler())
	})

	deps.Mux.Route("/v1", func(r chi.Router) {
		r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			// log := deps.Logger.With("uri", r.URL.String())

			log.Println("ping")
			w.Write([]byte("pong\n"))
		})
		r.HandleFunc("/user", deps.HandleUser)
		r.HandleFunc("/openai/*", deps.HandleOpenai)

	})
}
