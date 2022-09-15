package api

import (
	"encoding/json"
	random "math/rand"
	"net/http"
	"strings"

	"github.com/narayanprusty/average-blocks/api/auth"
	"github.com/narayanprusty/average-blocks/db"
	"github.com/narayanprusty/average-blocks/tracker"
)

type LoginResponse struct {
	JWTToken string `json:"jwtToken"`
}

type APIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

type RateResponse struct {
	Rate int `json:"rate"`
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
	username := r.Header["Username"][0]

	key := db.APIKey{}
	user := new(db.User)
	err := db.DB.Model(user).Where("username = ?", username).Select()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key.Key = randomString(32)
	key.UserId = user.Id

	_, err = db.DB.Model(&key).Insert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(APIKeyResponse{APIKey: key.Key})
}

func FetchRate(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header["Api-Key"][0]
	keyExists := new(db.APIKey)
	err := db.DB.Model(keyExists).Where("key = ?", apiKey).Select()
	if err != nil {
		http.Error(w, "invalid api key", http.StatusUnauthorized)
		return
	}

	rate := tracker.GetRate()
	json.NewEncoder(w).Encode(RateResponse{Rate: rate})
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[random.Intn(len(letters))]
	}
	return string(s)
}
