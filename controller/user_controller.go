package controller

import "net/http"

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}