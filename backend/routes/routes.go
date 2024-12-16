package routes

import (
	"imobiliaria_crm/backend/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/properties", controllers.GetProperties).Methods("GET")
	return router
}
