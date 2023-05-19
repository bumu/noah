package handlers

import (
	"log"
	"net/http"
	"strings"

	userv1 "apigw/gen/go/user/v1"

	"github.com/go-chi/render"
)

func (deps registerDeps) HandleUser(w http.ResponseWriter, r *http.Request) {
	log.Println("xxx")
	skey := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	token := skey

	key, err := deps.KeyRepo.First(r.Context(), token)

	if err == nil && key != nil {
		key.Token = "111"
	}

	resp := userv1.UserKeyResponse{}
	resp.Status = "1111"

	render.JSON(w, r, &resp)
}
