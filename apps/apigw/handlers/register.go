package handlers

import (
	"net/http"

	"apigw/apps/apigw/data/repos"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux
	//Logger  *slog.Logger
	KeyRepo            *repos.KeyRepo
	UserRepo           *repos.UserRepo
	SensitiveRepo      *repos.SensitiveRepo
	IpdbDataCenterRepo *repos.IpdbDataCenterRepo
	IpdbClientIpRepo   *repos.IpdbClientIpRepo
	UseragentRepo      *repos.UseragentRepo
	UseragentOSRepo    *repos.UseragentOSRepo

	// Checker *sensitivemod.Checker
}

func Register(deps registerDeps) {
	deps.Mux.Route("/", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to apigw service\n"))
		})
		// r.Handle("/metrics", promhttp.Handler())
		r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong\n"))
		})
	})

	// Internal management apis.
	deps.Mux.Route("/internal", func(r chi.Router) {
		r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok\n"))
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

		r.HandleFunc("/ipdb/datacenter", deps.GetIpdbDataCenter)
		r.HandleFunc("/ipdb/datacenter/list", deps.ListIpdbDataCenter)
		r.HandleFunc("/ipdb/datacenter/create", deps.CreateIpdbDataCenter)
		r.HandleFunc("/ipdb/datacenter/update", deps.UpdateIpdbDataCenter)

		r.HandleFunc("/ipdb/clientip", deps.GetIpdbClientIp)
		r.HandleFunc("/ipdb/clientip/list", deps.ListIpdbClientIp)
		r.HandleFunc("/ipdb/clientip/create", deps.CreateIpdbClientIp)
		r.HandleFunc("/ipdb/clientip/update", deps.UpdateIpdbClientIp)
	})
}
