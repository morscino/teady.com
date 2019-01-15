package models

import (
	"teady.com/utils"
	"net/http"
)
type Topic struct {
	ID      int
	Title   string
	Message string
	Author  string
	Created int
}

func TopicsPageLoad(w http.ResponseWriter, r *http.Request){
	utils.ExecuteTemplate(w,"topics.html",nil)
}

func TopicCreateLoad (w http.ResponseWriter, r *http.Request){
	utils.ExecuteTemplate(w,"createTopic.html",nil)
}
func TopicCreate(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
}