package main

import (
	"net/http"
	"fmt"
	"teady.com/models"
	"teady.com/routes"
	"teady.com/utils"
)

func main() {
	db, err := models.DatabaseConnect()
	if db != nil {
		fmt.Println("Connected to Database")
	}else if err != nil{
		fmt.Println("Connection to Database failed")
	}

	utils.LoadTemplate("templates/*.html")

	router := routes.NewRouter()

	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", fs))

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)

}
