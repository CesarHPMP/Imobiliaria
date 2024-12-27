package routes

import (
	"imobiliaria_crm/backend/internal/controllers"
	"imobiliaria_crm/backend/internal/middleware"
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
	protectedRouter := router.PathPrefix("/api/protected").Subrouter()
	protectedRouter.Use(middleware.AuthMiddleware)

	// Define routes
	protectedRouter.HandleFunc("/properties", controllers.GetProperties).Methods("GET")
	protectedRouter.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	protectedRouter.HandleFunc("/createProperty", controllers.CreateProperty).Methods("POST")
	router.HandleFunc("/api/addUsers", controllers.AddUser).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	return router, handler
}
