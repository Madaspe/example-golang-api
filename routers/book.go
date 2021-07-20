package routers

import (
	"encoding/json"
	"github.com/Madaspe/simple-messanger-go-api/models"
	"github.com/go-noodle/adapt/gorilla"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var books []*models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := gorilla.GetVars(r)
	id, _ := strconv.Atoi(vars["id"])

	neededBook := new(models.Book)
	for _, book := range books {
		if book.Id == int64(id) {
			neededBook = book
		}
	}

	if len(neededBook.Title) == 0 {
		err := json.NewEncoder(w).Encode(models.Message{Message: "not found book"})
		w.Header()
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	err := json.NewEncoder(w).Encode(neededBook)

	if err != nil {
		log.Fatal(err)
	}
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	book := new(models.Book)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	err := json.Unmarshal(body, &book)

	if err != nil {
		log.Fatal(err)
	}

	for _, bookExist := range books {
		if bookExist.Id == book.Id {
			err := json.NewEncoder(w).Encode(models.Message{Message: "already exist"})

			if err != nil {
				log.Fatal(err)
			}

			return
		}
	}

	books = append(books, book)

	err = json.NewEncoder(w).Encode(models.Message{Message: "ok"})
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := gorilla.GetVars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, book := range books {
		if book.Id == int64(id) {
			books = append(books[:index], books[index+1:]...)

			err := json.NewEncoder(w).Encode(models.Message{Message: "ok"})

			if err != nil {
				log.Fatal(err)
			}

			return
		}
	}

	err := json.NewEncoder(w).Encode(models.Message{Message: "not found"})

	if err != nil {
		log.Fatal(err)
	}
}

func PutBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	book := new(models.Book)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	err := json.Unmarshal(body, &book)

	if err != nil {
		log.Fatal(err)
	}

	for index, bookExist := range books {
		if bookExist.Id == book.Id {
			books[index] = book
			err := json.NewEncoder(w).Encode(models.Message{Message: "ok"})

			if err != nil {
				log.Fatal(err)
			}

			return
		}
	}

	err = json.NewEncoder(w).Encode(models.Message{Message: "not found"})
	if err != nil {
		log.Fatal(err)
	}
}
