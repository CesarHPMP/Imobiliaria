package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Property struct {
	ID      int    `json:"id"`
	Address string `json:"name"`
}

// Variável global de propriedades
var Properties = []Property{}

// Retorna propriedades para uma request GET de HTTP
func GetProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Properties)
}

// Inclui uma nova propriedade para uma request POST de HTTP
func CreateProperty(w http.ResponseWriter, r *http.Request) error {
	// Verifica se a requisição é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		return errors.New("invalid request method")
	}

	// Lê o corpo da requisição
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return err
	}

	// Defer close do body para evitar memória vazia
	defer r.Body.Close()

	// Faz variável auxiliar
	var property Property

	// Deserializa o JSON e em seguida faz o append da propriedade adquirida
	err = json.Unmarshal(body, &property)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON", http.StatusBadRequest)
		return err
	}

	Properties = append(Properties, property)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Property created successfully",
	})

	return nil
}
