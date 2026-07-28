package handlers

import (
	"noah/apps/company/data/repos"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type registerDeps struct {
	fx.In

	Mux *chi.Mux
	//Logger  *slog.Logger
	CompanyProfileRepo *repos.CompanyProfileRepo
	CompanyUploadRepo  *repos.CompanyUploadRepo
}

func Register(deps registerDeps) {
	deps.Mux.Post("/apis/company/profile/sg-business:upload", deps.UploadSGBusiness)

	deps.Mux.Route("/apis/v1/company", func(r chi.Router) {
		// User api management
		r.Get("/", deps.ListCompanyProfile)
		r.Post("/create", deps.CreateCompanyProfile)
	})
}
