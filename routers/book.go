package routers

import (
	f "fmt"
	"github.com/go-noodle/adapt/gorilla"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := gorilla.GetVars(r)
	_, err := f.Fprintf(w, "%s", vars["id"])

	if err != nil {
		return
	}
}
