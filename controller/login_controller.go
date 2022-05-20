package controller

import (
	"net/http"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/middleware"
	"github.com/rzldimam28/simple-notes/models/entity"
	"github.com/rzldimam28/simple-notes/models/web"
	"github.com/rzldimam28/simple-notes/service"
)

type LoginController struct {
	UserService *service.UserService
}

func (loginController *LoginController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: "Home Page",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (loginController *LoginController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u entity.Credential
	helper.ReadFromRequestBody(r, &u)

	// id := r.Header.Get("User_ID")
	// idInt, _ := strconv.Atoi(id)
	user := loginController.UserService.Get(u.Username)

	if u.Username != user.Username || u.Password != user.Password {
		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "Unauth",
			Data: nil,
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}
	token, err := middleware.GenerateToken(user.Id)
	if err != nil {
		webResponse := web.WebResponse{
			Code: http.StatusUnprocessableEntity,
			Status: err.Error(),
			Data: nil,
		}
		helper.WriteToResponseBody(w, webResponse)
		return
	}
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: token,
	}
	helper.WriteToResponseBody(w, webResponse)
}