package controllers

import (
	"encoding/json"
	"io"
	"net/http"
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
	// Verifica se a requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return
	}

	// Lê o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Defer close do body para evitar memória vazia
	defer r.Body.Close()

	// Faz variável auxiliar
	var user User

	// Deserializa o JSON e em seguida faz o append da propriedade adquirida
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusBadRequest)
		return
	}

	Users = append(Users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Property created successfully",
	})
}
