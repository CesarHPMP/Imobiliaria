package main

import (
	"imobiliaria_crm/backend/routes"
	"log"
	"net/http"
)

func main() {
	// Set up routes
	r := routes.NewRouter()

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
