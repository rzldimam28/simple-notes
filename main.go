package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rzldimam28/simple-notes/controllers"
)

func main() {

	r := mux.NewRouter()

	control := controllers.Handler{}

	r.HandleFunc("/users", control.Users)
	r.HandleFunc("/users/notes", control.Notes).Methods("GET")
	r.HandleFunc("/users/notes", control.Notes).Methods("POST")
	r.HandleFunc("/users/notes/{id}", control.Notes).Methods("PUT")
	r.HandleFunc("/users/notes/{id}", control.Notes).Methods("DELETE")
	r.HandleFunc("/users/notes/user/{id}", control.UserNote)

	fmt.Println("Server is running...")

	http.ListenAndServe("127.0.0.1:8080", r)
}