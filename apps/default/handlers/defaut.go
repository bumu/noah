package handlers

import (
	"fmt"
	"net"
	"net/http"
	"noah/pkg/configkit"
)

func (deps registerDeps) DefaultHandler(w http.ResponseWriter, r *http.Request) {
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

	w.Write([]byte(fmt.Sprintf("%s\n", cip)))

	if r.Header.Get("X-Debug") != "" {
		headers := r.Header
		w.Write([]byte(fmt.Sprintf("%s\n", headers)))
	}
}

func (deps registerDeps) ListApisHandler(w http.ResponseWriter, r *http.Request) {
	msg := "welcome to noah service\n"
	msg += "apps list: \n"

	for _, app := range configkit.Apps {
		msg += fmt.Sprintf("  - %s\n", app)
	}
	msg += "\n"

	msg += "Version /v1\n"

	w.Write([]byte(msg))
}
