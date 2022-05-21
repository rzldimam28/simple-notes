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
	
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)
	
	noteRepository := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(noteRepository, validate)
	noteController := controller.NewNoteController(noteService)
	
	r := mux.NewRouter()
	
	r.Use(middleware.PanicRecovery)
	
	// routing for users
	userRoutes := r.PathPrefix("/users").Subrouter()
	userRoutes.HandleFunc("/login", userController.Login).Methods("POST")
	userRoutes.HandleFunc("", userController.Create).Methods("POST")
	userRoutes.HandleFunc("", userController.List).Methods("GET")

	// routing for notes
	noteRoutes := r.PathPrefix("/notes").Subrouter()
	noteRoutes.Use(middleware.Auth)
	noteRoutes.HandleFunc("", noteController.FindAll).Methods("GET")
	noteRoutes.HandleFunc("", noteController.Create).Methods("POST")
	noteRoutes.HandleFunc("/{id}", noteController.FindById).Methods("GET")
	noteRoutes.HandleFunc("/{id}", noteController.Update).Methods("PUT")
	noteRoutes.HandleFunc("/{id}", noteController.Delete).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}