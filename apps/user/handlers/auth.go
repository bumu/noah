package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func (deps registerDeps) UserAuth(w http.ResponseWriter, r *http.Request) {

	//ip := r.Header.Get("Fly-Client-IP")
	query := r.URL.Query()
	code := query.Get("code")
	state := query.Get("state")

	user := r.Header.Get("X-Forwarded-User")
	email := r.Header.Get("X-Forwarded-Email")
	groups := r.Header.Get("X-Forwarded-Groups")

	// get all headers and render as json

	msg := fmt.Sprintf("code: %s, state: %s, user: %s, email: %s, groups: %s", code, state, user, email, groups)
	w.Write([]byte(msg))

	headers := r.Header
	render.JSON(w, r, headers)

	w.WriteHeader(http.StatusOK)

	/*

		if code == "" {
			http.Error(w, "Missing code in callback", http.StatusBadRequest)
			return
		}

		if state == "" {
			http.Error(w, "Missing state parameter", http.StatusBadRequest)
			return
		}

		// Exchange code for access token
		token, err := deps.OauthConfig.Exchange(r.Context(), code)
		if err != nil {
			http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
			return
		}

		// Get user info from Google
		client := deps.OauthConfig.Client(r.Context(), token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
		if err != nil {
			http.Error(w, "Failed to get user info", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var userInfo schema.GoogleOauth2
		if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
			http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Received code: %s and state: %s", code, state)

		render.JSON(w, r, userInfo)
	*/
}
