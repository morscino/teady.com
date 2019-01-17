package models

import (
	"sync"

	"github.com/globalsign/mgo"
	"teady.com/config"
)

var connection *mgo.Database
var err error

var once sync.Once

func StartDbSession() (*mgo.Session, error) {
	DbSession, err := mgo.Dial(config.URL)
	if err != nil {
		return nil, err
	}
	//defer DbSession.Close()
	return DbSession, nil
}
func initDatabase() (*mgo.Database, error) {
	DbSession, err := StartDbSession()
	if err != nil {
		return nil, err
	}

	DbConnect := DbSession.DB(config.DB)

	return DbConnect, nil
}

func DatabaseConnect() (*mgo.Database, error) {
	//This method handles connection to the mongo db database using the Singleton Pattern

	once.Do(func() {
		connection, err = initDatabase()
	})
	return connection, err
}

func InsertCollectionData(collection string, query *Topic) error {
	db, _ := DatabaseConnect()
	err = db.C(collection).Insert(query)
	return err
}

func FindCollectionData(collection, query string) []interface{} {
	var result []interface{}

	db, _ := DatabaseConnect()
	db.C(collection).Find(query).All(&result)

	return result
}
