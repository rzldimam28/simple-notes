package helper

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func GetUserId(w http.ResponseWriter, r *http.Request) int {
	tokenString := r.Header.Get("Auth")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.GetSigningMethod("HS256") {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret-key"), nil
	})
	PanicIfError(err)
	myToken := token.Claims.(jwt.MapClaims)
	userId := int(myToken["userId"].(float64))
	return userId
}