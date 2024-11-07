package handlers

import (
	"fmt"
	"net/http"
	"noah/pkg/configkit"
)

func (deps registerDeps) DefaultHandler(w http.ResponseWriter, r *http.Request) {

	cip := r.Header.Get("Fly-Client-IP")
	if cip == "" {
		cip = r.RemoteAddr
	}

	w.Write([]byte(fmt.Sprintf("%s\n", cip)))

	headers := r.Header
	w.Write([]byte(fmt.Sprintf("%s\n", headers)))
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
