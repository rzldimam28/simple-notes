package controller

import (
	"net/http"

	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/web"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: "Home Page",
	}
	helper.WriteToResponseBody(w, webResponse)
}