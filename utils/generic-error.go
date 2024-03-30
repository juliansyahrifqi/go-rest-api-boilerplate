package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func ResponseWithError(w http.ResponseWriter, message string, statusCode int) {
	response := ResponseError{
		Message:    message,
		StatusCode: statusCode,
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func ResponseCustomJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
