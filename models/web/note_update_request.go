package web

type NoteUpdateRequest struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Title     string  `json:"title" validate:"required"`
	Content   string  `json:"content" validate:"required"`
}