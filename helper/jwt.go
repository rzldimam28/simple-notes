package helper

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int) (string, error) {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))

	claims := sign.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return "", errors.New("can not generate token")
	}
	return token, nil
}

// func GetUserID(tokenString string) int {
// 	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		if t.Method != jwt.GetSigningMethod("HS256") {
// 			return nil, errors.New("unexpected signing method")
// 		}
// 		return []byte("secret"), nil
// 	})
// 	PanicIfError(err)
// 	myToken := token.Claims.(jwt.MapClaims)
// 	userId := int(myToken["userId"].(float64))
// 	return userId
// }