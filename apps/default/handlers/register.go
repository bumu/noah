package handlers

import (
	"fmt"
	"net/http"
	"noah/pkg/configkit"

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
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to noah service!\n"))
		})

		r.HandleFunc("/apis", func(w http.ResponseWriter, r *http.Request) {
			msg := "welcome to noah service\n"
			msg += "apps list: \n"

			for _, app := range configkit.Apps {
				msg += fmt.Sprintf("  - %s\n", app)
			}
			msg += "\n"

			msg += "Version /v1\n"

			w.Write([]byte(msg))
		})
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
