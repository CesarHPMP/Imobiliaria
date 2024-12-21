package main

import (
	"log"
	"net/http"

	"imobiliaria_crm/backend/config"
	"imobiliaria_crm/backend/controllers"
	"imobiliaria_crm/backend/database"
	"imobiliaria_crm/backend/routes"
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
	router := routes.NewRouter()

	// Define routes
	router.HandleFunc("/api/properties", controllers.GetProperties).Methods("GET")
	router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/addUsers", controllers.AddUser).Methods("POST")
	router.HandleFunc("/api/createProperty", controllers.CreateProperty).Methods("POST")

	/*
		Http request does not end after return, instead returns "Property created successfully" in response, which
		would skip errors in frontend (Major security issue)
	*/

	// Start the server
	log.Println("Server is running on port 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
