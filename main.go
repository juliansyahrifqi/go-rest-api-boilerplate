package main

import (
	"go-rest-api-boilerplate/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Connect DB
	// storage.InitDB()

	// Setup Router using Gorilla MUX
	router := mux.NewRouter()

	// Set Group/Prefix Route
	apiRouter := router.PathPrefix("/api").Subrouter()

	// List All Routes
	routes.SetHomeRoutes(apiRouter)

	// Run Server
	http.ListenAndServe(":8080", router)
}
