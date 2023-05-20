package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func (deps registerDeps) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func (deps registerDeps) ListUser(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 10

	users, err := deps.UserRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &users)
}

func (desp registerDeps) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create user"))
}

func (deps registerDeps) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}

func (deps registerDeps) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
