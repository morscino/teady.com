package main

import (
	"teady.com/utils"
	"net/http"

	"teady.com/routes"
)

func main() {
	utils.LoadTemplate("templates/*.html")
	router := routes.NewRouter()

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)

}
