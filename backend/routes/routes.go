package routes

import (
	"github.com/CesarHPMP/imobiliaria/backend/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/properties", controllers.GetProperties).Methods("GET")
	return router
}
