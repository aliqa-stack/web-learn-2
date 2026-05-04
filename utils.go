package main 

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"encoding/base64"
	"log"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(length int) (string, error) {
	bytes := make([]byte,length)
	if _, er := rand.Read(bytes); er != nil {
		log.Fatalf("failed to generate token: %v", er)
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}