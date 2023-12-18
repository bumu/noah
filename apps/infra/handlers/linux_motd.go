package handlers

import (
	"net/http"
	"noah/apps/infra/data/schema"

	"github.com/go-chi/chi/v5"
)

func (deps registerDeps) CreateVisitRecord(w http.ResponseWriter, r *http.Request) {

	hostname := chi.URLParam(r, "hostname")
	//err = deps.IpdbClientIpRepo.Create(r.Context(), &schema)

	schema := schema.InfraLinuxMotd{
		IP:       r.RemoteAddr,
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
