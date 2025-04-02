package user

import (
	"server-ids/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, m *middleware.Middleware, service *UserService) {
	handler := NewUserHandler(service)

	// r.HandleFunc("/user", m.ApplyMiddleware(handler.GetUser)).Methods("GET")
	r.HandleFunc("/users/{username}/role", m.ApplyMiddleware(handler.UpdateRole)).Methods("PATCH")
}
