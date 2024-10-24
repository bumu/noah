package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"noah/pkg/configkit"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

/*
type serverDeps struct {
	fx.In
}
*/

type Server struct {
	// deps *serverDeps

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

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"https://*", "http://*"},
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return mux
}

// Run the server
func (s *Server) Start() error {
	fmt.Println("server start ", s.server.Addr)

	return s.server.ListenAndServe()
}

// Stop the server
func (s *Server) Stop() error {
	return s.server.Shutdown(context.TODO())
}
