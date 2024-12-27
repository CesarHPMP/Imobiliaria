package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"imobiliaria_crm/backend/internal/database"
	"imobiliaria_crm/backend/internal/utils"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type login struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user login
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	transaction, err := db.PostgresDB.Begin(context.Background())
	if err != nil {
		http.Error(w, "Failed to start transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer transaction.Rollback(context.Background())

	query := `SELECT id, senha_hash FROM users WHERE email = $1`
	row := transaction.QueryRow(context.Background(), query, user.Email)

	var ID *int
	var storedHash *sql.NullString

	err = row.Scan(&ID, &storedHash)

	if err != nil {
		if err.Error() == "no rows in result set" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid credentials",
			})
		} else {
			http.Error(w, "Failed to scan row: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if storedHash.String == "" || !storedHash.Valid {
		http.Error(w, "failed to store scanned row", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash.String), []byte(user.Senha))
	if err != nil {
		println("passwords are different", err, storedHash.String, user.Senha)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateJWT(*ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "User logged in",
		"userToken": token,
	})
}
