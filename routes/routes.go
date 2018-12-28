package routes

import (
	"github.com/gorilla/mux"
	"teady.com/models"
)

func NewRouter() *mux.Router {
	
	//This is the fuction that handles all of the application's routing
	router := mux.NewRouter()

	router.HandleFunc("/", models.TestFunc)

	return router
}
