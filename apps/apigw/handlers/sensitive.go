package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"apigw/apps/apigw/data/schema"
	sensitivev1 "apigw/gen/go/apigw/v1"

	"github.com/go-chi/render"
)

func (deps registerDeps) GetSensitive(w http.ResponseWriter, r *http.Request) {
	ret, err := deps.SensitiveRepo.Get(r.Context(), 1)

	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &ret)
}

func (deps registerDeps) ListSensitive(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 10

	ret, err := deps.SensitiveRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("list key error\n"))
	}

	render.JSON(w, r, &ret)
}

func (deps registerDeps) CreateSensitive(w http.ResponseWriter, r *http.Request) {
	p, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("create sensitive error" + err.Error()))
		return
	}

	println(err == nil, p)

	var ret sensitivev1.SensitiveRequest

	err = json.Unmarshal(p, &ret)
	if err != nil {
		w.Write([]byte("create sensitive error" + err.Error()))
		return
	}

	fmt.Printf("create sensitive +%v\n", ret.Category)
	schema := schema.Sensitive{
		Category:    ret.Category,
		CountryCode: ret.CountryCode,
		Content:     ret.Content,
	}

	err = deps.SensitiveRepo.Create(r.Context(), &schema)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	w.Write([]byte("create sensitive"))
}

func (deps registerDeps) UpdateSensitive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update sensitive"))
}

func (deps registerDeps) DeleteSensitive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete sensitive"))
}
