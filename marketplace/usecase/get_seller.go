package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/rwirdemann/bffdashboard/marketplace/domain"
	"net/http"
	"regexp"
)

var sampleSecretKey2 = []byte("SecretYouShouldHide")

func SellerHandler(writer http.ResponseWriter, request *http.Request) {
	if !isValidJWT(extractJwtFromHeader(request.Header)) {
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	seller := domain.Seller{Name: "Ralf"}
	jsonString, _ := json.Marshal(seller)
	writer.Write(jsonString)
}

func extractJwtFromHeader(header http.Header) (jwt string) {
	var jwtRegex = regexp.MustCompile(`^Bearer (\S+)$`)

	if val, ok := header["Authorization"]; ok {
		for _, value := range val {
			if result := jwtRegex.FindStringSubmatch(value); result != nil {
				jwt = result[1]
				return
			}
		}
	}

	return
}

func isValidJWT(tokenString string) bool {
	vars, err := godotenv.Read(".env")
	if err != nil {
		println("Failed to load .env file")
		return false
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretKey := vars["SECRET_KEY"]
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}
