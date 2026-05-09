package main

import (
	"log"
	"os"
	"net/http"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

type Login struct {
	ID 		 int     	`bson: "_id,omtempty"`
	Username string  	`bson:"username"`
	Hash     string  	`bson:"hash"`
	SessionToken string	`bson:"session_token"`
    CSRFToken string	`bson:"csrf_token"`
}

var collection *mongo.Collection
var users = map[string]Login{}
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	mongoURI := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(option.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}


	_ = client.Ping(ctx, readpref.Primary())

	fmt.Println("connected to db")

	collecltions = client.Database("login site").Collection("users")

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

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "cannot find body", err)
		return
	}

	ctx, cancel := contect.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUsers User
	err := collection.FindOne(ctx, bson.M{"Username": creds.Username}).Decode(&existingUser)
	if err != nil {
		http.Error(w, "invalid username or password", err)
		return
	}

	if len(username) <8 || len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "username and password must be at least 8 characters long", er)
		return
}

    var sameUser User
	 err := collection.FindOne(ctx, bson.M{"Username": creds.Username}).Decode(&sameUser)
	 if err == nil  {
	
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

func Protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "invalid method", er)
		return
	}

	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "unauthorized", er)
		return
	}

	username := r.FormValue("username")
	fmt.Fprintf(w, "Welcome to the protected area, %s!", username)
}

func logout(w http.ResponseWriter, r *http.Request) { 
	if err:= Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "unauthorized", er)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name: "csrf_token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	username := r.FormValue("username")
	user, _ := users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	users[username] = user



}

