package controllers

import (
	"net/http"
	"noah/apps/infra/data/schema"

	"github.com/go-chi/chi/v5"
)

func (deps registerDeps) CreateVisitRecord(w http.ResponseWriter, r *http.Request) {

	hostname := chi.URLParam(r, "hostname")
	ip := r.Header.Get("Fly-Client-IP")

	schema := schema.InfraLinuxMotd{
		IP:       ip,
		Hostname: hostname,
	}

	err := deps.InfraLinuxRepo.Create(r.Context(), &schema)
	if err != nil {
		w.Write([]byte("create client ip error\n"))
	}

	w.Write([]byte("record visit success"))
	// set http status = 200
	w.WriteHeader(http.StatusOK)
}
