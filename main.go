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
	noteService := service.NoteService{NoteRepository: &noteRepo, UserRepository: &userRepo, Validate: validate}
	noteController := controller.NoteController{NoteService: &noteService}

	r := mux.NewRouter()

	r.Use(PanicRecovery)
	
	// routing for users
	r.HandleFunc("/users", userController.List).Methods("GET")
	r.HandleFunc("/users", userController.Create).Methods("POST")
	// routing for notes
	r.HandleFunc("/users/notes", noteController.FindAll).Methods("GET")
	r.HandleFunc("/users/notes", noteController.Create).Methods("POST")
	r.HandleFunc("/users/notes/{id}", noteController.FindById).Methods("GET")
	r.HandleFunc("/users/notes/{id}", noteController.Update).Methods("PUT")
	r.HandleFunc("/users/notes/{id}", noteController.Delete).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

func PanicRecovery(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Fprintln(w, err)
			}
		}()

		h.ServeHTTP(w,r)
	})
}

// func Auth(h http.Handler) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		userId := r.Header.Get("User_ID")
// 		password := r.Header.Get("Password")
		
// 		if userId == "1"{
// 			if password != "imam"{
// 				fmt.Fprintln(w,"Salah Password")
// 				return
// 			}
// 		}
// 		h.ServeHTTP(w,r)
// 	})
// }