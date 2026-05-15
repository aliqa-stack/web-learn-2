package main 

import (
	"errors"
	"net/http"
		"context"
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Authorize checks if the request has valid session and CSRF tokens


var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	username := r.FormValue("username")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	var user Login 
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
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