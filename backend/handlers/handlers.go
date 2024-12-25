package handlers

import (
	"encoding/json"
	"net/http"

	"imobiliaria_crm/backend/utils"
)

// Mock user authentication
var users = map[string]int{"user1": 1, "user2": 2} // username -> userID

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate user (replace with actual authentication logic)
	userID, exists := users[credentials.Username]
	if !exists || credentials.Password != "password" { // Simple check; use hashed passwords in prod
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(userID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Respond with token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
