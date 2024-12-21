package main

import (
	"encoding/json"
	"imobiliaria_crm/backend/config"
	"imobiliaria_crm/backend/database"
	"imobiliaria_crm/backend/routes"
	"net/http"
)

// Property represents a property structure
type Property struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Sample data
var properties = []Property{
	{ID: 1, Name: "Property One"},
	{ID: 2, Name: "Property Two"},
}

// GetProperties handles GET requests to fetch properties
func GetProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(properties)
}

func main() {
	configs := config.LoadConfig()

	database.Connect(configs)
	router := routes.NewRouter()
	router.HandleFunc("/api/properties", GetProperties).Methods("GET")

	// Start the server
	http.ListenAndServe(":8080", router)
}
