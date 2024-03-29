package handlers

import (
	"encoding/json"
	"go-rest-api-boilerplate/models"
	"net/http"
)

// Example return JSON
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Home{
		Message: "Hello World",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
