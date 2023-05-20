package handlers

import (
	"math/rand"
	"net/http"

	"apigw/apps/apigw/data/schema"
	userv1 "apigw/gen/go/user/v1"

	"github.com/go-chi/render"
)

func (deps registerDeps) GenerateNewKeys(w http.ResponseWriter, r *http.Request) {
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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
