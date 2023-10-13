package handlers

import (
	"math/rand"
	"net/http"

	"noah/apps/apigw/data/schema"
	userv1 "noah/gen/go/user/v1"

	"github.com/go-chi/render"
)

// Refer: https://google.aip.dev/131
func (desp registerDeps) GetKey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get key"))
}

func (deps registerDeps) ListKeys(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 10

	keys, err := deps.KeyRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	//resp := userv1.ListKeysResponse{Keys: make([]*userv1.UserKey, len(keys))}
	render.JSON(w, r, &keys)
}

func (deps registerDeps) CreateKeys(w http.ResponseWriter, r *http.Request) {
	skey := "sk-" + randStr(32)

	entity := schema.Key{
		Title:  "test",
		Vendor: "openai",
		Token:  skey,
	}

	deps.KeyRepo.Create(r.Context(), &entity)

	resp := userv1.UserKeyResponse{}
	resp.Status = skey

	render.JSON(w, r, &resp)
}

func (deps registerDeps) UpdateKey(w http.ResponseWriter, r *http.Request) {
}

func (deps registerDeps) DeleteKey(w http.ResponseWriter, r *http.Request) {
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
