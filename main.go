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
	r := mux.NewRouter()

	n := mw.Default(gorilla.Vars)

	r.Handle("/", n.Then(routers.Index)).Methods("GET")

	r.Handle("/book/{id}", n.Then(routers.GetBook)).Methods("GET")
	r.Handle("/book", n.Then(routers.PostBook)).Methods("POST")
	r.Handle("/book/{id}", n.Then(routers.DeleteBook)).Methods("DELETE")
	r.Handle("/book", n.Then(routers.PutBook)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
