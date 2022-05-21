package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/entity"
)

type NoteRepositoryImpl struct {
	DB *sql.DB
}

func NewNoteRepository(DB *sql.DB) NoteRepository {
	return &NoteRepositoryImpl{
		DB: DB,
	}
}

func (noteRepo *NoteRepositoryImpl) Save(ctx context.Context, note entity.Note) entity.Note {
	SQL := "INSERT INTO notes(user_id, title, content, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	result, err := noteRepo.DB.ExecContext(ctx, SQL, note.UserId, note.Title, note.Content, note.CreatedAt, note.UpdatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	note.Id = int(id)
	return note
}

func (noteRepo *NoteRepositoryImpl) Update(ctx context.Context, note entity.Note) entity.Note {
	SQL := "UPDATE notes SET title = ?, content = ?, updated_at = ? WHERE id = ?"
	_, err := noteRepo.DB.ExecContext(ctx, SQL, note.Title, note.Content, note.UpdatedAt, note.Id)
	helper.PanicIfError(err)
	return note
}

func (noteRepo *NoteRepositoryImpl) Delete(ctx context.Context, note entity.Note) {
	SQL := "DELETE FROM notes WHERE id = ?"
	_, err := noteRepo.DB.ExecContext(ctx, SQL, note.Id)
	helper.PanicIfError(err)
}

func (noteRepo *NoteRepositoryImpl) FindById(ctx context.Context, id int) (entity.Note, error) {
	SQL := "SELECT id, user_id, title, content, created_at, updated_at FROM notes WHERE id = ?"
	rows, err := noteRepo.DB.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var note entity.Note
	if rows.Next() {
		err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
		helper.PanicIfError(err)
		return note, nil
	} else {
		return note, errors.New("could not find note by id")
	}
}

func (noteRepo *NoteRepositoryImpl) FindAll(ctx context.Context) []entity.Note {
	SQL := "SELECT id, user_id, title, content, created_at, updated_at FROM notes"
	rows, err := noteRepo.DB.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var notes []entity.Note
	for rows.Next() {
		var note entity.Note
		err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
		helper.PanicIfError(err)
		notes = append(notes, note)
	}

	return notes
}