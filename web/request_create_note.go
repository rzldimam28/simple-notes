package web

type NotesRequestBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
