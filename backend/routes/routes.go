package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func EnableCORS(router *mux.Router) http.Handler {
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)
	return handler
}

func NewRouter() (*mux.Router, http.Handler) {
	router := mux.NewRouter()
	handler := EnableCORS(router)
	return router, handler
}
