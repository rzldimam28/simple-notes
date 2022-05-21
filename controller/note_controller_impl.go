package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/service"
)

type NoteControllerImpl struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{
		NoteService: noteService,
	}
}

func (noteController *NoteControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	// id := r.Header.Get("User_ID")
	// userId, _ := strconv.Atoi(id)

	userId := r.Context().Value("userId").(int)

	noteCreateRequest := web.NoteCreateRequest{}
	helper.ReadFromRequestBody(r, &noteCreateRequest)

	noteResponse := noteController.NoteService.Create(r.Context(), noteCreateRequest, userId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["id"]
	noteId, _ := strconv.Atoi(idString)

	noteUpdateRequest := web.NoteUpdateRequest{}
	helper.ReadFromRequestBody(r, &noteUpdateRequest)

	noteUpdateRequest.Id = noteId

	noteResponse := noteController.NoteService.Update(r.Context(), noteUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idString := params["id"]
	noteId, _ := strconv.Atoi(idString)

	noteController.NoteService.Delete(r.Context(), noteId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: nil,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	noteId, _ := strconv.Atoi(id)

	noteResponse := noteController.NoteService.FindById(r.Context(), noteId)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (noteController *NoteControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	noteResponses := noteController.NoteService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: noteResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}
