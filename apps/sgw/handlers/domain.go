package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func (deps registerDeps) ListDomain(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 50

	domains, err := deps.GatewayDomainRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &domains)
}
