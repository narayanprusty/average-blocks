package auth

import (
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
	claims["exp"] = time.Now().Add(60 * time.Minute)

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			http.Error(w, "Token header not found", http.StatusUnauthorized)
			return
		}

		signingKey := []byte(config.Config.JWTSecret)
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return signingKey, nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["username"])
			endpointHandler(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	})
}
