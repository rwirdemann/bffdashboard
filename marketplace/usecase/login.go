package usecase

import (
	"github.com/golang-jwt/jwt"
	"net/http"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	tokenString, _ := generateJWT()
	writer.Write([]byte(tokenString))
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
