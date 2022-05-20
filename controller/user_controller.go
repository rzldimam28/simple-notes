package controller

import (
	"net/http"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/service"
)

type UserController struct {
	UserService *service.UserService
}

func (userController *UserController) Create(w http.ResponseWriter, r *http.Request) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(r, &userCreateRequest)

	userResponse := userController.UserService.Create(userCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusCreated,
		Status: "Success Create New User",
		Data: userResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	userResponses := userController.UserService.List()
	webResponses := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Get All Users",
		Data: userResponses,
	}
	helper.WriteToResponseBody(w, webResponses)
}