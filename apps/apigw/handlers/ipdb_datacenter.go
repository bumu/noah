package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"apigw/apps/apigw/data/schema"
	ipinfov1 "apigw/gen/go/ipinfo/v1"

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

func (deps registerDeps) CreateIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
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

	schema := schema.IpdbDataCenter{
		Name:      ret.Name,
		Domain:    ret.Domain,
		IpStart:   ret.IpStart,
		IpEnd:     ret.IpEnd,
		CreatedTs: ret.CreatedTs,
		UpdatedTs: ret.UpdatedTs,
	}

	fmt.Println("===", schema)

	/*
		err = deps.IpdbDataCenterRepo.Create(r.Context(), &schema)
		if err != nil {
			w.Write([]byte("get key error\n"))
		}
	*/

	w.Write([]byte("create datacenter"))
}

func (deps registerDeps) UpdateIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update datacenter"))
}

func (deps registerDeps) DeleteIpdbDataCenter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete datacenter"))
}
