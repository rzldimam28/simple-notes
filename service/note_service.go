package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rzldimam28/simple-notes/models"
	"github.com/rzldimam28/simple-notes/web"
)

var Notes []models.Note

type NoteService struct {
}

func (n *NoteService) Create(newNote web.NotesRequestBody, userId string) (models.Note, error) {
	id := uuid.NewString()
	note := models.Note{
		Id:        id,
		UserId:    userId,
		Title:     newNote.Title,
		Content:   newNote.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	Notes = append(Notes, note)

	return note, nil
}

func (n *NoteService) List() []models.Note {
	return Notes
}

func (n *NoteService) Get(id string) (models.Note, error) {
	for _, note := range Notes {
		if note.Id == id {
			return note, nil
		}
	}
	return models.Note{}, errors.New("could not find note")
}

func (n *NoteService) Update(id string, updateNote web.NotesRequestBody) (models.Note, error) {
	for index, _ := range Notes {
		if Notes[index].Id == id {
			Notes[index].Title = updateNote.Title
			Notes[index].Content = updateNote.Body
			Notes[index].UpdatedAt = time.Now()
			return Notes[index], nil
		}
	}
	return models.Note{}, errors.New("could not update note")
}

func (n *NoteService) Delete(id string) {
	for index, note := range Notes {
		if note.Id == id {
			Notes = append(Notes[:index], Notes[index+1:]...)
			return
		}
	}
}