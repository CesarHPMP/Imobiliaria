package main

import (
	"log"
	"net/http"

	"imobiliaria_crm/backend/config"
	"imobiliaria_crm/backend/controllers"
	"imobiliaria_crm/backend/database"
	"imobiliaria_crm/backend/middleware"
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
	router, handler := routes.NewRouter()

	protectedRouter := router.PathPrefix("/api/protected").Subrouter()
	protectedRouter.Use(middleware.AuthMiddleware)

	// Define routes
	protectedRouter.HandleFunc("/properties", controllers.GetProperties).Methods("GET")
	protectedRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	protectedRouter.HandleFunc("/createProperty", controllers.CreateProperty).Methods("POST")
	router.HandleFunc("/api/addUsers", controllers.AddUser).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	// Start the server
	log.Println("Server is running on port 8080...")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
