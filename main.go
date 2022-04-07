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

var(
	Notes = []Note{}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})		

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case "POST":
			newNote := RequestBody{}
			if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil{
				WriteToResponse(w, 500, err.Error(), nil)
				return 
			}
			id := uuid.NewString()

			note := Note{
				Id: id,
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
			
			for i, note := range Notes {
				if note.Id == id {
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
			for i, note := range Notes {
				if note.Id == id {
					Notes = append(Notes[:i], Notes[i+1:]...)
					WriteToResponse(w, 200,"Success Delete", id)
					return
				} 
			}	
			WriteToResponse(w, http.StatusNotFound, "not found", nil)
			return	
		}
		
	})
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
