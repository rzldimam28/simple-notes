package web

type NoteCreateRequest struct {
	UserId int `json:"user_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}