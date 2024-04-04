package handlers

import (
	"encoding/json"
	"go-rest-api-boilerplate/models"
	"net/http"
)

type authHandler struct{}

func NewAuthHandler() *authHandler {
	return &authHandler{}
}

// Example return JSON
func (h *authHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Home{
		Message: "Hello World",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
