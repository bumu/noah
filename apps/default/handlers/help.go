package handlers

import (
	"fmt"
	"net/http"
	"noah/pkg/configkit"
)

func (deps registerDeps) ListApisHandler(w http.ResponseWriter, r *http.Request) {
	msg := "welcome to noah service\n"
	msg += "Version /v1\n"
	msg += "\n\n"

	msg += "apps list: \n"

	for _, app := range configkit.Apps {
		msg += fmt.Sprintf("  - %s\n", app)
	}
	msg += "\n"

	w.Write([]byte(msg))
}
