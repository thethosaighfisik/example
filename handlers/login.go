package handlers

import (
	"encoding/json"
	"net/http"
	"auth_service/db"
	"auth_service/auth"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	database, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	token, err := auth.AuthenticateUser(database, credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
