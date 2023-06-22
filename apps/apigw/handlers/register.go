package handlers

import (
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
	KeyRepo       *repos.KeyRepo
	UserRepo      *repos.UserRepo
	SensitiveRepo *repos.SensitiveRepo
	// Checker *sensitivemod.Checker
}

func Register(deps registerDeps) {
	deps.Mux.Route("/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to apigw service\n"))
		})
		r.Handle("/metrics", promhttp.Handler())
		r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong\n"))
		})
	})

	// Refer: https://google.aip.dev/158
	// method order: get, list, create, update, delete.
	deps.Mux.Route("/v1", func(r chi.Router) {
		// User api management
		r.Get("/user", deps.GetUser)
		r.Get("/user/keys", deps.GetHostKey)
		r.Get("/user/list", deps.ListUser)
		r.Post("/user/create", deps.CreateUser)
		r.Post("/user/update", deps.UpdateUser)
		r.Post("/user/delete", deps.DeleteUser)

		// Key api management
		r.Get("/user/key", deps.GetKey)
		r.Get("/user/key/list", deps.ListKeys)
		r.Post("/user/key/create", deps.CreateKeys)
		r.Post("/user/key/update", deps.UpdateKey)
		r.Post("/user/key/delete", deps.DeleteKey)

		r.Get("/sensitive", deps.GetSensitive)
		r.Get("/sensitive/list", deps.ListSensitive)
		r.Post("/sensitive/create", deps.CreateSensitive)
		r.Post("/sensitive/update", deps.UpdateSensitive)
		r.Post("/sensitive/delete", deps.DeleteSensitive)

		r.HandleFunc("/openai/*", deps.HandleOpenai)
	})
}
