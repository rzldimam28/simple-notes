package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models"
	"github.com/rzldimam28/simple-notes/service"
	"github.com/rzldimam28/simple-notes/web"
)

type Handler struct {
	UserService *service.UserService
	NoteService *service.NoteService
}

var(
	// Notes = []models.Note{}
	NotesById = []models.Note{}
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
		newUser := models.User{}
		err := json.NewDecoder(r.Body).Decode(&newUser)
		//todo check validate header/body
		if err != nil {
			helper.WriteToResponse(w, 400, "Could not create User", nil)
			return
		}
		user, err  := h.UserService.Create(newUser)

		if err != nil {
			helper.WriteToResponse(w, 500, "Server error", nil)
			return
		}

		helper.WriteToResponse(w, 201, "Success Create User", user)
		return
}

func (h *Handler) ListsUser(w http.ResponseWriter, r *http.Request) {
	users := h.UserService.Lists()
	helper.WriteToResponse(w, 201, "Success Create User", users)
	return
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	newNote := web.NotesRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&newNote); err != nil{
		helper.WriteToResponse(w, 500, err.Error(), nil)
		return 
	}
	userId := r.Header.Get("User_ID")
	note, err := h.NoteService.Create(newNote, userId)
	if err != nil {
		helper.WriteToResponse(w, 500, err.Error(), nil)
		return 
	}
	helper.WriteToResponse(w, 201,"Success Created New Note", note)
	return
}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	notes := h.NoteService.List()
	helper.WriteToResponse(w, 200, "Success Get All Notes", notes)
	return
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	note, err := h.NoteService.Get(id)
	if err != nil {
		helper.WriteToResponse(w, 500, err.Error(), nil)
		return
	}
	helper.WriteToResponse(w, 200, "Success Get Note", note)
	return
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	updateNoteRequest := web.NotesRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&updateNoteRequest)
	if err != nil {
		helper.WriteToResponse(w, 400, "Could Not Update Note", nil)
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	userId := r.Header.Get("User_ID")
	updatedNote, err := h.NoteService.Update(id, updateNoteRequest)
	if err != nil {
		helper.WriteToResponse(w, 500, err.Error(), nil)
	}
	if updatedNote.UserId != userId {
		helper.WriteToResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	helper.WriteToResponse(w, http.StatusOK, "Success Update Note", updatedNote)
	return
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userId := r.Header.Get("User_ID")
	note, err := h.NoteService.Get(id)
	if err != nil {
		helper.WriteToResponse(w, 400, err.Error(), nil)
		return
	}
	if note.UserId != userId {
		helper.WriteToResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	h.NoteService.Delete(id)
	helper.WriteToResponse(w, http.StatusOK, "Success Delete Note", nil)
	return
}