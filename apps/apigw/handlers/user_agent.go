package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"apigw/apps/apigw/data/schema"
	ipinfov1 "apigw/gen/go/ipinfo/v1"

	"github.com/go-chi/render"
)

func (deps registerDeps) GetUseragent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func (deps registerDeps) ListUseragent(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 10

	dcList, err := deps.UseragentRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &dcList)
}

func (deps registerDeps) CreateUseragent(w http.ResponseWriter, r *http.Request) {
	p, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("create ipdb datacenter error" + err.Error()))
		return
	}

	println(err == nil, p)

	var ret ipinfov1.IPDataCenterRequest

	err = json.Unmarshal(p, &ret)
	if err != nil {
		w.Write([]byte("create ipdb datacenter error" + err.Error()))
		return
	}

	schema := schema.Useragent{}

	err = deps.UseragentRepo.Create(r.Context(), &schema)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	w.Write([]byte("create datacenter"))
}

func (deps registerDeps) UpdateUseragent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update datacenter"))
}

func (deps registerDeps) DeleteUseragent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete datacenter"))
}
