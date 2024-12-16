package controllers

import (
	"encoding/json"
	"net/http"
)

type Property struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var properties = []Property{
	{ID: 1, Name: "Property 1"},
	{ID: 2, Name: "Property 2"},
}

func GetProperties(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(properties)
}
