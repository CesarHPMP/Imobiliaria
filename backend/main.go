package main

import (
	"log"
	"net/http"

	"imobiliaria_crm/backend/internal/config"
	"imobiliaria_crm/backend/internal/database"
	"imobiliaria_crm/backend/internal/routes"
)

func main() {
	// Load configuration
	configs := config.LoadConfig()

	// Connect to the database
	err := database.Connect(configs)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize router
	_, handler := routes.NewRouter()

	// Start the server
	log.Println("Server is running on port 8080...")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
