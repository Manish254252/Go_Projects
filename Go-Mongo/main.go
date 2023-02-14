package main

import (
	"net/http"

	"example.com/Go-Mongo/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUsercontroller(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb+srv://root:root@manish.yyx5jvu.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		panic(err)
	}
	return s

}
