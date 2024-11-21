package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux

	// OauthConfig *oauth2.Config

	// GatewayDomainRepo *repos.GatewayDomainRepo
}

func DefaultRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to user service\n"))
}

// Refer: https://google.aip.dev/158
func Register(deps registerDeps) {
	deps.Mux.Route("/apis/users", func(r chi.Router) {
		r.HandleFunc("/", DefaultRouter)

		r.HandleFunc("/oauth2/v1/callback", deps.UserAuth)

	})

	deps.Mux.HandleFunc("/apis/oauth2/v1/callback", deps.UserAuth)
}
