package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func GetUser(headerToken string) (string, string) {
	token, err := jwt.Parse(headerToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return []byte("Super secret key"), nil
	})

	if err != nil {
		return "", "Error occured"
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var email string = fmt.Sprint(claims["email"])
		return email, ""
	}

	return "", "Token Expired"
}
