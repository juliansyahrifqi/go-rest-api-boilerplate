package routes

import (
	"go-rest-api-boilerplate/handlers"

	"github.com/gorilla/mux"
)

// List routing for home
func SetHomeRoutes(apiRouter *mux.Router) {
	homeHandler := handlers.NewAuthHandler()

	apiRouter.HandleFunc("/home", homeHandler.HomeHandler).Methods("GET")
}
