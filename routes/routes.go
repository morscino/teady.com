package routes

import (
	"github.com/gorilla/mux"
	"teady.com/models"
)

func NewRouter() *mux.Router {

	//This is the fuction that handles all of the application's routing
	router := mux.NewRouter()

	router.HandleFunc("/", models.HomePageLoad).Methods("GET")
	router.HandleFunc("/Contact", models.HomeContactLoad).Methods("GET")
	
	router.HandleFunc("/register/user", models.RegisterPageLoad).Methods("GET")
	router.HandleFunc("/user/login", models.LoginPageLoad).Methods("GET")
	router.HandleFunc("/topics", models.TopicsPageLoad).Methods("GET")
	router.HandleFunc("/create/topic", models.TopicCreateLoad).Methods("GET")
	

	router.HandleFunc("/RegisterUser", models.UserRegister).Methods("POST")

	return router
}
