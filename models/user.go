package models

import (
	"net/http"
	"time"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"teady.com/sessions"
	)

type User struct{
	//ID int
	Lastname string		`json:"lastname" bson:"lastname"`
	Firstname string	`json:"firstname" bson:"firstname"`
	Username string		`json:"username" bson:"username"`
	Password string		`json:"password" bson:"password"`
	Email	string		`json:"email" bson:"email"`
	DOB string		`json:"dob" bson:"dob"`
	Created time.Time		`json:"created" bson:"created"`
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	lastname := r.PostForm.Get("lastname")
	firstname := r.PostForm.Get("firstname")
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	email := r.PostForm.Get("email")
	dob := r.PostForm.Get("dob")
	created := time.Now()
	
	user := &User{
		Lastname : lastname,
		Firstname : firstname,
		Username : username,
		Password : password,
		Email : email,
		DOB : dob,
		Created : created,
	}
	//check if username has already been used
	//check if email has already been used
	//check if password matches
	db, _ := DatabaseConnect()
	err = db.C("users").Insert(user)

	fmt.Println(err)
	http.Redirect(w,r,"/user/login",302)
}

func UserLogin(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	user := User{}

	//Authenticate User
	db, _ := DatabaseConnect()
	err := db.C("users").Find(bson.M{"username":username}).One(&user)

	if err != nil{
		http.Redirect(w,r,"/user/login",302)
	}

	if password != user.Password{
		http.Redirect(w,r,"/user/login",302)
	}


	//Create User Session
	session,_ := sessions.Store.Get(r,"session")
	session.Values["username"] = username
	session.Save(r,w)

	fmt.Println("User has been Logged in")
	http.Redirect(w,r,"/",302)
}
