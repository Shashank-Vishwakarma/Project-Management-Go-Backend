package routes

import (
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/handlers"
	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.UserLoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", handlers.UserRegistrationHandler).Methods(http.MethodPost)
	router.HandleFunc("/logout", handlers.UserLogoutHandler).Methods(http.MethodPost)
}
