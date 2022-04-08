package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rzldimam28/simple-notes/models"
)

type Handler struct {
}

type RequestBody struct{
	Title string `json:"title"`
	Body string `json:"body"`
}

var(
	Notes = []models.Note{}
	NotesById = []models.Note{}
	Users = []models.User{}
)

func (h *Handler) Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		newUser := models.User{}
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			WriteToResponse(w, 400, "Could not create User", nil)
			return
		}
		id := uuid.NewString()

		user := models.User{
			Id:        id,
			FirstName: newUser.FirstName,
			LastName:  newUser.LastName,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		Users = append(Users, user)

		WriteToResponse(w, 201, "Success Create User", user)
	case "GET":
		WriteToResponse(w, 200, "Success Get User", Users)
		return
	}
}

func (h *Handler) Notes(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case "POST":
			newNote := RequestBody{}
			if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil{
				WriteToResponse(w, 500, err.Error(), nil)
				return 
			}
			id := uuid.NewString()
			userId := r.Header.Get("User_ID")

			note := models.Note{
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
}


func (h *Handler) UserNote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// do something here
		var userId string
		userId = r.URL.Query().Get("id")
		notes := []models.Note{}
		for i,v := range Notes{
			if v.UserId == userId{
				notes = append(notes, Notes[i])
			}
		}
		WriteToResponse(w, 200, "Success get all notes", notes)
	}
}

