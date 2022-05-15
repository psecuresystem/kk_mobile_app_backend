package guards

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			http.Error(w, "No token found", 401)
			return
		}

		var mySigningKey = []byte("Super secret key")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			http.Error(w, "Your Token has been expired", 401)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r.Header.Set("Role", "user")
			handler.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Not Authorized", 401)
		return
	}
}
