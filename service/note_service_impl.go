package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/repository"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	Validate *validator.Validate
}

func NewNoteService(noteRepository repository.NoteRepository, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		Validate: validate,
	}
}

func (noteService *NoteServiceImpl) Create(ctx context.Context, request web.NoteCreateRequest, userId int) web.NoteResponse {	
	err := noteService.Validate.Struct(request)
	helper.PanicIfError(err)

	note := entity.Note{
		UserId: userId,
		Title: request.Title,
		Content: request.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newNote := noteService.NoteRepository.Save(ctx, note)
	return helper.ToNoteResponse(newNote)
}

func (noteService *NoteServiceImpl) Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse {
	err := noteService.Validate.Struct(request)
	helper.PanicIfError(err)

	note, err := noteService.NoteRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)
	
	note.Title = request.Title
	note.Content = request.Content
	note.UpdatedAt = time.Now()
	
	updatedNote := noteService.NoteRepository.Update(ctx, note)
	return helper.ToNoteResponse(updatedNote)
}

func (noteService *NoteServiceImpl) Delete(ctx context.Context, noteId int) {
	note, err := noteService.NoteRepository.FindById(ctx, noteId)
	helper.PanicIfError(err)

	noteService.NoteRepository.Delete(ctx, note)
}

func (noteService *NoteServiceImpl) FindById(ctx context.Context, noteId int) web.NoteResponse {
	note, err := noteService.NoteRepository.FindById(ctx, noteId)
	helper.PanicIfError(err)

	return helper.ToNoteResponse(note)
}

func (noteService *NoteServiceImpl) FindAll(ctx context.Context) []web.NoteResponse {
	notes := noteService.NoteRepository.FindAll(ctx)

	return helper.ToNoteResponses(notes)
}