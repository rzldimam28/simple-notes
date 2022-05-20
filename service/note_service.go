package service

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/repository"
)

type NoteService struct {
	NoteRepository *repository.NoteRepository
	Validate *validator.Validate
}

func (noteService *NoteService) Create(request web.NoteCreateRequest, userId int) web.NoteResponse {	
	err := noteService.Validate.Struct(request)
	helper.PanicIfError(err)

	note := entity.Note{
		UserId: userId,
		Title: request.Title,
		Content: request.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newNote := noteService.NoteRepository.Save(note)
	return helper.ToNoteResponse(newNote)
}

func (noteService *NoteService) Update(request web.NoteUpdateRequest) web.NoteResponse {
	err := noteService.Validate.Struct(request)
	helper.PanicIfError(err)

	note, err := noteService.NoteRepository.FindById(request.Id)
	helper.PanicIfError(err)
	
	note.Title = request.Title
	note.Content = request.Content
	note.UpdatedAt = time.Now()
	
	updatedNote := noteService.NoteRepository.Update(note)
	return helper.ToNoteResponse(updatedNote)
}

func (noteService *NoteService) Delete(noteId int) {
	note, err := noteService.NoteRepository.FindById(noteId)
	helper.PanicIfError(err)

	noteService.NoteRepository.Delete(note)
}

func (noteService *NoteService) FindById(noteId int) web.NoteResponse {
	note, err := noteService.NoteRepository.FindById(noteId)
	helper.PanicIfError(err)

	return helper.ToNoteResponse(note)
}

func (noteService *NoteService) FindAll() []web.NoteResponse {
	notes := noteService.NoteRepository.FindAll()

	return helper.ToNoteResponses(notes)
}