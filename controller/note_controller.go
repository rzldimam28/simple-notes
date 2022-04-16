package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/service"
)

type NoteController struct {
	NoteService *service.NoteService
}

func (noteController *NoteController) Create(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User_ID")
	userId, _ := strconv.Atoi(id)

	noteCreateRequest := web.NoteCreateRequest{}
	helper.ReadFromRequestBody(r, &noteCreateRequest)

	noteResponse := noteController.NoteService.Create(noteCreateRequest, userId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User_ID")
	userId, _ := strconv.Atoi(id)
	
	params := mux.Vars(r)
	idString := params["id"]
	noteId, _ := strconv.Atoi(idString)
	
	noteUpdateRequest := web.NoteUpdateRequest{}
	helper.ReadFromRequestBody(r, &noteUpdateRequest)

	noteUpdateRequest.Id = noteId

	noteResponse := noteController.NoteService.Update(noteUpdateRequest, userId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("User_ID")
	userId, _ := strconv.Atoi(id)

	params := mux.Vars(r)
	idString := params["id"]
	noteId, _ := strconv.Atoi(idString)

	noteController.NoteService.Delete(noteId, userId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteController) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	noteId, _ := strconv.Atoi(id)

	noteResponse := noteController.NoteService.FindById(noteId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteController) FindAll(w http.ResponseWriter, r *http.Request) {
	noteResponses := noteController.NoteService.FindAll()
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}
