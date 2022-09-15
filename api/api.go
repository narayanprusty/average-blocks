package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/narayanprusty/average-blocks/db"
)

type JsonResponse struct {
	Type    string      `json:"type"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
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

	userExists := &db.User{Username: user.Username}
	err = db.DB.Model(userExists).Select()
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

	response := JsonResponse{Type: "success"}

	json.NewEncoder(w).Encode(response)
}
