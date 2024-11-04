package controllers

import (
	"encoding/json"
	"net/http"

	"noah/apps/infra/data/schema"
	infrav1 "noah/gen/go/infra/v1"

	"github.com/go-chi/render"
)

func (deps registerDeps) CreateInfraHostHeartbeat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("method not allowed\n"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req infrav1.HeartbeatRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Write([]byte("invalid request body\n"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	schema := &schema.InfraHostHeartbeat{
		Hostname:   req.Hostname,
		IP:         req.Ip,
		Arch:       req.Arch,
		OS:         req.Os,
		Kernel:     req.Kernel,
		SystemInfo: req.SystemInfo,
	}

	jsonResp := infrav1.HeartbeatResponse{}

	err = deps.InfraHostHeartbeatRepo.Create(r.Context(), schema)
	if err != nil {
		jsonResp.Status = -1
		jsonResp.ErrMsg = "record error"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, jsonResp)
}
