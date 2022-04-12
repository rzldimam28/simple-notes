package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct{
	Code int `json:"code"`
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

func WriteToResponse(w http.ResponseWriter, codeStatus int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(codeStatus)

	response := Response{
		Code:   codeStatus,
		Status: message,
		Data:   data,
	}

	json.NewEncoder(w).Encode(response)
}