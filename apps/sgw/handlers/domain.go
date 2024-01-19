package handlers

import (
	"net/http"
)

func (deps registerDeps) ListDomain(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 50

	_, err := deps.GatewayDomainRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	w.Write([]byte("xxx\n"))
}
