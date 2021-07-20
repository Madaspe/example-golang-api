package main

import (
	routers "github.com/Madaspe/simple-messanger-go-api/routers"
	"github.com/go-noodle/adapt/gorilla"
	mw "github.com/go-noodle/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	n := mw.Default(gorilla.Vars)

	router.Handle("/", n.Then(routers.Index)).Methods("GET")

	router.Handle("/book/{id}", n.Then(routers.GetBook)).Methods("GET")
	router.Handle("/book", n.Then(routers.PostBook)).Methods("POST")
	router.Handle("/book/{id}", n.Then(routers.DeleteBook)).Methods("DELETE")
	router.Handle("/book", n.Then(routers.PutBook)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
