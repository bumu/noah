package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func (deps registerDeps) GetIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func (deps registerDeps) ListIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 10

	dcList, err := deps.IpdbDataCenterRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &dcList)
}

func (desp registerDeps) CreateIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create datacenter"))
}

func (deps registerDeps) UpdateIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update datacenter"))
}

func (deps registerDeps) DeleteIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete datacenter"))
}
