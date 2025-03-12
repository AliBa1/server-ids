package auth

// define the routes and calls the handlers

import (
	"server-ids/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, m *middleware.Middleware, service *AuthService) {
	handler := NewAuthHandler(service)

	r.HandleFunc("/auth", m.ApplyMiddleware(handler.GetAuth)).Methods("GET")

	r.HandleFunc("/auth/login", handler.PostLogin).Methods("POST")
	r.HandleFunc("/auth/register", handler.PostRegister).Methods("POST")
	r.HandleFunc("/auth/users", handler.GetUsers).Methods("GET")
}
