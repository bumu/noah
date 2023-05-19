package bootstrap

import (
	"apigw/pkg/configkit"
	"context"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

type serverDeps struct {
	fx.In
}

type Server struct {
	deps *serverDeps

	mux    *chi.Mux
	server *http.Server
}

func New(mux *chi.Mux) *Server {
	server := &Server{
		mux:    mux,
		server: &http.Server{Addr: configkit.GlobalConfig.Addr, Handler: mux},
	}

	return server
}

func NewMux() *chi.Mux {
	mux := chi.NewMux()
	mux.Use(middleware.Logger)
	mux.Use(render.SetContentType(render.ContentTypeHTML))

	return mux
}

// Run the server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Stop the server
func (s *Server) Stop() error {
	return s.server.Shutdown(context.TODO())
}
