package main

import (
	routers "awesomeProject1/routers"
	"github.com/go-noodle/adapt/gorilla"
	mw "github.com/go-noodle/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	n := mw.Default(gorilla.Vars)

	// Routers
	r.Handle("/", n.Then(routers.Index))
	r.Handle("/book/{id}", n.Then(routers.GetBook)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
