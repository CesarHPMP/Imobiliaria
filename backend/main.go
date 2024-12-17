package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Set up routes
	r := mux.NewRouter()

	// Start the server
	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
