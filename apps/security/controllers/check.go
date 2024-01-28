package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"noah/apps/security/data/schema"
	secv1 "noah/gen/go/security/v1"

	"github.com/go-chi/render"
)

func (deps registerDeps) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}

func (deps registerDeps) CheckSec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Write([]byte("method not allowed\n"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req secv1.SecCheckRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Write([]byte("invalid request body\n"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("debug===", &req)

	schema := schema.SecurityCheckRequest{
		RequestID: req.RequestId,
		// RemoteAddr:     req.RemoteAddr,
		UserAgent:      req.UserAgent,
		FingerprintJa3: req.FingerprintJa3,
		FingerprintH2:  req.FingerprintH2,
		Protocol:       req.Protocol,
		Method:         req.Method,
		CheckURL:       req.Url,
		CheckHeader:    req.ReqHeaders,
		CheckBody:      req.ReqBody,
	}

	err = deps.SecurityCheckRequestRepo.Create(r.Context(), &schema)
	if err != nil {
		w.Write([]byte("record error\n"))
	}

	w.Write([]byte("record success"))
	w.WriteHeader(http.StatusOK)
}

func (deps registerDeps) LastSec(w http.ResponseWriter, r *http.Request) {
	sec, err := deps.SecurityUserAgentRepo.Last(r.Context())
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &sec)
}
