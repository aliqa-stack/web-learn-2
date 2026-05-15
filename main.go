package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
    	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Login struct {
	ID           int    `bson:"_id,omitempty"`
	Username     string `bson:"username"`
	Hash         string `bson:"hash"`
	SessionToken string `bson:"session_token"`
	CSRFToken    string `bson:"csrf_token"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("cannot connect to MongoDB:", err)
	}

	fmt.Println("connected to db")


	collection = client.Database("login_site").Collection("users")
	 mux := http.NewServeMux()
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/protected", Protected)

	wrappedMux := enableCORS(mux)

	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		log.Fatal(err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// FIX #3: http.Error butuh int bukan error object
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "cannot parse body", http.StatusBadRequest)
		return
	}

	// FIX #4 & #5: "contect" -> "context", satu deklarasi ctx saja
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	if len(creds.Username) < 8 || len(creds.Password) < 8 {
		http.Error(w, "username and password must be at least 8 characters long", http.StatusNotAcceptable)
		return
	}

	var sameUser Login
	err := collection.FindOne(ctx, bson.M{"username": creds.Username}).Decode(&sameUser)
	// FIX #7: http.Error pakai status code int, bukan err
	if err == nil {
		http.Error(w, "username already exists", http.StatusConflict)
		return
	}

	// FIX #8: hashPassword(Password) -> hashPassword(creds.Password)
	hashedPassword, err := hashPassword(creds.Password)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	newUser := Login{
		Username: creds.Username,
		Hash:     hashedPassword,
	}

	// FIX #9: "Inserone" -> "InsertOne"
	// FIX #10: _, err := -> _, err = (err sudah dideklarasikan di atas)
	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "user registered successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		// FIX #11: `json: "username"` -> `json:"username"` (hapus spasi setelah titik dua)
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// FIX #12: http.Error pakai status code int, bukan err
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "cannot parse body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user Login
	// FIX #13: .Decode(user) -> .Decode(&user) — harus pakai pointer
	err := collection.FindOne(ctx, bson.M{"username": creds.Username}).Decode(&user)
	// FIX #14: Hash.User -> user.Hash
	if err != nil || !checkPasswordHash(creds.Password, user.Hash) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

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
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"username": creds.Username},
		bson.M{"$set": bson.M{
			"session_token": sessionToken,
			"csrf_token":    csrfToken,
		}},
	)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// FIX #15: username tidak terdefinisi -> creds.Username
	fmt.Fprintf(w, "Login successful for user: %s\n", creds.Username)
}

func Protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	if err := Authorize(r); err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	username := r.FormValue("username")
	fmt.Fprintf(w, "Welcome to the protected area, %s!", username)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	username := r.FormValue("username")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, _ = collection.UpdateOne(
		ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{
			"session_token": "",
			"csrf_token":    "",
		}},
	)

	fmt.Fprintln(w, "logged out successfully")
}

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Izinkan origin dari React kamu
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // atau "*" untuk semua
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}