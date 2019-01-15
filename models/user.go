package models

import "net/http"

type User struct{
	ID int
	Lastname string
	Firstname string
	Username string
	Password string
	DOB int
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// lastname := r.PostForm.Get("lastname")
	// firstname := r.PostForm.Get("firstname")
	// firstname := r.PostForm.Get("firstname")
	

}
