package controllers

import (
	"context"
	"encoding/json"
	"imobiliaria_crm/backend/database"
	"imobiliaria_crm/backend/utils"
	"io"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Endereco string `json:"endereco"`
	Numero   string `json:"numero"`
	Senha    string `json:"senha_hash"`
}

// Retorna Users para uma request POST de HTTP
func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	db := database.GetDB()

	query := "Select * from users;"

	row := db.PostgresDB.QueryRow(context.Background(), query, 1)

	var Users = []User{}
	err := row.Scan(Users)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(Users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	// Verify if the request is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Defer close of body to free memory
	defer r.Body.Close()

	// Auxiliary variable
	var user User

	// Deserialize JSON
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusBadRequest)
		return
	}

	// Hash the password before storing it
	hashedPassword, hashErr := utils.HashPassword(user.Senha)
	if hashErr != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Senha = hashedPassword

	// Get database instance
	db := database.GetDB()

	db.PostgresDB.QueryRow(context.Background(), "SELECT")
	// Start transaction
	transaction, err := db.PostgresDB.Begin(context.Background())
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	// Rollback if there is an error
	defer transaction.Rollback(context.Background())

	// SQL command for insertion
	query := `INSERT INTO users (nome, email, endereco, numero, senha_hash) VALUES ($1, $2, $3, $4, $5);`

	// Execute query
	_, err = transaction.Exec(context.Background(), query, user.Nome, user.Email, user.Endereco, user.Numero, user.Senha)
	if err != nil {
		http.Error(w, "Failed to execute query: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Commit transaction
	if commitErr := transaction.Commit(context.Background()); commitErr != nil {
		http.Error(w, "Failed to commit transaction: "+commitErr.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}
