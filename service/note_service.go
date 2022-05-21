package service

import (
	"context"

	"github.com/rzldimam28/simple-notes/models/web"
)

type NoteService interface {
	Create(ctx context.Context, request web.NoteCreateRequest, userId int) web.NoteResponse
	Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse
	Delete(ctx context.Context, noteId int)
	FindById(ctx context.Context, noteId int) web.NoteResponse
	FindAll(ctx context.Context) []web.NoteResponse
}