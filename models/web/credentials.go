package web

type Credentials struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}