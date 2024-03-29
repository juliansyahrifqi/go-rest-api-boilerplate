package routes

import (
	"go-rest-api-boilerplate/handlers"

	"github.com/gorilla/mux"
)

// List routing for home
func SetHomeRoutes(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/home", handlers.HomeHandler).Methods("GET")
}
