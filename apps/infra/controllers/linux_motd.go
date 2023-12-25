package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"noah/apps/infra/data/schema"
	infrav1 "noah/gen/go/infra/v1"

	"github.com/go-chi/chi/v5"
)

func (deps registerDeps) CreateVisitRecord(w http.ResponseWriter, r *http.Request) {

	hostname := chi.URLParam(r, "hostname")
	ip := r.Header.Get("Fly-Client-IP")

	action := ""
	switch r.Method {
	case http.MethodGet:
		action = r.URL.Query().Get("action")
	case http.MethodPost:
		p, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("create ipdb datacenter error" + err.Error()))
			return
		}

		println(string(p))

		var ret infrav1.InfraLinuxRequest

		err = json.Unmarshal(p, &ret)
		if err != nil {
			w.Write([]byte("unmarshal post json data error" + err.Error()))
			return
		}

		action = ret.Action
	}

	schema := schema.InfraLinuxMotd{
		IP:       ip,
		Hostname: hostname,
		Action:   action,
	}

	fmt.Println(&schema)

	err := deps.InfraLinuxRepo.Create(r.Context(), &schema)
	if err != nil {
		w.Write([]byte("create client ip error\n"))
	}

	w.Write([]byte("record visit success"))
	// set http status = 200
	w.WriteHeader(http.StatusOK)
}

func (deps registerDeps) Count(w http.ResponseWriter, r *http.Request) {
	cnt, err := deps.InfraLinuxRepo.Count(r.Context())
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	w.Write([]byte(fmt.Sprintf("login count: %d", cnt)))
}
