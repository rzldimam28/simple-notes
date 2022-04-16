package service

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/repository"
)

// import (
// 	"errors"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/rzldimam28/simple-notes/models"
// 	"github.com/rzldimam28/simple-notes/web"
// )

// var Notes []models.Note

// type NoteService struct {
// }

// func (n *NoteService) Create(newNote web.NotesRequestBody, userId string) (models.Note, error) {
// 	id := uuid.NewString()
// 	note := models.Note{
// 		Id:        id,
// 		UserId:    userId,
// 		Title:     newNote.Title,
// 		Content:   newNote.Body,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	Notes = append(Notes, note)

// 	return note, nil
// }

// func (n *NoteService) List() []models.Note {
// 	return Notes
// }

// func (n *NoteService) Get(id string) (models.Note, error) {
// 	for _, note := range Notes {
// 		if note.Id == id {
// 			return note, nil
// 		}
// 	}
// 	return models.Note{}, errors.New("could not find note")
// }

// func (n *NoteService) Update(id string, updateNote web.NotesRequestBody) (models.Note, error) {
// 	for index, _ := range Notes {
// 		if Notes[index].Id == id {
// 			Notes[index].Title = updateNote.Title
// 			Notes[index].Content = updateNote.Body
// 			Notes[index].UpdatedAt = time.Now()
// 			return Notes[index], nil
// 		}
// 	}
// 	return models.Note{}, errors.New("could not update note")
// }

// func (n *NoteService) Delete(id string) {
// 	for index, note := range Notes {
// 		if note.Id == id {
// 			Notes = append(Notes[:index], Notes[index+1:]...)
// 			return
// 		}
// 	}
// }

type NoteService struct {
	UserRepository *repository.UserRepository
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

func (noteService *NoteService) Update(request web.NoteUpdateRequest, userId int) web.NoteResponse {
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

func (noteService *NoteService) Delete(noteId int, userId int) {
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