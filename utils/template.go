package utils

import (
	"net/http"
	"html/template"
	
)

var templates *template.Template
func LoadTemplate (pattern string){
	//This is the method that loads all template
	templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate (w http.ResponseWriter, tpl string, data interface{}){
	templates.ExecuteTemplate(w,tpl,data)
}