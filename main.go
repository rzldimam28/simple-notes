package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
	Title     string `json:"title"`
	Content      string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RequestBody struct{
	Title string `json:"title"`
	Body string `json:"body"`
}

type Response struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

type User struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var(
	Notes = []Note{}
	NotesById = []Note{}
	Users = []User{}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})		

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "POST":
			newUser := User{}
			err := json.NewDecoder(r.Body).Decode(&newUser)
			if err != nil {
				WriteToResponse(w, 400, "Could not create User", nil)
				return
			}
			id := uuid.NewString()
			
			user := User{
				Id: id,
				FirstName: newUser.FirstName,
				LastName: newUser.LastName,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			Users = append(Users, user)

			WriteToResponse(w, 201, "Success Create User", user)
		case "GET":
			WriteToResponse(w, 200, "Success Get User", Users)
			return
		}
	})

	http.HandleFunc("/users/notes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case "POST":
			newNote := RequestBody{}
			if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil{
				WriteToResponse(w, 500, err.Error(), nil)
				return 
			}
			id := uuid.NewString()
			userId := r.Header.Get("User_ID")

			note := Note{
				Id: id,
				UserId: userId,
				Title: newNote.Title,
				Content: newNote.Body,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			Notes = append(Notes, note)

			WriteToResponse(w, 201,"Success Created New Note", note)
			return
		case "GET":
			WriteToResponse(w, 200, "Success Get All Notes", Notes)
			return
		case "PUT":
			updateNote := RequestBody{}
			if err := json.NewDecoder(r.Body).Decode(&updateNote); err != nil{
				WriteToResponse(w, 500, err.Error(), nil)
				return
			}
			id := r.URL.Query().Get("id")
			userId := r.Header.Get("User_ID")
			
			for i, note := range Notes {
				if note.Id == id {
					if note.UserId != userId{
						WriteToResponse(w, 400, "Bad Request", nil)
						return
					}
					Notes[i].Title = updateNote.Title
					Notes[i].Content = updateNote.Body			
					WriteToResponse(w, 200,"Success Update Note", Notes[i])
					return
				} 
			}

			WriteToResponse(w, http.StatusNotFound, "not found", nil)
			return
		case "DELETE":
			id := r.URL.Query().Get("id")	
			userId := r.Header.Get("User_ID")
			for i, note := range Notes {
				if note.Id == id {
					if note.UserId != userId{
						WriteToResponse(w, 400, "Bad Request", nil)
						return
					}
					Notes = append(Notes[:i], Notes[i+1:]...)
					WriteToResponse(w, 200,"Success Delete", id)
					return
				} 
			}	
			WriteToResponse(w, http.StatusNotFound, "not found", nil)
			return	
		}
		
	})

	http.HandleFunc("/users/notes/user", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case "GET":
			// do something here
			var userId string
			userId = r.URL.Query().Get("id")
			notes := []Note{}
			for i,v := range Notes{
				if v.UserId == userId{
					notes = append(notes, Notes[i])
				}
			}
			WriteToResponse(w, 200, "Success get all notes", notes)
		}
	})

	// http.HandleFunc("/users/notes", func(w http.ResponseWriter, r *http.Request){
	// 	switch r.Method {
	// 	case "POST":
	// 		var userId string
	// 		userId = r.URL.Query().Get("userid")

	// 		for _, user := range Users {
	// 			if user.Id == userId && len(userId) > 0 {
	// 				newNote := RequestBody{}
	// 				if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil{
	// 					WriteToResponse(w, 500, err.Error(), nil)
	// 					return 
	// 				}
	// 				noteId := uuid.NewString()
		
	// 				note := Note{
	// 					Id: noteId,
	// 					Title: newNote.Title,
	// 					Content: newNote.Body,
	// 					CreatedAt: time.Now(),
	// 					UpdatedAt: time.Now(),
	// 					UserId: userId,
	// 				}
		
	// 				Notes = append(Notes, note)
		
	// 				WriteToResponse(w, 201,"Success Created New Note", note)
	// 				return
	// 			}
	// 		}
	// 		WriteToResponse(w, http.StatusUnauthorized, "Please Create an Account First", nil)
	// 		return
	// 	case "GET":
	// 		// do something here
	// 		WriteToResponse(w, 200, "Success get all notes", Notes)
	// 	}
	// })

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func WriteToResponse(w http.ResponseWriter, codeStatus int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codeStatus)
	
	response := Response{
		Code: codeStatus,
		Status: message,
		Data: data,
	}

	json.NewEncoder(w).Encode(response)
}


