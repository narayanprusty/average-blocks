package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/narayanprusty/average-blocks/api/auth"
	"github.com/narayanprusty/average-blocks/db"
)

type LoginResponse struct {
	JWTToken string `json:"jwtToken"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user *db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "missing username or password", http.StatusBadRequest)
		return
	}

	userExists := new(db.User)
	err = db.DB.Model(userExists).Where("username = ?", user.Username).Select()
	if err != nil {
		if !strings.Contains(err.Error(), "no rows") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if userExists.Password != "" {
		http.Error(w, "user exists", http.StatusConflict)
		return
	}

	_, err = db.DB.Model(user).Insert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user *db.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "missing username or password", http.StatusBadRequest)
		return
	}

	userExists := new(db.User)
	err = db.DB.Model(userExists).Where("username = ?", user.Username).Select()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Password != userExists.Password {
		http.Error(w, "invalid credentails", http.StatusBadRequest)
		return
	}
	jwt, err := auth.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(LoginResponse{JWTToken: jwt})
}

func CreateKeyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("success")
}
