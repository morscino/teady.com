package models

import (
	"teady.com/utils"
	//"fmt"
	"net/http"
)

func TestFunc(w http.ResponseWriter, request *http.Request) {
// This is just a test function
	//fmt.Fprint(w, "HELLO APP")
	utils.ExecuteTemplate(w, "index.html",nil)
}
