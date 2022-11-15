package mongodb

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func New(url string) *mgo.Database {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("You are connected to your mongo database")
	return session.DB("mandaro")
}
