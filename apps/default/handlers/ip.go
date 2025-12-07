package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"noah/pkg/ipkit"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type IpResponse struct {
	ClientIp string      `json:"client_ip"`
	SearchIp string      `json:"search_ip,omitempty"`
	Headers  http.Header `json:"headers"`
	Body     string      `json:"body,omitempty"`
	Method   string      `json:"method,omitempty"`
	Url      string      `json:"url,omitempty"`
	Args     url.Values  `json:"args,omitempty"`
	// IpInfo   interface{} `json:"ip_info"`
	IpInfo IpInfo `json:"ip_info,omitempty"`
}

type IpInfo struct {
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
	ISP     string `json:"isp"`
}

func (deps registerDeps) IpHandler(w http.ResponseWriter, r *http.Request) {
	// By default, remote address is the client ip.
	cip := ipkit.GetRealIP(r)

	// Use X-Real-Ip if it exists.
	xRealIp := r.Header.Get("X-Real-Ip")
	if xRealIp != "" {
		cip = xRealIp
	}

	// User direct access to the fly service.
	flyClientIp := r.Header.Get("Fly-Client-IP")
	if xRealIp == "" && flyClientIp != "" {
		cip = flyClientIp
	}

	// w.Write([]byte(fmt.Sprintf("%s\n", cip)))

	if r.Header.Get("X-Debug") != "" {
		headers := r.Header
		w.Write([]byte(fmt.Sprintf("%s\n", headers)))
	}

	searchIp := chi.URLParam(r, "search_ip")
	if searchIp == "" {
		searchIp = cip
	}

	// ipkit.IpdbInit()

	// ipinfo, _ := ipkit.IpdbFind(searchIp)

	msg := &IpResponse{
		ClientIp: cip,
		Headers:  r.Header,
		// Body:    configkit.GlobalConfig.Ipdb.Ipv4,
		// IpInfo:  ipinfo,
	}
	render.JSON(w, r, msg)
}

func (deps registerDeps) IpJsonHandler(w http.ResponseWriter, r *http.Request) {
	// By default, remote address is the client ip.
	cip := ipkit.GetRealIP(r)

	searchIp := chi.URLParam(r, "ip")
	if searchIp != "" {

	}

	msg := &IpResponse{
		ClientIp: cip,
		SearchIp: searchIp,
		Headers:  r.Header,
		// IpInfo:   "",
	}

	render.JSON(w, r, msg)
}
