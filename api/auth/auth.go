package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/narayanprusty/average-blocks/config"
)

func GenerateJWT(username string) (string, error) {
	signingKey := []byte(config.Config.JWTSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func fetchUsername(jwtToken string) (string, error) {
	var mySigningKey = []byte(config.Config.JWTSecret)

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%v", claims["username"]), nil
	}

	return "", errors.New("invalid token")
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			username, err := fetchUsername(request.Header["Token"][0])

			if err != nil {
				http.Error(writer, err.Error(), http.StatusUnauthorized)
				return
			}

			request.Header.Set("Username", username)
			endpointHandler(writer, request)
		} else {
			http.Error(writer, "Token header is missing", http.StatusUnauthorized)
		}
	})
}
