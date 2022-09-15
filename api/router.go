package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/narayanprusty/average-blocks/api/auth"
)

func StartAPIServer() {
	router := mux.NewRouter()

	router.HandleFunc("/register", RegisterHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/key", auth.VerifyJWT(CreateKeyHandler)).Methods("GET")
	router.HandleFunc("/rate", FetchRate).Methods("GET")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))

}
