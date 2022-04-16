package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rzldimam28/simple-notes/app"
	"github.com/rzldimam28/simple-notes/controller"
	"github.com/rzldimam28/simple-notes/repository"
	"github.com/rzldimam28/simple-notes/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	userRepo := repository.UserRepository{DB: db}
	userService := service.UserService{UserRepository: &userRepo, Validate: validate}
	userController := controller.UserController{UserService: &userService}

	noteRepo := repository.NoteRepository{DB: db}
	noteService := service.NoteService{NoteRepository: &noteRepo, Validate: validate}
	noteController := controller.NoteController{NoteService: &noteService}

	r := mux.NewRouter()
	
	r.HandleFunc("/users", userController.List).Methods("GET")
	// r.HandleFunc("/users/{id}", userController.Get).Methods("GET")
	r.HandleFunc("/users", userController.Create).Methods("POST")
	r.HandleFunc("/users/notes", noteController.FindAll).Methods("GET")
	r.HandleFunc("/users/notes", noteController.Create).Methods("POST")
	r.HandleFunc("/users/notes/{id}", noteController.FindById).Methods("GET")
	r.HandleFunc("/users/notes/{id}", noteController.Update).Methods("PUT")
	r.HandleFunc("/users/notes/{id}", noteController.Delete).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}