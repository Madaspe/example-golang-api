package routers

import (
	f "fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, err := f.Fprintln(w, "Hello world")

	if err != nil {
		return
	}
}
