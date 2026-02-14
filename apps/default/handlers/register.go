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
		r.Get("/help", deps.ListApisHandler)
		r.Get("/apis", deps.ListApisHandler)
		r.Get("/info", deps.EchoHandler)
		r.Post("/upload", deps.UploadHandler)
		r.Get("/downloads/", deps.DownloadHandler)
		r.Get("/dl/", deps.DownloadHandler)

		// 兼容 IPv4 + IPv6 的路由参数（chi 支持复合正则）
		r.Get("/{ip:([0-9.]+)|([0-9a-fA-F:]+)}/json", deps.IpJsonHandler)
		r.Get("/json", deps.IpJsonHandler)

		r.Get("/echo", deps.EchoHandler)
		r.Post("/echo", deps.EchoHandler)
		r.Get("/*", deps.NotFoundHandler)
		r.Post("/*", deps.NotFoundHandler)

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
