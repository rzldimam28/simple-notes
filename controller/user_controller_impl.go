package controller

import (
	"net/http"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/service"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (userController *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(r, &userCreateRequest)
	userResponse := userController.UserService.Create(r.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: userResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (userController *UserControllerImpl) List(w http.ResponseWriter, r *http.Request) {
	userResponses := userController.UserService.Lists(r.Context())
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: userResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}