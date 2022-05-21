package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/web"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Auth")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.GetSigningMethod("HS256") {
				return nil, errors.New("unexpected signing method")
			}
			return []byte("secret"), nil
		})		
		if token != nil && err == nil {
			myToken := token.Claims.(jwt.MapClaims)["userId"]
			userId := int(myToken.(float64))
			ctx := r.Context()
			userCtx := context.WithValue(ctx, "userId", userId)
			next.ServeHTTP(w, r.WithContext(userCtx))
		} else {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauth",
				Data:   nil,
			}
			helper.WriteToResponseBody(w, webResponse)
		}
	})
}
