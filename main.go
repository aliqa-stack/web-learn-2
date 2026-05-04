package main

import (
	"log"
	"net/http"
	"fmt"
	"time"
)

type Login struct {
	Hash     string
	SessionToken string
    CSRFToken string
}

var users = map[string]Login{}
func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	//http.HandleFunc("/logout", logout)
	//http.HandleFunc("/protected", Protected)
	http.ListenAndServe(":8080", nil)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}


}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "invalid method", er)
		return 
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) <8 || len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "username and password must be at least 8 characters long", er)
		return
}


	if _ , ok := users[username]; ok {
		er := http.StatusConflict
		http.Error(w, "username already exists", er)
		return

	}

	hashPassword, err := hashPassword(password)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	users[username] = Login {
		Hash: hashPassword,
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "invalid method", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := users[username]
	if !ok || !checkPasswordHash(password, user.Hash) {
		er := http.StatusUnauthorized
		http.Error(w, "invalid credentials", er)
		return
	}

    // Generate session and CSRF tokens
	sessionToken, err := generateToken(32)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	csrfToken, err := generateToken(32)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name: "csrf_token",
		Value: csrfToken,
		Expires: time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	fmt.Printf("Login successful for user: %s\n", username)
	
}

