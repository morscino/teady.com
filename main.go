package main

import (
	"teady.com/utils"
	"net/http"

	"teady.com/routes"
)

func main() {
	utils.LoadTemplate("templates/*.html")
	router := routes.NewRouter()

	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static").Handler(http.StripPrefix("/static",fs))
	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)

}
