package repository

import (
	"context"

	"github.com/rzldimam28/simple-notes/models/entity"
)

type NoteRepository interface {
	Save(ctx context.Context, note entity.Note) entity.Note
	Update(ctx context.Context, note entity.Note) entity.Note
	Delete(ctx context.Context, note entity.Note)
	FindById(ctx context.Context, id int) (entity.Note, error)
	FindAll(ctx context.Context) []entity.Note
}