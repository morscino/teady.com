package main

import (
	"net/http"

	"teady.com/routes"
)

func main() {
	router := routes.NewRouter()

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)

}
