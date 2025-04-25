package main

import (
	"go-auth-jwt/configs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.Path("/api").Subrouter()

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
