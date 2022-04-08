package main

import (
	"fmt"
	"net/http"

	"github.com/rzldimam28/simple-notes/controllers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	control := controllers.Handler{}

	http.HandleFunc("/users", control.Users)
	http.HandleFunc("/users/notes", control.Notes)
	http.HandleFunc("/users/notes/user", control.UserNote)

	http.ListenAndServe("127.0.0.1:8080", nil)
}