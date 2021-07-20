package routers

import (
	"encoding/json"
	"github.com/Madaspe/simple-messanger-go-api/models"
	"github.com/go-noodle/adapt/gorilla"
	"log"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := gorilla.GetVars(r)
	id, _ := strconv.Atoi(vars["id"])

	book := models.Book{Id: int64(id)}
	err := json.NewEncoder(w).Encode(book)

	if err != nil {
		log.Fatal(err)
	}
}
