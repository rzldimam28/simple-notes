package web

type UserCreateRequest struct {
	Username string `validate:"required" json:"username"`
	Password  string `validate:"required" json:"password"`
}