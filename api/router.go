package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartAPIServer() {
	router := mux.NewRouter()

	router.HandleFunc("/register", RegisterHandler).Methods("POST")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
