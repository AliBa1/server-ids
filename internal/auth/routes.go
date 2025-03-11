package auth

// define the routes and calls the handlers

import (
	"server-ids/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, m *middleware.Middleware, service *AuthService) {
	handler := NewAuthHandler(service)

	r.HandleFunc("/auth", m.ApplyMiddleware(handler.GetAuth)).Methods("GET")

	// r.HandleFunc("/login", handler.GetUsers).Methods("POST")
	// r.HandleFunc("/register", handler.GetUsers).Methods("POST")
}