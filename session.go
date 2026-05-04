package main 

import (
	"errors"
	"net/http"
)

// Authorize checks if the request has valid session and CSRF tokens


var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) Error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return AuthError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return AuthError

	}

	csrf :=r.Header.Get("X-CSRF-Token")
	if csrf == "" || csrf != user.CSRFToken {
		return AuthError
	}
	return nil
}