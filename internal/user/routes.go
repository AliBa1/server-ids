package user

import (
	"server-ids/internal/middleware"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, m *middleware.Middleware, service *UserService) {
	handler := NewUserHandler(service)

	// r.HandleFunc("/user", m.ApplyMiddleware(handler.GetUser)).Methods("GET")
	r.HandleFunc("/users/{username}/role", m.Authorization(m.ApplyMiddleware(handler.UpdateRole), handler.service.db.Sessions)).Methods("PATCH")
}
