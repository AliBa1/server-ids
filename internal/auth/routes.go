package auth

// define the routes and calls the handlers

import (
	"server-ids/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, m *middleware.Middleware, service *AuthService) {
	handler := NewAuthHandler(service)

	r.HandleFunc("/auth", m.ApplyMiddleware(handler.GetAuth)).Methods("GET")
	r.HandleFunc("/auth/login", m.ApplyMiddleware(handler.PostLogin)).Methods("POST")
	r.HandleFunc("/auth/register", m.ApplyMiddleware(handler.PostRegister)).Methods("POST")
	// maybe change this to users and move it to user routes
	r.HandleFunc("/auth/users", m.ApplyMiddleware(handler.GetUsers)).Methods("GET")
}
