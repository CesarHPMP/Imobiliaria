package controllers

import (
	"context"
	"encoding/json"
	"imobiliaria_crm/backend/database"
	"imobiliaria_crm/backend/utils"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// tabela de usuários
type User struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Endereco string `json:"endereco"`
	Numero   string `json:"numero"`
	Senha    string `json:"senha_hash"`
}

type login struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

type Property struct {
	ID      int    `json:"id"`
	Address string `json:"name"`
}

// Variável global de propriedades
var Properties = []Property{}

// Variável global de usuários
var Users = []User{}

// Retorna propriedades para uma request GET de HTTP
func GetProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Properties)
}

// Retorna Users para uma request POST de HTTP
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

// Cria uma nova propriedade de acordo com o request
func CreateProperty(w http.ResponseWriter, r *http.Request) {
	// Verifica se a requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return // Finaliza o fluxo aqui
	}

	// Lê o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return // Finaliza o fluxo aqui
	}
	defer r.Body.Close()

	// Faz variável auxiliar
	var property Property

	// Deserializa o JSON
	err = json.Unmarshal(body, &property)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusBadRequest)
		return // Finaliza o fluxo aqui
	}

	// Faz o append na lista global
	Properties = append(Properties, property)

	// Responde com sucesso
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Property created successfully",
	})
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

	var ID, storedHash string
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

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(user.Senha))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid credentials",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User logged in",
		"userID":  ID,
	})
}
