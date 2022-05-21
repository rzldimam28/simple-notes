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

func (userController *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var u web.Credentials
	helper.ReadFromRequestBody(r, &u)
	user := userController.UserService.GetByUsername(r.Context(), u.Username)
	if u.Username != user.Username || u.Password != user.Password {
		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "Wrong Username of Password",
			Data: nil,
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}
	token, err := helper.GenerateToken(user.Id)
	if err != nil {
		webResponse := web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: "Can Not Generate Token",
			Data: nil,
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "Success Generate Token",
		Data: token,
	}
	helper.WriteToResponseBody(w, webResponse)
}