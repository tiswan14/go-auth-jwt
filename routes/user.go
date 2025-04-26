package routes

import (
	"go-auth-jwt/controllers"
	"go-auth-jwt/middlewares"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middlewares.Auth)

	router.HandleFunc("/me", controllers.Me).Methods("GET")
}
