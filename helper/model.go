package helper

import (
	"github.com/rzldimam28/simple-notes/models/entity"
	"github.com/rzldimam28/simple-notes/models/web"
)

func ToUserResponse(user entity.User) web.UserResponse {
	return web.UserResponse{
		Id: user.Id,
		Username: user.Username,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []entity.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToNoteResponse(note entity.Note) web.NoteResponse {
	return web.NoteResponse{
		Id: note.Id,
		UserId: note.UserId,
		Title: note.Title,
		Content: note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ToNoteResponses(notes []entity.Note) []web.NoteResponse {
	var noteResponses []web.NoteResponse
	for _, note := range notes {
		noteResponses = append(noteResponses, ToNoteResponse(note))
	}
	return noteResponses
}