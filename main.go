package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rzldimam28/simple-notes/controllers"
	"github.com/rzldimam28/simple-notes/service"
)

func main() {
	r := mux.NewRouter()

	control := controllers.Handler{
		UserService: &service.UserService{},
	}

	r.HandleFunc("/users", control.ListsUser).Methods("GET")
	r.HandleFunc("/users", control.CreateUser).Methods("POST")
	r.HandleFunc("/users/notes", control.ListNotes).Methods("GET")
	r.HandleFunc("/users/notes", control.CreateNote).Methods("POST")
	r.HandleFunc("/users/notes/{id}", control.GetNote).Methods("GET")
	r.HandleFunc("/users/notes/{id}", control.UpdateNote).Methods("PUT")
	r.HandleFunc("/users/notes/{id}", control.DeleteNote).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}