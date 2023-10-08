package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"apigw/apps/apigw/data/schema"
	ipinfov1 "apigw/gen/go/ipinfo/v1"
	"apigw/pkg/ipkit"

	"github.com/go-chi/render"
)

func (deps registerDeps) GetIpdbClientIp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func (deps registerDeps) ListIpdbClientIp(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 10

	dcList, err := deps.IpdbClientIpRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &dcList)
}

func (deps registerDeps) CreateIpdbClientIp(w http.ResponseWriter, r *http.Request) {
	p, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("create ipdb datacenter error" + err.Error()))
		return
	}

	println(err == nil, p)

	var ret ipinfov1.IpdbClientIpRequest

	err = json.Unmarshal(p, &ret)
	if err != nil {
		w.Write([]byte("create ipdb datacenter error" + err.Error()))
		return
	}

	schema := schema.IpdbClientIp{
		Ip:        ret.Ip,
		IspDomain: ret.IspDomain,
		IspName:   ret.IspName,
		IpCity:    ret.IpCity,
		Cnt:       ret.Cnt,
	}

	ip := ipkit.Ip2Int(ret.Ip)
	dc, err := deps.IpdbDataCenterRepo.GetDataCenter(r.Context(), ip)
	if err != nil {
		// w.Write([]byte("get datacenter error\n"))
		// return
		fmt.Println("errror", err)
	}

	// if datacenter is not 0, then means it is a datacenter ip.
	if dc.ID != 0 {
		schema.IsDataCenter = true
	}

	fmt.Println("insert to db===", schema)
	err = deps.IpdbClientIpRepo.Create(r.Context(), &schema)
	if err != nil {
		w.Write([]byte("create client ip error\n"))
	}

	w.Write([]byte("create client ip success"))
}

func (deps registerDeps) UpdateIpdbClientIp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update datacenter"))
}

func (deps registerDeps) DeleteIpdbClientIp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete datacenter"))
}
