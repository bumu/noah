package handlers

import (
	"net/http"
	"noah/pkg/configkit"
	"strings"

	// "noah/apps/apigw/data/repos"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux
}

func Register(deps registerDeps) {
	deps.Mux.Route("/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			msg := "welcome to noah service\n"
			msg += "Version /v1\n"
			msg += "\n\n\n"

			msg += "apps list: " + strings.Join(configkit.Apps, ",")

			w.Write([]byte(msg))
		})
		// r.Handle("/metrics", promhttp.Handler())
		r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong\n"))
		})

		//r.HandleFunc("/debug/pprof/", middleware.Profiler().ServeHTTP)
		r.Handle("/debug/pprof/*", middleware.Profiler())
		r.Handle("/debug/*", middleware.Profiler())
	})

	// Internal management apis.
	deps.Mux.Route("/internal", func(r chi.Router) {
		r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok\n"))
		})
	})
}
