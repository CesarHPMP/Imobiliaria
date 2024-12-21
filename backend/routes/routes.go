package routes

import (
	"net/http"

	"imobiliaria_crm/backend/controllers"

	"github.com/gorilla/mux"
)

// Essa função permite requests de origens diferentes
func enableCORS(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			next.ServeHTTP(w, r)
		})
	})
}

// Gera um router novo com CORS aplicado
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	enableCORS(router) // Enable CORS
	router.HandleFunc("/api/properties", controllers.GetProperties).Methods("GET")
	return router
}
