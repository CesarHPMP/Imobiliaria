package controllers

import (
	"encoding/json"
	"imobiliaria_crm/backend/internal/utils"
	"io"
	"net/http"
)

type Property struct {
	ID      int    `json:"id"`
	Address string `json:"name"`
}

var Properties = []Property{}

// Cria uma nova propriedade de acordo com o request
func CreateProperty(w http.ResponseWriter, r *http.Request) {
	// Verifica se a requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return // Finaliza o fluxo aqui
	}

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	_, err := utils.ValidateJWT(authToken)

	if err != nil {
		http.Error(w, "Invalid JWT", http.StatusUnauthorized)
		return
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

// Retorna propriedades para uma request GET de HTTP
func GetProperties(w http.ResponseWriter, r *http.Request) {
	Properties = []Property{
		{ID: 1, Address: "123 Main St, Springfield"},
		{ID: 2, Address: "456 Elm St, Metropolis"},
		{ID: 3, Address: "789 Oak St, Gotham"},
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	_, err := utils.ValidateJWT(authToken)

	if err != nil {
		http.Error(w, "Invalid JWT"+err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Properties)
}
