package models

import (
	"teady.com/utils"
	//"fmt"
	"net/http"
)

func HomePageLoad(w http.ResponseWriter, request *http.Request) {
	// This is just a test function
	//fmt.Fprint(w, "HELLO APP")
	utils.ExecuteTemplate(w, "index.html", nil)
}

func HomeContactLoad(w http.ResponseWriter, request *http.Request) {
	// This is just a test function
	//fmt.Fprint(w, "HELLO APP")
	utils.ExecuteTemplate(w, "contact.html", nil)
}

func RegisterPageLoad (w http.ResponseWriter, request *http.Request){
	utils.ExecuteTemplate(w, "register.html", nil)
}

func LoginPageLoad (w http.ResponseWriter, request *http.Request){
	utils.ExecuteTemplate(w, "login.html", nil)
}
