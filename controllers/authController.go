package controllers

import (
	"auth-api/models"
	"auth-api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// var users = map[string]string{}

func SignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("Headers:", r.Header)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// users[user.Email] = user.Password

	_, err = utils.AddUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// if pwd, exists := users[user.Email]; !exists || pwd != user.Password {
	// 	http.Error(w, "Invalid email or password", http.StatusUnauthorized)
	// 	return
	// }

	dbUser, err := utils.FindUserByEmail(user.Email)
	if err != nil || dbUser.Password != user.Password {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func Protected(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	// Validate the token
	_, err := utils.ValidateToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Welcome to the protected route"))
}

func RevokeToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	err := utils.RevokeToken(token)
	if err != nil {
		http.Error(w, "Failed to revoke token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Token revoked successfully"))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")
	newToken, err := utils.RefreshToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(newToken))
}
