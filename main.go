package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rzldimam28/simple-notes/app"
	"github.com/rzldimam28/simple-notes/controller"
	"github.com/rzldimam28/simple-notes/middleware"
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

	loginController := controller.LoginController{UserService: &userService}

	r := mux.NewRouter()

	r.Use(middleware.PanicRecovery)

	// home
	r.HandleFunc("/", controller.HomeHandler).Methods("GET")

	// login
	r.HandleFunc("/login", loginController.LoginHandler).Methods("POST")

	// routing for users
	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", userController.FindAll).Methods("GET")
	userRouter.HandleFunc("", userController.Create).Methods("POST")

	// routing for notes
	notesRouter := r.PathPrefix("/notes").Subrouter()
	notesRouter.Use(middleware.Auth)
	notesRouter.HandleFunc("", noteController.FindAll).Methods("GET")
	notesRouter.HandleFunc("", noteController.Create).Methods("POST")
	notesRouter.HandleFunc("/{id}", noteController.FindById).Methods("GET")
	notesRouter.HandleFunc("/{id}", noteController.Update).Methods("PUT")
	notesRouter.HandleFunc("/{id}", noteController.Delete).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}