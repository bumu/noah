package handlers

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"noah/pkg/configkit"
	"noah/pkg/ipkit"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type IpResponse struct {
	Cip      string      `json:"cip"`
	SearchIp string      `json:"search_ip"`
	Headers  http.Header `json:"headers"`
	Body     string      `json:"body"`
	Method   string      `json:"method"`
	Url      string      `json:"url"`
	Args     url.Values  `json:"args"`
	IpInfo   interface{} `json:"ip_info"`
}

func (deps registerDeps) IpHandler(w http.ResponseWriter, r *http.Request) {
	// By default, remote address is the client ip.
	cip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Invalid remote address", http.StatusInternalServerError)
		return
	}

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

	ipkit.IpdbInit()

	ipinfo, _ := ipkit.IpdbFind(searchIp)

	msg := &IpResponse{
		Cip:     cip,
		Headers: r.Header,
		Body:    configkit.GlobalConfig.Ipdb.Ipv4,
		IpInfo:  ipinfo,
	}
	render.JSON(w, r, msg)
}
