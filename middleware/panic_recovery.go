package middleware

import (
	"fmt"
	"net/http"
)

func PanicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Fprintln(w, err)
			}
		}()

		h.ServeHTTP(w, r)
	})
}

// func Auth(h http.Handler) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		userId := r.Header.Get("User_ID")
// 		password := r.Header.Get("Password")

// 		if userId == "1"{
// 			if password != "imam"{
// 				fmt.Fprintln(w,"Salah Password")
// 				return
// 			}
// 		}
// 		h.ServeHTTP(w,r)
// 	})
// }
