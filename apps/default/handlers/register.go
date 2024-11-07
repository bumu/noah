package handlers

import (
	"net/http"

	// "noah/apps/apigw/data/repos"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux
}

func Register(deps registerDeps) {
	deps.Mux.Route("/", func(r chi.Router) {
		r.Get("/", deps.DefaultHandler)
		r.Get("/apis", deps.ListApisHandler)
	})

	// Internal management apis.
	deps.Mux.Route("/internal", func(r chi.Router) {
		// r.Handle("/metrics", promhttp.Handler())
		r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok\n"))
		})

		//r.HandleFunc("/debug/pprof/", middleware.Profiler().ServeHTTP)
		r.Handle("/debug/pprof/*", middleware.Profiler())
		r.Handle("/debug/*", middleware.Profiler())
	})
}
