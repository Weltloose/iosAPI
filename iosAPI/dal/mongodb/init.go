package mongodb

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database
var userInfoC *mgo.Collection
var groupInfoC *mgo.Collection
var MainHub *Hub

func init() {
	var err error
	var session *mgo.Session
	session, err = mgo.Dial("mongo:27017")
	if err != nil {
		fmt.Println("dial error, ", err)
	}
	db := session.DB("ios")
	userInfoC = db.C("userInfo")
	groupInfoC = db.C("groupInfo")
	MainHub = newHub()
}
