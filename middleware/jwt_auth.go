package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rzldimam28/simple-notes/helper"
	"github.com/rzldimam28/simple-notes/models/web"
)

func GenerateToken(userId int) (string, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	// create payload
	claims := sign.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token, err := sign.SignedString([]byte("secret-key"))
	if err != nil {
		return "", errors.New("Could not Generate Token")
	}
	return token, nil
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Auth")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.GetSigningMethod("HS256") {
				return nil, errors.New("Unexpected Signing Method")
			}
			return []byte("secret-key"), nil
		})
		if token != nil && err == nil {
			next.ServeHTTP(w, r)
		} else {
			webResponse := web.WebResponse{
				Code: http.StatusUnauthorized,
				Status: "Unauthorized",
				Data: nil,
			}
			helper.WriteToResponseBody(w, webResponse)
		}

	})
}