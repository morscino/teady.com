package models

import (
	"fmt"
	"teady.com/utils"
	"net/http"
	//"github.com/globalsign/mgo/bson"
	"time"
)
type Topic struct {
	//ID      bson.ObjectId	`json:"id" bson:"_id"`
	Title   string			`json:"title" bson:"title"`
	Message string			`json:"message" bson:"message"`
	Author  string			`json:"author" bson:"author"`
	Created time.Time		`json:"created" bson:"created"`
}

func TopicsPageLoad(w http.ResponseWriter, r *http.Request){
	utils.ExecuteTemplate(w,"topics.html",nil)
}

func TopicCreateLoad (w http.ResponseWriter, r *http.Request){
	utils.ExecuteTemplate(w,"createTopic.html",nil)
}

func TopicInsert(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	title := r.PostForm.Get("title")
	message := r.PostForm.Get("message")
	author := "admin"
	created := time.Now()

	topic :=&Topic{
		Title : title,
		Message : message,
		Author : author,
		Created : created,
	}

	db, _ := DatabaseConnect()
	// c := db.C("topics")
	// count,err := c.Find(bson.M{"title":topic.Title}).Count()
	//if count > 0 {
		//return with some error...Title already exist
	//}
	err = db.C("topics").Insert(topic)
	//err := InsertCollectionData("topics",topic)
	fmt.Println(err)

	
	http.Redirect(w,r,"/create/topic",302)
}